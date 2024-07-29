package maps

import "fmt"

func Range() {
	cities := []string{"Ankara", "İstanbul", "İzmir"}

	for i := 0; i < len(cities); i++ {
		fmt.Println(cities[i])
	}

	for _, city := range cities {
		fmt.Println(city)
	}

	for i, city := range cities {
		fmt.Print(i)
		fmt.Println(city)
	}

	// 1-10 arasındaki sayılardan tek olanları toplayan program

	numbers := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	sum := 0

	for _, number := range numbers {
		if number%2 == 1 {
			sum = sum + number
		}
	}
	fmt.Println("sum of the odds", sum)

	dictionary := map[string]string{"book": "kitap", "table": "masa"}

	for k, v := range dictionary {
		fmt.Println(k)
		fmt.Println(v)
	}
}
