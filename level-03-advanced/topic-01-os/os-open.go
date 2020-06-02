/*
A simple program to demonstrate reading files using os package

Reference: https://golang.org/pkg
*/

package main

import (
	"fmt"
	"log"
	"os"
)

func main() {
	//open file for read access
	file, err := os.Open(".gitignore")

	//log details if error is returned
	if err != nil {
		log.Fatal(err)
	}

	//initialize a byte array
	data := make([]byte, 1000)

	//read data into byte array and assign number of bytes to count
	count, err := file.Read(data)

	//close the file
	//this will be executed after all operations are done
	defer file.Close()

	//log details if error is returned
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("read %d bytes: %q\n", count, data[:count])

}
