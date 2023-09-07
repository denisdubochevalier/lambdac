package lexer

import (
	"reflect"
	"testing"

	"github.com/denisdubochevalier/monad"
	"github.com/leanovate/gopter"
	"github.com/leanovate/gopter/gen"
	"github.com/leanovate/gopter/prop"
	"github.com/stretchr/testify/require"
)

func TestFinalizeIdentifierTokenBasic(t *testing.T) {
	t.Parallel()

	t.Run("when content is empty", func(t *testing.T) {
		t.Parallel()
		is := require.New(t)

		l := New().WithContent("")
		result, newLexer := finalizeIdentifierToken(l)

		// test that newLexer is equal to l
		is.Equal(monad.None[Token](), result)
		assertEqualLexer(is, l, newLexer)
	})

	t.Run("when content is not empty", func(t *testing.T) {
		t.Parallel()
		is := require.New(t)

		// Test when content is not empty
		l := New().WithContent("abc")
		result, newLexer := finalizeIdentifierToken(l)
		token, ok := result.(monad.Just[Token])

		is.True(ok) // Expecting Some(Token)
		is.Equal(token.Value().Type(), IDENT)

		// test that newLexer is equal to l
		assertEqualLexer(is, l, newLexer)
	})
}

func TestFinalizeIdentifierTokenPropertyBased(t *testing.T) {
	t.Parallel()

	t.Run(
		"Property 1: finalizeIdentifierToken should return None for an empty content string",
		func(t *testing.T) {
			t.Parallel()

			parameters := gopter.DefaultTestParameters()
			properties := gopter.NewProperties(parameters)

			properties.Property("Should return None for empty content", prop.ForAll(
				func(dummy bool) bool {
					l := New().WithContent("")
					result, _ := finalizeIdentifierToken(l)
					_, ok := result.(monad.Nothing[Token])
					return ok // Expecting None
				},
				gen.Const(true), // Dummy generator
			))
			properties.TestingRun(t)
		},
	)

	t.Run(
		"Property 2: finalizeIdentifierToken should return Some(Token) for a non-empty content string",
		func(t *testing.T) {
			t.Parallel()

			parameters := gopter.DefaultTestParameters()
			properties := gopter.NewProperties(parameters)

			properties.Property("Should return Some(Token) for non-empty content", prop.ForAll(
				func(str string) bool {
					if str == "" {
						return true // Skip empty string, already tested in Property 1
					}
					l := New().WithContent(str)
					result, _ := finalizeIdentifierToken(l)
					_, ok := result.(monad.Just[Token])
					return ok // Expecting Some(Token)
				},
				gen.AlphaString(), // Generates non-empty strings
			))
			properties.TestingRun(t)
		},
	)

	t.Run("Property 3: The token should always be of type IDENT", func(t *testing.T) {
		t.Parallel()

		parameters := gopter.DefaultTestParameters()
		properties := gopter.NewProperties(parameters)

		properties.Property("Token type should always be IDENT", prop.ForAll(
			func(str string) bool {
				if str == "" {
					return true // Skip empty string, as it will return None
				}
				l := New().WithContent(str)
				result, _ := finalizeIdentifierToken(l)
				token, ok := result.(monad.Just[Token])
				return ok && token.Value().Type() == IDENT
			},
			gen.AlphaString(), // Generates non-empty strings
		))
		properties.TestingRun(t)
	})
}

func TestFinalizeIdentifierTokenReferentialTransparency(t *testing.T) {
	t.Parallel()
	is := require.New(t)
	l := New().WithContent("abc")
	result1, l1 := finalizeIdentifierToken(l)
	result2, l2 := finalizeIdentifierToken(l)
	is.Equal(result1, result2) // Should produce the same output given the same input

	assertEqualLexer(is, l1, l2)
}

// TestHigherOrderFinalizeIdentifierToken tests the higher-order function behavior
// of finalizeIdentifierToken by using higher-order functions to encapsulate test scenarios.
func TestHigherOrderFinalizeIdentifierToken(t *testing.T) {
	t.Parallel()

	// Define test scenarios using higher-order functions
	testCases := []struct {
		name          string
		lexer         Lexer
		expectedToken monad.Maybe[Token]
	}{
		{
			name:          "With empty content",
			lexer:         New().WithContent(""),
			expectedToken: monad.None[Token](),
		},
		{
			name:          "With non-empty content",
			lexer:         New().WithContent("abc"),
			expectedToken: monad.Some(Token{IDENT, StartPosition(), Literal("")}),
		},
		{
			name:          "With space content",
			lexer:         New().WithContent(" \t"),
			expectedToken: monad.Some(Token{IDENT, StartPosition(), Literal("")}),
		},
		{
			name:          "With operators",
			lexer:         New().WithContent(":="),
			expectedToken: monad.Some(Token{IDENT, StartPosition(), Literal("")}),
		},
	}

	for _, testCase := range testCases {
		testCase := testCase
		t.Run(testCase.name, func(t *testing.T) {
			t.Parallel()
			is := require.New(t)

			// Execute function
			token, lexer := finalizeIdentifierToken(testCase.lexer)

			// Validate results
			is.Equal(testCase.expectedToken, token)

			assertEqualLexer(is, testCase.lexer, lexer)
		})
	}
}

func TestIdLexRecursivelyBasic(t *testing.T) {
	t.Parallel()

	t.Run("when content is empty", func(t *testing.T) {
		t.Parallel()
		is := require.New(t)

		l := New().WithContent("")
		result, newLexer := idLexRecursively(l)

		is.Equal(monad.None[Token](), result)
		assertEqualLexer(is, l, newLexer)
	})

	t.Run("when content contains valid identifier", func(t *testing.T) {
		t.Parallel()
		is := require.New(t)

		l := New().WithContent("abc123")
		result, _ := idLexRecursively(l)
		token, ok := result.(monad.Just[Token])

		is.True(ok)
		is.Equal(Literal("abc123"), token.Value().Literal())
	})

	t.Run("when content contains valid identifier with unicode", func(t *testing.T) {
		t.Parallel()
		is := require.New(t)

		l := New().WithContent("👋🏻abc123")
		result, _ := idLexRecursively(l)
		token, ok := result.(monad.Just[Token])

		is.True(ok)
		is.Equal(Literal("👋🏻abc123"), token.Value().Literal())
	})

	t.Run("when content contains valid identifier with symbols", func(t *testing.T) {
		t.Parallel()
		is := require.New(t)

		l := New().WithContent("abc-123")
		result, _ := idLexRecursively(l)
		token, ok := result.(monad.Just[Token])

		is.True(ok)
		is.Equal(Literal("abc-123"), token.Value().Literal())
	})

	t.Run(
		"when content contains valid identifier with composite operator mid way",
		func(t *testing.T) {
			t.Parallel()
			is := require.New(t)

			l := New().WithContent("abc->123")
			result, _ := idLexRecursively(l)
			token, ok := result.(monad.Just[Token])

			is.True(ok)
			is.Equal(Literal("abc"), token.Value().Literal())
		},
	)

	t.Run(
		"when content contains valid identifier with simple operator mid way",
		func(t *testing.T) {
			t.Parallel()
			is := require.New(t)

			l := New().WithContent("abc|123")
			result, _ := idLexRecursively(l)
			token, ok := result.(monad.Just[Token])

			is.True(ok)
			is.Equal(Literal("abc"), token.Value().Literal())
		},
	)
}

func TestIdLexRecursivelyPropertyBased(t *testing.T) {
	t.Parallel()

	t.Run("End of string should return None", func(t *testing.T) {
		t.Parallel()

		parameters := gopter.DefaultTestParameters()
		properties := gopter.NewProperties(parameters)

		properties.Property("End of string should return None", prop.ForAll(
			func(dummy bool) bool {
				l := New().WithContent("")
				result, _ := idLexRecursively(l)
				_, ok := result.(monad.Nothing[Token])
				return ok
			},
			gen.Const(true),
		))
		properties.TestingRun(t)
	})

	t.Run("Valid identifier should be constructed", func(t *testing.T) {
		t.Parallel()

		parameters := gopter.DefaultTestParameters()
		properties := gopter.NewProperties(parameters)

		properties.Property("Valid identifier should be constructed", prop.ForAll(
			func(str string) bool {
				if str == "" || !isValidIdentifier(str) {
					return true
				}
				l := New().WithContent(str)
				result, _ := idLexRecursively(l)
				token, ok := result.(monad.Just[Token])
				return ok && token.Value().Literal() == Literal(str)
			},
			gen.AlphaString(),
		))
		properties.TestingRun(t)
	})
}

func TestIdLexRecursivelyReferentialTransparency(t *testing.T) {
	t.Parallel()
	is := require.New(t)

	l := New().WithContent("foo")
	result1, l1 := idLexRecursively(l)
	result2, l2 := idLexRecursively(l)

	is.Equal(result1, result2)
	assertEqualLexer(is, l1, l2)
}

// Helper function to assert that two Lexers are equal
func assertEqualLexer(is *require.Assertions, l1, l2 Lexer) {
	is.Equal(l1.content, l2.content)
	is.Equal(l1.position, l2.position)
	nlf1 := reflect.ValueOf(l1.nextLexerFunc)
	nlf2 := reflect.ValueOf(l2.nextLexerFunc)
	is.Equal(nlf1.Pointer(), nlf2.Pointer())
}

// Helper function to check if a string is a valid identifier according to
// the rules of your language
func isValidIdentifier(str string) bool {
	for _, c := range str {
		if !isValidIdentifierChar(c) {
			return false
		}
	}
	// Replace this with actual logic
	return true
}

func TestMergeLiterals(t *testing.T) {
	// Test Case 1: When both existing Token and newChar are empty
	t.Run("Both existing Token and newChar are empty", func(t *testing.T) {
		t.Parallel()
		is := require.New(t)

		existing := monad.None[Token]()
		newChar := rune(0)
		result := mergeLiterals(existing, newChar)
		token, ok := result.(monad.Just[Token])
		is.True(ok)
		is.Equal(Literal(newChar), token.Value().Literal())
	})

	// Test Case 2: When existing Token is non-empty and newChar is empty
	t.Run("Existing Token nostringn-empty, newChar empty", func(t *testing.T) {
		t.Parallel()
		is := require.New(t)

		existing := monad.Some(Token{literal: Literal("abc")})
		newChar := rune(0)
		result := mergeLiterals(existing, newChar)
		token, ok := result.(monad.Just[Token])
		is.True(ok)
		is.Equal(Literal("abc"), token.Value().Literal())
	})

	// Test Case 3: When existing Token is empty and newChar is non-empty
	t.Run("Existing Token empty, newChar non-empty", func(t *testing.T) {
		t.Parallel()
		is := require.New(t)

		existing := monad.None[Token]()
		newChar := 'a'
		result := mergeLiterals(existing, newChar)
		token, ok := result.(monad.Just[Token])
		is.True(ok)
		is.Equal(Literal(newChar), token.Value().Literal())
	})

	// Test Case 4: When both existing Token and newChar are non-empty
	t.Run("Both existing Token and newChar are non-empty", func(t *testing.T) {
		t.Parallel()
		is := require.New(t)

		existing := monad.Some(Token{literal: Literal("abc")})
		newChar := 'a'
		result := mergeLiterals(existing, newChar)
		token, ok := result.(monad.Just[Token])
		is.True(ok)
		is.Equal(Literal("aabc"), token.Value().Literal())
	})
}

func TestUpdateLexerForRecursion(t *testing.T) {
	t.Parallel() // Ensures the tests run concurrently where possible

	t.Run("Initial Position", func(t *testing.T) {
		t.Parallel()
		is := require.New(t)

		l := New().WithPosition(Position{row: 1, col: 1})
		size := 3
		xs := "foo"

		newLexer := updateLexerForRecursion(l, size, xs)

		is.Equal(1, newLexer.position.row)
		is.Equal(4, newLexer.position.col) // Should be 1 + size

		// Lexer should not be mutated
		assertEqualLexer(is, l, New().WithPosition(Position{row: 1, col: 1}))
	})

	t.Run("Advanced Position", func(t *testing.T) {
		t.Parallel()
		is := require.New(t)

		l := New().WithPosition(Position{row: 2, col: 5})
		size := 2
		xs := "bar"

		newLexer := updateLexerForRecursion(l, size, xs)

		is.Equal(2, newLexer.position.row)
		is.Equal(7, newLexer.position.col) // Should be 5 + size

		// Lexer should not be mutated
		assertEqualLexer(is, l, New().WithPosition(Position{row: 2, col: 5}))
	})

	t.Run("Content Update", func(t *testing.T) {
		t.Parallel()
		is := require.New(t)

		l := New().WithContent("initial")
		size := 0 // size shouldn't affect content
		xs := "updated"

		newLexer := updateLexerForRecursion(l, size, xs)

		is.Equal("updated", newLexer.content)

		// Lexer should not be mutated
		assertEqualLexer(is, l, New().WithContent("initial"))
	})

	t.Run("Next Lexer Function", func(t *testing.T) {
		t.Parallel()
		is := require.New(t)

		l := New().WithNextLexerFunc(identifierLexer)
		size := 0 // size shouldn't affect nextLexerFunc
		xs := ""

		newLexer := updateLexerForRecursion(l, size, xs)

		nlf1 := reflect.ValueOf(eofLexer)
		nlf2 := reflect.ValueOf(newLexer.nextLexerFunc)
		is.Equal(nlf1.Pointer(), nlf2.Pointer())

		// Lexer should not be mutated
		assertEqualLexer(is, l, New().WithNextLexerFunc(identifierLexer))
	})

	t.Run("Zero Size", func(t *testing.T) {
		t.Parallel()
		is := require.New(t)

		l := New().WithPosition(Position{row: 1, col: 1})
		size := 0
		xs := "zero"

		newLexer := updateLexerForRecursion(l, size, xs)

		is.Equal(1, newLexer.position.row)
		is.Equal(1, newLexer.position.col) // Should be unchanged

		// Lexer should not be mutated
		assertEqualLexer(is, l, New().WithPosition(StartPosition().advanceCol()))
	})
}

func TestIsValidIdentifierChar(t *testing.T) {
	t.Parallel()

	t.Run("Basic ASCII", func(t *testing.T) {
		t.Parallel()
		is := require.New(t)

		validASCII := "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789_"
		for _, x := range validASCII {
			is.True(isValidIdentifierChar(rune(x)))
		}

		invalidASCII := "\\ .()|"
		for _, x := range invalidASCII {
			is.False(isValidIdentifierChar(rune(x)))
		}
	})

	t.Run("Whitespace", func(t *testing.T) {
		t.Parallel()
		is := require.New(t)

		whitespace := "\t \n\r\v\f"
		for _, x := range whitespace {
			is.False(isValidIdentifierChar(rune(x)))
		}
	})

	t.Run("UTF-8 and Special Characters", func(t *testing.T) {
		t.Parallel()
		is := require.New(t)

		validUTF8 := "åçñü👍🚀"
		for _, x := range validUTF8 {
			is.True(isValidIdentifierChar(rune(x)))
		}
	})

	t.Run("Control Characters", func(t *testing.T) {
		t.Parallel()
		is := require.New(t)

		// Testing range 0x00 to 0x1F which includes ASCII control characters
		for x := rune(0x00); x <= rune(0x1F); x++ {
			is.False(isValidIdentifierChar(x))
		}
	})
}

func TestCheckCompositeOps(t *testing.T) {
	t.Parallel()
	is := require.New(t)

	// Test Case 1: Valid composite operator ":="
	t.Run("Valid composite operator :=", func(t *testing.T) {
		t.Parallel()
		x := ':'
		xs := "="
		is.True(checkCompositeOps(x, xs))
	})

	// Test Case 2: Valid composite operator "->"
	t.Run("Valid composite operator ->", func(t *testing.T) {
		t.Parallel()
		x := '-'
		xs := ">"
		is.True(checkCompositeOps(x, xs))
	})

	// Test Case 3: Invalid composite operator
	t.Run("Invalid composite operator", func(t *testing.T) {
		t.Parallel()
		x := ':'
		xs := "-"
		is.False(checkCompositeOps(x, xs))
	})

	// Test Case 4: Invalid character but valid remaining string
	t.Run("Invalid character but valid remaining string", func(t *testing.T) {
		t.Parallel()
		x := '+'
		xs := "="
		is.False(checkCompositeOps(x, xs))
	})

	// Test Case 5: Valid character but invalid remaining string
	t.Run("Valid character but invalid remaining string", func(t *testing.T) {
		t.Parallel()
		x := ':'
		xs := "+"
		is.False(checkCompositeOps(x, xs))
	})

	// Test Case 6: Both invalid character and invalid remaining string
	t.Run("Both invalid character and invalid remaining string", func(t *testing.T) {
		t.Parallel()
		x := '&'
		xs := "$"
		is.False(checkCompositeOps(x, xs))
	})

	// Test Case 7: Empty remaining string
	t.Run("Empty remaining string", func(t *testing.T) {
		t.Parallel()
		x := ':'
		xs := ""
		is.False(checkCompositeOps(x, xs))
	})

	// Test Case 8: Valid composite operator with extra characters
	t.Run("Valid composite operator with extra characters", func(t *testing.T) {
		t.Parallel()
		x := ':'
		xs := "=abc"
		is.True(checkCompositeOps(x, xs))
	})

	// Test Case 9: UTF-8 valid characters
	t.Run("UTF-8 valid characters", func(t *testing.T) {
		t.Parallel()
		x := '→'
		xs := "abc"
		is.False(checkCompositeOps(x, xs))
	})
}

func TestIdentifierLexer(t *testing.T) {
	t.Parallel()

	t.Run("when content is empty", func(t *testing.T) {
		t.Parallel()
		is := require.New(t)

		l := New().WithContent("")
		result, newLexer := identifierLexer(l)

		is.Equal(monad.None[Token](), result)
		assertEqualLexer(is, l, newLexer)
	})

	t.Run("when content contains valid identifier", func(t *testing.T) {
		t.Parallel()
		is := require.New(t)

		l := New().WithContent("abc123")
		result, _ := identifierLexer(l)
		token, ok := result.(monad.Just[Token])

		is.True(ok)
		is.Equal(Literal("abc123"), token.Value().Literal())
	})

	t.Run("when content contains valid identifier with unicode", func(t *testing.T) {
		t.Parallel()
		is := require.New(t)

		l := New().WithContent("👋🏻abc123")
		result, _ := identifierLexer(l)
		token, ok := result.(monad.Just[Token])

		is.True(ok)
		is.Equal(Literal("👋🏻abc123"), token.Value().Literal())
	})

	t.Run("when content contains valid identifier with symbols", func(t *testing.T) {
		t.Parallel()
		is := require.New(t)

		l := New().WithContent("abc-123")
		result, _ := identifierLexer(l)
		token, ok := result.(monad.Just[Token])

		is.True(ok)
		is.Equal(Literal("abc-123"), token.Value().Literal())
	})

	t.Run(
		"when content contains valid identifier with composite operator mid way",
		func(t *testing.T) {
			t.Parallel()
			is := require.New(t)

			l := New().WithContent("abc->123")
			result, _ := identifierLexer(l)
			token, ok := result.(monad.Just[Token])

			is.True(ok)
			is.Equal(Literal("abc"), token.Value().Literal())
		},
	)

	t.Run(
		"when content contains valid identifier with simple operator mid way",
		func(t *testing.T) {
			t.Parallel()
			is := require.New(t)

			l := New().WithContent("abc|123")
			result, _ := identifierLexer(l)
			token, ok := result.(monad.Just[Token])

			is.True(ok)
			is.Equal(Literal("abc"), token.Value().Literal())
		},
	)
}
