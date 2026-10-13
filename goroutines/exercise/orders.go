package main

import (
	"fmt"
	"math/rand"
	"time"
)

type Order struct{
	ID int
	Status string
}

func processOrders(orders []Order) {
	for _, order := range orders {
		time.Sleep(time.Duration(rand.Intn(500)) * time.Millisecond)
		fmt.Printf("Processing order %d\n", order.ID)
	}
}

func updateStatusOrders(orders []Order){
	for i := range orders {
		orders[i].Status = []string{"Pending", "Delivered", "Finished"}[rand.Intn(3)]
		
		time.Sleep(time.Duration(rand.Intn(500)) * time.Millisecond)
		fmt.Printf("Processing order %d, Status %s\n", orders[i].ID, orders[i].Status)
	}
}

func generateOrder(count int) []Order {
	orders := make([]Order, count)

	for i := range count {
		orders[i] = Order{ID:i, Status: "Pending"}
	}

	return orders
}

func main() {
	orders := generateOrder(20)

	processOrders(orders)
	updateStatusOrders(orders)

	fmt.Println(orders)
}
