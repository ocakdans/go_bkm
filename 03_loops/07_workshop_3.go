package loops

import "fmt"

func Workshop3() {
	// take a number from the user
	// girdiği  asaldır ya da değildir i dön.

	var usersNumber int
	fmt.Println("Enter a number")
	fmt.Scanln(&usersNumber)
	isAsal := true

	for i := 2; i < usersNumber; i++ {
		if usersNumber%i == 0 {
			//fmt.Println("Asal Sayi değil !!!")
			isAsal = false
		}
	}
	if !isAsal {
		fmt.Println("Asal Değil")
	} else {
		fmt.Println("Asal sayı")
	}

}
