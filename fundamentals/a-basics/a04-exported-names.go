/*
A basic program to demonstrate usage of exported values

Reference: https://tour.golang.org
*/

package main

import (
	"fmt"
	"math"
)

func main() {
	//Pi is an exported value in math package
	fmt.Println("Value of Pi: ", math.Pi)
}
