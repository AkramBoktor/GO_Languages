package main

import (
	"fmt"
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
func describe(t *T) {
	fmt.Printf("the describe is type %T and value is %v \n", t, t.S)
}

func main() {
	var testStruct = &T{"Test interfaces pointer value"}
	// print value
	fmt.Printf("the value of the pointer of interface is %v \n", testStruct.S)
	// print string
	testStruct.M()
	// print type and value
	describe(testStruct)
}
