package lexer

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestPosition(t *testing.T) {
	t.Parallel() // Run all tests in this suite in parallel for efficiency

	t.Run("StartPosition", func(t *testing.T) {
		t.Parallel() // Parallel execution within the parent suite
		is := require.New(t)

		pos := StartPosition()

		is.Equal(1, pos.Row()) // Should start at row 1
		is.Equal(0, pos.Col()) // Should start at column 0
	})

	t.Run("Row", func(t *testing.T) {
		t.Parallel()
		is := require.New(t)

		pos := Position{row: 5, col: 7}
		is.Equal(5, pos.Row())
	})

	t.Run("Col", func(t *testing.T) {
		t.Parallel()
		is := require.New(t)

		pos := Position{row: 5, col: 7}
		is.Equal(7, pos.Col())
	})

	t.Run("newRow", func(t *testing.T) {
		t.Parallel()
		is := require.New(t)

		pos := Position{row: 1, col: 1}
		newPos := pos.newRow()

		is.Equal(2, newPos.Row())
		is.Equal(0, newPos.Col())
	})

	t.Run("advanceCol", func(t *testing.T) {
		t.Parallel()
		is := require.New(t)

		pos := Position{row: 1, col: 1}
		newPos := pos.advanceCol()

		is.Equal(1, newPos.Row())
		is.Equal(2, newPos.Col())
	})

	t.Run("advanceColBy", func(t *testing.T) {
		t.Parallel()
		is := require.New(t)

		pos := Position{row: 1, col: 1}
		newPos := pos.advanceColBy(3)

		is.Equal(1, newPos.Row())
		is.Equal(4, newPos.Col())
	})
}
