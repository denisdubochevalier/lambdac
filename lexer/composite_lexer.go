package lexer

import (
	"unicode/utf8"

	"github.com/denisdubochevalier/monad"
)

// compositeLexer functions as a specialized LexerFunc whose sole purview is the lexing of composite operators within the lambda calculus language.
// This function is invoked when a potential composite operator is encountered, to either confirm its composite nature or relegate it as part of another construct.
// For operators like ':=' and '->', the function inspects the Unicode rune following the initial character to ascertain whether they together form a composite operator.
//
// If they do, a Token of the corresponding type is generated, encapsulated in a monad.Maybe, and returned along with an updated Lexer instance.
// If the sequence does not form a recognized composite operator, the function returns a monad.None and delegates the remaining lexing task
// to the identifierLexer, ensuring a smooth transition in edge cases.
//
// By focusing exclusively on composite operators, compositeLexer maintains a clean separation of concerns and enhances the modularity and readability
// of the lexer subsystem.
func compositeLexer(l Lexer) (monad.Maybe[Token], Lexer) {
	x, size := utf8.DecodeRuneInString(l.content)
	xs := l.content[size:]

	// Check for composite operators like ':=' and '->'
	if x == ':' || x == '-' {
		if len(xs) > 0 {
			x2, size := utf8.DecodeRuneInString(xs)
			xs := xs[size:]
			if x == ':' && x2 == '=' {
				return monad.Some(Token{ASSIGN, l.position, Literal([]rune{x, x2})}), l.
					WithPosition(l.position.advanceCol()).
					WithContent(xs).
					WithNextLexerFunc(eofLexer)
			}
			if x == '-' && x2 == '>' {
				return monad.Some(Token{NSDEREF, l.position, Literal([]rune{x, x2})}), l.
					WithPosition(l.position.advanceCol()).
					WithContent(xs).
					WithNextLexerFunc(eofLexer)
			}
		}
	}

	// Fall back to identifierLexer for edge cases
	return monad.None[Token](), l.WithNextLexerFunc(identifierLexer)
}
