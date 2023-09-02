+++

title = 'Syntax Specification'

date = 2023-09-02T20:00:35+02:00

draft = false

+++

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
