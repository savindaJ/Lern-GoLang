package main

import (
	"fmt"
	"sync"
)

func sayHello(name string) {
	fmt.Println("Hello from Goroutine!")
}

func main() {
	go sayHello("Savinda")
	fmt.Println("Main function is finished")

	// chanel
	ch := make(chan string)
	go func() {
		ch <- "data is saved!"
	}()
	msg := <-ch
	fmt.Println(msg)

	// wait group

	var wg sync.WaitGroup
	wg.Add(2) // add 1 to the wait group
	go func() {
		defer wg.Done() // decrement the wait group
		fmt.Println("data is saved!")
	}()
	go func() {
		defer wg.Done() // decrement the wait group
		fmt.Println("data is saved! 2")
	}()
	wg.Wait() // wait for the wait group to be 0
	fmt.Println("all data is saved!")

	var mu sync.Mutex
	count := 0

	go func() {
		mu.Lock() // lock the mutex
		count++ // increment the count
		mu.Unlock() // unlock the mutex
	}()
	go func() {
		mu.Lock() // lock the mutex
		count++ // increment the count
		mu.Unlock() // unlock the mutex
	}()
	wg.Wait() // wait for the wait group to be 0
	fmt.Println("count: ", count)
}
