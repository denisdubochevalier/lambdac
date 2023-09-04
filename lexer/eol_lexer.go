package lexer

import (
	"unicode/utf8"

	"github.com/denisdubochevalier/monad"
)

// eolLexer is responsible for recognizing and handling End-Of-Line (EOL) characters.
// This function is called immediately after eofLexer, if the content is not empty.
//
// Parameters:
//   - l: The current Lexer object.
//
// Returns:
//   - monad.Maybe[Token]: A Maybe monad encapsulating the token (EOL), if applicable.
//   - Lexer: A new Lexer object with a potentially modified state.
//
// Upon encountering an EOL character, this function returns an EOL Token,
// advances the row position, and continues the lexer process by setting the
// next lexer function to eofLexer.
//
// If an EOL character is not encountered, it delegates the responsibility to spaceLexer.
func eolLexer(l Lexer) (monad.Maybe[Token], Lexer) {
	x, size := utf8.DecodeRuneInString(l.content)
	xs := l.content[size:]

	if x == '\n' {
		return monad.Some(Token{EOL, l.position, ""}), l.
			WithPosition(l.position.newRow()).
			WithContent(xs).
			WithNextLexerFunc(eofLexer)
	}

	return spaceLexer(l)
}
