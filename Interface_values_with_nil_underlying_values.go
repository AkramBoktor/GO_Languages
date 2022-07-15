/*Interface values with nil underlying values
If the concrete value inside the interface itself is nil, the method will be called with a nil receiver.

Important Hint
In some languages this would trigger a null pointer exception, but in Go it is common to write methods that gracefully handle being called with a nil receiver (as with the method M in this example.)

Note that an interface value that holds a nil concrete value is itself non-nil.
*/
package main

import (
	"fmt"
)

type I interface {
	implemntPrintString()
}
type emptyInterface interface {
}
type T struct {
	s string
}

func (t *T) implemntPrintString() {
	if t == nil {
		fmt.Println("<nil>")
		return
	}
	fmt.Printf("the value of string = %v \n", t.s)
}

func describe(v interface{}) {
	fmt.Printf("the type %T  value = %v \n", v, v)
}
func main() {

	// can't do that 	var checkString I = *T
	var checkString I
	var structValue *T
	checkString = structValue

	// here in the most of language will Exception null
	checkString.implemntPrintString()

	/***********   Empty interface      ***********/
	var checkEmpty interface{}
	// the same is ->	var checkEmpty emptyInterface

	describe(checkEmpty)
	checkEmpty = 3
	describe(checkEmpty)
	checkEmpty = " Hi "
	describe(checkEmpty)

}
