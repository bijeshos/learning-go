/*
A basic program to demonstrate usage of structs
*/

package main

import "fmt"

type person struct {
	firstName string
	lastName  string
	age       int
}

func main() {
	fmt.Println(person{"John", "Smith", 30})
	fmt.Println(person{firstName: "Jane", lastName: "Smith", age: 28})

	s := person{"Mike", "Smith", 10}
	fmt.Println(s.firstName)
	fmt.Println(s.lastName)
	fmt.Println(s.age)

}
