package lexer

// Position represents a 2D coordinate within a source text,
// where 'row' refers to the line number and 'col' refers to the
// column number. These coordinates are used to identify the
// location of tokens within the lexical structure of the input.
type Position struct {
	row int // row represents the current row in the text.
	col int // col represents the current column in the text.
}

// StartPosition initializes a new Position instance with its row set
// to 1 and column set to 0. This is generally used as the starting
// point for lexical analysis, marking the beginning of the source text.
func StartPosition() Position {
	return Position{1, 0}
}

// Row returns the row (line number) of a Position instance.
// This provides an immutable way to access the row, upholding
// the principle of encapsulation.
func (p Position) Row() int {
	return p.row
}

// Col returns the column number of a Position instance.
// Similar to Row(), this method provides an immutable way to access
// the column, ensuring data integrity.
func (p Position) Col() int {
	return p.col
}

// newRow advances the Position to the start of a new row (line),
// resetting the column to 0. This method is commonly invoked
// upon encountering a newline character during lexical analysis.
func (p Position) newRow() Position {
	p.row++
	p.col = 0
	return p
}

// advanceCol increments the column number of the current Position instance by 1.
// This method returns a new Position instance with the updated column, adhering to
// the principle of immutability. This function is typically invoked when the lexer
// processes a character that does not alter the row count, thereby advancing the position
// within the same line of text.
func (p Position) advanceCol() Position {
	return p.advanceColBy(1)
}

// advanceColBy increments the column number of the current Position instance by a given integer value 'n'.
// This method returns a new Position instance with the updated column count, adhering to the principle
// of immutability. This function is particularly useful when the lexer processes a sequence of characters
// that do not affect the row count but do advance the position within the same line by more than one column.
func (p Position) advanceColBy(n int) Position {
	p.col += n
	return p
}
