package main

import (
	"fmt"
	"math/cmplx"
)

func main() {
	// v using as list having multiple types
	var (
		check   bool       = true
		number  int        = 2
		complex complex128 = cmplx.Sqrt(-5 + 12i)
	)
	var x, y int = 4, 5
	// T means type and V means Value
	fmt.Printf("Type %T and Value %v \n", check, check)
	fmt.Printf("Type %T and Value %v \n", number, number)
	fmt.Printf("Type %T and Value %v \n", complex, complex)
	fmt.Printf("value x %v \n", x)
	fmt.Printf("value y %v \n", y)

	// declear constant varibale constant can't declear with :=
	const constValue int = 5
	//constValue = 6 * can't assign constValue to variables
	fmt.Println(constValue)
	// knowing type of the variables
	checktype := 5
	fmt.Printf("the type of checktype variable is =  %T \n", checktype)

	// you can used := in the main function only

	variablesInt := 5

	fmt.Printf("The type of variables %T and the values is %v", variablesInt, variablesInt)

}
