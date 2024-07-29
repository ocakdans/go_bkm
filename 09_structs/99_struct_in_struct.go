package structs

import "fmt"

func main() {
	//v1
	fmt.Println("Kullanıcı Oluşturma v1")
	user1 := &User{
		ID:        1,
		FirstName: "Selim",
		LastName:  "Ocakdan",
		UserName:  "ocakdans",
		Age:       30,
		Pay: &Payment{
			Bonus:  555,
			Salary: 5500,
		},
	}
	fmt.Println(user1.Pay.Salary)
	//v2
	fmt.Println("Kullanıcı Oluşturma v2")
	user2 := NewUser()
	user2.FirstName = "Feyza"
	user2.LastName = "Ocakdan"
	user2.Pay = &Payment{Salary: 10000, Bonus: 3000}
	fmt.Println(user2.Pay.Salary)
}

type User struct {
	ID        int
	FirstName string
	LastName  string
	UserName  string
	Age       int
	Pay       *Payment
}

type Payment struct {
	Salary float64
	Bonus  float64
}

func NewUser() *User {
	u := new(User)
	u.Pay = NewPayment()
	return u
}

func NewPayment() *Payment {
	p := new(Payment)
	return p
}

func (u User) GetFullName() string {
	return u.FirstName + " " + u.LastName
}

func (u User) GetUserName() string {
	return u.UserName
}

func (u User) GetPayment() float64 {
	return u.Pay.Salary + u.Pay.Bonus
}
