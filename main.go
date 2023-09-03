// Package main implements the lambdac compiler entry point.
package main

import (
	// importing ragel to force presence in go.mod
	_ "github.com/db47h/ragel/v2"
)

//go:generate ragel -Z -G2 -o lex.go lex.rl
func main() {}
