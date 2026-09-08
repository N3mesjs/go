# Pointers in Go

Go has pointers, allowing you to pass references to values and records within your program.

> [!WARNING]
> **No Pointer Arithmetic:** Unlike C, Go does **not** allow pointer arithmetic (e.g., `ptr++` is a compilation error). Pointers exist only for referencing and dereferencing data safely.

## 1. Syntax and Zero Value
A pointer type is declared with an asterisk `*` followed by the base type:
```go
var p *int // Declares a pointer to an int. Zero value is nil.
```

## 2. Deferenciation and & operator
the & operator is used to get the address of a 
variable to assign to a pointer variable.

We use ***** with a pointer variable to get the value 
of the pointed variable