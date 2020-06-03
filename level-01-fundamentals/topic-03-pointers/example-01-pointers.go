/*
A basic program to demonstrate usage of pointers

Reference: https://tour.golang.org
*/

package main

import "fmt"

func main() {
	i, j := 10, 40

	p := &i         // points to i
	fmt.Println(*p) // read i through the pointer

	*p = 21        // set i through the pointer
	fmt.Println(i) // see the new value of i

	p = &j         // point to j
	*p = *p / 5    // divide j through the pointer
	fmt.Println(j) // see the new value of j
}
