+++

title = 'Design Goals'

date = 2023-09-02T20:00:28+02:00

draft = false

+++

Lambda calculus is a formal system in mathematical logic for expressing
computation based on function abstraction and application using variable binding
and substitution. This project aims to provide a minimalist compiler for untyped
lambda calculus, with a simple and elegant syntax, and a standard library
implementing common data types and operations.

I am designing the compiler to be as minimal as possible, focusing on the core
principles of lambda calculus, and will built it with the following features:

- A simple syntax close to the initial lambda calculus, using `\` as a
  replacement for the lambda symbol
- A standard library including common combinators (K, S, C, ...) and data types
  (Church Literals, Pairs, Lists, Real and Complex Numbers, ...).
- Module import functionality based on a folder hierarchy, allowing the
  inclusion of custom modules.
- Eventually, a self-hosting compiler re-implementation in the language itself.
