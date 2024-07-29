package variables

import (
	"fmt"
	"strconv"
)

func TypeCasting() {
	myNumber := "1"
	number, _ := strconv.Atoi(myNumber)
	fmt.Println("Result:" + strconv.Itoa(number))

	fmt.Println(number)

	//you have to write the type casting. implicit e yazmadan müsade etmiyor.
	number2 := 12
	var myFloat float64 = float64(number2)
	fmt.Println(myFloat)

}
