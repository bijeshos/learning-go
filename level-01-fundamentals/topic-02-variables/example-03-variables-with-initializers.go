/*
A basic program to demonstrate variable declarations with initializers

Reference: https://tour.golang.org
*/

package main

import "fmt"

var i, j int = 1, 2

func main() {
	var a, b, c = true, false, "no!"
	fmt.Println(i, j, a, b, c)
}
