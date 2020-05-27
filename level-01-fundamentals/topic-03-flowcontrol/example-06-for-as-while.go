/*
A basic program to demonstrate using 'for' keyword for 'while' functtionality

Reference: https://tour.golang.org
*/

package main

import "fmt"

func main() {
	sum := 1
	//go has no while loop. while behavious can be achieved as follows
	for sum < 1000 {
		sum += sum
	}
	fmt.Println(sum)
}
