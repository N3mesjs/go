# Stringer Interface

In go inside the fmt package there is 
a Stringer interface like this:
```go
type Stringer interface{
    String() string
}
```

The fmt package uses it to look for this 
interface to print values.

For example the **fmt.Println** takes the any
type but if the variable implements the
**Stringer** interface, it will use its method
to print.

`Also many other packages uses it to see
the print method of the specific type
that implements that interface`