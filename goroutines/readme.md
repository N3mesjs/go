# Goroutines
Goroutines are a core feature of the language that
make developer chose go.
This lets us use go with thread that
execute at the same time.

see /index.go file.

## How to use them
we invoke a function with the keyword go:

```go
go count("fish")
```

In this way the function will execute immediatly
without waiting for other function to finish.

Notice that it will stop whenever the main
function finishes its execution
