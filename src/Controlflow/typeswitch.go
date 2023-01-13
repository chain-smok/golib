package main

import (
	"fmt"
)

func main() {
	var i interface{} = [2]int{12, 34}
	switch i.(type) {
	case int:
		fmt.Println("i is an int")
	case float64:
		fmt.Println("i is a float64")
	case string:
		fmt.Println("i is string")
	case [2]int:
		fmt.Println("i is [2]int")
	default:
		fmt.Println("i is another type")

	}
}
