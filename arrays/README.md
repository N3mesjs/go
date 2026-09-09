# Arrays in Go
Arrays are pretty simple we have multiple ways to declare them
but they are limited to the given size.
They come with the default method `len(arrayname)`

## Zeros initialization
When declaring an array if we only declare the variable it will automatically
be set to all zeros:
```go
var a [5]int //result [0,0,0,0,0]
```

## `...` notation
the three dots tells the compiler to calculate the size of the
array by the given elements

## Index insertion
When declaring an array and writing its element we can specify on which index
to put the element and the elements before that aren't declared will be set to
zeros:
```go
arr := [5]int{100, 3:400, 500} // [100,0,0,400,500]
```

## Two dimensional arrays
You can also have 2D arrays by using two square brackets like this `[][]`

