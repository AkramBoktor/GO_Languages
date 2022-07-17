package main

import (
	"fmt"
)

// function check types of interface
func checktype(i interface{}) {

	switch check := i.(type) {
	case string:
		fmt.Printf("the value is %v , with length is %v \n", check, len(check))
	case int:
		fmt.Printf("the value is %v , with multiple is %v \n", check, check*check)
	default:
		fmt.Println("No type is incorrect")

	}
}
func main() {

	// type asserstion
	// make i  equal empty interface reference
	var i interface{} = "Hello"

	//hello
	s := i.(string)
	fmt.Println(s)

	// hello true
	s, ok := i.(string)
	fmt.Println(s, ok)

	// 0  flase
	f, ok := i.(float64)
	fmt.Println(f, ok)

	// f = i.(float64) // panic
	// fmt.Println(f)

	// switch case type asserstion
	checktype(22)
	checktype("Test")
	checktype('c')
}
