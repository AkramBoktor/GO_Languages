package main

import (
	"fmt"
)

type Vertex struct {
	X int
	Y int
}

func main() {

	// pointer
	i, j := 10, 20

	// make p having referenec for address i
	p := &i
	fmt.Printf("The type (i) %T and the valu %v value of pinter %v \n", p, i, p)
	// change the value of address i to a value
	*p = 25
	fmt.Printf("The type (i) %T and the valu %v value of pinter %v \n", p, i, p)

	p = &j
	fmt.Printf("The type (j) %T and the valu %v value of pinter %v \n", p, j, p)
	*p = *p + 25
	fmt.Printf("The type (j) %T and the valu %v value of pinter %v \n", p, j, p)
	// struct
	//A struct is a collection of fields.
	fmt.Println(Vertex{1, 2})
	var v = Vertex{2, 4}
	fmt.Printf("the struct value are access by . %v \n", v.X)
	// pointer to struct
	pointerTOStruct := &Vertex{5, 6}
	pointerTOStruct.X = 1
	fmt.Printf("the struct pointer value are access by . %v \n", pointerTOStruct.X)
	// return The special prefix & returns a pointer to the struct value.

	fmt.Println(pointerTOStruct)
	println("*******************************************")
	// Arrays
	var arrayInt [10]int
	for i := 0; i < len(arrayInt); i++ {
		arrayInt[i] = i
		fmt.Printf("The element %v is equal %v \n", i, arrayInt[i])
	}

	// function of the array Slice [low:high]
	//0,1,2,3,4
	arraySlice := [5]int{1, 2, 3, 4, 5}
	var prime []int = arraySlice[1:4]
	// 2,3,4
	fmt.Printf("the array using slice %v \n", prime)
	// slice it's work as reference type when you change the value
	// it's will change of the main value working as call by reference
	prime[0] = 3
	// 1,3,3,4,5
	fmt.Printf("the array after modified value is %v", arraySlice)

}
