package main

import "fmt"

func main() {
	n := average(43, 56, 65, 12, 24, 45)
	fmt.Println(n)
}

func average(sf ...float64) float64 {
	var total float64

	fmt.Println(sf)
	fmt.Printf("%T \n", sf)

	for _, v := range sf {
		total += v
	}

	return total / float64(len(sf))
}
