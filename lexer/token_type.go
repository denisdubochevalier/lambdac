package lexer

// TokenType represents the type of a token
type TokenType int

const (
	ILLEGAL TokenType = iota // ILLEGAL represents an illegal character.
	EOF                      // EOF (End Of File) represents the end of the file that is being parsed.
	EOL                      // EOL (End Of Line) represents the end of a line.
	IDENT                    // IDENT represents an identifier, which could be a variable or a function name (e.g., m, recursive_fact, fact, ...).
	ASSIGN                   // ASSIGN represents the assignment operator (:=).
	MODULE                   // MODULE represents the module operator (|).
	NSDEREF                  // NSDEREF represents the namespace dereference operator (->).
	LAMBDA                   // LAMBDA represents the lambda operator (\).
	DOT                      // DOT represents the dot operator (.).
	STRING                   // STRING represents a string literal (e.g., "github.com/foo/bar", "text/lexer", ...).
	LPAREN                   // LPAREN represents the left parenthesis (().
	RPAREN                   // RPAREN represents the right parenthesis ()).
)

var values = []string{
	ILLEGAL: "ILLEGAL",
	EOF:     "EOF",
	EOL:     "EOL",
	IDENT:   "IDENT",
	STRING:  "STRING",
	LAMBDA:  "\\",
	DOT:     ".",
	LPAREN:  "(",
	RPAREN:  ")",
	MODULE:  "|",
	NSDEREF: "->",
	ASSIGN:  ":=",
}

// String returns a string representation of the TokenType.
// This is helpful for debugging and logging purposes,
// as it allows you to print out a human-readable version of the token type.
func (t TokenType) String() string {
	// handle unknown values
	if int(t)+1 > len(values) || int(t) < 0 {
		return "UNKNOWN"
	}
	return values[t]
}
