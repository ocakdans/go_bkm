package variables

import "fmt"

func Variable() {
	//1. Defining the variables
	var welcomeMessage string = "Hello World!"
	fmt.Println(welcomeMessage)

	var kdv int = 20
	fmt.Println("in Turkiye, kdv is", kdv)

	var weight float32 = 85.3
	fmt.Println("My weight is", weight)

	// 2. Defining variables in := format
	var myAge int = 39
	height := 189
	fmt.Println("My height is", height, "and my age is ", myAge)

	// 3. print the data type of the variable with Printf
	fmt.Printf("data type of height is : %T\n", height)

	var verification bool
	var text1 string = "selim"
	var text2 string = "feyza"
	// 4. check if they are same "=="
	// 5. check if they are NOT same "!="
	verification = text1 != text2
	fmt.Println("the status of this verification is", verification)
	fmt.Println("the status of this verification is", !verification) // prints opposite of the verification
}

func Hello(name string) string {
	return "Hello, " + name
}
