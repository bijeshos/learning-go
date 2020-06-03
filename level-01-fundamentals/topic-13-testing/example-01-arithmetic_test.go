/*
A simple program to demonstrate testing.
This file contains unit and benchmarking tests.

a) how to execute the test:
- open a terminal, traverse to to this directory and execute the following
- $ go test
- $ go test -cover # to know the coverage as well

b) to visualize coverage, execute the following:
- $ go test -cover -coverprofile=c.out
- $ go tool cover -html=c.out -o coverage.html

c) to perform benchmarking, execute the following:
- $ go test -bench=.

*/

package main

import (
	"fmt"
	"testing"
)

//unit test
func TestAdd(t *testing.T) {
	total := Add(2, 4)
	if total != 6 {
		t.Errorf("sum was incorrect, expected: %d, actual: %d", 6, total)
	}
}

//benchmarking test
func BenchmarkAdd(b *testing.B) {
	total := Add(2, 4)
	fmt.Println("total:", total)
}
