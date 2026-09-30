package main

import "fmt"

func main() {
	defer fmt.Println("This is printed last")
	fmt.Println("This is printed first")

	for i:=0; i<5; i++ {
		defer fmt.Println(i)
	}
}
