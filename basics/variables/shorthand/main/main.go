package main

import (
	"fmt"

	"github.com/grantsru/udemy/basics/variables/shorthand/vars"
)

func main() {
	a := vars.AInt
	b := vars.BStr
	c := vars.CFlt
	d := vars.DBool

	fmt.Printf("%T \t %v \n", a, a)
	fmt.Printf("%T \t %v \n", b, b)
	fmt.Printf("%T \t %v \n", c, c)
	fmt.Printf("%T \t %v \n", d, d)
}
