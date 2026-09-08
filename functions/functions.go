package main

import (
	"fmt"
	"math"
)

func add(x, y int) int {
	return x+y
}

func swap(str1, str2 string) (string, string) {
	return str2, str1
}

/**
 * Named Return Values & Naked Returns:
 * By naming the return variable in the function signature (e.g., 'res float64'),
 * Go automatically declares and initializes it with its zero value.
 *
 * We can use a "naked return" (just 'return') and Go returns the current value of 'res'.
 * Best practice: keep naked returns for short, simple functions.
 * Nothing stops us from using an explicit return ('return res') for better readability!
 */

func exp(base, exp float64) (res float64) {
	res = math.Pow(base, exp)
	return
	// return res
}

func main() {
	fmt.Println("the sum of 43 + 12 is:", add(43, 12))

	//var str1 string = "ciao amici!"
	// var str1 = "ciao amici!"
	str1 := "ciao amici"
	str2 := "followeres!"
	//fmt.Println(str1 + str2) creates a third string
	fmt.Println(str1, str2)
	str1, str2 = swap(str1, str2)
	fmt.Println(str1, str2)

	fmt.Println(exp(3, 2))
}