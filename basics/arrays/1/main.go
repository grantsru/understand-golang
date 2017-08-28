package main

import "fmt"

func main() {
	arrayOne()
	arrayTwo()
	arrayThree()
}

func arrayOne() {
	var x [58]string
	fmt.Println(x)
	fmt.Println(len(x))
	fmt.Println(x[42])

	for i := 65; i <= 122; i++ {
		x[i-65] = string(i)
	}

	fmt.Println(x)
	fmt.Println(len(x))
	fmt.Println(x[42])
}

func arrayTwo() {
	var x [256]int
	fmt.Println(len(x))
	fmt.Println(x[42])
	for i := 0; i < 256; i++ {
		x[i] = i
	}

	for i, v := range x {
		fmt.Printf("%v - %T - %b\n", v, v, v)
		if i > 50 {
			break
		}
	}
}

func arrayThree() {
	var x [256]byte
	fmt.Println(len(x))
	fmt.Println(x[0])
	for i := 0; i < 256; i++ {
		x[i] = byte(i)
	}
	for i, v := range x {
		fmt.Printf("%v - %T - %b\n", v, v, v)
		if i < 50 {
			break
		}
	}
}
