package main

import "fmt"

func main(){
	// var m map[string]int nil map, non initialized map
	var m map[string]int = map[string]int{}
	fmt.Println("Empty map:", m)
	m["one"] = 1
	fmt.Println("Map with one element:", m)
}
