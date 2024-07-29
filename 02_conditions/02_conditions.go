package conditions

import "fmt"

func Condition() {
	var accountBalance float64 = 1000
	var withDrawAmount float64 = 900

	if withDrawAmount > accountBalance {
		fmt.Println("Balance is invalid")
	}

	//fmt.Sprintf to format the print
	// "%f" --> float 100.000000 şeklinde
	// "%v" --> value 100 şeklinde

	if withDrawAmount <= accountBalance {
		fmt.Println("Balance is ready, nw!!")
		accountBalance = accountBalance - withDrawAmount
		//fmt.Println("your currnet balance is " + fmt.Sprintf("%v", accountBalance))
		fmt.Printf("your current balance is: %v", accountBalance)
	}
}
