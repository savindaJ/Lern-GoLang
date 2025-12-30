package main

import (
	"fmt"
)

func main() {
	/*
		Visual representation:
		Underlying array: [0, 0, 0, _, _]
		                   ^     ^     ^
		                   |     |     |
		                 start  len   cap
	*/

	slice := make([]int, 3, 5)
	fmt.Println(len(slice))
	slice = append(slice, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15, 16, 17, 18, 19, 20)
	fmt.Println(slice)
	// fmt.Println(cap(slice))
	fmt.Println(slice[cap(slice)-1])

	// slice = append(slice, 1, 2, 3, 4, 5)
	// fmt.Println(len(slice))
	// fmt.Println(cap(slice))
	// fmt.Println(slice)
}
