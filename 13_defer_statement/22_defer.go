package deferstatement

import "fmt"

func A() {
	fmt.Println("A is in progress")
}

func B() {
	defer A() // A her zaman çalışacak demek.
	//önce B'yi çalıştırır. Sonra A'yı çalıştırır. A'yı B'nin bitiminde çalıştırmak demek
	defer C() // en üstteki defer en son çalışır.
	// yani bu durumda önce B, sonra C, sonra A çalışır.
	fmt.Println("B is in progress")
}

func C() {
	fmt.Println("C is in progress")
}

func E(sayi int) string {
	defer DeferFunc() // return eder, sonra defer ü çalıştırır. sonra ekrana yazdırır.

	if sayi%2 == 0 {
		return "Çift sayıdır."
	}

	if sayi%2 == 1 {
		return "Tek sayıdır."
	}
	return "not defined"
}

func Test() {
	result := E(10)
	fmt.Println(result)
}

func DeferFunc() {
	fmt.Println("Bitti")
}
