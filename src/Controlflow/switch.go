package main

import (
	"fmt"
)

func main() {
	i := 10
	//break has actually implied
	switch {
	case i <= 10:
		fmt.Println("one")
		fallthrough
	case i >= 20:
		fmt.Println("two")
	default:
		fmt.Println("not one or two")
	}
}
