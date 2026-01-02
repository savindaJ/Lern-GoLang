package main

import (
    // "encoding/xml"
    // "fmt"
)

type MyPerson2 struct {
    Name  string `xml:"name"`
    Age   int    `xml:"age"`
    Email string `xml:"email"`
}

// func main() {
//     xmlData := `
//     <MyPerson>
//         <name>John</name>
//         <age>30</age>
//         <email>john@example.com</email>
//     </MyPerson>`

//     var person MyPerson2

//     err := xml.Unmarshal([]byte(xmlData), &person)
    
//     if err != nil {
//         fmt.Println("Error:", err)
//         return
//     }

//     fmt.Println("After unmarshalling:")
//     fmt.Printf("Name: %s, Age: %d, Email: %s\n", person.Name, person.Age, person.Email)
// }