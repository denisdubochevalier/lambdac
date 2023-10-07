package parser

import (
	"github.com/denisdubochevalier/lambdac/lexer"
)

// ASTNode represents a node within the Abstract Syntax Tree (AST) that is
// constructed during the parsing phase. An AST is a hierarchical, tree-like
// representation of the syntactic structure of the source code, and each
// ASTNode serves as a building block of this tree.
//
// The ASTNode struct holds essential information that describes both the
// semantics and the syntax of individual constructs in the language.
//
// Fields:
//   - NodeType: Enumerated type that categorizes the node into various possible
//     syntactic constructs (e.g., Identifier, Lambda Expression, etc.)
//   - Token:    The lexical token corresponding to the node, encapsulating the
//     token's type, position in the source code, and its literal value.
//   - Children: An array of child ASTNodes, enabling the formation of the
//     tree-like hierarchical structure of the AST. For example, in the case of
//     a Lambda expression, Children might include pointers to the parameter and
//     body of the lambda.
//
// The NodeType allows for quick syntactic categorization, making it easier to
// traverse and manipulate the AST during subsequent phases like optimization
// and code generation.
//
// The Token field provides finer-grained information about the node, such as
// its exact location in the source code, which is invaluable for tasks like
// error reporting and source-to-source transformations.
//
// The Children field enables the recursive nature of the AST, allowing for
// complex expressions to be composed of simpler ones in a nested fashion.
//
// Example usage:
//
//		identifierNode := ASTNode{NodeType: lexer.IDENT, Token: identToken}
//		lambdaNode := ASTNode{
//	   NodeType: lexer.LAMBDA,
//	   Token: lambdaToken,
//	   Children: []ASTNode{identifierNode}
//	 }
type ASTNode struct {
	nodeType NodeType
	token    lexer.Token
	children []ASTNode
}

// newASTNode is a factory function that assembles and returns a fresh
// instance of ASTNode, configured with the specified NodeType and Token.
// It serves as a constructor for ASTNode, abstracting away the initialization
// details and providing a controlled interface for object instantiation.
//
// By employing this function, you decouple the rest of your codebase from
// the underlying representation of ASTNode, thereby adhering to principles
// of encapsulation and modular design.
//
// Parameters:
//   - nodeType: An enumerated value that specifies the type of the AST node.
//   - token: A lexer.Token that contains the lexical metadata associated with
//     the node.
//
// Returns:
//   - A newly instantiated ASTNode object populated with the provided
//     nodeType and token, and initialized with an empty slice for children.
func newASTNode(nodeType NodeType, token lexer.Token) ASTNode {
	return ASTNode{
		nodeType: nodeType,
		token:    token,
		children: []ASTNode{},
	}
}

// appendChild generates a new ASTNode by appending the provided child ASTNode
// to the children slice of the current instance. This method facilitates
// the incremental construction of the abstract syntax tree, extending its
// hierarchy with each new syntactic entity encountered during parsing.
//
// Note that this method is designed to be functionally pure: it returns a new
// ASTNode with the appended child rather than modifying the original instance
// in-place. This is consistent with the overarching functional programming
// paradigm employed in the codebase, where immutability and statelessness are
// prioritized.
//
// Parameters:
// - astNode: The ASTNode to be appended as a child.
//
// Returns:
// - A new ASTNode instance featuring the appended child in its children slice.
func (a ASTNode) appendChild(astNode ASTNode) ASTNode {
	a.children = append(a.children, astNode)
	return a
}
