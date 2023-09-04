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
//	l := lexer.New().WithContent(input)
//
//	for {
//	  tok := l.Next()
//	  if tok.Type == EOF {
//	    break
//	  }
//	  // Process tok
//	}
package lexer

import "github.com/denisdubochevalier/monad"

// Lexer serves as a state container for lexical analysis, holding essential
// information such as the text content being processed, the current position
// within that content, and the next lexing function to be executed. This design
// adheres to an immutability-oriented programming model, aiding in the
// clarity and predictability of the code.
type Lexer struct {
	position      Position
	content       string
	nextLexerFunc lexerFunc
}

// New initializes and returns a new Lexer instance with its position set to the starting point.
// This constructor adheres to the principle of sensible defaults, automatically setting the position
// to the beginning of the source text, thereby allowing for immediate utilization for lexical analysis.
// Additional state variables like content and nextLexerFunc are left uninitialized and can be set
// using their respective methods.
func New() Lexer {
	return Lexer{position: StartPosition(), nextLexerFunc: dispatcherLexer}
}

// WithPosition updates the Lexer's Position, returning a new Lexer with the
// updated state. This method allows for a functional approach to updating state,
// supporting easier debugging and more reasoned program flow.
func (l Lexer) WithPosition(p Position) Lexer {
	l.position = p
	return l
}

// WithContent updates the content of the Lexer. This approach sustains
// the immutable state model of the Lexer, reducing the possibility of
// unexpected side-effects and enhancing code readability.
func (l Lexer) WithContent(c string) Lexer {
	l.content = c
	return l
}

// WithNextLexerFunc updates the lexer function in the Lexer instance.
// This dynamic adjustment permits a more nuanced navigation through
// the stages of lexical analysis.
func (l Lexer) WithNextLexerFunc(f lexerFunc) Lexer {
	l.nextLexerFunc = f
	return l
}

// Next serves as a higher-order function that delegates the task of tokenization to the
// lexerFunc stored in the Lexer instance it receives. In doing so, it adheres to the
// Single Responsibility Principle by limiting its own role and thereby simplifying its
// behavior. Essentially, Next treats lexerFunc as a strategy for lexing and forwards
// the responsibility of generating the next Token to it. The function returns a monad.Result
// wrapping the Token, thereby offering a unified approach to handle both success and failure states.
func (l Lexer) Next() (monad.Maybe[Token], Lexer) {
	return l.nextLexerFunc(l)
}
