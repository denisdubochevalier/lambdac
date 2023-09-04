package lexer

import (
	"github.com/denisdubochevalier/monad"
)

// eofLexer serves as the entry point for the lexer pipeline and is responsible for
// handling the End-Of-File (EOF) condition. This function adheres to the principle
// of returning a Maybe monad of Token and a new Lexer, encapsulating the optional
// nature of the token in question.
//
// Parameters:
//   - l: The current Lexer object, encapsulating the state of the lexer.
//
// Returns:
//   - monad.Maybe[Token]: A Maybe monad encapsulating the token (EOF), if applicable.
//   - Lexer: A new Lexer object with potentially modified state, primed for the next lexer function.
//
// When the lexer reaches the end of the input content (l.content is empty),
// this function returns an EOF Token and sets the next lexer function to nil,
// signaling the termination of the lexer process.
//
// If the end of the content has not been reached, this function delegates to eolLexer.
func eofLexer(l Lexer) (monad.Maybe[Token], Lexer) {
	if len(l.content) == 0 {
		return monad.Some(Token{EOF, l.position, ""}), l.WithNextLexerFunc(nil)
	}
	return eolLexer(l)
}
