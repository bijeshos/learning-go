/*
A simple program to demonstrate usage of creating new error

Reference: https://golang.org/pkg
*/

package main

import (
	"fmt"
)

func main() {
	const name, id = "test user", 17
	err := fmt.Errorf("user %q with id %d not found", name, id)
	if err != nil {
		fmt.Print(err)
	}
}
