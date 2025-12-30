// package variable
package main

import (
	"context"
	"errors"
	"fmt"
	"sync"
	"time"
)

type Person struct {
	Name  string
	Age   int
	Email string
}

func initVariable() {
	var a int = 10
	var b int = 20
	const c int = 30
	fmt.Println(a + b + c)
}

func initVariableWithType() {
	var a int = 10
	var b int = 20
	const c int = 30
	fmt.Println(a + b + c)
}

func initVariableWithoutType() {
	a := 10
	b := 20
	c := 30
	fmt.Println(a + b + c)
}

func initVariableWithTypeAndShortDeclaration() {
	a := 10
	b := 20
	c := 30
	fmt.Println(a + b + c)
}

func initVariableWithTypeAndShortDeclarationAndConst() {
	a := 10
	b := 20
	c := 30
	fmt.Println(a + b + c)
}

func initVariableWithTypeAndShortDeclarationAndConstAndFloat() {
	a := 10.0
	b := 20.0
	c := 30.0
	fmt.Println(a + b + c)
}

func StringVariable() {
	a := "Hello, World!"
	fmt.Println(a)
}

func BoolVariable() {
	a := true
	b := false
	fmt.Println(a && b)
}

func ErrorVariable() {
	// a := errors.New("error")
	// fmt.Println(a)
}

func TimeVariable() {
	a := time.Now()
	fmt.Println(a)
}

func DurationVariable() {
	a := time.Duration(10 * time.Second)
	fmt.Println(a)
}

func TimezoneVariable() {
	a := time.Now().Location()
	fmt.Println(a)
}

func LocationVariable() {
	a := time.Now().Location()
	fmt.Println(a)
}

func MapVariable() {
	a := map[string]string{
		"name": "John",
		"age":  "20",
	}
	fmt.Println(a)
}

func SliceVariable() {
	a := []string{"John", "Jane", "Jim"}
	fmt.Println(a)
}

func ArrayVariable() {
	a := [3]string{"John", "Jane", "Jim"}
	fmt.Println(a)
}

func StructVariable() {
	a := struct {
		name string
		age  int
	}{
		name: "John",
		age:  20,
	}
	fmt.Println(a)
}

func InterfaceVariable() {
	a := interface{}(Person{Name: "John", Age: 20, Email: "john@example.com"})
	fmt.Println(a)
}

func PointerVariable() {
	a := &Person{Name: "John", Age: 20, Email: "john@example.com"}
	fmt.Println(a)
}

/*
Channel is a communication mechanism that allows you to send and receive values between goroutines.
It is commonly used to communicate between goroutines.
*/

func ChannelVariable() {
	a := make(chan string, 1) // buffered channel with capacity 1
	a <- "Hello, World!"
	fmt.Println(<-a)
}

/*
WaitGroup is a synchronization primitive that allows you to wait for a group of goroutines to finish executing.
It is commonly used to wait for a group of goroutines to finish before proceeding with the execution of the main function.
*/

func WaitGroupVariable() {
	var wg sync.WaitGroup
	wg.Add(1)
	go func() {
		defer wg.Done()
		fmt.Println("Hello, World!")
	}()
	wg.Wait()
}

/*
Mutex
Mutex is a mutual exclusion lock.
It is used to protect a shared resource from concurrent access.
*/
func MutexVariable() {
	var mu sync.Mutex
	mu.Lock()
	defer mu.Unlock()
	fmt.Println("Hello, World!")
	// mu.Unlock() is handled by defer above
}

/*
Context is a mechanism that allows you to pass values between goroutines.
It is commonly used to pass values between goroutines.
*/

func ContextVariable() {
	ctx := context.Background()
	ctx = context.WithValue(ctx, "HELLO", "WORLD")
	fmt.Println(getContextValue(ctx))
	defer ctx.Done()
}

func getContextValue(ctx context.Context) string {
	value := ctx.Value("HELLO")
	if value == nil {
		return ""
	}
	return value.(string)
}

func ErrorVariableWithNew() {
	a := errors.New("error with new error")
	fmt.Println(a)
}

func main() {
	// initVariable()
	// initVariableWithType()
	// initVariableWithoutType()
	// initVariableWithTypeAndShortDeclaration()
	// initVariableWithTypeAndShortDeclarationAndConst()
	// initVariableWithTypeAndShortDeclarationAndConstAndFloat()
	// StringVariable()
	// BoolVariable()
	// ErrorVariable()
	// TimeVariable()
	// DurationVariable()
	// TimezoneVariable()
	// LocationVariable()
	// MapVariable()
	// SliceVariable()
	// ArrayVariable()
	// StructVariable()
	// InterfaceVariable()
	// PointerVariable()
	// ChannelVariable()
	// WaitGroupVariable()
	// MutexVariable()
	ContextVariable()
}
