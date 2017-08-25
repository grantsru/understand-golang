package main

import "fmt"

func main() {
	halved := func(i int) {
		divided := i / 2
		boolean := (divided % 2) == 0

		fmt.Println(divided, boolean)
	}
	halved(8)
}
