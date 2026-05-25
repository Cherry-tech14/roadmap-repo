package main

import (
	"fmt"
)

func main() {
	day := 1

	switch day {
	case 1:
		fmt.Println("monday")

	case 2:
		fmt.Println("tuesday")
	case 3:
		fmt.Println("wednesday")
	default:
		fmt.Println("unknown day")
	}
}
