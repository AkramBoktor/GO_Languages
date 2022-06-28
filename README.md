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

