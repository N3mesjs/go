package main

import (
	"fmt"
	"sync"
	"time"
	"math/rand"
)

type Order struct {
	id int
	status string
}

func main() {
	var wg sync.WaitGroup
	wg.Add(2)

	var orderChan chan *Order = make(chan *Order)
	go func(){
		for _, order := range generateOrders(20){
			orderChan <- order
		}

		close(orderChan)
		wg.Done()
	}()

	go func(){
		processOrders(orderChan, &wg)
	}()

	wg.Wait()

	fmt.Println("All the processes done!")
}

func generateOrders(count int) (orders []*Order){
	orders = make([]*Order, count)
	for i := 0; i < count; i++ {
		orders[i] = &Order{id: i, status: "pending"}
	}

	return
}

func processOrders(orderChan <-chan *Order, wg *sync.WaitGroup){
	defer wg.Done()
	for order := range orderChan {
		time.Sleep(
			time.Duration(rand.Intn(500)) *
			time.Millisecond,
		)
		fmt.Printf("Processing order %d\n", order.id)
	}

	//same syntax but range manages automatically the
	// ok status variable
	// for {
	// 	order , ok := <-orderChan
	// 	if !ok {
	// 		break
	// 	}
	// 	...
	// }
}

func reportOrderStatuses(orders []*Order){
	fmt.Println("----- Report of Orders -----")
	for _, order := range orders {
		fmt.Printf("ID: %d, Status: %s\n",
			order.id, order.status,
		)
	}
	fmt.Println("---------------")
}
