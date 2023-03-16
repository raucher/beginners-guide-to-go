package main

import "runtime"

func main() {
	customerHeight := 180
	customerAge := 21

	// if customerAge >= 18 || customerHeight >= 150 {
	// 	println("can access ride")

	// } else if customerHeight >= 120 {
	// 	println("can access children ride")

	// } else {
	// 	println("customer is too small")
	// }

	switch {
	case customerAge >= 18 || customerHeight >= 150:
		println("can access ride")
	case customerHeight >= 120:
		println("can access children rides")
	default:
		println("customer is too small")
	}

	os := runtime.GOOS

	switch os {
	case "darwin":
		println("OS X")
	case "linux":
		println("Tux spotted")
	default:
		println("Something wrong, are you really a developer ?:=")
	}

	switch n := runtime.NumCPU(); n {
	default:
		println(n)
	}
}
