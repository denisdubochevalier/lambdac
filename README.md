![λ.c](/lambdac_logo.png)

# λ.c

[![License: BSD 2-Clause + Charity](https://img.shields.io/badge/License-BSD%202--Clause%20%2B%20Charity-blue)](LICENSE)
![GitHub go.mod Go version (subdirectory of monorepo)](https://img.shields.io/github/go-mod/go-version/denisdubochevalier/lambdac)
[![GoDoc](https://godoc.org/github.com/denisdubochevalier/lambdac?status.svg)](https://pkg.go.dev/github.com/denisdubochevalier/lambdac)
![Build Status](https://github.com/denisdubochevalier/lambdac/actions/workflows/go.yml/badge.svg)
![Lint Status](https://github.com/denisdubochevalier/lambdac/actions/workflows/golangci-lint.yml/badge.svg)
[![Go report](https://goreportcard.com/badge/github.com/denisdubochevalier/lambdac)](https://goreportcard.com/report/github.com/denisdubochevalier/lambdac)
[![Coverage](https://img.shields.io/codecov/c/github/denisdubochevalier/lambdac)](https://codecov.io/gh/denisdubochevalier/lambdac)

Crafted with an unwavering focus on simplicity and elegance, λ.c is more than
just a minimalist compiler. It's a vision of what computing can be at its most
elemental.

**🚨 Notice**: We're in the throes of initial development. The codebase is
experimental—perfect for the brave of heart and curious of mind. Come, be a
pioneer!
[See what we're building →](https://denisdubochevalier.github.io/lambdac/post)

## Table of Contents

- [Current Features](#current-features)
- [Roadmap](#roadmap)
- [Documentation](#documentation)
- [Installation](#installation)
- [Contribute](#contribute)
- [A Note on Kindness](#a-note-on-kindness)
- [License](#license)

## Current Features

### The Gateway: Lexer

As of now, our journey into untyped lambda calculus is heralded by a fully
operational lexer. This initial offering brings with it:

- Tokenization of basic lambda calculus syntax: variables, lambdas, and
  parentheses.
- Robust error handling for unexpected tokens.
- A glimpse into the kind of meticulous craftsmanship that will define future
  stages of this project.

## Roadmap

### AST Generation (In Progress)

- Create a robust and extensible AST to represent lambda calculus expressions.
- Implement parsing logic that translates token streams into the AST.

### Semantic Analysis (Upcoming)

- Validate the logical coherence of expressions.
- Introduce scope resolution and perform alpha-renaming as necessary.

### Intermediate Representations (Future)

- Translate the AST into one or more intermediate forms for optimization.
- Experiment with graph-based representations like Single Static Assignment
  (SSA).

### Optimization Techniques (Future)

- Implement constant folding, dead code elimination, and other classical
  optimization strategies.
- Explore lambda calculus-specific optimizations.

### Backend Generation (Future)

- Translate intermediate forms into target machine code or another high-level
  language.
- Experiment with generating LLVM IR or WebAssembly for broader platform
  support.

### Runtime Environment (Future)

- If applicable, develop a minimal runtime to manage program execution.
- Consider garbage collection or other memory management strategies.

### Documentation and Examples (Ongoing)

- Continue to document the architecture, algorithms, and design patterns used.
- Create example programs to demonstrate features and educational use-cases.

### Community Building (Ongoing)

- Engage with early users to gather feedback and iterate.
- Develop contributor guidelines and good first issues to welcome new
  developers.

## Documentation

Dive deep into the inner workings of λ.c through our
[detailed documentation](https://denisdubochevalier.github.io/lambdac). Take
your understanding from zero to hero—no calculus textbook required!

## Installation

Follow the path of enlightenment:

```sh
$ git clone git@github.com:denisdubochevalier/lambdac.git
$ cd lambdac
$ make && make install
```

## Contribute

Dare to contribute? It's simple:

1. Fork this repository.
2. Check out your own feature branch:
   `git checkout -b feature/wondrous-feature`.
3. Craft your magic: `git commit -am 'Add my wondrous feature'`.
4. Share with the world: `git push origin feature/wondrous-feature`.
5. Open a pull request and join the ranks of the enlightened.

For even more details, consult our [Contributing Guide](/CONTRIBUTING.md).

## A Note on Kindness

Though it costs nothing to use λ.c, consider extending your kindness by donating
to the [Electronic Frontier Foundation (EFF)](https://www.eff.org). It's not
required, but it's a way to contribute to the fight for civil liberties in the
digital world.

## License

λ.c is released under the BSD 2-Clause License. See the [LICENSE](/LICENSE) file
for details.
