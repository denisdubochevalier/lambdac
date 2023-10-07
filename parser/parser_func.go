package parser

import "github.com/denisdubochevalier/monad"

// parserFunc defines the signature for specialized parsing functions that
// operate within the parsing framework. A parserFunc is responsible for taking
// an existing parser state and performing a specific parsing operation based on
// the tokens available in that state.
//
// Each parserFunc function aims to accomplish the following:
//   - Inspect the current state's token(s) and determine whether it can handle
//     them.
//   - Generate a new abstract syntax tree node (ASTNode) if the token(s) are
//     recognized, encapsulating it within a Result monad to signify successful
//     parsing.
//   - Update the parser state, usually by advancing the position to point to
//     the next token for subsequent parsing functions.
//   - Return an error within the Result monad if it encounters an unexpected
//     token or another issue that prevents successful parsing.
//
// Parameters:
//   - State: The current parser state, which includes the tokens to be parsed
//     and the current position within that list.
//
// Returns:
//   - A Result monad that encapsulates either a successfully parsed ASTNode or
//     an error detailing what went wrong.
//   - An updated parser state that is passed on to the next parserFunc for
//     further processing.
type parserFunc func(State) (monad.Result[ASTNode, error], State)
