/*
A basic program to demonstrate global variable declaration

Reference: https://tour.golang.org
*/

package main

import "fmt"

var a, b, c bool

func main() {
	var i int
	fmt.Println(i, a, b, c)
}
