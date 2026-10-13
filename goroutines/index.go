package main

import (
	"fmt"
	"time"
)

func count(s string) {
	for i := range 10 {
		fmt.Println(s, i)
		time.Sleep(time.Millisecond * 500)
	}
}

func main() {
	go count("fish")
	count("cat")
}
