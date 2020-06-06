/*
A simple program to demonstrate function with multiple results

Reference: https://tour.golang.org
*/

package main

import "fmt"

//here, two strings are returned
func swap(x, y string) (string, string) {
	return y, x
}

func main() {
	a, b := swap("hello", "world")
	fmt.Println(a, b)
}
