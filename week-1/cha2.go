package main

import (
	"fmt"
)

func main() {
	score := 44

	if score >= 70 {
		fmt.Println("Excellent")
	} else if score >= 50 && score <= 69 {
		fmt.Println("pass")
	} else {
		fmt.Println("fail")
	}
}
