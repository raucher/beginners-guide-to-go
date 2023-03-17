package main

import (
	"fmt"
	"math/rand"
)

var foo = func() {
	println("Hello World!")
}

func RandInt(valueChanel chan int) {
	randNum := rand.Intn(42)
	fmt.Printf("Random num = %d\n", randNum)
	valueChanel <- randNum
}

func main() {
	valueChannel := make(chan int)

	go RandInt(valueChannel)

	// wait until first value comes from the channel
	value := <-valueChannel

	fmt.Printf("first val in channel is %d\n", value)
}
