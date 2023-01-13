package main

import (
	"fmt"
	"reflect"
)

type Animal struct {
	Name   string `required max:"100"`
	origin string
}

// Bird struct has an Animal
// struct,which is composition
type Bird struct {
	Animal
	SpeedKPH float32
	CanFly   bool
}

func main() {
	t := reflect.TypeOf(Animal{})
	b := Bird{
		Animal:   Animal{Name: "Emu", origin: "Australia"},
		SpeedKPH: 30,
		CanFly:   false,
	}
	field, _ := t.FieldByName("Name")
	fmt.Println(b)
	fmt.Println(field.Tag)
}
