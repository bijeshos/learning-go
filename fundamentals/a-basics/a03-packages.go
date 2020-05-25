/*
A basic program to demonstrate package main and import packages

Reference: https://tour.golang.org
*/

package main

import (
	"fmt"
	"math/rand"
)

func main() {
	//Though the package imported is math/rand, we can use rand directly.
	fmt.Println("Here is a random number: ", rand.Intn(100))
}
