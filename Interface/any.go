package interface

import "fmt"

var i interface{} = 42

var num uint = i.(uint) // type assertion

var str string = i.(string) // will panic at runtime since i does not hold a string
var str, ok := i.(string) // type assertion with check

// switch type assertion

switch v := i.(type) {
	case int:
		fmt.Println("Integer:", v)
	case string:
		fmt.Println("String:", v)
	default:
		fmt.Println("Unknown type")
}