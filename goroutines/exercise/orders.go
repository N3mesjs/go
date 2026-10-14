package main

import (
	"fmt"
	"math/rand"
	"sync"
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

func reportOrders(orders []Order){
	for i:=0; i<5; i++{
		time.Sleep(time.Second * 2)
		fmt.Printf("---- Order Status Report ----\n")
		for _, val := range orders {
			fmt.Printf("Order ID: %d, Status: %s\n", val.ID, val.Status)
		}
		fmt.Printf("---------------------------\n")
	}
}

func main() {
	var wg sync.WaitGroup;
	wg.Add(3)
	orders := generateOrder(20)

	go processOrders(orders)
	go updateStatusOrders(orders)
	go reportOrders(orders)

	wg.Wait()
	fmt.Println("All operations done!")
}
