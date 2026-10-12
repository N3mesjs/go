# Generics
Generics is a way to define function in a way to 
accept a parameter that can be of any type:

```go
func foo(x []T) T {}
```

In this case we have as parameter a variable of type
`T` and it will also return a value that is `T`.

## Comparable interface
There is an interface that we can use for generics
to make them comparable, so we can use `!=` or `==`

here is an example:
```go
func foo[T comparable](x []T, val T) T {}
```

in this way we can for example compare a value of `x`
with val.