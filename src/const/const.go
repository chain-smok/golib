package main

import (
	"fmt"
)

// iota ia s counter
const (
	a = iota
	b
	c
)
const (
	a2 = iota
	a3
	a4
)

func main() {
	fmt.Printf("%v\n", a)
	fmt.Printf("%v\n", b)
	fmt.Printf("%v\n", c)
	fmt.Printf("%v\n", a2)
	fmt.Printf("%v\n", a3)
	fmt.Printf("%v\n", a4)
}
