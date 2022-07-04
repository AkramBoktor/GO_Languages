package main

import (
	"fmt"
	"math"
	"runtime"
)

func main() {

	// for loops
	for i := 0; i < 6; i++ {

		fmt.Println(i)
	}
	println("******************************************")
	// while loops
	sum := 1
	for sum < 100 {
		println(sum)
		sum += 1
	}
	println("******************************************")

	// if statment
	// not needed for () parentheses  but {} braces  are required
	for i := 0; i < 10; i++ {
		if i == 3 {
			fmt.Printf("The value of condition if is %v", i)
		}
	}

	// if statment shorcut
	fmt.Printf("shortcut of using if %v \n", Pow(3, 3, 20))
	fmt.Printf("shortcut of using if %v \n", Pow(3, 2, 10))

	// switch case
	/*
	 the break statement that is needed at the end of
	 each case in those languages is provided automatically in Go.
	 Another important difference is that Go's switch cases need not be constants,
	 and the values involved need not be integers.
	*/
	switch os := runtime.GOOS; os {

	case "darwin":
		fmt.Println("OS X.")
	case "linux":
		fmt.Println("Linux")
	default:
		// freebsd, openbsd,
		// plan9, windows...
		fmt.Printf("%s.\n", os)
	}

	// Stacking defers

	/*
		Stacking defers
		Deferred function calls are pushed onto a stack.
		When a function returns, its deferred calls are executed in last-in-first-out order.
	*/
	fmt.Println("First Print Counting")
	//because it's like stack so it's will be last in first out
	for count := 0; count < 5; count++ {
		defer fmt.Println(count)
	}

	fmt.Println("Last Print")
}

func Pow(x float64, y float64, limit float64) float64 {
	if v := math.Pow(x, y); v < limit {
		return v
	}
	return limit
}
