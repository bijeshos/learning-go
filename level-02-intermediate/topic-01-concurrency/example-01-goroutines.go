/*
A simple program to demonstrate goroutines
*/

package main

import (
	"fmt"
)

func display(n int) {
	for i := 0; i < 5; i++ {
		fmt.Println(n + i + 1)
	}
	fmt.Println("---")
}

func main() {
	//call the function directly
	display(10)

	//invoke function as a goroutine
	go display(20)

	//This is added to halt the execution till next key-stroke
	var input string
	fmt.Scanln(&input)
}
