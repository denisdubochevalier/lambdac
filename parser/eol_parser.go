package parser

import (
	"fmt"

	"github.com/denisdubochevalier/monad"
)

// eolParser is a specialized parser function focused on handling end-of-line
// (EOL) tokens within the parsing process. Unlike other parser functions that
// may create new AST nodes, eolParser does not add a new AST node for EOL tokens.
// Instead, it focuses on validating and updating the parser state for the
// subsequent parsing tasks.
//
// The eolParser function undertakes the following steps:
//   - Validates that the parser has not reached the end of the token list.
//     If it has, an "unexpected end of input" error is returned.
//   - Checks the type of the current token. If it is an EOL, the parser state is
//     simply advanced to the next token, bypassing the creation of a new AST node.
//   - Delegates control to spaceParser for additional parsing if the current token
//     is not of type EOL.
//
// Parameters:
//   - State: The current state of the parser, including the list of tokens to
//     be parsed and the current position within that list.
//
// Returns:
//   - A Result monad that either encapsulates the current AST or returns an error
//     detailing any parsing failure.
//   - An updated parser State, advanced to the next token for subsequent parsing.
func eolParser(state State) (monad.Result[ASTNode, error], State) {
	if state.done() {
		return monad.Fail[ASTNode, error](
			fmt.Errorf("unexpected end of input"),
		), state
	}

	return monad.Fail[ASTNode, error](
		fmt.Errorf("not implemented"),
	), state
}
