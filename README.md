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

# Methods continued
You can declare a method on non-struct types, too.
In this example we see a numeric type MyFloat with an Abs method.
You can only declare a method with a receiver whose type is defined in the same package as the method. You cannot declare a method with a receiver whose type is defined in another package (which includes the built-in types such as int)

![image](https://user-images.githubusercontent.com/35446384/178640361-25b98cb7-29a4-4028-a906-dda4d80c91ce.png)

# Pointer receivers
You can declare methods with pointer receivers.
This means the receiver type has the literal syntax *T for some type T. (Also, T cannot itself be a pointer such as *int.)
For example, the Scale method here is defined on *Vertex.
Methods with pointer receivers can modify the value to which the receiver points (as Scale does here). Since methods often need to modify their receiver, pointer receivers are more common than value receivers.
Try removing the * from the declaration of the Scale function on line 16 and observe how the program's behavior changes.

![image](https://user-images.githubusercontent.com/35446384/178640400-ad50f399-98da-4d7c-892f-6fa10ff38ace.png)


# Methods and pointer indirection

Comparing the previous two programs, you might notice that functions with a pointer argument must take a pointer:

var v Vertex

ScaleFunc(v, 5)  // Compile error!

ScaleFunc(&v, 5) // OK

while methods with pointer receivers take either a value or a pointer as the receiver when they are called:

var v Vertex

v.Scale(5)  // OK

p := &v

p.Scale(10) // OK

For the statement v.Scale(5), even though v is a value and not a pointer, the method with the pointer receiver is called automatically. That is, as a convenience, Go 

interprets the statement v.Scale(5) as (&v).Scale(5) since the Scale method has a pointer receiver.

![image](https://user-images.githubusercontent.com/35446384/178645092-7113406a-82af-4a52-901c-09c0b63487d5.png)

# Methods and pointer indirection (2)

The equivalent thing happens in the reverse direction.

Functions that take a value argument must take a value of that specific type:

var v Vertex

fmt.Println(AbsFunc(v))  // OK

fmt.Println(AbsFunc(&v)) // Compile error!

while methods with value receivers take either a value or a pointer as the receiver when they are called:

var v Vertex

fmt.Println(v.Abs()) // OK

p := &v

fmt.Println(p.Abs()) // OK

In this case, the method call p.Abs() is interpreted as (*p).Abs().

![image](https://user-images.githubusercontent.com/35446384/178645141-7572374e-4aee-4100-b6c6-19395c73d68e.png)



# Choosing a value or pointer receiver
There are two reasons to use a pointer receiver.

The first is so that the method can modify the value that its receiver points to.

The second is to avoid copying the value on each method call. This can be more efficient if the receiver is a large struct, for example.

In this example, both Scale and Abs are with receiver type *Vertex, even though the Abs method needn't modify its receiver.

In general, all methods on a given type should have either value or pointer receivers, but not a mixture of both. (We'll see why over the next few pages.)

![image](https://user-images.githubusercontent.com/35446384/178648517-7ee72b97-7e84-4c7c-b7c8-1745f4d97f50.png)

# Interfaces are implemented implicitly
A type implements an interface by implementing its methods. There is no explicit declaration of intent, no "implements" keyword.

Implicit interfaces decouple the definition of an interface from its implementation, which could then appear in any package without prearrangement.

![image](https://user-images.githubusercontent.com/35446384/178891712-58a8b45e-6b1b-4491-bbac-f689e28df65a.png)


# Interface values
Under the hood, interface values can be thought of as a tuple of a value and a concrete type:

(value, type)
An interface value holds a value of a specific underlying concrete type.

Calling a method on an interface value executes the method of the same name on its underlying type.
![image](https://user-images.githubusercontent.com/35446384/178895693-182727c6-5c7b-4e23-91fb-00e7b131f2d5.png)

![image](https://user-images.githubusercontent.com/35446384/178898158-1519b546-716f-4ad1-a7d7-edd25404f378.png)

# Interface values with nil underlying values
If the concrete value inside the interface itself is nil, the method will be called with a nil receiver.

In some languages this would trigger a null pointer exception, but in Go it is common to write methods that gracefully handle being called with a nil receiver (as with the method M in this example.)

Note that an interface value that holds a nil concrete value is itself non-nil.

![image](https://user-images.githubusercontent.com/35446384/178919895-29dbec37-400e-42be-b85b-2808a5089feb.png)

# Nil interface values

A nil interface value holds neither value nor concrete type.

Calling a method on a nil interface is a run-time error because there is no type inside the interface tuple to indicate which concrete method to call.

![image](https://user-images.githubusercontent.com/35446384/178930173-937fb231-1497-4e3d-8f0d-f2695d95e6b6.png)

# The empty interface
The interface type that specifies zero methods is known as the empty interface:

interface{}

An empty interface may hold values of any type. (Every type implements at least zero methods.)

Empty interfaces are used by code that handles values of unknown type. For example, fmt.Print takes any number of arguments of type interface{}.

![image](https://user-images.githubusercontent.com/35446384/179142642-1ba9d6a7-afa7-4125-92e8-639194b428dc.png)


Type assertions
A type assertion provides access to an interface value's underlying concrete value.

t := i.(T)

This statement asserts that the interface value i holds the concrete type T and assigns the underlying T value to the variable t.

If i does not hold a T, the statement will trigger a panic.

To test whether an interface value holds a specific type, a type assertion can return two values: the underlying value and a boolean value that reports whether the assertion succeeded.

t, ok := i.(T)

If i holds a T, then t will be the underlying value and ok will be true.

If not, ok will be false and t will be the zero value of type T, and no panic occurs.

Note the similarity between this syntax and that of reading from a map.

![image](https://user-images.githubusercontent.com/35446384/179388739-2e159c2d-3a3f-4948-9784-6c9cb8da3c58.png)

Type switches
A type switch is a construct that permits several type assertions in series.
A type switch is like a regular switch statement, but the cases in a type switch specify types (not values), and those values are compared against the type of the value held by the given interface value.
switch v := i.(type) {
case T:
    // here v has type T
case S:
    // here v has type S
default:
    // no match; here v has the same type as i
}
The declaration in a type switch has the same syntax as a type assertion i.(T), but the specific type T is replaced with the keyword type.

This switch statement tests whether the interface value i holds a value of type T or S. In each of the T and S cases, the variable v will be of type T or S

respectively and hold the value held by i. In the default case (where there is no match), the variable v is of the same interface type and value as i.

![image](https://user-images.githubusercontent.com/35446384/179388768-01163054-ded0-42f8-b8a1-5b179bc5892b.png)

# Stringers

One of the most ubiquitous interfaces is Stringer defined by the fmt package.

type Stringer interface {
    String() string
}
A Stringer is a type that can describe itself as a string. The fmt package (and many others) look for this interface to print values.

![image](https://user-images.githubusercontent.com/35446384/180149548-026df1fb-ff7d-4f87-91fb-1992fcf8c87f.png)

 # Console Input (Bufio Scanner) & Type Conversion
 Read input from user keyboard , make **scan** for it and print
 
 **bufio** , **os** , **fmt**
 
 ![image](https://user-images.githubusercontent.com/35446384/180376011-710bd5e4-0ffb-446a-b470-87ed968591ed.png)
 
 # convert string to int 
 ![image](https://user-images.githubusercontent.com/35446384/180597221-6e3a52af-d519-4a9b-b819-c9bc8bde2a14.png)

#  Chained Conditionals (AND, OR, NOT)
**the operator you have !  ||  &&**

**true || false -> true**

**true && false -> false**

# if , else , if else

func main() {

	fmt.Println("Please enter your age !")
	age := bufio.NewScanner(os.Stdin)
	age.Scan()
	ageConvert, _ := strconv.ParseInt(age.Text(), 10, 64)
	if ageConvert > 10 {
		fmt.Printf("age is greater than 10 and it's value %v", ageConvert)
	} else if ageConvert == 10 {
		fmt.Printf("your age is equal to 10")
	} else {
		fmt.Printf("your age less than 10 and it's value 10 it's equal %v", ageConvert)

	}

}
