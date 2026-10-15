package main

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

type Order struct {
	ID     int
	Status string
	mu     sync.Mutex
}

var (
	totalUpdates int 
	updateMutex sync.Mutex
)

func main() {
	var wg sync.WaitGroup
	wg.Add(4)
	orders := generateOrder(20)

	// go func() {
	// 	defer wg.Done()
	// 	processOrders(orders)
	// }()

	for i := 0; i < 3; i++ {
		go func() {
			defer wg.Done()
			for order := range orders {
				updateOrderStatus(&orders[order])
			}
		}()
	}

	wg.Wait()

	reportOrdersStatuses(orders)
	fmt.Println("All operations done!")
	fmt.Println(totalUpdates)
}

func processOrders(orders []Order) {
	for _, order := range orders {
		time.Sleep(time.Duration(rand.Intn(500)) * time.Millisecond)
		fmt.Printf("Processing order %d\n", order.ID)
	}

}

func updateOrderStatus(order *Order) {
	order.mu.Lock()
	time.Sleep(
		time.Duration(time.Duration(rand.Intn(500)) *
			time.Millisecond,
		))
	status := []string{"Processing", "Shipped", "Delivered"}[rand.Intn(3)]
	order.Status = status
	fmt.Printf("Updated order ID: %d, status: %s\n",
		order.ID, order.Status,
	)

	order.mu.Unlock()


	// With this method some goroutines
	// Will not make any modifications due to 
	// race conditions, it will work wrong without
	// mutedx
	updateMutex.Lock()
	defer updateMutex.Unlock()
	currentUpdates := totalUpdates
	time.Sleep(5 * time.Millisecond)
	totalUpdates = currentUpdates + 1
}

func generateOrder(count int) []Order {
	orders := make([]Order, count)

	for i := range count {
		orders[i] = Order{ID: i, Status: "Pending"}
	}

	return orders
}

func reportOrdersStatuses(orders []Order) {
	fmt.Printf("---- Order Status Report ----\n")
	for _, val := range orders {
		fmt.Printf("Order ID: %d, Status: %s\n", val.ID, val.Status)
	}
	fmt.Printf("---------------------------\n")

}
