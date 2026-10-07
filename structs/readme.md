# Methods

Methods are functions that have and additional field
that is the `receiver` field.

syntax:
```go
func (v Vertex) Abs() float64 {
	return math.Sqrt(v.X*v.X + v.Y*v.Y)
}
```
in this example the receivers is a **Vertex** that 
in this case in the function is called **v**.

So we can call this method with the dot notation.

## Implicit pointer receivers
We can also have as receiver a pointer
so we can access the values of the type
and we can modify it withoud having copies of the same
type.

When you call a method with a pointer receiver on an
addressable value, Go automatically takes its address.

or if our functions requires a normal variable but we
have a pointer we can omit the following syntax: `(*varname).Method()`

here's an example:
```go
func (v *Vertex) Scale(factor int){
	v.X *= factor
	v.Y *= factor
}

point := Vertex{1,2}
point.Scale(2) //implicit (&point).Scale

var p *Vertex = &Vertex{1,2}
p.Scale(34) //implicit (*p)
```
Go will manage that for us!
For readaility it would be more effective to do:

```go
point := &Vertex{1,2}
point.Scale(2)
```