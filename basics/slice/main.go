package main

import "fmt"

func one() {
	// This is a slice of type []int - it is read as "slice int"
	mySlice := []int{1, 3, 5, 7, 8, 11}
	fmt.Printf("%T\n", mySlice)
	fmt.Println(mySlice)
}

func two() {
	// This is a slice of type []string - it is read as "slice string"
	mySlice := []string{"a", "b", "c", "d", "e", "f", "g"}
	fmt.Println(mySlice)
	fmt.Println(mySlice[2:4])  // slicing a slice
	fmt.Println(mySlice[2])    // index access
	fmt.Println("myString"[2]) // index access
}

func three() {
	mySlice := make([]int, 0, 5)
	fmt.Println(mySlice)
	fmt.Println(len(mySlice))
	fmt.Println(cap(mySlice))
	fmt.Println("------")

	for i := 0; i < 80; i++ {
		mySlice = append(mySlice, i)
		fmt.Println("Len: ", len(mySlice), "Cap: ", cap(mySlice), "Val: ", mySlice[i])
	}
}

func four() {
	greeting := []string{
		"Good morning",
		"Bonjour",
		"Buenos dias",
		"Bongiorno",
		"Ohayo",
		"Selemat pagi",
		"Gutten morgen"}

	fmt.Println("[1:2] ", greeting[1:2])
	fmt.Println("[:2] ", greeting[:2])
	fmt.Println("[5:] ", greeting[5:])
	fmt.Println("[:] ", greeting[:])
}

func five() {
	greeting := make([]string, 3, 5)
	greeting[0] = "Good morning"
	greeting[1] = "Bonjour"
	greeting[2] = "Buenos dias"
	// greeting[3] = "Suprabadham" // This is an index out of range
	greeting = append(greeting, "Suprabadham")

	fmt.Println(greeting[2])
}

func six() {
	slice1 := []int{1, 2, 3, 4, 5}
	slice2 := []int{6, 7, 8, 9, 10}

	slice3 := append(slice1, slice2...)
	fmt.Println(slice3)
}

func seven() {
	slice1 := []string{"Mon", "Tues"}
	slice2 := []string{"Wed", "Thurs", "Fri"}

	slice1 = append(slice1, slice2...)
	fmt.Println(slice1)

	slice1 = append(slice1[:2], slice1[3:]...)
	fmt.Println(slice1)
}

func eight() {
	records := make([][]string, 0)
	values := []string{"Foster", "Nathan", "100,00", "74.00", "Gomez", "Lisa", "92.00", "96.00"}

	// Student 1
	student1 := make([]string, 4)
	student1 = append(student1, values[:3]...)
	records = append(records, student1)

	// Student 2
	student2 := make([]string, 4)
	student2 = append(student2, values[3:]...)
	records = append(records, student2)

	fmt.Println(records)
}

func main() {
	one()
	two()
	three()
	four()
	five()
	six()
	seven()
	eight()
}
