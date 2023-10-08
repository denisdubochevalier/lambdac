package parser

import (
	"github.com/denisdubochevalier/monad"

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

// replaceLastChild aims to substitute the last child node of an ASTNode
// instance (`a`) with another ASTNode (`node`). This method plays a crucial
// role in the parser's capacity to modify the evolving Abstract Syntax Tree
// (AST) in an immutable manner.
//
// The function returns a monad.Maybe[ASTNode] that:
//   - Encloses the modified ASTNode (`a`) with the replaced child in a
//     monad.Some if the operation succeeds.
//   - Yields a monad.None to indicate failure, especially if `a` has no
//     children nodes to replace.
//
// Operational Schema:
//   - Initially, the function assesses the existence of children nodes in
//     the calling ASTNode (`a`).
//   - If there are no children, the method returns a monad.None[ASTNode].
//   - Otherwise, the last child node of `a` is replaced with the given
//     ASTNode (`node`), and a monad.Some wrapping the modified `a` is
//     returned.
//
// Monadic Design Note:
// The use of a Maybe monad elegantly sidesteps the need for exception
// handling, in turn, adhering to functional programming principles.
func (a ASTNode) replaceLastChild(node ASTNode) monad.Maybe[ASTNode] {
	if len(a.children) == 0 {
		return monad.None[ASTNode]()
	}

	a.children[len(a.children)-1] = node
	return monad.Some(a)
}

// lastChild retrieves the last child node of the calling ASTNode instance
// (`a`) in the context of constructing or traversing an Abstract Syntax Tree
// (AST). This method is typically invoked during the parser's recursive
// descent to access the most recently added child for potential
// modification or validation.
//
// The function returns a monad.Maybe[ASTNode] which:
//   - Encases the last child node in a monad.Some if `a` contains children.
//   - Returns a monad.None if `a` has no children, thereby signaling the
//     absence of a target node.
//
// Operational Schema:
// - The method inspects the length of the `children` slice of `a`.
// - If `a` has no children, it returns a monad.None[ASTNode].
// - Otherwise, it returns the last child in a monad.Some wrapper.
//
// Monadic Design Note:
// This method, like its counterparts, employs the Maybe monad to elegantly
// handle the absence of a value without resorting to error handling, aligning
// with functional programming norms.
func (a ASTNode) lastChild() monad.Maybe[ASTNode] {
	if len(a.children) == 0 {
		return monad.None[ASTNode]()
	}
	return monad.Some(a.children[len(a.children)-1])
}

// NodeType is a straightforward accessor method that returns the type of the
// calling ASTNode (`a`) as defined by the NodeType enumeration. This method is
// chiefly utilized for type validation and conditional branching during the
// recursive descent parsing process.
//
// The function serves multiple roles in ensuring the coherence and accuracy of
// the parser:
//   - In the discriminative phases, it is invoked to match tokens with expected
//     node types to maintain syntactic correctness.
//   - In AST manipulations, it is used to programmatically decide how to adapt
//     the tree based on token types.
//   - It also has diagnostic utility, aiding in error reporting or logging by
//     providing the type of the node under consideration.
//
// Design Note:
// The method embodies the encapsulation principle by providing controlled
// access to an object's internal state, aligning with Object-Oriented
// Programming best practices.
//
// Complexity: O(1)
func (a ASTNode) NodeType() NodeType {
	return a.nodeType
}
