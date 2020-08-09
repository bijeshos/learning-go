/*
A simple program to demonstrate closing channels and using range to loop through channel values

Note: Usually channels need not be closed. Closing is only necessary when the receiver must be told
there are no more values coming. If the receiver uses a range loop, it will help them to loop values accordingly.
Reference: https://tour.golang.org
*/
package main

import (
	"fmt"
)

func fibonacci(n int, c chan int) {
	x, y := 0, 1
	for i := 0; i < n; i++ {
		c <- x
		x, y = y, x+y
	}
	close(c) // This closes the channel. Only sender should close the channel.
}

func main() {
	c := make(chan int, 10) //change this values to 12 to see change in range
	go fibonacci(cap(c), c) // capacity of the channel is passed along with the channel
	//range will get executed till channel has values
	for i := range c {
		fmt.Println(i)
	}
}
