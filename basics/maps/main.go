package main

import "fmt"

func one() {
	m := make(map[string]int)

	m["k1"] = 7
	m["k2"] = 13
	fmt.Println("map:", m)

	v1 := m["k1"]
	fmt.Println("v1: ", v1)
	fmt.Println("len:", len(m))

	delete(m, "k2")
	fmt.Println("map:", m)
}

func two() {
	var greet1 = make(map[string]string)
	greet1["Tim"] = "Good morning."
	greet1["Jenny"] = "Bonjour."

	fmt.Println(greet1)
}

func three() {
	greet := map[int]string{
		0: "Good morning",
		1: "Bonjour",
		2: "Buenos dias",
		3: "Bongiorno",
	}

	if val, exists := greet[2]; exists {
		delete(greet, 2)
		fmt.Println("Val: ", val)
		fmt.Println("Exists: ", exists)
	} else {
		fmt.Println("That value doesn't exist.")
		fmt.Println("Val: ", val)
		fmt.Println("Exists: ", exists)
	}
}

func main() {
	one()
	two()
	three()
}
