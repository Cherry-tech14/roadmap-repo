package main

import (
	"fmt"
)

func main() {
	grade := "A"

	switch grade {
	case "A":
		fmt.Println("Excellent")
	case "B":
		fmt.Println("Good")
	case "C":
		fmt.Println("Average")
	default:
		fmt.Println("Poor")

	}
}
