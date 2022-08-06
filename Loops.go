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
	fmt.Printf("Array values is %v \n", arr)

	fmt.Println("2D array Loops")

	arr2d := [2][3]int{{1, 3, 4}, {5, 6, 7}}

	fmt.Printf("Array 2d values is %v \n", arr2d[0][1])
}
