/*
A basic program to demonstrate default values for variables

Reference: https://tour.golang.org
*/

package main

import (
	"fmt"
)

func main() {
	var a bool
	var i int
	var f float32
	var c complex64

	const constant = "PI"

	fmt.Println("boolean:", a, ", int:", i, ", float: ", f, ", complex: ", c)
	fmt.Println("constant:", constant)
}
