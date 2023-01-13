package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

// Defer:LIFO
func main() {
	res, err := http.Get("http://www.google.com/robots.txt")
	//nil is default,if function failed,
	//it will return not default value to error
	if err != nil {
		log.Fatal(err)
	}
	defer res.Body.Close()
	robots, err := ioutil.ReadAll(res.Body)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("%s", robots)
}
