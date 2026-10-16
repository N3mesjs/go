package main

import (
	"golang.org/x/tour/tree"
	"fmt"
)

func main() {
	var ch chan int = make(chan int)
	go Walk(tree.New(1), ch)
	
	for val := range ch {
		fmt.Println(val)
	}
	
	fmt.Println(Same(tree.New(1), tree.New(1)))
	fmt.Println(Same(tree.New(1), tree.New(2)))
}

// Walk walks the tree t sending all values
// from the tree to the channel ch.
func Walk(t *tree.Tree, ch chan int){
	var walker func(*tree.Tree)
	walker = func(node *tree.Tree){
		if(node == nil){
			return
		}
		walker(node.Left)
		ch <- node.Value
		walker(node.Right)
	}

	walker(t)
	close(ch)
}

// Same determines whether the trees
// t1 and t2 contain the same values.
func Same(t1, t2 *tree.Tree) bool {
	cht1 := make(chan int)
	cht2 := make(chan int)

	go Walk(t1, cht1)
	go Walk(t2, cht2)

	for {
		n1, ok1 := <- cht1
		n2, ok2 := <- cht2

		if ok1 != ok2 || n1 != n2 {
			return false
		}
		if !ok1 {
			break
		}
	}
	return true
}