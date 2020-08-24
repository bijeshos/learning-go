/*
A simple program to demonstrate creating directory structure

Reference: https://golang.org/pkg
*/

package main

import (
	"fmt"
	"log"
	"os"
)

func main() {

	path := "/home/bos/tmp/go-test/os-test"

	//create directory and parents as needed.
	//ensure that the go process has permissions to perform read/write actions
	err := os.MkdirAll(path, os.ModePerm)
	if err != nil {
		log.Fatal(err)
	} else {
		fmt.Println("directory created")
	}
}
