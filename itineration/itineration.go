package main

import "fmt"

func main(){
	for i:=0; i<3; i++ {
		fmt.Println(i)
	}

	var j int

	for j = 0; j<3; j++ {
		fmt.Println(j)
	}

	for i := range 6 {
		fmt.Println("range", i)
	}
}
