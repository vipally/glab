package main

import "C"
import "fmt"

//go build -buildmode=c-shared -o exportgo.dll exportgo.go

type ExportType struct {
	A int
	B string
}

//export NewT
func NewT() ExportType {
	return ExportType{}
}

//export PrintBye
func PrintBye() {
	fmt.Println("From DLL: Bye!")
}

//export Sum
func Sum(a int, b int, c string) int {
	return a + b
}

func main() {
	// Need a main function to make CGO compile package as C shared library
}
