/*
A simple program to demonstrate usage of os/exec command and run options

Reference: https://golang.org/pkg
*/

package main

import (
	"log"
	"os/exec"
)

func main() {
	cmd := exec.Command("sleep", "1")
	log.Printf("Running command and waiting for it to finish...")
	err := cmd.Run()
	log.Printf("Command finished with error: %v", err)
}
