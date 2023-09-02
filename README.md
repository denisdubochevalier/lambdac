![λ.c](/lambdac_logo.png)

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

- [Documentation](#documentation)
- [Installation](#installation)
- [Contributing](#contributing)
- [Charity Ware](#charity-ware)
- [License](#license)

## Documentation

I am documenting the project
[on the official website](https://denisdubochevalier.github.io/lambdac).

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
[CONTRIBUTING.md](/CONTRIBUTING.md) file.

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
extension. For more information, please refer to the [LICENSE](/LICENSE) file.
