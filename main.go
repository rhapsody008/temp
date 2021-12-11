package main

import (
	"fmt"

	"rsc.io/quote"
)

func main() {

	temp()

	fmt.Println(quote.Go())

	var inta float32
	var intb float32
	var selectresult string

	fmt.Print("Enter first number: ")
	fmt.Scanln(&inta)
	fmt.Print("Enter second number: ")
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
		if intb == 0 {
			fmt.Println("Error! Second Number cannot be ZERO when executing division!")
		} else {
			fmt.Println("the result is ", inta/intb)
		}
	default:
		fmt.Println("Please choose between a, b, c and d.")

	}

}

func temp() {
	var a int
	type hotdog int
	var b hotdog

	a = 42
	b = 43
	fmt.Println(a)
	fmt.Printf("%T\n", a)
	fmt.Println(b)
	fmt.Printf("%T\n", b)

}
