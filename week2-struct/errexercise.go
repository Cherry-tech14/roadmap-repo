// package main

// import (
// 	"fmt"
// )

// func divide(a int, b int) (int, error) {
// 	if b == 0 {
// 		return 0, fmt.Errorf("cannot divide by zero")

// 	}
// 	return a / b, nil
// }
// func main() {
// 	result, err := divide(10, 0)
// 	if err != nil {
// 		fmt.Println(err)
// 		return
// 	}
// 	fmt.Println(result)
// }

// package main

// import (
// 	"fmt"
// )

// func checkPassword(password string) error {
// 	if len(password) < 8 {
// 		return fmt.Errorf("password too short")
// 	}
// 	return nil
// }
// func main() {
// 	err := checkPassword("10456896")
// 	if err != nil {
// 		fmt.Println(err)
// 		return
// 	}
// 	fmt.Println("password accepted")

// }

// package main

// import (
// 	"fmt"
// )

// func withdraw(balance int, amount int) (int, error) {
// 	if amount > balance {
// 		return 0, fmt.Errorf("insufficient balance")
// 	}

// 	return balance - amount, nil
// }
// func main() {
// 	newBalance, err := withdraw(2000, 3000)
// 	if err != nil {
// 		fmt.Println(err)
// 		return
// 	}
// 	fmt.Println(newBalance)
// }

package main

import (
	"fmt"
)

func login(username string, password string) error {
	if username != "admin" {
		return fmt.Errorf("invalid login")

	}
	if password != "12345" {
		return fmt.Errorf("invalid login")
	}
	return nil
}

func main() {
	err := login("admin", "123")
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println("login successful")
}
