# error interface
In the standard language there is the 
standard error interface, like this:

```go
type error interface {
    Error() string
}
```

Every Error that implements this interface becomes an error.

Notice that like the **fmt.Stringer**, `fmt.Println` for example 
will see if our variable is an **Error** and will use its Error()
method to use the format specified to print!

`This one is the same behaviour of Stringer`

Take a look also to Stringer, it has the capital letter because
its exported from the **fmt library**, but the error interface
is not exported, but its part of the language