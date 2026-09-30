package main

import "fmt"

func main() {
	var a rune = 'a'
	var b rune = 'b'

	println(a, b)

	s := "caffè"
  	fmt.Println(len(s)) // 6: conta i byte, non i caratteri
	const sample = "\xbd\xb2\x3d\xbc\x20\xe2\x8c\x98"
	fmt.Println(sample)
	for i, r := range sample {
		fmt.Printf("%d: %U\n", i, r)
	}

	str := "⌘"

	fmt.Println(len(str))
	fmt.Println([]byte(str))
}
