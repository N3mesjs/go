package main

import "fmt"

type Vertex struct {
	X int
	Y int
}

func vertex_Init(x, y int) (v Vertex) {
	v = Vertex{x, y}
	return
}

func (v Vertex) String() string {
	return fmt.Sprintf("(%d, %d)", v.X, v.Y)
}

// Notice that if we have a pointer to a struct
// and try to access the values inside it, we can omit
// the following syntax (*ptr).x, go will automatically
// manage it for us!

func main(){
	p1 := vertex_Init(3, 5)

	ptr := &p1
	fmt.Printf("%v %d %d %v\n", ptr, ptr.X, ptr.Y, *ptr)
	fmt.Print(p1.String())
}