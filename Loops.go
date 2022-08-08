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

	fmt.Println("Range Element")
	var arrayRange = [5]int{1, 3, 4, 5, 2}
	for i, ele := range arrayRange {

		fmt.Printf("index %d value %d \n", i, ele)
	}
	//anonymous variables

	for _, ele := range arrayRange {

		fmt.Printf(" value %d \n", ele)
	}

	//get duplicated element in the array
	var arrayDuplicated = [7]int{1, 3, 2, 1, 3, 5, 6}
	for i, element := range arrayDuplicated {

		for j := i + 1; j < len(arrayDuplicated); j++ {
			if element == arrayDuplicated[j] {
				fmt.Printf("Duplicated Values is %d \n", element)
			}
		}
	}
}
