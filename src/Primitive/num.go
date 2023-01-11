package main

import (
	"fmt"
)

func main() {
	//var n uint16 = 42
	//fmt.Printf("%v,%T", n, n)
	//a := 10
	//b := 3
	//fmt.Println(a + b)
	//fmt.Println(a - b)
	//fmt.Println(a * b)
	//fmt.Println(a / b)
	//fmt.Println(a % b)
	//fmt.Println(a & b)
	//fmt.Println(a | b)
	//fmt.Println(a ^ b)
	//fmt.Println(a & ^b) //^b=~b
	a := 8
	fmt.Println(a << 3) //2^3 *2^3
	fmt.Println(a >> 3) //2^3 / 2^3
	var n complex128 = complex(5, 12)
	fmt.Println(n)
	fmt.Printf("%v, %T\n", real(n), real(n))
	fmt.Printf("%v, %T\n", imag(n), imag(n))
	s := "this is a string"
	fmt.Printf("%v, %T\n", string(s[2]), s[2])
	var r rune = 'a'
	fmt.Printf("%v, %T\n", r, r)

}
