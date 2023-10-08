package parser

import (
	"fmt"

	"github.com/denisdubochevalier/monad"

	"github.com/denisdubochevalier/lambdac/lexer"
)

// identParser is a specialized function responsible for parsing identifiers in
// the λ.c language. It operates on a given State, which encapsulates the parser's
// current position and AST (Abstract Syntax Tree).
//
// This function returns a tuple consisting of two elements:
//  1. A `monad.Result[ASTNode, error]` which contains either the parsed ASTNode
//     corresponding to the identifier or an error.
//  2. An updated State, which reflects any changes made during the parsing
//     process, such as advancing the token position or modifying the AST.
//
// Workflow:
//   - Firstly, it checks if the parser has reached the end of the token list,
//     returning an "unexpected end of input" error if so.
//   - If the current token is an identifier (lexer.IDENT), it appends a new ASTNode
//     to the AST encapsulated within the State. It then delegates the parsing task
//     to `eofParser` after updating the State, to restart the parsing loop.
//   - If the current token is not an identifier, it delegates the task to the
//     `moduleParser` function.
//
// Note: The function utilizes monads to encapsulate the inherent duality of
// parsing, capturing both the successful parsing outcome and any potential
// errors in a functional programming paradigm.
func identParser(state State) (monad.Result[ASTNode, error], State) {
	if state.done() {
		return monad.Fail[ASTNode, error](
			fmt.Errorf("unexpected end of input"),
		), state
	}

	if state.currentToken().Type() == lexer.IDENT {
		ast := state.ast().appendChild(newASTNode(lexer.IDENT, state.currentToken()))
		return eofParser(state.withAST(ast).advance())
	}

	return moduleParser(state)
}
