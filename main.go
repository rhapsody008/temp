package main

import (
	"fmt"

	"rsc.io/quote"
)

func main() {
	fmt.Println(quote.Go())

	var inta float32
	var intb float32
	var selectresult string

	fmt.Println("Enter first number:")
	fmt.Scanln(&inta)
	fmt.Println("Enter second number:")
	fmt.Scanln(&intb)

	fmt.Println("Would you like: a.add, b.substract, c.multiply or d.divide? (Please specify a, b, c or d)")
	fmt.Scanln(&selectresult)

	switch selectresult {
	case "a":
		fmt.Println("the result is ", inta+intb)
	case "b":
		fmt.Println("the result is ", inta-intb)
	case "c":
		fmt.Println("the result is ", inta*intb)
	case "d":
		fmt.Println("the result is ", inta/intb)
	}

}
