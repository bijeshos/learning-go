/*
A simple program to demonstrate xxx

Reference: https://tour.golang.org
*/

package main

import "fmt"

func main() {
	s := []int{2, 3, 5, 7, 11, 13}

	//takes values from index 1 to 3
	s = s[1:4]
	fmt.Println(s)

	//takes values from index 0 to 1
	s = s[:2]
	fmt.Println(s)

	//takes values from index 1 to length of slice (1)
	s = s[1:]
	fmt.Println(s)
}
