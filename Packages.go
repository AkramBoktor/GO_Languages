package main

import (
	"fmt"
	"math"
	"math/rand"
)

func main() {
	fmt.Println("My favorite number is", rand.Intn(100))
	fmt.Println((rand.Float64()*5)+5, ",")
	fmt.Println("Number Sqrt ", math.Sqrt(9))
	fmt.Printf("Now you have %g problems.\n", math.Sqrt(7))
}
