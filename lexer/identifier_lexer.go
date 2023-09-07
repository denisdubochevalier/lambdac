package lexer

import (
	"strings"
	"unicode"
	"unicode/utf8"

	"github.com/denisdubochevalier/monad"
)

// identifierLexer serves as an entry point for the identification and
// extraction of IDENT type tokens from a given string. It is a LexerFunc,
// a specific type of function that complies with the lexer's requirements
// for token recognition functions. This function delegates the primary
// work of token identification to the idLexRecursively function.
//
// Parameters:
// l: Lexer instance containing the current state, such as the remaining
//
//	content to be lexed and the current position within that string.
//
// Returns:
//  1. An optional Token, encapsulated in a Maybe monad. This will be None if
//     no valid identifier can be constructed, or Some(Token) if an identifier
//     has been recognized.
//  2. A new Lexer instance with the state updated based on the recognized token.
//
// The function performs the following operations:
//  1. Invokes idLexRecursively with the current lexer state, thereby commencing
//     the recursive descent parsing process to identify the IDENT token.
//  2. The actual identification and construction of the token are performed
//     in idLexRecursively. This function thus acts as a thin wrapper or a gateway,
//     streamlining the call to the recursive function.
//
// As a member of the LexerFunc family, identifierLexer can be integrated into
// a chain of lexer functions, thereby contributing to the lexer's modular and
// extensible architecture.
func identifierLexer(l Lexer) (monad.Maybe[Token], Lexer) {
	return idLexRecursively(l)
}

// checkCompositeOps checks for composite operators like ":=" and "->"
func checkCompositeOps(x rune, xs string) bool {
	if len(xs) > 0 {
		x2, _ := utf8.DecodeRuneInString(xs)
		return (x == ':' && x2 == '=') || (x == '-' && x2 == '>')
	}
	return false
}

// isValidIdentifierChar validates if a rune can be part of an identifier.
func isValidIdentifierChar(x rune) bool {
	reservedOps := "\\.()|"
	return unicode.IsGraphic(x) && !unicode.IsSpace(x) && x != '\n' &&
		!strings.ContainsRune(reservedOps, x)
}

// updateLexerForRecursion prepares a new Lexer instance for the next recursion level.
func updateLexerForRecursion(l Lexer, size int, xs string) Lexer {
	return New().
		WithPosition(l.position.advanceColBy(size)).
		WithContent(xs).
		WithNextLexerFunc(eofLexer)
}

// mergeLiterals combines an existing Token with a new rune to form a new Token. The function employs monadic abstractions
// to ensure composability and error safety.
//
// Parameters:
// - existing: A monad.Maybe[Token] containing an optional existing Token. If present, this Token will be combined with the new rune.
// - newChar: A rune that will be combined with the existing Token's literal.
//
// Returns:
//   - monad.Maybe[Token]: A new monadic Maybe instance encapsulating the resultant Token. If an existing Token is present,
//     its literal is extended with the new rune at the beginning; otherwise, a new Token is formed solely from the new rune.
//
// Example:
// Assuming existing Token has a literal "bcd", and newChar is 'a',
// the returned Token will have a literal "abcd".
func mergeLiterals(existing monad.Maybe[Token], newChar rune) monad.Maybe[Token] {
	// Handle the edge case when both existing and newChar are empty.
	if existing == monad.None[Token]() && newChar == rune(0) {
		return monad.Some(Token{literal: Literal(newChar)})
	}

	// If newChar is empty but existing is not None, return existing.
	if newChar == rune(0) {
		return existing
	}

	// If existing is Just, merge the literals.
	if existingToken, ok := existing.(monad.Just[Token]); ok {
		t := existingToken.Value()
		t.literal = Literal(newChar) + t.Literal()
		return monad.Some(t)
	}

	// If existing is None and newChar is not empty, create a new Token.
	return monad.Some(Token{literal: Literal(newChar)})
}

// idLexRecursively is a recursive descent function that identifies tokens of
// type 'IDENT' within a string. The function receives a Lexer instance that
// contains the current state, including the string to be processed and
// the current position within that string. It returns two values:
//  1. An optional Token encapsulated in a Maybe monad, which will be None
//     if no valid identifier can be constructed from the current position.
//  2. A new Lexer instance with updated state, including a string with the
//     processed characters removed and potentially an updated position.
//
// The function applies the following logic to identify tokens:
//  1. It checks for end-of-string conditions to terminate recursion.
//  2. If the current character (rune) can be part of an identifier, it recursively
//     continues lexing the remaining string.
//  3. During the recursive backtracking, it accumulates the characters into the
//     Token's literal.
//  4. Handles specific composite operators, ensuring they are not mistaken for identifiers.
//  5. If it encounters a character that cannot be part of an identifier, it finalizes
//     the token assembly.
//
// The function adheres to functional programming paradigms and ensures immutability
// by creating new instances of Lexer with updated state rather than modifying the
// existing one.
func idLexRecursively(l Lexer) (monad.Maybe[Token], Lexer) {
	if len(l.content) == 0 {
		return monad.None[Token](), l
	}

	x, size := utf8.DecodeRuneInString(l.content)
	xs := l.content[size:]

	if isValidIdentifierChar(x) {
		// If the rune might be part of a composite operator, we finalize
		if checkCompositeOps(x, xs) {
			return finalizeIdentifierToken(l)
		}

		// Continue with the recursion, consuming the character
		nextLexer := updateLexerForRecursion(l, size, xs)
		nextToken, remainingLexer := idLexRecursively(nextLexer)
		mergedToken := mergeLiterals(nextToken, x)
		return monad.Some(Token{IDENT, l.position, mergedToken.Value().Literal()}), remainingLexer
	}

	// Finalize the token when we reach an invalid character for an identifier
	return finalizeIdentifierToken(l)
}

// finalizeIdentifierToken is a helper function that finalizes the process of
// constructing an 'IDENT' type token. It is called when idLexRecursively
// encounters a character that can't be part of an identifier or reaches
// the end of the string.
//
// The function receives a Lexer instance that contains the current state,
// including the string to be processed and the current position within that
// string. It returns two values:
//  1. An optional Token encapsulated in a Maybe monad. If the content string
//     of the Lexer is zero-length, it returns None; otherwise, it constructs
//     an IDENT type token.
//  2. A new Lexer instance with updated state, including the processed string
//     and potentially an updated position.
//
// This function serves as the termination condition for the recursion in idLexRecursively
// and decides whether an identifier token can be constructed based on the current state.
func finalizeIdentifierToken(l Lexer) (monad.Maybe[Token], Lexer) {
	if len(l.content) == 0 {
		return monad.None[Token](), l
	}
	return monad.Some(Token{IDENT, l.position, Literal("")}), l
}
