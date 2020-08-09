/*
A simple program to demonstrate goroutines and channels

Reference: https://tour.golang.org
*/

package main

import (
	"fmt"
	"time"
)

func prepareMessage(msg string, c chan string) {
	fmt.Println("inside prepareMessage")
	for i := 0; ; i++ { //note that there is no conditional check here
		fmt.Println("pushing to channel: ", msg, i)
		c <- fmt.Sprintf("%s %d", msg, i) //this blocks if channel is not empty
		time.Sleep(time.Duration(10 * time.Millisecond))
	}
}

func main() {
	c := make(chan string)
	go prepareMessage("hello", c)

	for i := 0; i < 2; i++ {
		fmt.Printf("taking from channel: %q \n", <-c)
	}
	fmt.Println("leaving main")

	//This is added to halt the execution till next key-stroke
	var input string
	fmt.Scanln(&input)
}
