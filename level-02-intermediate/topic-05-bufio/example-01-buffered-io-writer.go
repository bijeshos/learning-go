/*
A simple program to demonstrate usage of buffered io

Reference: https://golang.org/pkg
*/

package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	//create a new buffered writer with stdout as the target
	w := bufio.NewWriter(os.Stdout)

	//write to buffered writer
	fmt.Fprint(w, "Hello, ")
	fmt.Fprint(w, "there!")

	//flush writer
	w.Flush()

	//this won't get printed since there is no flush after this statment
	fmt.Fprint(w, "how are you?")
}
