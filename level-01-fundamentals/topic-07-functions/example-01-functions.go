/*
A simple program to demonstrate basic function

Reference: https://tour.golang.org
*/
package main

import "fmt"

func add(x int, y int) int {
	return x + y
}

func subtract(a, b int) int {
	return a - b
}

func main() {
	fmt.Println("after adding: ", add(10, 20))
	fmt.Println("after subtracting: ", subtract(80, 10))
}
