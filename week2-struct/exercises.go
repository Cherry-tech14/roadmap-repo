// package main

// import (
// 	"fmt"
// )

// type Car struct {
// 	Brand string
// 	Model string
// 	Year  int
// 	Color string
// }

// func main() {
// 	myCar := Car{
// 		Brand: "Toyota",
// 		Model: "Camry",
// 		Year:  2020,
// 		Color: "Black",
// 	}
// 	fmt.Println(myCar.Brand)
// 	fmt.Println(myCar.Model)
// 	fmt.Println(myCar.Year)
// 	fmt.Println(myCar.Color)
// }

package main

import (
	"fmt"
)

type Student struct {
	Name  string
	Age   int
	Email string
	Score int
}

func main() {
	student1 := Student{
		Name:  "Deborah",
		Age:   20,
		Email: "www.deborah@gmail.com",
		Score: 50,
	}
	student2 := Student{
		Name:  "John",
		Age:   17,
		Email: "www.john23@gmail.com",
		Score: 75,
	}
	fmt.Println(student1.Name)
	fmt.Println(student1.Age)
	fmt.Println(student1.Email)
	fmt.Println(student1.Score)
	fmt.Println(student2.Name)
	fmt.Println(student2.Age)
	fmt.Println(student2.Email)
	fmt.Println(student2.Score)
}
