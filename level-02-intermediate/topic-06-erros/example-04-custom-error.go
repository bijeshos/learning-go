/*
A simple program to demonstrate custom error creation

Reference: https://golang.org/pkg
*/

package main

import (
	"fmt"
	"time"
)

type CustomError struct {
	Time    time.Time
	Message string
}

func (e *CustomError) Error() string {
	return fmt.Sprintf("Error occurred at %v, with message '%s'",
		e.Time, e.Message)
}

func raiseError() error {
	return &CustomError{
		time.Now(),
		"custom error message",
	}
}

func main() {
	if err := raiseError(); err != nil {
		fmt.Println(err)
	}
}
