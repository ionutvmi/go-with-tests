package main

import (
	"fmt"
	"os"

	"go-with-tests/first-project/dependencyinjection"
	"go-with-tests/first-project/mocking"
	"go-with-tests/first-project/numbers"
)

type Person struct {
	firstName string
	lastName  string
}

func main() {
	var a = 44

	var p1 = Person{}

	p1.firstName = "Mihai"
	p1.lastName = "Snow"

	if a > 15 {
		fmt.Println("hello there")
	}

	fmt.Println("hello there")

	fmt.Println("this is p1 = ", p1)

	numbers.ShowNumber()

	dependencyinjection.GreetWithDI(os.Stdout, "Mihai")

	mocking.Countdown(os.Stdout, mocking.DefaultSleeper{})
}
