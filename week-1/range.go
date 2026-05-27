package main

import "fmt"

// func main() {
// 	str := "Hello"

// 	for index, char := range str {
// 		fmt.Println(index, string(char))
// 	}

// }

func main() {
	str := "hello"
	for _, char := range str {
		fmt.Println(string(char))
	}
}
