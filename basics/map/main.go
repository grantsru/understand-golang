package main

import "fmt"

func one() {
	attended := map[string]bool{
		"Ann":  true,
		"Mark": false,
	}

	if attended["Ann"] {
		fmt.Println("Ann was at the meeting")
	}
}

func two() {
	multiplier := 60 * 60
	timezone := map[string]int{
		"UTC": 0 * multiplier,
		"EST": 1 * multiplier,
		"CST": 2 * multiplier,
		"MST": 3 * multiplier,
		"PST": 4 * multiplier,
	}

	offset := timezone["EST"]
	fmt.Println(offset)
}

func main() {
	one()
	two()
}
