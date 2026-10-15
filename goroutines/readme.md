# Goroutines
Goroutines are a core feature of the language that
make developer chose go.
This lets us use go with thread that
execute at the same time.

see /index.go file.

## How to use them
we invoke a function with the keyword go:

```go
go count("fish")
```

In this way the function will execute immediatly
without waiting for other function to finish.

Notice that it will stop whenever the main
function finishes its execution

## How does it work?
The go routines are managed by the go `runtime`, 
that is based on a scheduler that puts these routines
inside the pc threads

## Work Groups
The sync package gives us access to `sync.WorkGroup`
its a struct that gives us the chance to 
block the main function and wait for the goroutines
to finish, so basically we invoke the `sync.WorkGroup.Add()`
and we say the amout of goroutines, `.Done()` with this
function it will decrement the counter.

There is also `.Wait()` it will block the main function to 
wait for the execution of the goroutines, when they will
finish the mail will continue its execution

See the exercise

## Mutex
We can also have mutex for race conditions
to lock a certain block of code and make 
possible for every instance to access it

## Channels
Channels are like pipes that go routines can use 
to transfer data in realtime between each other
so basically see the example, its like a blocking
function, so the reciever will stay in a waiting
stauts until the sender closes it with `close(channel name)`

also in the parameter we can say if our function is
a channel reciever or sender:

```go
func foo(meat <-chan string) //reciever
func foo(meat chan<- string) //sender
```
