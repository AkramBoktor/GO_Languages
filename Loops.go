package main

import (
	"fmt"
)

func main() {
	fmt.Println("For Loops")

	for i := 1; i < 3; i++ {

		fmt.Printf("The value of i %v \n", i)
	}

	fmt.Println("array Loops")

	var arr [5]int
	for i := 0; i < len(arr); i++ {

		arr[i] = i
	}
	fmt.Printf("Array values is %v", arr)
}
