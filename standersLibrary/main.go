package main

import (
	"fmt"
	"time"
)

func main() {
	time.Sleep(5 * time.Second)
	now := time.Now()
	fmt.Println(now)
}
