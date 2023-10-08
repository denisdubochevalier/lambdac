package parser

import (
	"fmt"

	"github.com/denisdubochevalier/monad"

	"github.com/denisdubochevalier/lambdac/lexer"
)

// stringParser is tasked with parsing string literals within the λ.c
// programming language. This function accepts an immutable State object
// encompassing the parser's present circumstances, which include the token
// stream and the evolving Abstract Syntax Tree (AST).
//
// The function yields a tuple of:
//  1. A `monad.Result[ASTNode, error]` encapsulating either a successfully
//     parsed ASTNode for the string literal or an error object.
//  2. An updated State object manifesting the consequences of the parsing
//     action.
//
// Operational Schema:
//   - Initiates by scrutinizing for the end of the token list and throws an
//     error if this is the case.
//   - If the current token does not correspond to a string literal, it
//     returns a failure monad detailing the type of the unexpected token.
//
// Syntactic Semantics:
//   - Stipulates that a string literal should occur immediately after a 'module'
//     operator. In the absence of this condition, an error is emitted.
//   - Utilizes a sequence of monadic operations (FlatMap) to:
//     a. Append a new AST node for the string literal to the last node in the
//     AST.
//     b. Replace the last child node in the AST with the newly created composite
//     node.
//
// Monadological Note: The employment of monadic operations maintains functional
// purity, mitigating error handling complexity and endowing the codebase with
// greater clarity.
func stringParser(state State) (monad.Result[ASTNode, error], State) {
	if state.done() {
		return monad.Fail[ASTNode, error](
			fmt.Errorf("unexpected end of input"),
		), state
	}

	if state.currentToken().Type() != lexer.STRING {
		return monad.Fail[ASTNode, error](
			fmt.Errorf("unexpected token type: %s", state.currentToken().Type()),
		), state
	}

	if result := state.ast().lastChild(); result.Nothing() ||
		(result.Just() && result.Value().NodeType() != lexer.MODULE) {
		return monad.Fail[ASTNode, error](
			fmt.Errorf("string token not after a module operator"),
		), state
	}

	if result := state.ast().lastChild().FlatMap(
		func(node ASTNode) monad.Maybe[ASTNode] {
			return monad.Some(node.appendChild(newASTNode(lexer.STRING, state.currentToken())))
		},
	).FlatMap(
		func(node ASTNode) monad.Maybe[ASTNode] {
			return state.ast().replaceLastChild(node)
		},
	); result.Just() {
		return eofParser(state.withAST(result.Value()).advance())
	}

	return monad.Fail[ASTNode, error](
		fmt.Errorf("inserting string token into ast"),
	), state
}
