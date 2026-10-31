package main

import (
	"encoding/json/v2"
	"fmt"
	"time"
)

type Message struct {
	Name string
	Body string
	Time time.time	
}

func main(){
	m := message{Name: "Alice", Body: "Hello", Time: time.Now}
	fmt.Println(m)
}
