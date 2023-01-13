package main

import (
	"fmt"
	"math"
)

func main() {
	myNum := 0.123456789
	if math.Abs(myNum/math.Pow(math.Sqrt(myNum), 2)-1) < 0.01 {
		fmt.Println("These are the same")
	} else {
		fmt.Println("These are different")
	}
	statePopulations := make(map[string]int)
	statePopulations = map[string]int{
		"Taipei":     1000000,
		"New Taipei": 1233444,
		"Taichung":   2344578,
		"Huanlin":    3457980,
		"America":    12444553,
		"Russia":     1278909,
	}
	if pop, ok := statePopulations["America"]; ok {
		fmt.Println(pop)
	}

	if false {
		fmt.Println("mom fuck")
	}
}
