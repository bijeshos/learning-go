/*
A basic program to demonstrate the following:
- main package
- import of other built-in packages
- usage of exported functions and values

Reference: https://tour.golang.org
*/

package main

import (
	"fmt"
	"math"
	"math/rand"
)

func main() {

	//math is one of the imported packages and Sqrt() is a function within math package
	fmt.Printf("Square root of 7 is: %g \n", math.Sqrt(7))

	//Pi is an exported value in math package
	fmt.Println("Value of Pi: ", math.Pi)

	//Though the package imported is math/rand, we can use rand directly.
	fmt.Println("Here is a random number: ", rand.Intn(100))

}
