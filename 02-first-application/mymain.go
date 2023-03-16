package main

func main() {
	var greeting string = "Hello World!"

	const answer int8 = 42

	const K = "Kelvin" // untyped constants

	var C = "Celsius" // untyped variable

	D := "Diameter" // but so is simplier

	println(greeting, C, D)

	var m int16 = 42
	n := m
	n++
	n += 1

	println(greeting, m, n)

	// const PI float32 = 3.14
	const PI = 3.14 // leaving constant untyped
	println(PI)

	var diameter float64 = 20

	println(diameter * PI) // compiler casts untyped constants to type needed
}
