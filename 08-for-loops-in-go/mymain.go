package main

func main() {
	ages := []int8{11, 21, 54, 42, 86, 9}

	for i := 0; i < len(ages); i++ {
		println(ages[i])
	}

	// 	for i := 0; i < len(ages); print(i) {
	// 		println(ages[i])
	// 		i++
	// 	}

	for index, value := range ages {
		println(index, value)
	}

	for _, value := range ages {
		println(value)
	}

}
