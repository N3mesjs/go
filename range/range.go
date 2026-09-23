package main

import "fmt"

/*
Range: is a special keyword accessible only in for loops. It returns
two values, the index and the value of the element at that index.
It will work with any iterable data structure such as arrays, slices,
strings, maps ..
*/

func main(){
	m := map[string]int{
		"apple": 1,
		"banana": 2,
		"cherry": 3,
	}

	fmt.Println("Map:", m)

	for i := range m {
		fmt.Println("Key:", i, "Value:", m[i])
	}

	// or 

	for key, value := range m {
		fmt.Println("Key:", key, "Value:", value)
	}

	// clear the map
	clear(m)
	fmt.Println("Cleared map:", m)
}
