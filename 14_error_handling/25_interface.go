package errorhandling

import "fmt"

func Dogrula(i interface{}) {
	sayi, ok := i.(int) // int doğrulaması yapıyoruz burada
	fmt.Println(sayi, ok)
}

func TestDogrula() {
	var sayi1 interface{} = 5 // herşey bir interface aslında.
	Dogrula(sayi1)

	var sayi2 interface{} = "engin" // herşey bir interface aslında. // default olarak da 0 atıyor ve o şekilde ilerliyor.
	Dogrula(sayi2)
}
