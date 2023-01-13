package main

import (
	"fmt"
)

func main() {
	//var a int = 42
	//var b *int = &a
	//fmt.Println(&a, *b)
	var ms *myStruct
	fmt.Println(ms)
	ms = new(myStruct)
	ms.foo = 42
	fmt.Println(ms.foo)

}

type myStruct struct {
	foo int
}
