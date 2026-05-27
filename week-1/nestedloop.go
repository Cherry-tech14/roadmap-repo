package main

import (
	"fmt"
)

// func main() {
// 	for i := 1; i <= 3; i++ {
// 		for j := 1; j <= 2; j++ {
// 			fmt.Println("i =", i, "j =", j)
// 		}
// 	}
// }

// func main() {
// 	for i := 1; i <= 5; i++ {
// 		fmt.Print("*")
// 	}
// }

func main() {
	for i := 1; i <= 5; i++ {
		for j := 1; j <= i; j++ {
			fmt.Print("*")
		}
		fmt.Println()
	}
}
