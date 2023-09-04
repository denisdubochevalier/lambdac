package lexer

import (
	"unicode/utf8"

	"github.com/denisdubochevalier/monad"
)

// operatorMap is a map of rune to TokenType for operators.
var operatorMap = map[rune]TokenType{
	'\\': LAMBDA,
	'.':  DOT,
	'(':  LPAREN,
	')':  RPAREN,
	'|':  MODULE,
}

// operatorLexer serves as a LexerFunc exclusively devoted to lexing simple operators within the lambda calculus language.
// This function meticulously handles terminal symbols that represent simple operations or serve as delimiters in the language syntax.
// Utilizing a pre-defined 'operatorMap', it swiftly maps the Unicode rune encountered to its corresponding TokenType.
// If the rune corresponds to a simple operator, the function generates a Token, encapsulates it in a monad.Maybe, and returns it
// along with a Lexer instance reflecting the updated state.
//
// Should the function encounter a rune that does not exist in the 'operatorMap', it delegates the responsibility of lexing
// to the compositeLexer function. This segregation allows operatorLexer to focus solely on simple operators, thereby adhering
// to the Single Responsibility Principle and contributing to a modular and maintainable codebase.
func operatorLexer(l Lexer) (monad.Maybe[Token], Lexer) {
	x, size := utf8.DecodeRuneInString(l.content)
	xs := l.content[size:]

	// Check if the rune is a simple operator
	if tokenType, exists := operatorMap[x]; exists {
		return monad.Some(Token{tokenType, l.position, Literal(x)}), l.
			WithPosition(l.position.advanceCol()).
			WithContent(xs).
			WithNextLexerFunc(eofLexer)
	}

	// Fall back to compositeLexer for potentially composite operators
	return compositeLexer(l)
}
