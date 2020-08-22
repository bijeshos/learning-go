/*
A simple program to demonstrate reading directory content using ioutil

Reference: https://golang.org/pkg
*/

package main

import (
	"fmt"
	"io/ioutil"
	"log"
)

func main() {

	//file names are returned in a sorted order
	files, err := ioutil.ReadDir(".")
	if err != nil {
		log.Fatal(err)
	}

	for _, file := range files {
		fmt.Println(file.Name())
	}
}
