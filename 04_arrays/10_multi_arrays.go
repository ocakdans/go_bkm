package arrays

import "fmt"

func MultiArrays() {
	var numbers [2][3]int
	numbers[0][0] = 1
	numbers[0][1] = 2
	numbers[0][2] = 3
	numbers[1][0] = 4
	numbers[1][1] = 5
	numbers[1][2] = 6

	//fmt.Println(len(numbers))
	//fmt.Println(len(numbers[1]))

	//fmt.Println(numbers)
	for i := 0; i < len(numbers); i++ {
		for j := 0; j < len(numbers[1]); j++ {
			fmt.Println(numbers[i][j])
		}
	}

}
