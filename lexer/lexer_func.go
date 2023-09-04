package lexer

import "github.com/denisdubochevalier/monad"

// lexerFunc is a type alias for a function that embodies the core behavior of the lexer.
// It defines a lexical operation as a state transition: taking an existing Lexer instance
// and returning a new Lexer state, along with the processed Token encapsulated in a monad.Result.
// This architecture promotes a functional approach by encapsulating state transitions as
// pure functions, enhancing testability, and making the behavior more predictable.
// The use of a monad.Result allows for a robust way to handle both successful tokenization
// and errors without relying on Go's idiomatic error-handling mechanisms.
type lexerFunc func(l Lexer) (monad.Maybe[Token], Lexer)
