package main

import "fmt"

// Package-level scope is accessible in the entire project
var x = 42

func main() {
	fmt.Println(x)
	foo()
}

func foo() {
	fmt.Println(x)
}
