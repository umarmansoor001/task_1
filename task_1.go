package task_1

import "fmt"

type Person struct {
	name   string
	age    int
	job    string
	salary int
}

func add_person(name string, age int, job string, salary int) Person {
	var p Person
	p.name = name
	p.age = age
	p.job = job
	p.salary = salary
	return p
}

func Display_person(p Person) {
	fmt.Println("Name: ", p.name)
	fmt.Println("Age: ", p.age)
	fmt.Println("Job: ", p.job)
	fmt.Println("Salary: ", p.salary)
}
