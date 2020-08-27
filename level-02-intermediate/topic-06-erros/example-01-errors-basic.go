/*
A simple program to demonstrate usage of creating new error

Reference: https://golang.org/pkg
*/
package main

import (
	"errors"
	"fmt"
)

func main() {
	err := errors.New("a sample custom error")
	if err != nil {
		fmt.Print(err)
	}
}
