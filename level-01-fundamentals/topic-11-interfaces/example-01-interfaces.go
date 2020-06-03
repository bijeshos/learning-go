/*
A simple program to demonstrate usage of interface
*/

package main

import "fmt"

//define an interface
type shape interface {
	area() float64
	perimeter() float64
}

//define a type
type rectangle struct {
	width, height float64
}

//define another type
type square struct {
	length float64
}

//implement shape interface on rectangle
func (r rectangle) area() float64 {
	return r.width * r.height
}
func (r rectangle) perimeter() float64 {
	return 2 * (r.width + r.height)
}

//implement shape interface on sqaure
func (s square) area() float64 {
	return 2 * s.length
}
func (s square) perimeter() float64 {
	return 4 * s.length
}

// print shape's details
func details(s shape) {
	fmt.Println(s)
	fmt.Println(s.area())
	fmt.Println(s.perimeter())
}

func main() {
	r := rectangle{width: 20, height: 10}
	s := square{length: 5}

	details(r)
	details(s)
}
