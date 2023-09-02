+++

title = 'Launching λ.c: A Journey into the Heart of Computation'

date = 2023-09-02T21:16:50+02:00

draft = false

+++

Welcome to the first official blog post for λ.c (lambda c), a new programming
language that aims to bring the elegance and power of untyped lambda calculus to
modern developers. I am excited to share the journey of this project, from its
inception to its current status, and how you can contribute to its development.

## Why λ.c?

Lambda calculus, a mathematical framework developed by Alonzo Church in the
1930s, is the smallest universal programming language, capable of expressing any
computable function. It forms the theoretical foundation of many popular
programming languages today - like Common Lisp and Haskell -, yet it is rarely
used directly due to its minimalistic and abstract nature.

I have always been fascinated by the simplicity and expressive power of lambda
calculus. While playing with my usual work and toy projects, I realized the
potential of creating a language that stays true to the principles of lambda
calculus while adding just enough conveniences to make it practical for
real-world use.

λ.c is born out of this realization. It aims to provide a minimalistic syntax
and feature set, staying true to its roots while making it accessible and
practical for modern developers. Everything is a function in λ.c, and functions
are first-class citizens. You can define functions, pass them as arguments, and
return them as results. The language uses a mix of Polish notation (prefix) and
currying to make the syntax as clean and intuitive as possible.

## Project Goals

The primary goal of λ.c is to create a minimalistic, yet practical programming
language based on untyped lambda calculus. Here are some specific objectives:

1. **Minimalistic Syntax**: Implement a syntax that is as close as possible to
   the initial lambda calculus, with minor modifications for practicality.
2. **Common Combinators**: Automatically include a prefix with the definition of
   all common combinators (like K, S, C, ...) in any program.
3. **Module Imports**: Implement a folder-based hierarchy for importing modules
   that implement common data types like Church Literals, Pairs, Lists, Real and
   Complex Numbers arithmetic, but also IO and Networking.
4. **No Initial Type System**: Start without a type system, possibly adding it
   as an extension or module later on.
5. **Self-Hosting Compiler**: Eventually reimplement the compiler in λ.c itself.

## Current Status

As of now, λ.c is in its infancy. The project is in the planning and
specification stage, with initial drafts of the syntax and feature set being
developed. The goal is to start with a very minimal implementation and
incrementally build upon it, documenting the process along the way.

The project is hosted on GitHub, and I am currently working on setting up the
repository, drafting the initial specifications, and planning the development
roadmap.

## How to Contribute

Contributions are warmly welcomed and greatly appreciated. Here are some ways
you can contribute to the project:

1. **Feedback**: Provide
   [feedback](https://github.com/denisdubochevalier/lambdac/issues) on the
   language design, syntax, and feature set. Your insights and suggestions can
   help shape the direction of the project.
2. **Documentation**: Help with
   [documenting](https://github.com/denisdubochevalier/lambdac/tree/main/site)
   the language, its features, and the development process. Good documentation
   is crucial for the success of any open-source project.
3. **Development**: Contribute to the development of the language, its compiler,
   and standard library. Whether you are an experienced developer or just
   starting, there are
   [many ways](https://github.com/denisdubochevalier/lambdac/blob/main/CONTRIBUTING.md)
   to contribute to the codebase.
4. **Testing**: Help with testing the language and its features. Testing is a
   critical aspect of the development process and ensures the language's
   reliability and stability. If you are interested in contributing, please
   check out the
   [GitHub repository](https://github.com/denisdubochevalier/lambdac) for more
   information on how to get started.

## Conclusion

I am thrilled to embark on this journey and excited to see where it leads.
Developing a programming language is a challenging and rewarding endeavor, and I
am grateful for the opportunity to contribute to the community in this way.

Thank you for your interest in λ.c and for joining me on this adventure. Stay
tuned for more updates, and in the meantime, feel free to join the discussion on
our [GitHub repository](https://github.com/denisdubochevalier/lambdac).
