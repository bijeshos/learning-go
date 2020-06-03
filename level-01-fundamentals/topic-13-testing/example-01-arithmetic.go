/*
A simple program to demonstrate testing.
This is the file to be tested. A separate test file is available to test functionalities present in this file.
*/

package main

import "fmt"

//Add is the function to be tested
func Add(a int, b int) int {
	return a + b
}

//Subtract does not have any unit test cases
func Subtract(a int, b int) int {
	return a - b
}

func main() {
	fmt.Println("sum of 2 and 4 is: ", Add(2, 4))
}
