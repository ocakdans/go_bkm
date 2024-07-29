package conditions

import "fmt"

func Workshop1() {
	// Define 3 int variables
	// Dsiplay the biggest one to the scree

	var a, b, c int = 10, 5, 18

	var bigOne int

	if a > b {
		bigOne = a
	} else if b <= a {
		bigOne = b
	}

	if c > bigOne {
		bigOne = c
	}

	fmt.Println("the biggest number is", bigOne)

}
