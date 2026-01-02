package main

import (
	"errors"
	"fmt"
)

func main() {
	result, err := divide(10, 0)
	if err != nil {
		fmt.Println("Error: ", err)
	}
	fmt.Println("Result: ", result)

	err = &MyError{Code: 404, Message: "Not Found"}
	fmt.Println(err.Error())

	defer handlePanic()
	panic("panic error")
}

func handlePanic() {
    if r := recover(); r != nil {
        fmt.Println("panic error: ", r)
    }
}

// func main() {
//     defer handlePanic()
//     panic("බරපතල දෝෂයක්!") // මෙතනදී වැඩසටහන නතර නොවී recover වේ
// }

func divide(a, b float64) (float64, error) {
	if b == 0 {
		return 0, errors.New("cannot divide by zero")
	}
	return a / b, nil
}

type MyError struct {
    Code    int
    Message string
}

func (e *MyError) Error() string {
    return fmt.Sprintf("Error %d: %s", e.Code, e.Message)
}