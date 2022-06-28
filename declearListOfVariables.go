package main

import (
	"fmt"
)

var a, b, c bool = true, false, true

func main() {
	var i int = 2
	fmt.Println(a, b, c, i)
}

// short decleartion variables
/*
Short variable declarations
Inside a function,
the := short assignment statement can be used in place of a var declaration with implicit type.

Outside a function, every statement begins with
a keyword (var, func, and so on) and so the := construct is not available.

*/
