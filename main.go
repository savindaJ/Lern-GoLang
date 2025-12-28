package main

import "fmt"

type Person struct {
	Name string
	Age int
	Email string
}

func (p Person) String() string {
	return fmt.Sprintf("Name: %s, Age: %d, Email: %s", p.Name, p.Age, p.Email)
}

func sayHello(name string) {
	fmt.Println("Hello, " + name)
}

func runAsync() {
    ch := make(chan string)
    go func() {
		sayHello("Savinda")
		person := Person{Name: "Savinda", Age: 20, Email: "savinda@gmail.com"}
		fmt.Println(person)
        ch <- "done"
    }()
    fmt.Println(<-ch)
}



// func main() {
	
// 	runAsync()
// }

