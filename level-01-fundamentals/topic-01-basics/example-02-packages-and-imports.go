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
	"time"
)

func main() {

	//fmt is an imported package and Println() is a function within fmt package
	fmt.Println("This is a test message")

	//time is another imported package and Now() is a function within time package. Now() does not have any arguments
	fmt.Println("Current time is: ", time.Now())

	//math is one of the imported packages and Sqrt() is a function within math package. Sqrt() expects an input argument
	fmt.Printf("Square root of 7 is: %g \n", math.Sqrt(7))

	//Pi is an exported value in math package
	fmt.Println("Value of Pi: ", math.Pi)

	//Though the package imported is math/rand, we can use rand directly.
	fmt.Println("Here is a random number: ", rand.Intn(100))

}
