/*
A simple program to demonstrate usage of scanner

Reference: https://golang.org/pkg
*/

package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	//create a scanner with stdin as input source
	scanner := bufio.NewScanner(os.Stdin)
	fmt.Println("Press Ctrl c to exit")
	fmt.Println("Enter a value:")
	//start scanning for input
	for scanner.Scan() {
		fmt.Println("Entered value is: ", scanner.Text())
		fmt.Println("Enter another value:")
	}
	if err := scanner.Err(); err != nil {
		fmt.Fprintln(os.Stderr, "reading standard input:", err)
	}
}

//reference: https://golang.org/pkg
