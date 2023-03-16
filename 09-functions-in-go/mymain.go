package main

import "fmt"

func main() {
	Foo("Max", 40)

	println(sumSub(4, 5))
}

func Foo(name string, age uint) {
	println(MakeGreeting(name, age))
}

func MakeGreeting(name string, age uint) string {
	return fmt.Sprintf("Hello %s of age %d", name, age)
}

func sumSub(a int, b int) (int, int) {
	return a + b, a - b
}
