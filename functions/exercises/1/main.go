package main

import "fmt"

func main() {
	fmt.Println(half(1))
	fmt.Println(half(2))
}

func half(i int) (int, bool) {
	divided := i / 2
	boolean := (divided % 2) == 0

	return divided, boolean
}
