/*
A simple program to demonstrate reading files using ioutil

Reference: https://golang.org/pkg
*/

package main

import (
	"fmt"
	"io/ioutil"
)

func main() {

	//This will return the content of the file
	data, err := ioutil.ReadFile("./.gitignore")
	if err != nil {
		fmt.Println("File reading error", err)
		return
	}

	fmt.Println("Contents of file:", string(data))
}
