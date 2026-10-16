# Goroutines

Goroutines are lightweight units of execution managed by the Go runtime. They make it easy to run functions concurrently, which is one of the reasons Go is a popular choice for concurrent programs.

See [`index.go`](index.go) for a simple example.

## Starting a goroutine

Use the `go` keyword to start a function in a new goroutine:

```go
go count("fish")
```

The call returns immediately, so the current function can continue while `count` runs. A goroutine is not guaranteed to finish before `main` returns: when `main` exits, the program exits too. Use a synchronization tool when you need to wait for goroutines.

## How goroutines run

The Go runtime schedules goroutines onto operating-system threads. Many goroutines can be handled by a smaller number of threads, and the runtime can move goroutines between threads as they run.

## Waiting with `sync.WaitGroup`

`sync.WaitGroup` lets a function wait until a group of goroutines has finished:

```go
var wg sync.WaitGroup

wg.Add(1)
go func() {
	defer wg.Done()
	work()
}()

wg.Wait()
```

- `Add` increases the counter (often once for each goroutine).
- `Done` decreases it when a goroutine finishes. `defer` helps ensure it is called.
- `Wait` blocks until the counter returns to zero.

See the exercise for another example.

## Mutexes

A mutex protects shared data when multiple goroutines may access it at the same time. Lock it before accessing the data, then unlock it when finished. This prevents simultaneous access to the protected section; it does not by itself make unrelated code safe.

```go
mu.Lock()
sharedValue++
mu.Unlock()
```

In real code, `defer mu.Unlock()` is often useful when a function has multiple return paths.

## Channels

Channels let goroutines send and receive values. An unbuffered channel synchronizes the sender and receiver: a send waits until a receiver is ready, and a receive waits until a value is sent.

```go
messages := make(chan string)

go func() {
	messages <- "hello"
}()

message := <-messages
fmt.Println(message)
```

You can restrict a channel parameter to one direction:

```go
func receive(messages <-chan string) {} // receive only
func send(messages chan<- string) {}    // send only
```

Closing a channel signals that no more values will be sent. A receiver does not need to wait for the channel to close before receiving: it can receive values as they arrive. After closure, any values already sent can still be received; later receives return the zero value immediately. Usually, the sender is responsible for closing the channel.

## Buffered channels

A buffered channel can hold a limited number of values without requiring a receiver for each send immediately:

```go
messages := make(chan string, 2)
messages <- "one"
messages <- "two"
close(messages)

for message := range messages {
	fmt.Println(message)
}
```

Here the buffer has a capacity of two. A send blocks when the buffer is full, and a receive blocks when it is empty. After the channel is closed, a `range` loop continues until all buffered values have been received.

## Selecting between channel operations

The `select` statement lets a goroutine wait on several channel operations at once. Each `case` describes a send or receive. Go runs one case whose operation is ready; if multiple cases are ready, it chooses one pseudo-randomly. If none is ready, `select` blocks until one can proceed.

This is useful when a worker must handle incoming values while also reacting to a timeout. The following example follows the order-processing flow in [`channels/channel.go`](channels/channel.go): it receives processed orders from one channel and stops if no order arrives before the timeout.

```go
for {
	select {
	case order, ok := <-processChan:
		if !ok {
			// The sender closed the channel: there will be no more orders.
			return
		}
		fmt.Printf("Processed order %d: %s\n", order.id, order.status)

	case <-time.After(3 * time.Second):
		fmt.Println("Timed out waiting for an order")
		return
	}
}
```

The receive uses the two-value form, `order, ok := <-processChan`. `ok` is `false` only after the channel has been closed and all buffered values have been received. Checking it prevents treating the zero value of `*Order` (`nil`) as a real order.

`time.After` returns a channel that receives a value after the specified duration. Since it is inside the loop, this example starts a fresh three-second timeout on each iteration. That makes it an **inactivity timeout**: each received order starts a new wait period. If the timeout should apply to the whole operation instead, create the timer before the loop and select on the same timer channel each time.

### Non-blocking `select`

Adding a `default` case makes `select` non-blocking. If no channel operation is ready, Go runs `default` immediately:

```go
select {
case message := <-messages:
	fmt.Println("Received:", message)
default:
	fmt.Println("No message available right now")
}
```

Use this when you want to check a channel without waiting. Avoid a tight loop around a `default` case, since it can repeatedly consume CPU while doing no useful work.
