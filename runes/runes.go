package main

import "fmt"

func main() {
	var a rune = 'a'
	var b rune = 'b'

	println(a, b)

	s := "caffè"
  	fmt.Println(len(s)) // 6: conta i byte, non i caratteri
}
