![λ.c](https://github.com/denisdubochevalier/lambdac/raw/main/lambdac_logo.png)

# λ.c

[![License: BSD 2-Clause + Charity](https://img.shields.io/badge/License-BSD%202--Clause%20%2B%20Charity-blue)](LICENSE)
![GitHub go.mod Go version (subdirectory of monorepo)](https://img.shields.io/github/go-mod/go-version/denisdubochevalier/lambdac)
[![GoDoc](https://godoc.org/github.com/denisdubochevalier/lambdac?status.svg)](https://pkg.go.dev/github.com/denisdubochevalier/lambdac)
![Build Status](https://github.com/denisdubochevalier/lambdac/actions/workflows/go.yml/badge.svg)
![Lint Status](https://github.com/denisdubochevalier/lambdac/actions/workflows/golangci-lint.yml/badge.svg)
[![Go report](https://goreportcard.com/badge/github.com/denisdubochevalier/lambdac)](https://goreportcard.com/report/github.com/denisdubochevalier/lambdac)
[![Coverage](https://img.shields.io/codecov/c/github/denisdubochevalier/lambdac)](https://codecov.io/gh/denisdubochevalier/lambdac)

A minimalist compiler written in go for untyped lambda calculus with a focus on
simplicity and elegance.

**Project Status: First Stages of Development**

Please note that this project is still in its first stages of development and
should not get used in production environments. The initial versions of the
compiler are experimental and primarily intended for testing, feedback, and
contributions from the developer community. We appreciate your interest and
support in the project and encourage you to check back for updates.

## Table of Contents

- [Introduction](#introduction)
- [Installation](#installation)
- [Contributing](#contributing)
- [Charity Ware](#charity-ware)
- [License](#license)

## Introduction

Welcome to the introductory post of λ.c (lambdac), a new programming language
designed to dive deep into the foundations of computation and bring the elegance
of untyped lambda calculus to the modern developer.

Lambda calculus, developed by Alonzo Church in the 1930s, is a framework of
expressing computations based on function abstraction and application. It is the
smallest universal programming language, meaning that it is capable of
expressing any computable function. While it forms the theoretical backbone of
many popular programming languages today, lambda calculus itself is rarely used
directly due to its minimalistic and abstract nature.

That’s where λ.c comes in. This language aims to bring the power and simplicity
of the untyped lambda calculus to your fingertips. It provides a minimalistic
syntax and feature set, staying true to its roots while adding just enough
conveniences to make it practical for real-world use.

In λ.c, everything is a function, and functions are first-class citizens. You
can define functions, pass them as arguments, and return them as results. The
language uses a mix of Polish notation (prefix) and currying to make the syntax
as clean and intuitive as possible. For example, the addition of two numbers can
be expressed as `+ 1 2`, which is equivalent to `(add 1) 2` in a curried
notation.

This site will serve as a documentation, a journey log, and a platform for
discussion. Throughout the coming [posts](lambdac/post), we will delve into the
language's design, its implementation, and explore various topics related to
lambda calculus, functional programming, and compiler construction.

Thank you for embarking on this journey with us. Stay tuned for more updates,
and in the meantime, feel free to join the discussion on our
[GitHub repository](https://github.com/denisdubochevalier/lambdac).

## Installation

To install the compiler, follow these instructions:

```sh
$ git clone git@github.com:denisdubochevalier/lambdac.git
$ cd lambdac
$ make && make install
```

## Contributing

Contributions are welcome! Please the following these steps to contribute:

1. Fork the repository.
2. Create your feature branch: `git checkout -b my-new-feature`.
3. Commit your changes `git commit -am 'Add some feature'`.
4. Push to the branch: `git push origin my-new-feature`.
5. Submit a pull request.

For more detailed contributing guidelines, please refer to the
[CONTRIBUTING.md](https://github.com/denisdubochevalier/lambdac/blob/main/CONTRIBUTING.md)
file.

## Charity Ware

As an act of kindness, if you find this software useful, please consider making
a donation to the Electronic Frontier Foundation (EFF), a leading nonprofit
organization defending civil liberties in the digital world. EFF champions user
privacy, free expression, and innovation through impact litigation, policy
analysis, grassroots activism, and technology development. Donations are not
required to use the software, but are greatly appreciated and will help support
a good cause. For more information on the EFF or to make a donation, please
visit EFF's Website: https://www.eff.org/.

## License

I am licensing this project under the BSD 2-Clause License with a Charity
extension. For more information, please refer to the [LICENSE](lambdac/license).
