package main

import "fmt"

func main() {
	numbers := [5]int8{1, 2, 3}

	// println(numbers)

	fmt.Println(numbers) // [1 2 3 0 0]

	// slice is something ala List in Go
	mySlice := []int8{1, 2, 3}
	// append returns new slice
	mySlice = append(mySlice, 11, 44) // [1 2 3 11 44]

	fmt.Println(mySlice)
}
