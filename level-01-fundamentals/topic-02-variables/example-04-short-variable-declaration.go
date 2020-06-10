/*
A basic program to demonstrate short variable declaration

Reference: https://tour.golang.org
*/

package main

import "fmt"

func main() {
	var i, j int = 1, 2
	k := 3
	a, b, c := true, false, "no!"

	fmt.Println(i, j, k, a, b, c)
}
