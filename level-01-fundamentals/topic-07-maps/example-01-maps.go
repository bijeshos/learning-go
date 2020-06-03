/*
A simple program to demonstrate map

Reference: https://tour.golang.org
*/

package main

import "fmt"

func main() {

	//create an empty map. usage: make(map[key-type]val-type)
	m := make(map[string]int)

	m["key-1"] = 10
	fmt.Println("value for key-1:", m["key-1"])

	m["key-2"] = 20
	fmt.Println("value for key-2:", m["key-2"])

	m["key-3"] = 30
	fmt.Println("value for key-3:", m["key-3"])

	fmt.Println("map:", m)
	fmt.Println("len:", len(m))

	delete(m, "key-3")

	_, present := m["key-3"]
	fmt.Println("is key-3 present:", present)

	fmt.Println("map:", m)
	fmt.Println("len:", len(m))

	v, ok := m["key-1"]
	fmt.Println("key-1 value:", v, "Present?", ok)
}
