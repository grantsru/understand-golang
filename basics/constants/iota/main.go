package main

import "fmt"

const (
	// _ - 0
	_ = iota
	// KB - KB value
	KB = 1 << (iota * 10)
	// MB - MB value
	MB = 1 << (iota * 10)
	// GB - GB value
	GB = 1 << (iota * 10)
)

func main() {
	fmt.Printf("%b\t", KB)
	fmt.Printf("%d\n", KB)
	fmt.Printf("%b\t", MB)
	fmt.Printf("%d\n", MB)
	fmt.Printf("%b\t", GB)
	fmt.Printf("%d\n", GB)
}
