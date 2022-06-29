package main

import (
	"fmt"
	"math/cmplx"
)

func main() {
	var (
		check   bool       = true
		number  int        = 2
		complex complex128 = cmplx.Sqrt(-5 + 12i)
	)
	fmt.Printf("Type %T and Value %v", check, check)
	fmt.Printf("Type %T and Value %v", number, number)
	fmt.Printf("Type %T and Value %v", complex, complex)
}
