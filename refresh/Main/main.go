package main

import "fmt"
import "rsc.io/quote"
import "github.com/N3mesjs/golang/greetings"

func main() {
	fmt.Println("Hello, World!")
	fmt.Println(quote.Glass())
	fmt.Println(quote.Hello())
	fmt.Println(greetings.Hello("Nemesjs"))
}