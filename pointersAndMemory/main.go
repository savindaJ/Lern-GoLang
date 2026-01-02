package main

import "fmt"

func main() {
	x := 10
	p := &x         // x's address is assigned to p
	fmt.Println(p)  // Output: 0xc000012088 (Memory Address)
	fmt.Println(*p) // Output: 10 (Value at that address)

	name := "John"
	updateName(&name)
	fmt.Println(name)
}

func updateName(n *string) {
    *n = "Kamal"
}
