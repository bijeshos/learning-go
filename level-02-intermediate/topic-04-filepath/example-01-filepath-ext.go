/*
A simple program to demonstrate usage of filepath.Ext()

Reference: https://golang.org/pkg
*/

package main

import (
	"fmt"
	"path/filepath"
)

func main() {

	//returns extension of the file; returns empty if no extension present
	fmt.Printf("Value returned for input with no dots: %q\n", filepath.Ext("index"))
	fmt.Printf("Value returned for input with one dot: %q\n", filepath.Ext("index.js"))
	fmt.Printf("Value returned for input with two dots: %q\n", filepath.Ext("main.test.js"))
}
