package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

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
