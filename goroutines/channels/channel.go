package main

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

type Order struct {
	id int
	status string
}

func main() {
	var wg sync.WaitGroup
	wg.Add(3)

	var orderChan chan *Order = make(chan *Order, 20)
	var processChan chan *Order = make(chan *Order, 20)
	go func(){
		defer wg.Done()
		for _, order := range generateOrders(20){
			orderChan <- order
		}

		close(orderChan)
		fmt.Printf("Finished orders generation\n ")
	}()

	go func(){
		processOrders(orderChan, processChan, &wg)
	}()

	go func(){
		defer wg.Done()

		for {
			select {
			case processedOrder, ok := <- processChan:
				if !ok {
					return
				}

				fmt.Printf("Processed order no. %d, status %s\n",
					processedOrder.id, processedOrder.status,
				)
			
			case <-time.After(3*time.Second):
				fmt.Println("Timeout due to no orders")
				return
			}
		}
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

func processOrders(inChan <-chan *Order, outChan chan<- *Order, wg *sync.WaitGroup){
	defer func(){
		wg.Done()
		close(outChan)
	}()

	for order := range inChan {
		time.Sleep(
			time.Duration(rand.Intn(500)) *
			time.Millisecond,
		)
		order.status = "Processed"
		outChan <- order
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
