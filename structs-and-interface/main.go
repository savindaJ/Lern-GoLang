package main


import (
	"fmt"
)

type Person struct {
	Name string
	Age int
	Email string
}

type Employee struct {
	Person
	Salary int
	employeeID string
}

func main() {
	person := Person{Name: "John", Age: 30, Email: "john@example.com"}
	employee := Employee{Person: person, Salary: 100000, employeeID: "123456"}
	fmt.Println(employee)
}