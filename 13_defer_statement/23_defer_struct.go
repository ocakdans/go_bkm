package deferstatement

import "fmt"

type product struct {
	productName string
	unitPrice   int
}

func (p product) Save() {
	fmt.Println("Product is saved to db:", p.productName)
	defer Log()
}

func Log() {
	fmt.Println("logged")
}

func RunDefer() {
	p := product{productName: "Laptop", unitPrice: 5000}
	p = product{productName: "Mouse", unitPrice: 45}
	defer p.Save() // burada mouse u kaydeder ama laptop tan sonra olsaydı laptop olurdu.- bi üst line a taşısaydım.

	fmt.Println("işlem başarılı")
}
