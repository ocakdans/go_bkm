package loops

import "fmt"

func Workshop2() {
	// Take a number as a variable
	// guess it
	myNumber := 80
	userNumber := 0
	count := 0

	for userNumber != myNumber {
		fmt.Println("Enter a number: ")
		fmt.Scanln(&userNumber)
		fmt.Println("your number: ", userNumber)
		count++

		if userNumber > myNumber {
			fmt.Println("you said higher")
		} else if userNumber < myNumber {
			fmt.Println("you said less")
		} else {
			fmt.Println("Done, congrats!!!")
			if count <= 3 {
				fmt.Println("Super, only " + fmt.Sprintf("%v", count) + " trials")
			} else if count > 3 && count <= 6 {
				fmt.Println("Good, only " + fmt.Sprintf("%v", count) + " trials")
			} else {
				fmt.Println("Try hard, you did in " + fmt.Sprintf("%v", count) + " trials")
			}
		}
	}
}
