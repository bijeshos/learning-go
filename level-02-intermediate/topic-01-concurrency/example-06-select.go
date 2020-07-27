/*
A simple program to demonstrate select operation with channels

- select is similar to switch, but, each case is a communication
- all cases are evaluated
- selection blocks until one communication can proceed.
- if multiple can proceed, selection is done pseudo-randomly
- if default case is present, it gets executes if no channel is ready

Reference: https://tour.golang.org
*/

package main

import (
	"fmt"
	"time"
)

func fibonacci(c chan int, quit chan int) {
	x, y := 0, 1
	for {
		time.Sleep(time.Duration(10 * time.Millisecond))
		select {
		case c <- x:
			fmt.Println("pushed: ", x)
			x, y = y, x+y
		case <-quit:
			fmt.Println("quitting")
			return
			/*default:
			fmt.Println("no channel is ready")*/
		}
	}
}

func main() {
	c := make(chan int)
	quit := make(chan int)

	//execute following section as a goroutine
	go func() {
		for i := 0; i < 5; i++ {
			fmt.Println("taking: ", <-c)
		}
		quit <- 0
	}()

	fibonacci(c, quit)

	/*
		//This is added to halt the execution till next key-stroke
		var input string
		fmt.Scanln(&input)
	*/
}
