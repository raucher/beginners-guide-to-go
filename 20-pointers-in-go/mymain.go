package main

import "fmt"

var name = "Max"

type Engineer struct {
	Name string
	Age  int
}

func main() {
	ptr := &name
	// Address
	fmt.Println(ptr) // 0x524f00
	// Value
	fmt.Println(*ptr) // Max

	// Under the hood: var backup string
	backup := *ptr

	*ptr = "Bill"
	fmt.Println(name)   // Bill
	fmt.Println(backup) // Max

	// Under the hood: var engineer *Engineer
	engineer := &Engineer{
		Name: "Max",
		Age:  40,
	}
	fmt.Println(engineer) // &{Max 40}
}
