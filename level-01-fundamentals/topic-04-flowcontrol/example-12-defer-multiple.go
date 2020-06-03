/*
A basic program to demonstrate usage of multiple 'defer'

Reference: https://tour.golang.org
*/

package main

import "fmt"

func main() {
	fmt.Println("counting")

	for i := 0; i < 10; i++ {
		//these defers will stack up and execute in last-in-first-out fashion
		defer fmt.Println(i)
	}

	fmt.Println("done")
}
