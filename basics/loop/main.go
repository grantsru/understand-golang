package main

import "fmt"

func main() {
	// simple()
	// nested()
	// adv()
	remainder()
}

func simple() {
	for i := 0; i <= 100; i++ {
		fmt.Println(i)
	}
}

func nested() {
	for i := 0; i <= 100; i++ {
		for j := 0; j <= 100; j++ {
			fmt.Println(i, j)
		}
	}
}

func adv() {
	i := 0
	for {
		fmt.Println(i)
		if i >= 10 {
			break
		}
		i++
	}
}

func remainder() {
	i := 0
	for {
		i++
		if i%2 == 0 {
			continue
		}
		fmt.Println(i)
		if i >= 50 {
			break
		}
	}
}
