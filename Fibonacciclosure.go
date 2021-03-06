/*

Exercise: Fibonacci closure
Let's have some fun with functions.

Implement a fibonacci function that returns a function
(a closure) that returns successive fibonacci numbers (0, 1, 1, 2, 3, 5, ...).
*/
package main

import "fmt"

// fibonacci is a function that returns
// a function that returns an int.
func fibonacci() func(int) int {
	first := 0
	second := 1
	return func(fib int) int {
		if fib == 0 {
			return 0
		} //To return the first Fibonacci number
		if fib == 1 {
			return 1
		} //To return the second Fibonacci number
		sum := first + second
		second = sum
		first = second
		return sum
	}
}

func main() {
	f := fibonacci()
	for i := 0; i < 10; i++ {
		fmt.Println(f(i))
	}
}
