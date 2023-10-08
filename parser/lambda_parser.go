package parser

import (
	"fmt"

	"github.com/denisdubochevalier/monad"
)

func lambdaParser(state State) (monad.Result[ASTNode, error], State) {
	return monad.Fail[ASTNode, error](
		fmt.Errorf("not implemented"),
	), state
}
