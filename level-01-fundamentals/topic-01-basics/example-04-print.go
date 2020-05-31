/*
A basic program to demonstrate printing various values

Reference: https://tour.golang.org
*/

package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	fmt.Println("This is a test message")
	fmt.Println("Current time is: ", time.Now())
	fmt.Println("Here is a random number: ", rand.Intn(100))
}
