package main

import (
	"fmt"
)

func main() {
	number := 0

	if number > 0 {
		fmt.Println("positive")
	} else if number < 0 {
		fmt.Println("negative")

	} else {
		fmt.Println("Zero")
	}

}
