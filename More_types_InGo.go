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
	fmt.Printf("the array after modified value is %v \n", arraySlice)

	// slice literal with value
	//A slice literal is like an array literal without the length.
	// This is an array literal:
	arrayLiteral := []bool{true, false, true}
	fmt.Printf("Print array Literal value %v \n", arrayLiteral)

	s := []int{2, 3, 5, 7, 11, 13}

	s = s[1:4]
	fmt.Println(s)

	s = s[:2]

	fmt.Println(s)

	s = s[1:]
	fmt.Println(s)

	// nil slice
	// the zeo capacity and length is nil
	var arrayNilSlice []int
	fmt.Println("Array nil slice")
	fmt.Println(arrayNilSlice, len(arrayNilSlice), cap(arrayNilSlice))

	// make function you can create slice with make function to make
	// it's as dynamic make(type,size,capacity)
	arraySliceMake := make([]int, 3)
	fmt.Printf("the array created with make len=%d cpacity=%d array=%v \n",
		len(arraySliceMake), cap(arraySliceMake), arraySliceMake)

	// Slice nil
	// when create an array it's will be null
	var arrayNill []int
	fmt.Printf("nill array value %d len=%d cap=%d \n",
		arrayNill, len(arrayNill), cap(arrayNill))

	//apending a new element in the slice
	arrayAppending := []int{1, 2, 3, 4, 5}
	// apend new element
	arrayAppending = append(arrayAppending, 6)
	fmt.Printf("Append new element in the array %d \n", arrayAppending)
	arrayAppending = append(arrayAppending, 7, 8, 9)
	fmt.Printf("Append new elements more than one in the array %d \n", arrayAppending)

	// Range
	/*The range form of the for loop iterates over a slice or map.
	When ranging over a slice, two values are returned for each iteration.
	The first is the index,
	and the second is a copy of the element at that index.*/
	rangePowArray := []int{1, 2, 3, 4, 5, 6, 7}
	for i, v := range rangePowArray {
		fmt.Printf("the index value =%d and the value element = %d \n", i, v)
	}

	//map
	m := map[string]int{
		"Akram": 1,
	}
	fmt.Printf("The map of declear new list = %v \n", m)
	// map for intialize
	var arrayMap = make(map[int]int)
	arrayMap[1] = 1
	fmt.Printf("The map of declear new list  by using make = %v \n", arrayMap)
	// map literal
	var maparrayLiteral = map[string]Vertex{
		"One": {1, 2},
		"Two": {3, 4},
	}
	fmt.Printf("the map using array literal %v \n", maparrayLiteral)
	//Insert _ Update _ Delete Value from the Map
	//retrieve value of map
	var firstValueOfMap = maparrayLiteral["One"]
	fmt.Printf("Retrieve the value of map %d \n", firstValueOfMap)
	// Delete the map is by deleteing function by map name and key
	delete(maparrayLiteral, "One")
	fmt.Printf("map after delete an elements from it %v \n", maparrayLiteral)
}
