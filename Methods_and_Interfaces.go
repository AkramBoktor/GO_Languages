package main

import (
	"fmt"
	"math"
)

type Vertex struct {
	x, y float64
}

func Abs(v Vertex) float64 {
	return math.Sqrt(v.x*v.x + v.y*v.y)
}

// create  declare a method on non-struct types, too.
type myInt int

func (digit myInt) checkvalue() int {
	if digit < 0 {
		return 0

	}
	return 1
}

func (v *Vertex) ScalePointer(number int) {
	v.x = v.x * float64(number)
	v.y = v.y * float64(number)
}
func ScalePointer(v *Vertex, number int) {
	v.x = v.x * float64(number)
	v.y = v.y * float64(number)
}
func (v Vertex) Abspointer() int {
	return int(math.Sqrt(v.x*v.x + v.y*v.y))
}
func Abspointer(v Vertex) int {
	return int(math.Sqrt(v.x*v.x + v.y*v.y))
}
func main() {
	// Go language doesn't have classes but it's have method
	v := Vertex{3.0, 4.0}
	fmt.Printf("The multiple sqrt function is %v \n", Abs(v))

	// calll method with non struct types
	fmt.Println("Called Method of non struct types")
	var x myInt = 1
	fmt.Printf("create  declare a method on non-struct types is %v \n", x.checkvalue())

	// declear function with Pointer
	// Notes of you remove * from scale the value will be 5  9 + 16
	vPointer := Vertex{3, 4}
	vPointer.ScalePointer(10)
	fmt.Println("*Pointer here*")
	fmt.Printf("create  declare a method on non-struct types using Pointer is %v \n", vPointer.Abspointer())

	// second way eclear function with Pointer
	// Notes of you remove * from scale the value will be 5  9 + 16
	vPointer2 := Vertex{3, 4}
	ScalePointer(&vPointer2, 10)
	fmt.Println("*Pointer here second way*")
	fmt.Printf("create  declare a method on non-struct types using Pointer second way is %v \n", Abspointer(vPointer2))

}

// another solution way
/*
type Vertex struct {
	X, Y float64
}

func (v Vertex) Abs() float64 {
	return math.Sqrt(v.X*v.X + v.Y*v.Y)
}

func main() {
	v := Vertex{3,4}
	fmt.Println(v.Abs())
}
*/
