package task_1

import "fmt"

type Person struct {
	name   string
	age    int
	job    string
	salary int
}

func Display_person(p Person) {
	fmt.Println("Name: ", p.name)
	fmt.Println("Age: ", p.age)
	fmt.Println("Job: ", p.job)
	fmt.Println("Salary: ", p.salary)
}

