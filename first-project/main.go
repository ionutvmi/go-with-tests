package main

import "fmt"

type Person struct {
	firstName string
	lastName  string
}

func main() {
	var a = 44

	var p1 = Person{}

	p1.firstName = "Mihai"

	if a > 15 {
		fmt.Println("hello there")
	}

	fmt.Println("hello there")

	fmt.Println("this is p1 = ", p1)

}
