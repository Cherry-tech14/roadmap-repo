package main

import "fmt"

type Student struct {
	Name  string
	Age   int
	Score int
	Grade string
}

func (s Student) Display() {
	fmt.Println(s.Name, "scored", "and got grade", s.Grade)
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
		Score: 75,
		Grade: "B",
	}

	student3 := Student{
		Name:  "Aisha",
		Age:   21,
		Score: 85,
		Grade: "A",
	}
	students := []Student{
		student1,
		student2,
		student3,
	}
	for _, student := range students {
		student.Display()
	}

}
