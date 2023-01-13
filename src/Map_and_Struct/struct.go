package main

import (
	"fmt"
)

type Doctor struct {
	number     int
	actorname  string
	episodes   []string
	companions []string
}

func main() {
	aDoctor := Doctor{
		3,
		"Jon Pertwee",
		[]string{},
		[]string{
			"Liz Shaw",
			"Jo Grant",
			"Sarah Jane Smith",
		},
	}
	anotherdoctor := aDoctor
	anotherdoctor.actorname = "Tom Baker"
	fmt.Println(aDoctor.actorname)
	fmt.Println(anotherdoctor.actorname)
}
