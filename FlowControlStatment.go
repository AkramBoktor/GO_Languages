package main

import (
	"fmt"
)

func main() {

	// for loops
	for i := 0; i < 6; i++ {

		fmt.Println(i)
	}
	println("******************************************")
	// while loops
	sum := 1
	for sum < 100 {
		println(sum)
		sum += 1
	}
	println("******************************************")

	// if statment
	// not needed for () parentheses  but {} braces  are required
	for i := 0; i < 10; i++ {
		if i == 3 {
			fmt.Printf("The value of condition if is %v", i)
		}
	}
}
