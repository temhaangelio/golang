package main

import (
	"fmt"
)

func main() {
	var city string = "Bursa"
	if city == "Bursa" {
		fmt.Println("Hello ", city)
	} else {
		fmt.Println("I do not know you!")
	}
}
