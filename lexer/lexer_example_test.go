package lexer_test

import (
	"fmt"

	"github.com/denisdubochevalier/lambdac/lexer"
	"github.com/denisdubochevalier/monad"
)

func Example() {
	input := `maths | "github.com/foo/bar"

	Y := \f.(\x.f(x x)).(\x.f(x x))

	fact := Y maths.non_recursive_factorial

	5 := \f.\x.f f f f f x

	fact 5`
	l := lexer.New().WithContent(input)

	for {
		var token monad.Maybe[lexer.Token]
		token, l = l.Next()
		if token.Nothing() {
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
	// Type: IDENT, Position: 3 - 0, Literal: "Y"
	// Type: :=, Position: 3 - 2, Literal: ":="
	// Type: \, Position: 3 - 4, Literal: "\\"
	// Type: IDENT, Position: 3 - 5, Literal: "f"
	// Type: ., Position: 3 - 6, Literal: "."
	// Type: (, Position: 3 - 7, Literal: "("
	// Type: \, Position: 3 - 8, Literal: "\\"
	// Type: IDENT, Position: 3 - 9, Literal: "x"
	// Type: ., Position: 3 - 10, Literal: "."
	// Type: IDENT, Position: 3 - 11, Literal: "f"
	// Type: (, Position: 3 - 12, Literal: "("
	// Type: IDENT, Position: 3 - 13, Literal: "x"
	// Type: IDENT, Position: 3 - 15, Literal: "x"
	// Type: ), Position: 3 - 16, Literal: ")"
	// Type: ), Position: 3 - 17, Literal: ")"
	// Type: ., Position: 3 - 18, Literal: "."
	// Type: (, Position: 3 - 19, Literal: "("
	// Type: \, Position: 3 - 20, Literal: "\\"
	// Type: IDENT, Position: 3 - 21, Literal: "x"
	// Type: ., Position: 3 - 22, Literal: "."
	// Type: IDENT, Position: 3 - 23, Literal: "f"
	// Type: (, Position: 3 - 24, Literal: "("
	// Type: IDENT, Position: 3 - 25, Literal: "x"
	// Type: IDENT, Position: 3 - 27, Literal: "x"
	// Type: ), Position: 3 - 28, Literal: ")"
	// Type: ), Position: 3 - 29, Literal: ")"
	// Type: EOL, Position: 3 - 30, Literal: ""
	// Type: IDENT, Position: 5 - 0, Literal: "fact"
	// Type: :=, Position: 5 - 5, Literal: ":="
	// Type: IDENT, Position: 5 - 7, Literal: "Y"
	// Type: IDENT, Position: 5 - 9, Literal: "maths"
	// Type: ., Position: 5 - 14, Literal: "."
	// Type: IDENT, Position: 5 - 15, Literal: "non_recursive_factorial"
	// Type: EOL, Position: 5 - 38, Literal: ""
	// Type: IDENT, Position: 7 - 0, Literal: "5"
	// Type: :=, Position: 7 - 2, Literal: ":="
	// Type: \, Position: 7 - 4, Literal: "\\"
	// Type: IDENT, Position: 7 - 5, Literal: "f"
	// Type: ., Position: 7 - 6, Literal: "."
	// Type: \, Position: 7 - 7, Literal: "\\"
	// Type: IDENT, Position: 7 - 8, Literal: "x"
	// Type: ., Position: 7 - 9, Literal: "."
	// Type: IDENT, Position: 7 - 10, Literal: "f"
	// Type: IDENT, Position: 7 - 12, Literal: "f"
	// Type: IDENT, Position: 7 - 14, Literal: "f"
	// Type: IDENT, Position: 7 - 16, Literal: "f"
	// Type: IDENT, Position: 7 - 18, Literal: "f"
	// Type: IDENT, Position: 7 - 20, Literal: "x"
	// Type: EOL, Position: 7 - 21, Literal: ""
	// Type: IDENT, Position: 9 - 0, Literal: "fact"
	// Type: IDENT, Position: 9 - 5, Literal: "5"
	// Type: EOF, Position: 9 - 6, Literal: ""
}
