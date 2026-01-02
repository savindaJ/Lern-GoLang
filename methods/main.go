package main


import (
	"fmt"
)

func main() {
	person := Person{Name: "Savinda Jayasekara", Age: 20, Email: "savinda@gmail.com"}
	fmt.Println(person.String())
	fmt.Println(person.SayHello())
	person.SayGoodbye()
}

type Person struct {
	Name string
	Age int
	Email string
}

func (p Person) String() string {
	return fmt.Sprintf("Name: %s, Age: %d, Email: %s", p.Name, p.Age, p.Email)
}

func (p Person) SayHello() string {
	return "Hello, " + p.Name
}

func (p Person) SayGoodbye() {
	fmt.Println("Goodbye, " + p.Name)
}