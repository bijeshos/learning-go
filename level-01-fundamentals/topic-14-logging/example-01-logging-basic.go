/*
A simple program to demonstrate logging (to bytes buffer)
*/

package main

import (
	"bytes"
	"fmt"
	"log"
)

func main() {
	var (
		buf bytes.Buffer
		//this creates a logger which writes to buf and has a prefix 'my-logger'
		logger = log.New(&buf, "my-logger: ", log.Lshortfile)
	)

	logger.Print("Hello, log file!")

	fmt.Print(&buf)
	//this is used to log fatal errors
	//this follows os.Exit(1)
	logger.Fatal("fatal error")

	fmt.Print(&buf) //this will not get executed
}
