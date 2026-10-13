package main

import (
	"fmt"
	//"time"
)

type Order struct{
	ID int
	Status string
}

func generateOrder(count int) *[]Order {
	orders := make([]Order, count)

	for i := range count {
		orders[i] = Order{ID:i, Status: "Pending"}
	}

	return &orders
}

func main() {
	orders := generateOrder(20)

	fmt.Println(orders)
}
