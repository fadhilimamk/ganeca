package main

import "github.com/fadhilimamk/ganeca/src/students"
import "fmt"

func main() {
	var data []students.Student

	data = students.GetAllStudents()

	for _, student := range data {
		fmt.Println(student.ToString())
	}

	fmt.Println("finish")

}
