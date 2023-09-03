package lex

import (
	"fmt"
	"log"
	"strings"
)

func ExampleLex() {
	input := strings.NewReader(`m | "github.com/foo/bar\n"
	Y := \f.(\x.f(x x))(\x.f(x x))
	recursive_fact := Y m->fact`)

	tokens, err := Lex(input)
	if err != nil {
		log.Fatal(err)
	}
	for _, tok := range tokens {
		fmt.Println(tok)
	}
	// Output:
	// {IDENT m}
	// {MODULE |}
	// {STRING github.com/foo/bar\n}
	// {IDENT Y}
	// {ASSIGN :=}
	// {LAMBDA \}
	// {IDENT f}
	// {DOT .}
	// {LPAREN (}
	// {LAMBDA \}
	// {IDENT x}
	// {DOT .}
	// {IDENT f}
	// {LPAREN (}
	// {IDENT x}
	// {IDENT x}
	// {RPAREN )}
	// {RPAREN )}
	// {LPAREN (}
	// {LAMBDA \}
	// {IDENT x}
	// {DOT .}
	// {IDENT f}
	// {LPAREN (}
	// {IDENT x}
	// {IDENT x}
	// {RPAREN )}
	// {RPAREN )}
	// {IDENT recursive_fact}
	// {ASSIGN :=}
	// {IDENT Y}
	// {IDENT m}
	// {NSDEREF ->}
	// {IDENT fact}
	// {EOF }
}
