package main

import (
	"fmt"
)

func main() {

	fmt.Println("Hello there all")
	val := true
	fmt.Printf("%v \n", val)
	// the operator you have !  ||  &&
	// true || false -> true
	// true && false -> false
	valcheck := true || false && true
	fmt.Printf("%t  \n", valcheck)
}
