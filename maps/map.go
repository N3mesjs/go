package main

import "fmt"

func main(){
	// var m map[string]int nil map, non initialized map, its nil
	var m map[string]int = map[string]int{}
	fmt.Println("Empty map:", m)
	m["one"] = 1
	fmt.Println("Map with one element:", m)
	m["two"] = 2
	m["three"] = 3

	fmt.Println("third element:", m["three"])
	val, ok := m["three"]
	fmt.Println(val, ok)
	delete(m, "three")
	fmt.Println(m)
	val, ok = m["three"]
	fmt.Println(val, ok)

	clear(m)
	fmt.Println(m)
}
