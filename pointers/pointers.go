package main

import "fmt"

type Person struct {
	name string
	age int
}

type Node struct {
	val int
	next *Node
}

func loadPerson(p *Person) {
	p.name = "John"
	p.age = 30
}

func main() {
	var p *int

	i := 42
	p = &i

	fmt.Println(*p)

	// Example of the use of the new() function to create a pointer to a struct

	newPtr := new(Person)

	fmt.Println(*newPtr, newPtr.name, newPtr.age)

	// ---

	list := new(Node)
	list.val = 1
	list.next = new(Node)
	list.next.val = 2

	fmt.Println(list.val, list.next.val)

	// ----

	newPerson := new(Person)
	loadPerson(newPerson)
	fmt.Println(*newPerson)

	// or we can use the & operator to create a pointer to a struct literal
	
	newPerson2 := &Person{}
	loadPerson(newPerson2)
	fmt.Println(*newPerson2)
}
