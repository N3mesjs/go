package main

import "fmt"

func main() {
	var a [5]int
	fmt.Println("Empty array:", a)

	a[4] = 100
	fmt.Println(len(a), a[4])

	b:= [5]int{1,2,3,4,5}
	fmt.Println("lenght of b:", len(b))
	b = [...]int{1,2,3,4,5}
	fmt.Println("lenght of b:", len(b))
	b = [...]int{100, 3:400, 500}
	fmt.Println(b)

	// Chapter about slices

	//s:= make([]int, 3)
	s:= b[1:4]
	fmt.Println(s, len(s))
}