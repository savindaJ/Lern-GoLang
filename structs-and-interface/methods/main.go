package main


import (
	"fmt"
)

func main() {
	person := Person{Name: "John", Age: 30, Email: "john@example.com"}
	fmt.Println(person.String())

	animal := Animal{address: "123 Main St"}
	fmt.Println(animal.String())
}

type Person struct {
	Name string
	Age int
	Email string
}

type Animal struct {
	address string
}

func (a Animal) String() string {
	return fmt.Sprintf("Address: %s", a.address)
}

func (p Person) String() string {
	return fmt.Sprintf("Name: %s, Age: %d, Email: %s", p.Name, p.Age, p.Email)
}