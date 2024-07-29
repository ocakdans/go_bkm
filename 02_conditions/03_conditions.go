package conditions

import (
	"fmt"
	"time"
)

func ConditionTwo() {
	var accountBalance float64 = 1000
	var withDrawAmount float64 = 1000

	if withDrawAmount > accountBalance {
		fmt.Println("Balance is invalid")
	} else if withDrawAmount == accountBalance {
		fmt.Println("Your money is being prepared at equality")
	} else {
		fmt.Println("Your money is being prepared")
	}
}

func Continue() {
	//continue sadece o case için atlamayı sağlar, sonrasına işlem yaptırmaz ama döngüyü tamamiyle kırmaz, i + 1 için döngü yeniden başlar.
	// break ise döngüyü kırar tamamen, i+1 için de işle yapılmaz
	// her ikisi de kendinden sonrakini o koşul için çalıştırmamış olur.
	for i := 0; i < 7; i++ {
		if i == 3 {
			continue
		} else if i == 5 {
			break
		}
		fmt.Println("Çıktı : ", i)
	}
	fmt.Println("İşleme devam..")
}

func WhileInGo() {
	// normalde Java ya da C# da while kullanılacağı yerde burada sadece for ile benzer syntax ile halledilmiş oluyor.
	sum := 1
	for sum < 10 {
		sum += sum
		fmt.Println("sum is:", sum)
		time.Sleep(300 * time.Millisecond)
	}
}
