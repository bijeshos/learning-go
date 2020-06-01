/*
A simple program to demonstrate xxx

Reference: https://tour.golang.org
*/

package main

import "fmt"

func main() {
	//creates int array of lenth 5, capacity 5
	a := make([]int, 5)
	printSlice("a", a)

	//creates int array of lenth 0, capacity 5
	b := make([]int, 0, 5)
	printSlice("b", b)

	//assigning values from index 0 to 1
	c := b[:2]
	printSlice("c", c)

	//assigning values from index 2 to 4
	d := c[2:5]
	printSlice("d", d)
}

func printSlice(s string, x []int) {
	fmt.Printf("%s len=%d cap=%d %v\n",
		s, len(x), cap(x), x)
}
