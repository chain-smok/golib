package main

import (
	"fmt"
)

func main() {
	for i := 0; i < 5; i++ {
		func() {
			fmt.Println(i)
		}()
	}
	g := greeter{
		greeting: "Hello",
		name:     "Go",
	}
	g.greet()
	fmt.Println("The new name is:", g.name)
	/*d, err := divide(5.0, 0.0)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(d)*/
	//sayMessage("Hello go!")
	/*greeting := "Hello"
	name := "Stacey"
	sayGreeting(&greeting, &name)
	fmt.Println(name)*/
	//s := sum(1, 2, 3, 4, 5)
	//fmt.Println("The sum is ", *s)
}

/*
	func sayMessage(msg string) {
		fmt.Println(msg)
	}
*/
/*func sayGreeting(greeting, name *string) {
	fmt.Println(*greeting, *name)
	*name = "Ted"
	fmt.Println(*name)
}*/
func sum(values ...int) *int {
	fmt.Println(values)
	result := 0
	for _, v := range values {
		result += v
	}
	return &result
}

func divide(a, b float64) (float64, error) {
	if b == 0.0 {
		return 0.0, fmt.Errorf("Can't divide by 0")
	}
	return a / b, nil
}

type greeter struct {
	greeting string
	name     string
}

func (g *greeter) greet() {
	fmt.Println(g.greeting, g.name)
	g.name = ""
}
