package parser

import (
	"fmt"

	"github.com/denisdubochevalier/monad"

	"github.com/denisdubochevalier/lambdac/lexer"
)

// moduleParser is designed to parse the 'module' keyword and associated logic
// in the λ.c programming language. The function operates on a given State that
// includes the current state of the parser, including the token stream and the
// constructed Abstract Syntax Tree (AST).
//
// The function returns a tuple composed of two primary elements:
//  1. A `monad.Result[ASTNode, error]` encapsulating either the successfully parsed
//     ASTNode pertaining to the 'module' keyword or an error delineating exceptions.
//  2. A modified State reflecting the operations performed, e.g., token progression
//     or AST modifications.
//
// Operational Cadence:
//   - Initially checks if the token list has been exhausted; returns an error if true.
//   - If the current token is not 'module', the responsibility of parsing is transferred
//     to `lambdaParser`.
//   - Utilizes chained monadic operations (`FlatMap`) to:
//     a. Verify the last ASTNode as an identifier.
//     b. Append this identifier to a new 'module' ASTNode.
//     c. Replace the last child of the AST with this newly augmented 'module' node.
//   - On successful completion, the parsing task is delegated to `stringParser` after
//     updating the State.
//   - If the chain of monadic operations fails at any point, an error ("module operator
//     without previous ident") is emitted.
//
// Caveat: Monadic operations are judiciously leveraged for sequencing operations and
// encapsulating both the resultant ASTNode or any potential errors, thereby conferring
// functional purity to the parsing process.
func moduleParser(state State) (monad.Result[ASTNode, error], State) {
	if state.done() {
		return monad.Fail[ASTNode, error](
			fmt.Errorf("unexpected end of input"),
		), state
	}

	if state.currentToken().Type() != lexer.MODULE {
		return lambdaParser(state)
	}

	newNode := newASTNode(lexer.MODULE, state.currentToken())
	if result := state.ast().lastChild().FlatMap(
		func(node ASTNode) monad.Maybe[ASTNode] {
			if node.NodeType() != lexer.IDENT {
				return monad.None[ASTNode]()
			}
			return monad.Some(node)
		},
	).FlatMap(
		func(node ASTNode) monad.Maybe[ASTNode] {
			return monad.Some(newNode.appendChild(node))
		},
	).FlatMap(
		func(node ASTNode) monad.Maybe[ASTNode] {
			return state.ast().replaceLastChild(node)
		},
	); result.Just() {
		return stringParser(state.withAST(result.Value()).advance())
	}

	return monad.Fail[ASTNode, error](
		fmt.Errorf("module operator without previous ident"),
	), state
}
