/*
A basic program to demonstrate usage of 'defer'

Reference: https://tour.golang.org
*/

package main

import "fmt"

func main() {
	//this will get executed after main completes execution
	defer fmt.Println("world")

	fmt.Println("hello")
}
