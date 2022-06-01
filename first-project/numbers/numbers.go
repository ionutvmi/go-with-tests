// Copyright Some Company Corp.
// All Rights Reserved
// Here is where we explain the package.
// Some other stuff.

package numbers

import "fmt"

// This is a function
// and it has a nice description
func getNumber() (int, error) {
	return 1, nil
}

// this method is public because it starts with an uppercase
func ShowNumber() {
	n, err := getNumber()

	if err != nil {
		fmt.Println("Could not generate the number")
		return
	}

	fmt.Printf("The number is %d", n)
}
