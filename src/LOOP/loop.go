package main

import (
	"fmt"
)

func main() {
	/*for i, j := 0, 0; i < 5; i, j = i+1, j+2 {
		fmt.Println(i, j)
		if i%2 == 0 {
			i /= 2
		} else {
			i = 2*i + 1
		}
	}*/
	/*i := 0
	for {
		fmt.Println(i)
		i++
		if i == 5 {
			break
		}
	}*/
	/*for i := 1; i <= 3; i++ {
		for j := 1; j <= 3; j++ {
			fmt.Println(i * j)
		}

	}*/
	s := []int{1, 2, 3}

	//_:we don't care this key,we only want to value
	for _, v := range s {
		fmt.Println(v)
	}
}
