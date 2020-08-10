/*
A simple program to demonstrate mutual exclusion using mutex

Reference: https://tour.golang.org
*/

package main

import (
	"fmt"
	"sync"
	"time"
)

// counter is safe to use concurrently.
type counter struct {
	v   map[string]int
	mux sync.Mutex
}

// this method increments the counter for the given key.
func (c *counter) increment(key string) {
	c.mux.Lock()
	fmt.Println("Inside increment: locked | current key value: ", c.v[key])
	// Lock so only one goroutine at a time can access the map c.v.
	c.v[key]++
	fmt.Println("Inside increment: about to unlock after incrementing")
	c.mux.Unlock()
}

// this method returns current value of the counter for the given key.
func (c *counter) retrieve(key string) int {
	c.mux.Lock()
	fmt.Println("Inside retrieve: locked | current key value: ", c.v[key])
	// Lock so only one goroutine at a time can access the map c.v.
	defer c.mux.Unlock() //this will get executed after main completes execution
	return c.v[key]
}

func main() {
	c := counter{v: make(map[string]int)}
	for i := 0; i < 5; i++ {
		go c.increment("my-key")
	}

	time.Sleep(time.Second)
	fmt.Println("final key value: ", c.retrieve("my-key"))
}
