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

}
