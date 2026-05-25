package main

import (
	"fmt"
)

func main() {
	age := 16

	if age >= 18 {
		fmt.Println("adult")
	} else {
		fmt.Println("minor")
	}
}
