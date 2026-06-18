// package main

// import (
// 	"fmt"
// )

// type Student struct {
// 	Name  string
// 	Age   int
// 	Score int
// }

// func main() {
// 	student1 := Student{
// 		Name:  "Mariam",
// 		Age:   22,
// 		Score: 90,
// 	}

// 	fmt.Println(student1.Name)
// 	fmt.Println(student1.Age)
// 	fmt.Println(student1.Score)
// }

// Structs with slices:

// package main

// import (
// 	"fmt"
// )

// type Student struct {
// 	Name  string
// 	Age   int
// 	Email string
// 	Score int
// }

// func main() {
// 	student1 := Student{
// 		Name:  "Christy",
// 		Age:   67,
// 		Email: "www.Christy@gmail.com",
// 		Score: 50,
// 	}
// 	student2 := Student{
// 		Name:  "Joy",
// 		Age:   24,
// 		Email: "www.joy2345@gmail.com",
// 		Score: 45,
// 	}

// 	students := []Student{
// 		student1,
// 		student2,
// 	}
// 	for _, student := range students {
// 		fmt.Println(student.Name, student.Age, student.Email, student.Score)
// 	}
// }

package main

import "fmt"

type Student struct {
	Name  string
	Age   int
	Score int
	Grade string
}

func (s Student) Passed() {

	if s.Score >= 50 {
		fmt.Println(s.Name, "passed")
	} else {
		fmt.Println(s.Name, "failed")
	}

}

func main() {

	student1 := Student{
		Name:  "Mariam",
		Age:   22,
		Score: 90,
		Grade: "A",
	}

	student2 := Student{
		Name:  "John",
		Age:   20,
		Score: 40,
		Grade: "F",
	}

	student1.Passed()
	student2.Passed()
}
