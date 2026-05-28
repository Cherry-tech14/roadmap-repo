// package main

// import (
// 	"fmt"
// )

// func main() {
// 	sum := 0
// 	for i := 1; i <= 10; i++ {
// 		sum += i

// 	}
// 	fmt.Println(sum)

// }

// package main

// import "fmt"

// func main() {
// 	count := 0

// 	for i := 1; i <= 10; i++ {
// 		if i%2 == 0 {
// 			fmt.Println(i)
// 			count++
// 		}
// 	}
// 	fmt.Println("total even numbers:", count)

// }

// package main

// import (
// 	"fmt"
// )

// func main() {
// 	a := 15
// 	b := 22
// 	if a > b {
// 		fmt.Println(a, "is larger")
// 	} else {
// 		fmt.Println(b, "is larger")
// 	}
// }

package main

import (
	"fmt"
)

func main() {
	score := 73

	switch {
	case score >= 70:
		fmt.Println("excellent")

	case score >= 69:
		fmt.Println("pass")

	default:
		fmt.Println("fail")

	}
}
