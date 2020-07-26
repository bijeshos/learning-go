/*
A simple program to demonstrate goroutines and channels

The example code sums the numbers in a slice, distributing the work between two goroutines.
Once both goroutines have completed their computation, it calculates the final result.
By default, sends and receives block until the other side is ready. This allows goroutines to synchronize without
explicit locks or condition variables.

Reference: https://tour.golang.org
*/
package main

import "fmt"

func sum(s []int, c chan int) {
	sum := 0
	for _, v := range s {
		sum += v
	}
	c <- sum // send sum to c
}

func main() {
	numbers := []int{7, 2, 8, -9, 4, 0}

	sumChannel := make(chan int) //make a channel of int;

	go sum(numbers[:len(numbers)/2], sumChannel) // call sum by passing number array and channel
	go sum(numbers[len(numbers)/2:], sumChannel) // call sum one more time
	x, y := <-sumChannel, <-sumChannel           // receive from sumChannel

	fmt.Println("x:", x, ", y:", y, ", x+y: ", x+y)
}
