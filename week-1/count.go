// package main

// import (
// 	"fmt"
// )

// func main() {
// 	words := []string{
// 		"go",
// 		"is",
// 		"fun",
// 		"go",
// 		"is",
// 		"fast",
// 	}
// 	count := make(map[string]int)
// 	for _, w := range words {
// 		count[w]++
// 	}
// 	for index, value := range count {
// 		if value > 1 {
// 			fmt.Println(index, ":", value)

// 		}

// 	}
// }

package main

import (
	"fmt"
)

func main() {
	students := map[string]int{
		"Mariam": 90,
		"John":   75,
		"Aisha":  85,
	}
	for name, score := range students {

		fmt.Println(name, ":", score)
	}
	highestScore := 0
	highestStudent := ""

	for name, score := range students {
		if score > highestScore {
			highestScore = score
			highestStudent = name
		}
	}

	fmt.Println("Highest score:")
	fmt.Println(highestStudent, ":", highestScore)
}
