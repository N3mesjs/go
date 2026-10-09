package main

import (
	"fmt"
	"time"
)

type MyError struct {
	Time time.Time
	Msg  string
}

func (e *MyError) Error() string {
	return fmt.Sprintf("[ERROR]%v: %v", e.Time, e.Msg)
}

func run() error {
	return &MyError{
		time.Now(),
		"it didn't work",
	}
}

func main() {
	if err := run(); err != nil {
		fmt.Println(err)
	}
}
