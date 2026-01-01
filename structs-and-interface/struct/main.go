package main

import (
	"encoding/xml"
	"fmt"
)

type Person struct {
	Name  string `xml:"name"`
	Age   int    `xml:"age"`
	Email string `xml:"email"`
}

func main() {
	person := Person{Name: "John", Age: 30, Email: "john@example.com"}
	xml.Marshal(person)
	fmt.Println(person)
}
