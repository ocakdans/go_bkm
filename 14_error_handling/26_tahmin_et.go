package errorhandling

import (
	"errors"
	"fmt"
)

func TahminEt(number int) (string, error) {
	myNumber := 78
	if number < 1 || number > 100 {
		return "", errors.New("1-100 arasında sayı gir")
	}

	if number > myNumber {
		return "daha küçük bir sayı giriniz", nil
	}

	if number < myNumber {
		return "daha büyük bir sayı giriniz", nil
	}

	return "Bildiniz", errors.New("success")
}

func TahminTest() {
	fmt.Println(TahminEt(80))
	fmt.Println(TahminEt(70))
	fmt.Println(TahminEt(78))
}
