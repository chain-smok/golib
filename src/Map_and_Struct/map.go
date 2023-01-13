package main

import (
	"fmt"
)

func main() {
	statePopulations := make(map[string]int)
	statePopulations = map[string]int{
		"Taipei":     1000000,
		"New Taipei": 1233444,
		"Taichung":   2344578,
		"Huanlin":    3457980,
		"America":    12444553,
		"Russia":     1278909,
	}
	statePopulations["Georgia"] = 10310371
	delete(statePopulations, "Georgia")
	pop, ok := statePopulations["America"]
	fmt.Println(statePopulations["Taichung"])
	fmt.Println(statePopulations["Georgia"])
	fmt.Println(pop, ok)
	fmt.Println(len(statePopulations))
}
