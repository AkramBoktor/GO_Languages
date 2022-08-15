package main

import (
	"fmt"
)

func main() {
	a, b := swap("hello", "world")
	fmt.Println(a, b)
	fmt.Println(split(10))
	// function calllent function
	firstfunction(myfunct)

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

func firstfunction(myfunct func(int) int) {
	fmt.Printf("Function callled function %v", myfunct(5))
}
func myfunct(x int) int {
	return x
}
