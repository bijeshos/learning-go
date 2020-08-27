/*
A simple program to demonstrate error handling

Reference: https://golang.org/pkg
*/

package main

import (
	"errors"
	"fmt"
	"os"
)

func main() {
	if _, err := os.Open("non-existing-file"); err != nil {
		var pathError *os.PathError
		if errors.As(err, &pathError) {
			fmt.Println("Got error of type pathError while opening file. Error message:", pathError)
		} else {
			fmt.Println("Got error of type other than pathError while opening file.", err)
		}
	}

}
