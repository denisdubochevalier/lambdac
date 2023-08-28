# lambdac

![License: BSD 2-Clause + Charity](https://img.shields.io/badge/License-BSD%202--Clause%20%2B%20Charity-blue)

A minimalist compiler written in go for untyped lambda calculus with a focus on
simplicity and elegance.

**Project Status: First Stages of Development**

Please note that this project is still in its first stages of development and
should not get used in production environments. The initial versions of the
compiler are experimental and primarily intended for testing, feedback, and
contributions from the developer community. We appreciate your interest and
support in the project and encourage you to check back for updates.

## Table of Contents

- [Charity Ware](#charity-ware)
- [Background](#background)
- [Installation](#installation)
- [Usage](#usage)
- [Contributing](#contributing)
- [License](#license)

## Charity Ware

As an act of kindness, if you find this software useful, please consider making
a donation to the Electronic Frontier Foundation (EFF), a leading nonprofit
organization defending civil liberties in eht digital world. EFF champions user
privacy, free expression, and innovation through impact litigation, policy
analysis, grassroots activism, and technology development. Donations are not
required to use the software, but are greatly appreciated and will help support
a good cause. For more information on the EFF or to make a donation, please
visit EFF's Website: https://www.eff.org/.

## Background

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

## Installation

To install the compiler, follow these instructions:

```sh
$ git clone git@github.com:denisdubochevalier/lambdac.git
$ cd lambdac
$ make && make install
```

## Usage

The syntax of the language is close to the initial lambda calculus. For example:

```haskell
f = \x.\y.x
y = \f.(\x.f(x x))(\x.f(x x))
```

To call a function, you can use the following syntax:

```haskell
f (\x.\y.y)
```

To include a module, use the `|` symbol followed by the module's path:

```haskell
io | 'github.com/foo/fileio'
```

For more detailed usage instructions and examples, please refer to the
documentation.

## Contributing

Contributions are welcome! Please the following these steps to contribute:

1. Fork the repository.
2. Create your feature branch: `git checkout -b my-new-feature`.
3. Commit your changes `git commit -am 'Add some feature'`.
4. Push to the branch: `git push origin my-new-feature`.
5. Submit a pull request.

For more detailed contributing guidelines, please refer to the
[CONTRIBUTING.md](/blob/main/CONTRIBUTING.md) file.

## License

I am licensing this project under the BSD 2-Clause License with a Charity
extension. For more information, please refer to the
[LICENSE](/blob/main/LICENSE) file.
