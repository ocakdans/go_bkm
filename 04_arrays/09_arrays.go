package arrays

import "fmt"

func Array() {

	var numbers [5]int
	//fmt.Println(numbers) // [0 0 0 0 0]
	numbers[2] = 50
	//fmt.Println(numbers)    // [0 0 50 0 0]
	//fmt.Println(numbers[2]) // 50

	var cities [5]string
	cities[0] = "Ankara"
	cities[1] = "İstanbul"
	cities[2] = "İzmir"
	cities[3] = "Adana"
	cities[4] = "Ordu"
	//cities[5] = "D.Bakır"
	//fmt.Println(cities)

	for i := 0; i < len(cities); i++ {
		//fmt.Println(cities[i])
	}

	counts := [5]int{20, 30, 10, 42, 5} //you can assign while defining it
	// bir de [...] şeklinde verirsek autosize yapar, sayısını vermeden de ne kadar tanımlarsan ona göre set eder.
	var maxNumber int = 0
	for i := 0; i < len(counts); i++ {
		if counts[i] > maxNumber {
			maxNumber = counts[i]
		}
	}

	fmt.Println("En Büyük number:", maxNumber)

}
