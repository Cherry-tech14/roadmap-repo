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
// 		Name:  "Deborah",
// 		Age:   20,
// 		Email: "www.deborah@gmail.com",
// 		Score: 50,
// 	}
// 	student2 := Student{
// 		Name:  "John",
// 		Age:   17,
// 		Email: "www.john23@gmail.com",
// 		Score: 75,
// 	}
// 	fmt.Println(student1.Name)
// 	fmt.Println(student1.Age)
// 	fmt.Println(student1.Email)
// 	fmt.Println(student1.Score)
// 	fmt.Println(student2.Name)
// 	fmt.Println(student2.Age)
// 	fmt.Println(student2.Email)
// 	fmt.Println(student2.Score)
// }

// package main

// import (
// 	"fmt"
// )

// type Student struct {
// 	Name  string
// 	Age   int
// 	Score int
// 	Grade string
// }

// func main() {
// 	student1 := Student{
// 		Name:  "Mariam",
// 		Age:   22,
// 		Score: 90,
// 		Grade: "A",
// 	}
// 	student2 := Student{
// 		Name:  "John",
// 		Age:   20,
// 		Score: 75,
// 		Grade: "B",
// 	}
// 	student3 := Student{
// 		Name:  "Aisha",
// 		Age:   21,
// 		Score: 85,
// 		Grade: "A",
// 	}
// 	students := []Student{
// 		student1,
// 		student2,
// 		student3,
// 	}
// 	for _, student := range students {
// 		fmt.Println("Name:", student.Name, "Score:", student.Score, "Age:", student.Age, "Grade:", student.Grade)
// 		fmt.Println("Mariam", "scored 90", "and got grade", "A")

// 	}

// }

// package main

// import (
// 	"fmt"
// )

// type Student struct {
// 	Name  string
// 	Age   int
// 	Score int
// 	Grade string
// }

// func (s Student) Passed() {
// 	if s.Score >= 50 {
// 		fmt.Println(s.Name, "passed")

// 	} else {
// 		fmt.Println(s.Name, "failed")
// 	}

// }

// func main() {
// 	student1 := Student{
// 		Name:  "Mariam",
// 		Age:   22,
// 		Score: 90,
// 		Grade: "A",
// 	}

// 	student2 := Student{
// 		Name:  "John",
// 		Age:   20,
// 		Score: 75,
// 		Grade: "B",
// 	}

// 	student3 := Student{
// 		Name:  "Aisha",
// 		Age:   21,
// 		Score: 85,
// 		Grade: "A",
// 	}
// 	students := []Student{
// 		student1,
// 		student2,
// 		student3,
// 	}
// 	for _, student := range students {
// 		student.Passed()
// 	}

// }

package main

import (
	"fmt"
)

type Car struct {
	Brand string
	Model string
	Speed int
}

func (c Car) Drive() {
	if c.Speed > 100 {
		fmt.Println(c.Brand, "is driving fast")
	} else {
		fmt.Println(c.Brand, "is driving slowly")
	}
}
func main() {
	car1 := Car{
		Brand: "Toyota",
		Model: "Camry",
		Speed: 180,
	}

	car2 := Car{
		Brand: "Honda",
		Model: "Civic",
		Speed: 80,
	}
	car1.Drive()
	car2.Drive()
}
