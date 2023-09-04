+++

title = 'Writing a Lexer from Scratch'

date = 2023-09-03T06:48:06+02:00

draft = false

author = 'Denis Chevalier'

description = 'Project goals explanation'

tags = ['compiler','project','status','lexical analyzer', 'lexer']

categories = ['status', 'blog']

series = ['Project Advencement']

+++

## Introduction

Greetings to those who tread the fine line between artistry and engineering.
Today, we journey into the labyrinthine depths of compiler construction, with
our focus on the creation of a lexer from scratch for the minimalist lambda
calculus language, λ.c.

## The Genesis of λ.c

Before we venture further, it is vital to understand the ethos behind λ.c. This
language aims to synthesize mathematical beauty with computational pragmatism,
and in doing so, occupy a unique space within the spectrum of programming
paradigms. λ.c is both an intellectual exercise and a work of engineering craft.

## The Lexical Semiotics of λ.c

Our sojourn begins with identifying the atomic syntactic units of the
language—tokens. Identifiers, integers, strings, and composite operators like :=
and -> populate the lexical universe of λ.c. In particular, λ.c requires tokens
that echo its lambda calculus roots while remaining syntactically simple yet
expressively potent.

## The Creative Process

### The Conundrum of State

One of the initial challenges was managing the state during tokenization. A
naive approach could quickly devolve into a maze of conditional statements,
becoming an arduous affair both to write and debug. This led me to conceptualize
tokenization as a form of state management, where each state is represented as a
monad.

### Monad, the Epiphany

The idea to use a monadic approach to encapsulate state was a defining moment.
Monads in functional programming are paradigms for input-output transformation,
capturing both state and action in a singular functional construct. With this
approach, managing state became not just efficient, but elegant.

```go
// Example Code
type LexerFunc func(*Lexer) LexerFunc
```

### Pitfalls and Evasions

During development, I grappled with the challenge of efficiently handling
composite operators. Naive handling of such operators can lead to an explosion
of states. However, by employing a "lookahead" mechanism, I was able to
preemptively identify these operators, thus ensuring the lexer’s efficiency and
accuracy.

## Architectural Ingenuity in the Lexer

### A Functional Detour in Go

In most lexers built with Go, the state is generally maintained using struct
fields and methods, a pattern that could easily lead to an entanglement of
mutable state. Instead, I decided on a decidedly functional approach,
implementing each lexical state as a higher-order function returning another
function. This functional orientation deviates from Go's conventional
object-oriented style, representing a fascinating syntactic and semantic
crossover.

```go
// Example Code
func StartLexing(l *Lexer) LexerFunc {
    return LexRoot
}
```

### Higher-order States as Functions

This unique architecture is underpinned by the concept of higher-order states, a
fascinating diversion from common lexer designs. In λ.c's lexer, a `LexerFunc`
type represents a state and is a function that takes a lexer object and returns
the next state as a function.

```go
// Example Code
type LexerFunc func(*Lexer) LexerFunc
```

By maintaining the state in this way, the architecture is not only simplified
but gains a temporal coherency; each state knows what comes next without
requiring a complicated state machine or event loop.

### Immutable Elegance

The functional approach brings about immutability, a quality often associated
with functional programming languages rather than Go. This lack of mutable state
makes the lexer easier to test, debug, and reason about, while also embracing
the virtues of functional programming paradigms within the Go ecosystem.

### Synergy with Monads

The decision to employ a monadic architecture synergizes exceedingly well with
this functional construct, reinforcing the notion that monads and functional
programming paradigms can have a rightful place in Go—a language not
traditionally associated with these concepts.

## Lessons in Transdisciplinary Innovation

The lexer for λ.c stands at the crossroads of disciplines: borrowing functional
programming paradigms to improve upon the object-oriented disposition of Go,
harnessing the mathematical rigor of monads, and reconciling all these elements
in the realms of both computer science and software engineering.

## What's Next?

While the lexer has reached a functional state, several areas could be further
refined:

1. **Composite Operators**: There is a plan to extend the syntax tree to
   accommodate more complex expressions.

2. **Error Handling**: One of my next steps is to incorporate detailed
   diagnostic messages that enhance the developer experience.

3. **Unicode Support**: The decision to include or exclude Unicode is pending,
   as it could expand the language's versatility or detract from its
   minimalistic intent.

## The Road Ahead: Testing as the Next Frontier

It is important to note that this session has been singularly focused on the
architectural nuances of the lexer. As any seasoned engineer would attest,
architecture without testing is like a ship without a compass—technically
functional but directionally ambivalent. Therefore, the next chapter in the λ.c
saga will be dedicated to rigorous testing. This will involve creating a
comprehensive suite of unit tests and possibly even delving into property-based
testing to ensure that the lexer not only functions as expected but is also
robust under a myriad of conditions. After all, the true merit of any
architectural decision manifests most transparently under the scrutiny of
rigorous testing.

## Conclusion

The construction of the λ.c lexer has been an enlightening journey—a blend of
theoretical knowledge, inventive problem-solving, and painstaking engineering.
This endeavor has reinforced the delicate equilibrium between functionality and
elegance, between craft and science.

As λ.c continues to grow, so will our collective understanding of the
complexities and delights that accompany the process of crafting a compiler from
scratch. I invite you, dear reader, to join me in this grand adventure.
