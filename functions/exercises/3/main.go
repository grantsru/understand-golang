package main

import "fmt"

func main() {
	n := greatest(1, 2, 3, 4, 5, 6, 7, 8, 9, 10)
	fmt.Println(n)
}

func greatest(values ...int) int {
	var largest int
	for _, v := range values {
		if v > largest {
			largest = v
		}
	}

	return largest
}
