package parser

import (
	"fmt"

	"github.com/denisdubochevalier/lambdac/lexer"
	"github.com/denisdubochevalier/monad"
)

// eofParser is a specialized parser function responsible for handling the
// end-of-file (EOF) token within the parsing process. It is invoked as part
// of the overall parsing strategy to deal with the termination of the token stream.
//
// The eofParser function performs the following operations:
//   - Checks if the parser has reached the end of the token list (state.done()).
//     If it has, an "unexpected end of input" error is returned.
//   - Inspects the current token to see if it is of the type EOF. If it is,
//     a new ASTNode corresponding to this EOF token is created and appended to
//     the existing AST. The parser state is updated to include this new ASTNode.
//   - Advances the parser's position to the next token in preparation for
//     subsequent parsing operations.
//   - If the current token is not of type EOF, delegates to eolParser for
//     further parsing.
//
// Parameters:
//   - State: The current state of the parser, containing the list of tokens to
//     be parsed and the current position in that list.
//
// Returns:
//   - A Result monad that either encapsulates a successfully parsed ASTNode
//     representing the EOF or returns an error detailing what went wrong.
//   - An updated parser State that includes any newly created ASTNode and advances
//     the current position for subsequent parsing operations.
func eofParser(state State) (monad.Result[ASTNode, error], State) {
	if state.done() {
		return monad.Fail[ASTNode, error](
			fmt.Errorf("unexpected end of input"),
		), state
	}

	if state.currentToken().Type() == lexer.EOF {
		ast := state.ast().appendChild(newASTNode(lexer.EOF, state.currentToken()))
		return monad.Succeed[ASTNode, error](ast), state.withAST(ast).advance()
	}

	return eolParser(state)
}
