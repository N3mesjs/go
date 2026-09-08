package main

import "fmt"

func main() {
    /** 
    * if multiple variables share the same type
    * we can declare the variables and then assing the type
    */
    var var1, var2 int

    var1 = 12
    var2 = 13
    fmt.Println(var1, var2)

    var a = "initial"
    fmt.Println(a)

    var b, c int = 1, 2
    fmt.Println(b, c)

    var d = true
    fmt.Println(d)

    var e int
    fmt.Println(e)

    f := "apple"
    fmt.Println(f)

    const str string = "ciao amici!"

    //str = "bah" we cannot perform this operation!
}