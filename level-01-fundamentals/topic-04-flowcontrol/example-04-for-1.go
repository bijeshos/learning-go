/*
A basic program to demonstrate usage of simple for loop

Reference: https://tour.golang.org
*/

package main

import "fmt"

func main() {
	sum := 0
	//simple for loop, no paranthesis needed
	for i := 0; i < 10; i++ {
		sum += i
	}
	fmt.Println(sum)
}
