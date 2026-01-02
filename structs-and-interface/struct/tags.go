package main

import (
	"encoding/xml"
	"fmt"
)

type MyPerson struct {
	Name  string `xml:"name"`
	Age   int    `xml:"age"`
	Email string `xml:"email"`
}

func (p MyPerson) String() string {
	return fmt.Sprintf("Name: %s, Age: %d, Email: %s", p.Name, p.Age, p.Email)
}

func main() {
	person := MyPerson{Name: "John", Age: 30, Email: "john@example.com"}
	fmt.Println(person.String())
	fmt.Println(xml.Marshal(person))

	data, _ := xml.MarshalIndent(person, "", "    ")
	fmt.Println(string(data))
}
