package main

import (
	"fmt"
)

type Employee struct {
	name     string
	age      int
	isRemote bool
}

type Person struct {
	Name string
	age  int
}

func main() {
	employee1 := Employee{
		name:     "Alice",
		age:      32,
		isRemote: true,
	}

	fmt.Println("Employee name:", employee1.name)

	job := struct {
		title  string
		salary int
	}{
		title:  "salary",
		salary: 2000,
	}
	fmt.Println("Job", job)
}
