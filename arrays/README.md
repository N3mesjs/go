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

---

# Slices
slices is a dinamically sized array, but it doesn't store values
its self but it uses an array to store values.

We can create a slice starting from an existing array using
the following syntax:
```go
array := b[start:end] //the end is excluded, this choise was made to facilitate the lenght calculation
```

## Zeros initialization
For the slices if we initialize like `[]int` it will be empty with no 
elements, but we can give it an initial size and it will have all zeros
values like using the following syntax:
```go
array := make([]int, 3, 5)
```
The third parameters that is the **cap** so a limit for the size of
the underlying array, it is used for ottimization to use an array
of a size instead of using a bigger one for nothing.
So the initial lenght will be 3 but it can grow up to **5 elements**,
After it will create a new array for the slice that will be **bigger**