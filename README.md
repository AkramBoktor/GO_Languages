# GO_Languages
# install
Open the MSI file you downloaded and follow the prompts to install Go.
By default, the installer will install Go to Program Files or Program Files (x86). 
You can change the location as needed. After installing, you will need to close and 
reopen any open command prompts so that changes to the environment made by the installer are reflected at the command prompt.

Verify that you've installed Go.
In Windows, click the Start menu.
In the menu's search box, type cmd, then press the Enter key.
In the Command Prompt window that appears, type the following command:
**$ go version**
Confirm that the command prints the installed version of Go.

# Packages
Programs start running in package main.

This program is using the packages with import paths **"fmt" , "math/rand and math"**.

By convention, the package name is the same as the last element of the import path. For instance, the "math/rand" package comprises files that begin with the statement package rand.

# Exported names
In Go, a name is exported if it begins with a **capital letter**. For example, Pizza is an exported name, as is Pi, which is exported from the math package.

**pizza and pi** do not start with a capital letter, so they are not exported.

When importing a package, you can refer only to its exported names. Any "unexported" names are not accessible from outside the package.

Run the code. Notice the error message.

To fix the error, rename math.pi to math.Pi and try it again.

# Function 
A function can take zero or more arguments.

In this example, add takes two parameters of type int.

Notice that the type comes after the variable name.

# Naked Return Function
A return statement without arguments returns the named return values. This is known as a **"naked"** return.
used in short function
<br/>
package main

import "fmt"
<br/>
func split(sum int) (x, y int) {
	x = sum * 4 / 9
	y = sum - x
	return
}
<br/>
func main() {
	fmt.Println(split(17))
}

# Declares  List of variables 
**Var keyword** <br/>
package main

import "fmt"

var c, python, java bool

func main() {
	var i int
	fmt.Println(i, c, python, java)
}

# Short variable declarations
Inside a function, the **:= short assignment statement can be used in place of a var declaration with implicit type.

Outside a function, every statement begins with a keyword (var, func, and so on) and so the := construct is not available.
<br/>
package main

import "fmt"

func main() {
	var i, j int = 1, 2
	k := 3
	c, python, java := true, false, "no!"

	fmt.Println(i, j, k, c, python, java)
}

# Basic types
bool

string

int  int8  int16  int32  int64
uint uint8 uint16 uint32 uint64 uintptr

byte // alias for uint8

rune // alias for int32
     // represents a Unicode code point

float32 float64

complex64 complex128

**Examples**
<br/>
package main

import (
	"fmt"
	"math/cmplx"
)
<br/>

var (
	ToBe   bool       = false
	MaxInt uint64     = 1<<64 - 1
	z      complex128 = cmplx.Sqrt(-5 + 12i)
)

func main() {
	fmt.Printf("Type: %T Value: %v\n", ToBe, ToBe)
	fmt.Printf("Type: %T Value: %v\n", MaxInt, MaxInt)
	fmt.Printf("Type: %T Value: %v\n", z, z)
}

**Results**



Type: bool Value: false
Type: uint64 Value: 18446744073709551615
Type: complex128 Value: (2+3i)

# Constants
Constants are declared like variables, but with the const keyword.

Constants can be character, string, boolean, or numeric values.

Constants cannot be declared using the **:= syntax.**
// declear constant varibale constant can't declear with :=

const constValue int = 5

//constValue = 6 * can't assign constValue to variables
	
fmt.Println(constValue)

# Flow Control Statement 
# For
The basic **for loop** has three components separated by semicolons:

•	the init statement: executed before the first iteration

•	the condition expression: evaluated before every iteration

•	the post statement: executed at the end of every iteration
The init statement will often be a short variable declaration, and the variables declared there are visible only in the scope of the for statement.

The loop will stop iterating once the boolean condition evaluates to false.

**Note: Unlike other languages like C, Java, or JavaScript there are no parentheses surrounding the three components of the for statement and the braces { } are always required.**

# If
**Go's if statements are like its for loops; the expression need not be surrounded by parentheses ( ) but the braces { } are required.**

# Types

	// to know type - # type %t
	
	fmt.Printf("%T \n", 10)
	
	// to know the value %v
	
	fmt.Printf("%v \n", 10)
	
**bool**
	
	// to know the value of bool %t

	fmt.Printf("%t \n", true)
	
**integer**
	
	%b base 2
	
	// %o base 8
	
	// %x base 16
	
	// %d base 10
	
	fmt.Printf("%d \n", 10)
	
	fmt.Printf("%o \n", 10)
	
	fmt.Printf("%x \n", 10)
	
	fmt.Printf("%d \n", 10)

**floating point**
	
	// %e scientif notation
	
	// %f decimal
	
	// %g large exponent
	
	fmt.Printf("%e \n", 10.4)
	
	fmt.Printf("%f \n", 10.4)
	
	fmt.Printf("%g \n", 10.4)

**string**
	
	// %s
	
	// %q
	
	fmt.Printf("%s \n", "Hello")
	
	fmt.Printf("%q \n", "Hello")
