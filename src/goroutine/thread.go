package main

import (
	"fmt"
	"sync"
)

var wg = sync.WaitGroup{}
var counter = 0
var m = sync.RWMutex{}

func main() {
	for i := 0; i < 10; i++ {
		//two more goroutines
		wg.Add(2)
		m.RLock()
		go sayHello()
		//to mutatate data
		m.Lock()
		go increment()
	}
	//main func doesn't exist out early
	wg.Wait()
}
func sayHello() {
	fmt.Printf("Hello #%v\n", counter)
	m.RUnlock()
	wg.Done()
}
func increment() {
	counter++
	m.Unlock()
	wg.Done()
}
