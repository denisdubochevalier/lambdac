package lexer_test

import (
	"fmt"

	"github.com/denisdubochevalier/lambdac/lexer"
	"github.com/denisdubochevalier/monad"
)

func ExampleLexerNext() {
	input := `maths | "github.com/foo/bar"

	Y := \f.(\x.f(x x)).(\x.f(x x))

	fact := Y maths.non_recursive_factorial

	5 := \f.\x.f f f f f x

	fact 5`
	l := lexer.New().WithContent(input)

	for {
		token := monad.None[lexer.Token]()
		token, l = l.Next()
		if _, ok := token.(monad.Nothing[lexer.Token]); ok {
			continue
		}
		fmt.Printf(
			"Type: %s, Position: %d - %d, Literal: %q\n",
			token.Value().Type(),
			token.Value().Position().Row(),
			token.Value().Position().Col(),
			token.Value().Literal(),
		)

		if token.Value().Type() == lexer.EOF {
			break
		}
	}
	// Output:
	// Type: IDENT, Position: 1 - 0, Literal: "maths"
	// Type: |, Position: 1 - 6, Literal: "|"
	// Type: STRING, Position: 1 - 8, Literal: "github.com/foo/bar"
	// Type: EOL, Position: 1 - 26, Literal: ""
	// Type: EOL, Position: 2 - 0, Literal: ""
	// Type: IDENT, Position: 3 - 1, Literal: "Y"
	// Type: :=, Position: 3 - 3, Literal: ":="
	// Type: \, Position: 3 - 5, Literal: "\\"
	// Type: IDENT, Position: 3 - 6, Literal: "f"
	// Type: ., Position: 3 - 7, Literal: "."
	// Type: (, Position: 3 - 8, Literal: "("
	// Type: \, Position: 3 - 9, Literal: "\\"
	// Type: IDENT, Position: 3 - 10, Literal: "x"
	// Type: ., Position: 3 - 11, Literal: "."
	// Type: IDENT, Position: 3 - 12, Literal: "f"
	// Type: (, Position: 3 - 13, Literal: "("
	// Type: IDENT, Position: 3 - 14, Literal: "x"
	// Type: IDENT, Position: 3 - 16, Literal: "x"
	// Type: ), Position: 3 - 17, Literal: ")"
	// Type: ), Position: 3 - 18, Literal: ")"
	// Type: ., Position: 3 - 19, Literal: "."
	// Type: (, Position: 3 - 20, Literal: "("
	// Type: \, Position: 3 - 21, Literal: "\\"
	// Type: IDENT, Position: 3 - 22, Literal: "x"
	// Type: ., Position: 3 - 23, Literal: "."
	// Type: IDENT, Position: 3 - 24, Literal: "f"
	// Type: (, Position: 3 - 25, Literal: "("
	// Type: IDENT, Position: 3 - 26, Literal: "x"
	// Type: IDENT, Position: 3 - 28, Literal: "x"
	// Type: ), Position: 3 - 29, Literal: ")"
	// Type: ), Position: 3 - 30, Literal: ")"
	// Type: EOL, Position: 3 - 31, Literal: ""
	// Type: EOL, Position: 4 - 0, Literal: ""
	// Type: IDENT, Position: 5 - 1, Literal: "fact"
	// Type: :=, Position: 5 - 6, Literal: ":="
	// Type: IDENT, Position: 5 - 8, Literal: "Y"
	// Type: IDENT, Position: 5 - 10, Literal: "maths"
	// Type: ., Position: 5 - 15, Literal: "."
	// Type: IDENT, Position: 5 - 16, Literal: "non_recursive_factorial"
	// Type: EOL, Position: 5 - 39, Literal: ""
	// Type: EOL, Position: 6 - 0, Literal: ""
	// Type: IDENT, Position: 7 - 1, Literal: "5"
	// Type: :=, Position: 7 - 3, Literal: ":="
	// Type: \, Position: 7 - 5, Literal: "\\"
	// Type: IDENT, Position: 7 - 6, Literal: "f"
	// Type: ., Position: 7 - 7, Literal: "."
	// Type: \, Position: 7 - 8, Literal: "\\"
	// Type: IDENT, Position: 7 - 9, Literal: "x"
	// Type: ., Position: 7 - 10, Literal: "."
	// Type: IDENT, Position: 7 - 11, Literal: "f"
	// Type: IDENT, Position: 7 - 13, Literal: "f"
	// Type: IDENT, Position: 7 - 15, Literal: "f"
	// Type: IDENT, Position: 7 - 17, Literal: "f"
	// Type: IDENT, Position: 7 - 19, Literal: "f"
	// Type: IDENT, Position: 7 - 21, Literal: "x"
	// Type: EOL, Position: 7 - 22, Literal: ""
	// Type: EOL, Position: 8 - 0, Literal: ""
	// Type: IDENT, Position: 9 - 1, Literal: "fact"
	// Type: IDENT, Position: 9 - 6, Literal: "5"
	// Type: EOF, Position: 9 - 7, Literal: ""
}
