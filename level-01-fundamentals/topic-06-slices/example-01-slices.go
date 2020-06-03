/*
A simple program to demonstrate slices

Reference: https://tour.golang.org
*/

package main

import "fmt"

func main() {
	primes := [6]int{2, 3, 5, 7, 11, 13}

	// index 1 to index 3 (excluding 4)
	var s []int = primes[1:4]
	fmt.Println(s)
}
