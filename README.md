# GO_Languages
![image](https://user-images.githubusercontent.com/35446384/178409278-0bd06389-92e5-4f9e-8c95-3795dfd21147.png)
https://go.dev/tour/welcome/1
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

# Nil slices
The zero value of a slice is nil.

A nil slice has a length and capacity of 0 and has no underlying array.
![image](https://user-images.githubusercontent.com/35446384/178130538-10e7499e-ec22-47ea-9d15-39e0ce32e28e.png)


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
	
	
# Switch Case
A **switch** statement is a shorter way to write a sequence of if - else statements. It runs the first case whose value is equal to the condition expression

the **break** statement that is needed at the end of each case in those languages is provided automatically in Go. Another important difference is that Go's switch cases need not be constants, and the values involved need not be integers.

![image](https://user-images.githubusercontent.com/35446384/177084768-88676781-29ef-4ae2-b60c-c137c9dcebf9.png)


# Defer
**Stacking defers**
Deferred function calls are pushed onto a stack. When a function returns, its deferred calls are executed in last-in-first-out order.
![image](https://user-images.githubusercontent.com/35446384/177085660-91e68178-abf7-41a9-9b02-1789b364c9a9.png)
**Output** 

![image](https://user-images.githubusercontent.com/35446384/177085721-d31acfc9-0180-4ff2-b0eb-add429e12869.png)



# Pointers
**Go has pointers.** A pointer holds the memory address of a value.

The type *T is a pointer to a T value. Its zero value is nil.

var p *int

The & operator generates a pointer to its operand.

i := 42

p = &i;

![image](https://user-images.githubusercontent.com/35446384/177928665-de512947-7be1-41f2-8046-227b06a52138.png)

# Pointers to structs
Struct fields can be accessed through a struct pointer.

To access the field X of a struct when we have the struct pointer p we could write (*p).X. However, that notation is cumbersome, so the language permits us instead to write just p.X, without the explicit dereference.

![image](https://user-images.githubusercontent.com/35446384/177947741-b999665d-ce77-4256-ad7e-bcfdd58b973b.png)

# Array
The type [n]T is an array of n values of type T.

The expression

var a [10]int
declares a variable a as an array of ten integers.
An array's length is part of its type, so arrays cannot be resized. This seems limiting, but don't worry; Go provides a convenient way of working with arrays.

![image](https://user-images.githubusercontent.com/35446384/177979547-d1ba42a1-6dae-4d04-8ff3-0425c68097f4.png)

**Length of the array by len(arr)**
![image](https://user-images.githubusercontent.com/35446384/177979700-0d5c64f1-45b8-492f-a175-0ed192c5a3a6.png)

# Slice

An array has a fixed size. A slice, on the other hand, is a dynamically-sized, flexible view into the elements of an array. In practice, slices are much more common than arrays.

The type []T is a slice with elements of type T.

A slice is formed by specifying two indices, a low and high bound, separated by a colon:

**a[low : high]**
![image](https://user-images.githubusercontent.com/35446384/178087572-8160a9b7-4cc8-4647-a2bc-e6bfab63fc3d.png)

# Slice are like references to arrays

A slice does not store any data, it just describes a section of an underlying array.

Changing the elements of a slice modifies the corresponding elements of its underlying array.

![image](https://user-images.githubusercontent.com/35446384/178087888-5c78f3ee-bb9d-445d-b397-cb3ab2ad291a.png)

# Slice Literal : 
It's like array literal without length.
![image](https://user-images.githubusercontent.com/35446384/178129463-500de6db-6617-4704-b27a-7e67457596be.png)


# Nil slices
The zero value of a slice is nil.

A nil slice has a length and capacity of 0 and has no underlying array.

![image](https://user-images.githubusercontent.com/35446384/178130586-9e93e737-5c5f-4d28-9f39-0a0ead83c18c.png)

# Create Slice with make

Slices can be created with the built-in make function; this is how you create dynamically-sized arrays.

The make function allocates a zeroed array and returns a slice that refers to that array:

a := make([]int, 5)  // len(a)=5
To specify a capacity, pass a third argument to make:

b := make([]int, 0, 5) // len(b)=0, cap(b)=5
![image](https://user-images.githubusercontent.com/35446384/178131435-7cf02e97-b0e3-4cae-92c4-1227adb35918.png)

# Slice nil

The zero value of a slice is nil.

A nil slice has a length and capacity of 0 and has no underlying array.
![image](https://user-images.githubusercontent.com/35446384/178181093-8aded6f5-3ed9-4211-9676-366dd7b77412.png)
![image](https://user-images.githubusercontent.com/35446384/178181213-10680009-375a-40d7-b50b-10e165e29cbd.png)

# Appending to a slice
It is common to append new elements to a slice, and so Go provides a built-in append function. The documentation of the built-in package describes append.

func append(s []T, vs ...T) []T
The first parameter s of append is a slice of type T, and the rest are T values to append to the slice.

The resulting value of append is a slice containing all the elements of the original slice plus the provided values.

![image](https://user-images.githubusercontent.com/35446384/178181977-450a6582-c163-450d-b93a-a92032dc9894.png)
![image](https://user-images.githubusercontent.com/35446384/178181994-18acbc34-5bed-4307-beac-b971cf734e22.png)


# Range
The range form of the for loop iterates over a slice or map.

When ranging over a slice, two values are returned for each iteration. The first is the index, and the second is a copy of the element at that index.
![image](https://user-images.githubusercontent.com/35446384/178185406-1c773254-49d2-4bf6-a67f-f99a65a4debf.png)

# Map literals
Map literals are like struct literals, but the keys are required.

![image](https://user-images.githubusercontent.com/35446384/178229073-d9bd2e21-57b1-4655-8943-da204e2997a4.png)

**You can create map by using make function**
![image](https://user-images.githubusercontent.com/35446384/178230389-bb642ba3-53b0-4922-89ba-effe0a1d2946.png)

# Map  literal
# type Vertex struct {
	X int
	Y int
}
**map literal**
	var maparrayLiteral = map[string]Vertex{
		"One": {1, 2},
		"Two": {3, 4},
	}
	fmt.Printf("the map using array literal %v \n", maparrayLiteral)

![image](https://user-images.githubusercontent.com/35446384/178246404-6a961b52-acc7-42b1-8e6e-425703ff08bb.png)

# Insert _ Update _ Delete Value from the Map

Insert or update an element in map m:

m[key] = elem
Retrieve an element:

elem = m[key]
Delete an element:

delete(m, key)
Test that a key is present with a two-value assignment:

elem, ok = m[key]
If key is in m, ok is true. If not, ok is false.

If key is not in the map, then elem is the zero value for the map's element type.

Note: If elem or ok have not yet been declared you could use a short declaration form:

elem, ok := m[key]

![image](https://user-images.githubusercontent.com/35446384/178401838-24ca6828-064b-433e-b3fa-ba75b1bf00b5.png)

