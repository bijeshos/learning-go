/*
A simple program to demonstrate goroutines and buffered channels

To create buffered channels, provide the buffer length as the second argument to make to initialize a buffered channel.
Sends to a buffered channel block only when the buffer is full. Receives block when the buffer is empty.

Reference: https://tour.golang.org
*/

package main

import "fmt"

func main() {
	// this channel has a buffer size of 2
	ch := make(chan int, 2)
	ch <- 1
	ch <- 2
	//ch <- 3 //uncomment this line to see buffer overfill; increase buffer size to 3 to avoid overfill

	fmt.Println(<-ch)
	fmt.Println(<-ch)
	//fmt.Println(<-ch) //uncomment this line to see empty access on buffer

}
