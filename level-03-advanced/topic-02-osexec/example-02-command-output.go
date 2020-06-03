/*
A simple program to demonstrate usage of os/exec command and output options

Reference: https://golang.org/pkg
*/

package main

import (
	"fmt"
	"log"
	"os/exec"
)

func main() {
	out, err := exec.Command("date").Output()
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("The date is %s\n", out)
}
