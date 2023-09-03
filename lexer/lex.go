// Package lexer implements the lexical analyzer for λ.c
//
// Overview:
//
// This package encapsulates the functionality required for tokenizing the source code of a
// minimalist lambda calculus language. The lexing process begins with raw textual input
// and converts it into a sequence of tokens, which are the atomic units for the subsequent parsing stage.
//
// The Lexer type is the principal entity that performs lexical analysis. It maintains the current
// state, including the position in the source code and a queue of upcoming characters.
//
// Formal Syntax:
//
// The lexer formalizes the following language syntax through tokenization:
//
//	Term       ::= Identifier | "(" Term ")" | "\" Identifier "." Term
//	Expression ::= Term { Term }
//
// Tokens:
//
// The following token categories are defined for the lexer:
//
//   - IDENTIFIER:  Any sequence of Unicode graphical characters, excluding specific reserved characters.
//   - LAMBDA:      The backslash ("\") symbol representing the lambda function.
//   - DOT:         The dot (".") symbol used in lambda abstractions.
//   - LPAREN:      The left parenthesis ("(") symbol.
//   - RPAREN:      The right parenthesis (")") symbol.
//   - ILLEGAL:     Any unrecognized sequence of characters.
//   - EOF:         End of the file.
//
// Additionally, the lexer supports special constructs like strings with escape sequences, composite
// operators like ":=" and "->", and line breaks.
//
// Features:
//
//  1. Unicode Support: Identifiers can include any graphical Unicode character, except for the defined
//     reserved characters.
//  2. Monadic Parsing: Utilizes monads for optionally storing tokens, providing for cleaner code and
//     better error handling.
//  3. Recursive Lexing: Many lexer functions, such as stringLexer and identifierLexer, are implemented
//     using recursive techniques for simplicity and maintainability.
//
// Usage:
//
// Initialize a Lexer with the input source code, then iteratively call its NextToken method to
// retrieve the tokens one by one until an EOF token is returned.
//
// Example:
//
//	l := lexer.New(input)
//	for {
//	  tok := l.Next()
//	  if tok.Type == EOF {
//	    break
//	  }
//	  // Process tok
//	}
package lexer

import (
	"strings"
	"unicode"
	"unicode/utf8"

	"github.com/denisdubochevalier/monad"
)

// TokenType represents the type of a token
type TokenType int

const (
	ILLEGAL TokenType = iota // ILLEGAL represents an illegal character.
	EOF                      // EOF (End Of File) represents the end of the file that is being parsed.
	EOL                      // EOL (End Of Line) represents the end of a line.
	IDENT                    // IDENT represents an identifier, which could be a variable or a function name (e.g., m, recursive_fact, fact, ...).
	ASSIGN                   // ASSIGN represents the assignment operator (:=).
	MODULE                   // MODULE represents the module operator (|).
	NSDEREF                  // NSDEREF represents the namespace dereference operator (->).
	LAMBDA                   // LAMBDA represents the lambda operator (\).
	DOT                      // DOT represents the dot operator (.).
	STRING                   // STRING represents a string literal (e.g., "github.com/foo/bar", "text/lexer", ...).
	LPAREN                   // LPAREN represents the left parenthesis (().
	RPAREN                   // RPAREN represents the right parenthesis ()).
)

var values = []string{
	ILLEGAL: "ILLEGAL",
	EOF:     "EOF",
	EOL:     "EOL",
	IDENT:   "IDENT",
	STRING:  "STRING",
	LAMBDA:  "\\",
	DOT:     ".",
	LPAREN:  "(",
	RPAREN:  ")",
	MODULE:  "|",
	NSDEREF: "->",
	ASSIGN:  ":=",
}

// String returns a string representation of the TokenType.
// This is helpful for debugging and logging purposes,
// as it allows you to print out a human-readable version of the token type.
func (t TokenType) String() string {
	// handle unknown values
	if int(t)+1 > len(values) {
		return "UNKNOWN"
	}
	return values[t]
}

// Position represents a position in the text being parsed.
// It contains two fields: Row and Col, representing the current row and column, respectively.
type Position struct {
	Row int // Row represents the current row in the text.
	Col int // Col represents the current column in the text.
}

// NewRow increments the Row field of the Position by one and sets the Col field to 0.
// This represents moving to the start of a new row in the text.
func (p Position) NewRow() Position {
	p.Row++
	p.Col = 0
	return p
}

// Literal is a type used to store the value of a token.
// It is a simple wrapper around the string type.
type Literal string

// String returns the string representation of the Literal.
// This method is used to implement the fmt.Stringer interface.
func (l Literal) String() string {
	return string(l)
}

// Token represents a lexeme or a sequence of characters that have a collective meaning.
// It contains the type of the token (e.g. IDENT, ASSIGN, etc.), the position
// in the input where the token starts, and the literal value of the token.
type Token struct {
	TokenType TokenType
	Position  Position
	Literal   Literal
}

// Lexer is the main struct of the package responsible for tokenizing an input stream.
// It holds the current state of the lexing process on the buffer.
type Lexer struct {
	Position      Position
	Content       string
	NextLexerFunc LexerFunc
}

// New initializes a Lexer with a given input content and an initial lexical operation
// embodied by the startLexer function. The function sets the lexer's position to the
// beginning of the input buffer (Row 1, Col 0). This architecture leans into functional
// principles by treating lexical operations as state transitions, where each operation
// produces a new lexer state and an accompanying token.
//
// Parameters:
//   - content: The input string that will be tokenized.
//
// Returns:
//   - An initialized Lexer with the starting position and the initial lexical function (startLexer).
func New(content string) Lexer {
	return Lexer{
		Position: Position{
			Row: 1,
			Col: 0,
		},
		Content:       content,
		NextLexerFunc: dispatcherLexer,
	}
}

// LexerFunc is a type alias for a function that embodies the core behavior of the lexer.
// It defines a lexical operation as a state transition: taking an existing Lexer instance
// and returning a new Lexer state, along with the processed Token encapsulated in a monad.Result.
// This architecture promotes a functional approach by encapsulating state transitions as
// pure functions, enhancing testability, and making the behavior more predictable.
// The use of a monad.Result allows for a robust way to handle both successful tokenization
// and errors without relying on Go's idiomatic error-handling mechanisms.
type LexerFunc func(l Lexer) (monad.Maybe[Token], Lexer)

// Next serves as a higher-order function that delegates the task of tokenization to the
// LexerFunc stored in the Lexer instance it receives. In doing so, it adheres to the
// Single Responsibility Principle by limiting its own role and thereby simplifying its
// behavior. Essentially, Next treats LexerFunc as a strategy for lexing and forwards
// the responsibility of generating the next Token to it. The function returns a monad.Result
// wrapping the Token, thereby offering a unified approach to handle both success and failure states.
func Next(l Lexer) (monad.Maybe[Token], Lexer) {
	return l.NextLexerFunc(l)
}

// dispatcherLexer serves as a pivotal LexerFunc that orchestrates the lexical analysis process.
// It acts as a central dispatcher that delegates specific lexing tasks to other, more specialized, lexers.
// As a result, it encapsulates the mechanism that directs the flow of control based on the type of token
// encountered. It is, in essence, the nucleus of the lexer subsystem, making high-level decisions based
// on a first-pass scan of each rune.
//
// A noteworthy feature of this function is its handling of corner cases:
//  1. If the content to be lexed is empty, it returns an EOF (End-of-File) token, signaling the lexer
//     to cease operation.
//  2. If a newline character is encountered, it returns an EOL (End-of-Line) token and prepares the lexer
//     to commence scanning the new line.
//
// For whitespace characters, the function simply advances the lexer's position without producing a token.
//
// If any of the predefined operator characters are encountered (e.g., '.', '(', ')', '|', ':', '-'),
// the function delegates the task to the operatorLexer for further scrutiny.
//
// If a string delimiter '"' is encountered, the function transitions control to the stringLexer.
//
// For all other characters, it assumes that an identifier or another type of token is being formed
// and delegates this task to the identifierLexer.
//
// The crux of this function is the use of monad.Maybe and setting the NextLexerFunc field of the returned
// Lexer object. By doing so, it elegantly marries functional programming paradigms with the procedural
// nature of lexing, achieving a modular, maintainable, and robust design.
func dispatcherLexer(l Lexer) (monad.Maybe[Token], Lexer) {
	if len(l.Content) == 0 {
		return monad.Some(Token{EOF, l.Position, ""}), Lexer{
			Position:      l.Position,
			Content:       l.Content,
			NextLexerFunc: nil,
		}
	}

	x, size := utf8.DecodeRuneInString(l.Content)
	xs := l.Content[size:]

	if x == '\n' {
		return monad.Some(Token{EOL, l.Position, ""}), Lexer{
			Position:      Position{l.Position.Row + 1, 0},
			Content:       xs,
			NextLexerFunc: dispatcherLexer,
		}
	}

	if unicode.IsSpace(x) {
		return monad.None[Token](), Lexer{
			Position:      Position{l.Position.Row, l.Position.Col + 1},
			Content:       xs,
			NextLexerFunc: dispatcherLexer,
		}
	}

	if strings.ContainsAny(string(x), "\\.()|:-") {
		return monad.None[Token](), Lexer{
			Position:      l.Position,
			Content:       l.Content,
			NextLexerFunc: operatorLexer,
		}
	}

	if x == '"' {
		return monad.None[Token](), Lexer{
			Position:      l.Position,
			Content:       l.Content,
			NextLexerFunc: stringLexer,
		}
	}

	return monad.None[Token](), Lexer{
		Position:      l.Position,
		Content:       l.Content,
		NextLexerFunc: identifierLexer,
	}
}

// OperatorMap is a map of rune to TokenType for operators.
var OperatorMap = map[rune]TokenType{
	'\\': LAMBDA,
	'.':  DOT,
	'(':  LPAREN,
	')':  RPAREN,
	'|':  MODULE,
}

// operatorLexer serves as a LexerFunc dedicated to lexing operators within a lambda calculus language.
// Operators are terminal symbols that signify operations or serve as delimiters in the language syntax.
// This function is acutely attuned to the peculiarities of the lexicon, as it handles both single and
// composite operators.
//
// The lexing process first identifies the Unicode rune at the current position and checks its nature.
// For simple operators like '(', ')', or '\\', it uses a predefined OperatorMap to quickly resolve
// the type of the Token to be generated. The Token is then encapsulated in a monad.Maybe and returned
// along with a Lexer that captures the updated state.
//
// For composite operators such as ':=' or '->', the function inspects the subsequent rune to make a
// determination. In such cases, both characters are consumed, and the composite Token is returned.
//
// An intriguing aspect of this function is its seamless transition to the identifierLexer in edge cases.
// For example, a ':' or '-' not followed by '=' or '>' respectively are not considered operators but part
// of an identifier or another construct. In such cases, the function elegantly delegates the lexing task to
// identifierLexer by updating the NextLexerFunc field of the returned Lexer instance.
//
// This design ensures a modular, maintainable, and comprehensible codebase while adhering to the principle
// of single responsibility. Thus, operatorLexer only concerns itself with lexing operators, and nothing more.
func operatorLexer(l Lexer) (monad.Maybe[Token], Lexer) {
	x, size := utf8.DecodeRuneInString(l.Content)
	xs := l.Content[size:]

	// Check if the rune is a simple operator
	if tokenType, exists := OperatorMap[x]; exists {
		return monad.Some(Token{tokenType, l.Position, Literal(x)}), Lexer{
			Position:      Position{l.Position.Row, l.Position.Col + 1},
			Content:       xs,
			NextLexerFunc: dispatcherLexer,
		}
	}

	// Check for composite operators like ':=' and '->'
	if x == ':' || x == '-' {
		if len(xs) > 0 {
			x2, size := utf8.DecodeRuneInString(xs)
			xs := xs[size:]
			if x == ':' && x2 == '=' {
				return monad.Some(Token{ASSIGN, l.Position, Literal(x)}), Lexer{
					Position:      Position{l.Position.Row, l.Position.Col + 1},
					Content:       xs,
					NextLexerFunc: dispatcherLexer,
				}
			}
			if x == '-' && x2 == '>' {
				return monad.Some(Token{NSDEREF, l.Position, Literal(x)}), Lexer{
					Position:      Position{l.Position.Row, l.Position.Col + 1},
					Content:       xs,
					NextLexerFunc: dispatcherLexer,
				}
			}
		}
	}

	// Fall back to identifierLexer for edge cases.
	return monad.None[Token](), Lexer{
		Position:      l.Position,
		Content:       l.Content,
		NextLexerFunc: identifierLexer,
	}
}

// stringLexer is a specialized LexerFunc designed for tokenizing strings in a lambda calculus language.
// It utilizes a form of recursive descent to build up the string token incrementally. Each invocation
// of the function peels off one character from the front of the remaining input and decides on the next
// course of action based on that character and the state accumulated so far.
//
// This function faces unique challenges, such as handling escaped quotes within strings, and capturing
// illegal states like unclosed strings or strings containing newlines.
//
// Recursive descent is typically used in parsing rather than lexing, but in this unique context, it serves
// to modularize and isolate the logic required for string tokenization.
//
// The function uses a local higher-order function `strLex` that encapsulates the core recursion logic. The
// `strLex` function is tail-recursive and operates on an accumulating monad.Either[Token] to build up the token's
// Literal value. It offers early termination through monad.NewRVal in case of illegal states, thereby folding
// both the success and failure states into a unified approach.
//
// Finally, the function returns a monad.Maybe[Token] encapsulating the resulting token if a legal string is
// found or the illegal state otherwise. It also returns a new Lexer with an updated state to be used in
// subsequent lexing operations.
func stringLexer(l Lexer) (monad.Maybe[Token], Lexer) {
	// Invalid string with EOF after a single "
	if len(l.Content) < 1 {
		return monad.Some(Token{ILLEGAL, l.Position, Literal(l.Content)}), Lexer{
			Position:      l.Position,
			Content:       l.Content,
			NextLexerFunc: nil,
		}
	}

	// Skip first '"' character
	_, size := utf8.DecodeRuneInString(l.Content[1:])
	xs := l.Content[size+1:]

	var strLex func(t monad.Either[Token]) monad.Either[Token]
	strLex = func(t monad.Either[Token]) monad.Either[Token] {
		if _, ok := t.(monad.Left[Token]); !ok {
			return t
		}
		if len(t.Value().Literal) == 0 {
			return monad.NewRVal(Token{ILLEGAL, l.Position, ""})
		}
		x, size := utf8.DecodeRuneInString(t.Value().Literal.String())
		xs := t.Value().Literal[size:]

		if x == '\n' {
			return monad.NewRVal(Token{ILLEGAL, l.Position, ""})
		}
		if x == '"' {
			return monad.NewRVal(Token{STRING, l.Position, Literal(l.Content)})
		}

		if x == '\\' {
			if len(xs) > 1 { // Ensure that there are enough characters left.
				x2, _ := utf8.DecodeRuneInString(string(xs))
				if x2 == '"' {
					return strLex(t).Or(func(t Token) monad.Either[Token] {
						t.Literal = "\"" + t.Literal
						return monad.NewLVal(t)
					})
				}
				return strLex(t).Or(func(t Token) monad.Either[Token] {
					t.Literal = "\\" + t.Literal
					return monad.NewLVal(t)
				})
			}
			return monad.NewRVal(
				Token{ILLEGAL, l.Position, ""},
			) // Edge case: if '\\' is the last character.
		}

		return strLex(t).Or(func(t Token) monad.Either[Token] {
			t.Literal = Literal(x) + t.Literal
			return monad.NewLVal(t)
		})
	}

	val := strLex(monad.NewLVal(Token{STRING, l.Position, Literal(xs)}))
	content, _ := strings.CutPrefix(l.Content, val.Value().Literal.String())
	return monad.Some(val.Value()), Lexer{
		Position: Position{
			l.Position.Row,
			l.Position.Col + utf8.RuneCountInString(val.Value().Literal.String()),
		},
		Content:       content,
		NextLexerFunc: dispatcherLexer,
	}
}

// identifierLexer is responsible for lexing identifiers, which typically consist of a sequence of
// alphanumeric characters and underscores. An identifier must start with an alphabetic character. This
// function uses recursion to iteratively construct the identifier token from the provided input.
//
// The function begins by checking the first rune of the remaining content. If the rune adheres to the
// identifier character rules (alphabetic or underscore), the function recursively lexes the rest of the
// string and builds the identifier token. Once an invalid identifier rune or the end of the content is
// reached, the function returns the built identifier token and the updated Lexer state.
//
// Parameters:
//
//	l: Initial lexer state including the position and the remaining content to lex.
//
// Returns:
//  1. A monad.Maybe[Token] containing either the identifier Token or None.
//  2. An updated Lexer state for continued lexical analysis.
func identifierLexer(l Lexer) (monad.Maybe[Token], Lexer) {
	return idLexRecursively(l, "")
}

func idLexRecursively(l Lexer, acc string) (monad.Maybe[Token], Lexer) {
	if len(l.Content) == 0 {
		if acc == "" {
			return monad.None[Token](), l
		}
		return monad.Some(Token{IDENT, l.Position, Literal(acc)}), l
	}

	x, size := utf8.DecodeRuneInString(l.Content)
	xs := l.Content[size:]
	reservedOps := "\\.()|:-"
	if !unicode.IsSpace(x) && x != '\n' && !strings.ContainsRune(reservedOps, x) {
		// Check for composite operators like ':=' and '->'
		if x == ':' || x == '-' {
			if len(xs) > 0 {
				x2, _ := utf8.DecodeRuneInString(xs)
				if (x == ':' && x2 == '=') || (x == '-' && x2 == '>') {
					return finalizeIdentifierToken(l, acc)
				}
			}
		}

		newPos := Position{l.Position.Row, l.Position.Col + size}
		newContent := l.Content[size:]
		return idLexRecursively(Lexer{newPos, newContent, dispatcherLexer}, acc+string(x))
	}

	return finalizeIdentifierToken(l, acc)
}

func finalizeIdentifierToken(l Lexer, acc string) (monad.Maybe[Token], Lexer) {
	if acc == "" {
		return monad.None[Token](), l
	}
	return monad.Some(Token{IDENT, l.Position, Literal(acc)}), l
}
