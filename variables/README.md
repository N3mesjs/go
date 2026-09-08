# Variables in Go

Go is a statically and strictly typed language. Variables can be declared in different ways thanks to **type inference**.

## 1. Explicit Type (Zero Value)
`var a string`
We explicitly define the type. When no initial value is provided, Go initializes it with its **Zero Value** (`""` for strings, `0` for numbers, `false` for booleans, `nil` for pointers/slices/maps).

## 2. Type Inference with `var`
`var a = "initial"`
The type is inferred by the compiler at compile time based on the value on the right.
* **Scope:** Can be used **both inside and outside** functions (package level).

## 3. Short Variable Declaration (`:=`)
`a := "cacca"`
Shorthand syntax that declares, infers, and assigns in a single step.
* **Scope:** Works **ONLY inside functions**.
* **Bonus:** Allows redeclaration of an existing variable if at least one other variable in the tuple is new (common with `result, err := ...`).