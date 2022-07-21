package main

import (
	"fmt"
	"strings"
)

type person struct {
	Name string
	Age  int
}
type IPAddr [4]byte

func (ip IPAddr) String() string {
	IPAddrString := []string{}
	for _, num := range ip {
		IPAddrString = append(IPAddrString, fmt.Sprint(int(num)))
	}
	return strings.Join(IPAddrString, ".")
}
func main() {

	firstPerson := person{
		Name: "Akram",
		Age:  28,
	}

	fmt.Printf("The person value is name : %v age : %v \n", firstPerson.Name, firstPerson.Age)

	fmt.Println("***********************************")
	hosts := map[string]IPAddr{
		"loopback":  {127, 0, 0, 1},
		"googleDNS": {8, 8, 8, 8},
	}
	for name, ip := range hosts {
		fmt.Printf("%v: %v\n", name, ip)
	}

}
