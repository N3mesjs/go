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

	fmt.Println("------- Slices! -------") //Notice that end is not included
	s:= make([]int, 3) //make a slice with lenght 3, initialized with 0
	fmt.Println(s, len(s))
	fmt.Println(cap(s), cap(b))
	s = append(s, 1,2,3,4,5)
	fmt.Println(s, len(s), cap(s))

	t:= s[0:3] //Notice that end is not included
	t = append(t, 100)
	fmt.Println(t, len(t), cap(t))

	copy(t, s) //overwrite the firt min(len(t), len(s)) elements of t with s
	fmt.Println(t)
}