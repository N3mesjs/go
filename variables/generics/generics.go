package main

import "fmt"

func findIndex[T comparable](s []T, x T) int{
	for i, v := range s {
		if v == x {
			return i
		}
	}

	return -1
}

func main() {
		// Index works on a slice of ints
	si := []int{10, 20, 15, -10}
	fmt.Println(findIndex(si, 15))

	// Index also works on a slice of strings
	ss := []string{"foo", "bar", "baz"}
	fmt.Println(findIndex(ss, "hello"))
}
