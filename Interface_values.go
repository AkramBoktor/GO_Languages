package main

import (
	"fmt"
	"math"
)

// make interface here
type I interface {
	M()
}
type T struct {
	S string
}

func (t *T) M() {
	fmt.Println(t.S)
}
func describe(t I) {
	fmt.Printf("the describe is type %T and value is %v \n", t, t)
}

// using another type
type F float64

func (f F) M() {
	fmt.Println(f)
}

func main() {
	var testStruct = &T{"Test interfaces pointer value"}
	// print value
	fmt.Printf("the value of the pointer of interface is %v \n", testStruct.S)
	// print string
	testStruct.M()
	// print type and value
	describe(testStruct)

	// another type of float
	var testFloat = F(math.Pi)
	describe(testFloat)
	testFloat.M()
}
