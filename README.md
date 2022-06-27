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
