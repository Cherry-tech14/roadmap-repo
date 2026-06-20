package main

import (
	"fmt"
)

func checkAge(age int) error {
	if age < 18 {
		return fmt.Errorf("age is below 18")
	}
	return nil
}
func main() {
	err := checkAge(20)

	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println("Allowed")
}
