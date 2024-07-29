package structs

import "fmt"

type customer struct {
	firstName string
	lastName  string
	age       int
}

func (c customer) save() {
	fmt.Println("Eklendi : ", c.firstName)

}

func (c customer) update() {
	fmt.Println("Güncellendi : ", c.firstName)

}

func Struct2() {
	c := customer{firstName: "Selim", lastName: "Ocakdan", age: 39}
	c.save()
	c.update()
}

//v2
// c := new(customer)
// c.firstName = "Selim" de yazabilirsin mesela.

// ******Default Constructor******

// func NewCustomer() *customer {
// 	c := new(customer)
// 	return c
// }
// z := NewCustomer()

// ******Parametrized Constructor******

// func NewCustomerWithParams(firstName, lastName  string) *customer {
// 	c := new(customer)
//  c.firstName = firstName
//  c.lastName = lastName
// 	return c
// }
