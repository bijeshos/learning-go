/*
A simple program to demonstrate basic function

Reference: https://tour.golang.org
*/
package main

import "fmt"

//here, types of inputs are declared individually
func add(x int, y int) int {
	return x + y
}

//here, since both inputs are of same types, type is only declared for the final argument
func subtract(a, b int) int {
	return a - b
}

func main() {
	result := add(10, 20)
	fmt.Println("10 + 20 = ", result)
	result = subtract(80, 10)
	fmt.Println("80 - 10 = ", result)
}
