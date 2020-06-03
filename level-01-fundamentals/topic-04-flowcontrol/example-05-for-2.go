/*
A basic program to demonstrate usage of simple for loop

Reference: https://tour.golang.org
*/

package main

import "fmt"

func main() {
	sum := 1
	// for without initialization and increment
	//for ; sum < 1000; {
	for sum < 1000 {
		sum += sum
	}
	fmt.Println(sum)
}
