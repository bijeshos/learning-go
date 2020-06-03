/*
A simple program to demonstrate the following:
- basic methods
- pointer reciever
- value reciever

Reference: https://gobyexample.com/methods
*/
package main

import (
	"fmt"
)

type rectangle struct {
	width, height int
}

//this method has a receiver type of *rectangle (pointer reciever)
func (r *rectangle) area() int {
	return r.width * r.height
}

//this method has a receiver type of rectangle (value reciever)
func (r rectangle) perimeter() int {
	return 2 * (r.width + r.height)
}

func main() {
	r := rectangle{width: 20, height: 10}

	//conversion between values and pointers for method calls are handled automatically
	//use pointer reciever to avoid working on a copy
	fmt.Println("area: ", r.area())
	fmt.Println("perimeter:", r.perimeter())

	rp := &r
	fmt.Println("area: ", rp.area())
	fmt.Println("perimeter:", rp.perimeter())
}
