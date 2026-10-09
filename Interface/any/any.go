package main

import "fmt"

// switch type assertion
func do(i interface{}) {
	switch v := i.(type) {
	case uint:
		fmt.Println("Unsigned Integer:", v)
	case string:
		fmt.Println("String:", v)
	default:
		fmt.Println("Unknown type")
	}
}

func main() {
	var i interface{} = uint(42)

	var num uint = i.(uint) // type assertion
	fmt.Println(num)
	//var str string = i.(string) // will panic at runtime since i does not hold a string
	str, ok := i.(string) // type assertion with check
	fmt.Println("String:", str, "ok:", ok)
	do(i)
}
