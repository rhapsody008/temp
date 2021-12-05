package main

import (
	"fmt"

	"rsc.io/quote"
)

func main() {
	fmt.Println(quote.Go())

	fmt.Println("the result is ", foo())

}

func foo() int {
	result := 3 + 4
	return result
}
