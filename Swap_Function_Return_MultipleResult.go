package main

import (
	"fmt"
)

func main() {
	a, b := swap("hello", "world")
	fmt.Println(a, b)
	fmt.Println(split(10))
}

func swap(x, y string) (string, string) {
	return y, x
}

// called naked function
func split(num int) (x, y int) {
	x = num + 10
	y = num - 10
	return
}
