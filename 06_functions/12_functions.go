package functions

import "fmt"

func Sum(num1 int, num2 int) int {
	sum := num1 + num2
	return sum
}

func SayHello(name string) {
	fmt.Println("hey hey", name)
}

// multiple returns
func DortIslem(num1 int, num2 int) (int, int, int, float64) {
	sum := num1 + num2
	abstract := num1 - num2
	multiply := num1 * num2
	division := float64(num1 / num2)

	return sum, abstract, multiply, division
}

// named return
func NamedReturn(num1 int, num2 int) (sum, abstract, multiply int, division float64) {
	sum = num1 + num2
	abstract = num1 - num2
	multiply = num1 * num2
	division = float64(num1 / num2)

	return
}

// variadic functions

func SumVariadic(numbers ...int) int {
	sum := 0
	for i := 0; i < len(numbers); i++ {
		sum = sum + numbers[i]

	}
	return sum
}

//Anonim Fonksiyonlar
//birkez çalıştırılır ve bir kez kullanılırlar.
