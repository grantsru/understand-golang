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

func three() {
	greeting := map[int]string{
		0: "Hello",
		1: "Bonjour",
		2: "Ciao",
		3: "Buenos dias",
		4: "Bongiorno",
	}

	for key, val := range greeting {
		fmt.Println(key, " - ", val)
	}
}

func main() {
	one()
	two()
	three()
}
