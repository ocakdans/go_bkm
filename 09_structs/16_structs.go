package structs

import "fmt"

func Struct() {
	fmt.Println(product{"Computer", 500, "Casper", 10})
	fmt.Println(product{"Computer", 500, "", 5})
	fmt.Println(product{name: "Computer"}) // {Computer 0  0} bu şekilde gelir, sadece bir kısmını yazdırmak istememiz durumunda.
}

type product struct {
	name         string
	price        float64
	brand        string
	discountRate int
}
