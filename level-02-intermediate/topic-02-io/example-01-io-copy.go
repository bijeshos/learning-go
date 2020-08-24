/*
A simple program to demonstrate io.Copy()

Reference: https://golang.org/pkg
*/

package main

import (
	"io"
	"log"
	"os"
	"strings"
)

func main() {
	r := strings.NewReader("some random words coming from a reader\n")

	if _, err := io.Copy(os.Stdout, r); err != nil {
		log.Fatal(err)
	}

}
