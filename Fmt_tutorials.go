package main

import (
	"fmt"
)

func main() {
	// to know type - type %t
	fmt.Printf("%T \n", 10)
	// to know the value %v
	fmt.Printf("%v \n", 10)
	// to know the value of bool %t

	fmt.Printf("%t \n", true)
	// integer %b base 2
	// %o base 8
	// %x base 16
	// %d base 10
	fmt.Printf("%d \n", 10)
	fmt.Printf("%o \n", 10)
	fmt.Printf("%x \n", 10)
	fmt.Printf("%d \n", 10)

	// floating point
	// %e scientif notation
	// %f decimal
	// %g large exponent
	fmt.Printf("%e \n", 10.4)
	fmt.Printf("%f \n", 10.4)
	fmt.Printf("%g \n", 10.4)

	// string
	// %s
	// %q
	fmt.Printf("%s \n", "Hello")
	fmt.Printf("%q \n", "Hello")

}
