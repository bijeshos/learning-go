/*
A simple program to demonstrate goroutines
*/

package main

import (
	"fmt"
)

func print(n int) {
	for i := 0; i < 5; i++ {
		fmt.Println(n + i + 1)
	}
}

func main() {
	//call the function directly
	print(10)

	//invoke function as a goroutine
	go print(20)

	//This is added to halt the execution till next keystroke
	var input string
	fmt.Scanln(&input)
}
