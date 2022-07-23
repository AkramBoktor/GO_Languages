package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func main() {

	reader := bufio.NewScanner(os.Stdin)
	reader.Scan()
	fmt.Printf("Reader Input %s", reader.Text())

	fmt.Println("************************")
	fmt.Println("Enter the year you were born.!")
	birthYear := bufio.NewScanner(os.Stdin)
	birthYear.Scan()
	// value , type(decimal , hexa ,binary) , size of integer
	birthdataAsInt, _ := strconv.ParseInt(birthYear.Text(), 10, 64)
	fmt.Printf("you have %d years old", 2022-birthdataAsInt)

}
