package main

import (
	"fmt"
)

func main() {
	grades := [3]int{97, 85, 93}
	fmt.Printf("Grades: %v\n", grades)
	grades[0] = 10
	fmt.Printf("Grades: %v\n", grades)
	fmt.Printf("Num of Grades: %v\n", len(grades))
	//2D array
	/*var a [][]int = [][]int{
		0: {1, 0, 0},
		1: {0, 1, 0},
		2: {0, 0, 1},
	}*/
	//b := a
	//b[1] = []int{0, 2, 0}
	//fmt.Println(a)
	//fmt.Println(b)
	a := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	b := a[:]   //alice of all element
	c := a[3:]  //slice from 4 th element to end
	d := a[:6]  //slice first 6 elements
	e := a[3:6] //slice the 4th, 5th, and 6th elements
	a[5] = 42
	fmt.Println(a)
	fmt.Println(b)
	fmt.Println(c)
	fmt.Println(d)
	fmt.Println(e)
	a = append(a, []int{11, 12, 13, 14}...)
	fmt.Println(a)
}
