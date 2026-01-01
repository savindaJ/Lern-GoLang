package main

import (
	"fmt"
)

// Define interfaces
type Writer interface {
	Write(p []byte) (n int, err error)
	writing() string
}

type Reader interface {
	Read(p []byte) (n int, err error)
	reading() string
}

type ReadWriter interface {
	Reader // embedded interface
	Writer // embedded interface
}

// Concrete type that implements Writer
type FileWriter struct {
	filename string
}

func (f FileWriter) Write(p []byte) (n int, err error) {
	fmt.Printf("Writing %d bytes to %s\n", len(p), f.filename)
	return len(p), nil
}

func (f FileWriter) writing() string {
	return "FileWriter is writing"
}

// Concrete type that implements Reader
type FileReader struct {
	filename string
}

func (f FileReader) Read(p []byte) (n int, err error) {
	fmt.Printf("Reading from %s\n", f.filename)
	return 0, nil
}

func (f FileReader) reading() string {
	return "FileReader is reading"
}

// Concrete type that implements BOTH (ReadWriter)
type File struct {
	filename string
}

func (f File) Write(p []byte) (n int, err error) {
	fmt.Printf("File: Writing %d bytes to %s\n", len(p), f.filename)
	return len(p), nil
}

func (f File) writing() string {
	return "File is writing"
}

func (f File) Read(p []byte) (n int, err error) {
	fmt.Printf("File: Reading from %s\n", f.filename)
	return 0, nil
}

func (f File) reading() string {
	return "File is reading"
}

// Function that accepts any Writer
func doWrite(w Writer) {
	fmt.Println(w.writing())
	w.Write([]byte("hello"))
}

// Function that accepts any Reader
func doRead(r Reader) {
	fmt.Println(r.reading())
	r.Read(make([]byte, 100))
}

func main() {
	// Create instances
	fw := FileWriter{filename: "output.txt"}
	fr := FileReader{filename: "input.txt"}
	f := File{filename: "data.txt"}

	fmt.Println("=== Using Writer interface ===")
	doWrite(fw) // FileWriter implements Writer
	doWrite(f)  // File also implements Writer

	fmt.Println("\n=== Using Reader interface ===")
	doRead(fr) // FileReader implements Reader
	doRead(f)  // File also implements Reader

	fmt.Println("\n=== Using ReadWriter interface ===")
	var rw ReadWriter = f // Only File implements both!
	rw.Write([]byte("data"))
	rw.Read(make([]byte, 50))
}