/*
A basic program to demonstrate package import.

Reference: https://tour.golang.org
*/

package main

import (
	"fmt"
	"math"
)

func main() {
	//math is the package and Sqrt() is a function within the package
	fmt.Printf("Square root of 7 is: %g \n", math.Sqrt(7))
}
