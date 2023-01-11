package main

import (
	"fmt"
	"strconv"
)

func main() {
	//var j float32 = 27
	//k := 999
	//fmt.Printf("%v, %T", j, k)
	var i int = 12
	fmt.Println(i)

	var k string
	k = strconv.Itoa(i)
	fmt.Printf("%v ,%T\n", k, k)
}
