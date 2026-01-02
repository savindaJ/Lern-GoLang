package main

import (
	"fmt"
)

type Writer interface {
	Write(p []byte) (n int, err error)
	writing() string
}

type Reader interface {
	Read(p []byte) (n int, err error)
	reading() string
}

// Concrete type that implements Writer
type MyWriter struct{}
type MyReader struct{}

// reading implements [Reader].
func (r MyReader) reading() string {
	return "Hello, World!"
}

func (w MyWriter) Write(p []byte) (n int, err error) {
	fmt.Printf("Writing: %s\n", string(p))
	return len(p), nil
}

func (w MyWriter) writing() string {
	return "Hello, World!"
}

func (r MyReader) Read(p []byte) (n int, err error) {
	fmt.Printf("Reading: %s\n", string(p))
	return len(p), nil
}

// func main() {
// 	var writer Writer = MyWriter{}
// 	fmt.Println(writer.writing())
// 	writer.Write([]byte("test data"))

// 	var reader Reader = MyReader{}
// 	fmt.Println(reader.reading())
// 	reader.Read([]byte("Hello, World!"))
// }
