package main

import "fmt"

type Printable interface {
	Print()
}

type Person struct {
	Name string
	Age  uint8
}

// func (p *Person) Print() {
// 	fmt.Printf("%+v\n", p)
// }

func (p Person) Print() {
	fmt.Printf("%+v\n", p)
}

func PrintInfo(p Printable) {
	p.Print()
}

func main() {
	// me := &Person{Name: "Max", Age: 40} // if use a pointer-receiver, passed param should be pointer too
	me := Person{Name: "Max", Age: 40} // if use a pointer-receiver, passed param should be pointer too
	PrintInfo(me)
}
