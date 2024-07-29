package loops

import "fmt"

func Loops() {
	var text string = "Selim"

	//1
	for i := 0; i < 5; i++ {
		fmt.Println("text is,", text)
	}

	//2
	y := 1
	for y <= 3 {
		fmt.Println("text is nothing")
		y = y + 1
	}

}
