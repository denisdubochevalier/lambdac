package parser

import "github.com/denisdubochevalier/lambdac/lexer"

// State encapsulates the internal state of the parser during the parsing
// phase of the compilation process. The state comprises two critical
// components:
//
// Fields:
//   - Tokens: An array of lexer.Token instances representing the tokenized
//     source code. This array serves as the "input tape" for the parser and is
//     traversed linearly to construct the Abstract Syntax Tree (AST).
//   - Position: An integer that keeps track of the parser's current position
//     within the Tokens array. As parsing proceeds, Position is incremented to
//     advance through the token sequence.
//
// Together, these two fields allow the parser to maintain a snapshot of its
// current status, facilitating features like backtracking and error reporting.
type State struct {
	tokens   []lexer.Token
	position int
	astNode  ASTNode
}

// NewState is a constructor function for initializing the State structure that
// serves as the parser's state. The function accepts a list of tokens produced
// by the lexer and constructs a new State instance with the given tokens.
//
// The initialized State comprises:
//   - A list of lexer tokens (`tokens`) that are to be parsed.
//   - The current position (`position`) within that list, initially set to 0.
//   - An empty AST node (`astNode`), which serves as the starting point for
//     building the Abstract Syntax Tree (AST) during the parsing process.
//
// By centralizing the construction of the initial parser state, NewState
// enhances the modularity and reusability of the parsing subsystem.
//
// Parameters:
//   - tokens: An array of lexer.Token instances representing the lexical units
//     to be parsed.
//
// Returns:
//   - A newly initialized State instance, prepared for the commencement of the
//     parsing process.
func NewState(tokens []lexer.Token) State {
	return State{
		tokens:   tokens,
		position: 0,
		astNode:  ASTNode{},
	}
}

// advance is a method on the State struct that moves the current position
// forward by one unit. The function creates a new State instance with the
// updated position and returns it.
//
// The method is instrumental in progressing through the list of tokens during
// parsing, effectively serving as the stepping stone from one state to another
// in the state machine encapsulated by the State monad.
//
// While the original State remains immutable, the updated position reflects
// the incremental progress in parsing, thereby helping to maintain the
// integrity and traceability of the parsing process.
//
// Returns:
// - A new State instance with the position incremented by one.
func (s State) advance() State {
	s.position = s.position + 1
	return s
}

// withAST is a method on the State struct that returns a new State instance
// containing the updated AST node. The method achieves this without mutating
// the original State object, thereby adhering to the principles of functional
// programming and immutability.
//
// The updated AST node typically encapsulates the parsed elements up to the
// current position in the token stream. By incorporating the new AST node into
// the state, withAST effectively captures the parsing progress in a type-safe,
// traceable manner.
//
// Parameters:
// - ast: The new ASTNode to be incorporated into the State.
//
// Returns:
// - A new State instance containing the updated AST node.
func (s State) withAST(ast ASTNode) State {
	s.astNode = ast
	return s
}

// done is a method on the State struct that returns a boolean indicating
// whether the parser has reached the end of the token stream. This serves
// as a termination criterion for the parsing loop, signaling when no more
// tokens are left to be processed.
//
// The method compares the current position index against the length of the
// token array, effectively acting as a guard to prevent out-of-bounds access.
// This allows for a more robust parser design, making the code less prone to
// errors related to index overflows.
//
// Returns:
//   - true if the parser has consumed all tokens and reached the end of the
//     token stream.
//   - false otherwise.
func (s State) done() bool {
	return s.position >= len(s.tokens)
}

// currentToken is a method on the State struct that retrieves the lexer.Token
// currently under consideration by the parser. The method is implemented to
// offer a higher level of abstraction over direct array indexing operations,
// thereby enhancing code readability and maintainability.
//
// It returns the lexer.Token located at the current 'position' index within
// the 'tokens' array, effectively allowing the parser to assess and act upon
// the token's semantic value and syntactic role.
//
// Returns:
//   - lexer.Token representing the token at the current position within the
//     parser state.
//
// Note:
// Care must be taken to invoke this method only when the parser has not yet
// reached the end of the token stream, i.e., the 'done()' method should return
// false. Failing to do so may lead to index out-of-bounds errors.
func (s State) currentToken() lexer.Token {
	return s.tokens[s.position]
}

// ast is a method on the State struct responsible for retrieving the current
// Abstract Syntax Tree (AST) node being operated upon or constructed by the
// parser. This function serves as an encapsulated getter for the 'astNode'
// field of the State struct, adhering to the principle of information hiding.
//
// The method is chiefly implemented to isolate direct field access, thereby
// increasing modularity and providing a single point of interaction for
// obtaining the current AST node. This simplifies potential future changes
// to the internal structure or behavior related to AST manipulation.
//
// Returns:
//   - ASTNode representing the current node in the abstract syntax tree within
//     the parser state.
func (s State) ast() ASTNode {
	return s.astNode
}
