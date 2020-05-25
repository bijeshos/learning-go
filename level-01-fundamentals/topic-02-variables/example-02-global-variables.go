/*
A basic program to demonstrate global variable declaration

Reference: https://tour.golang.org
*/

package main

import "fmt"

var c, python, java bool

func main() {
	var i int
	fmt.Println(i, c, python, java)
}
