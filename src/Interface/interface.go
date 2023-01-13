package main

import (
	"fmt"
)

func main() {
	myInt := IntCounter(0)
	var inc Incrementer = &myInt
	for i := 0; i < 10; i++ {
		fmt.Println(inc.Increment())
	}
	var w Writer = ConsoleWriter{}
	a, b := w.Write([]byte("Hello GO!"))
	fmt.Println(a, b)
}

type Incrementer interface {
	Increment() int
}
type IntCounter int

func (ic *IntCounter) Increment() int {
	*ic++
	return int(*ic)
}

// interface don't describe data,
// interface describe method
type Writer interface {
	Write([]byte) (int, error)
}
type ConsoleWriter struct{}

// implement interface
func (cw ConsoleWriter) Write(data []byte) (int, error) {
	n, err := fmt.Println(string(data))
	return n, err
}
