package main

import (
	"fmt"
	"strconv"
)

func main() {
	// kullanıcı veri oluşturma
	//v1

	fmt.Println("Kullanıcı oluşturma v1")

	user1 := &User{
		Id:        1,
		FirstName: "Furkan",
		LastName:  "SARIKAYA",
		UserName:  "furkansarikaya",
		Age:       24,
		Pay: &Payment{
			Salary: 4250,
			Bonus:  202,
		},
	}
	fmt.Println(user1.GetFullName())
	fmt.Println(user1.GetUserName())
	fmt.Println(user1.GetPayment())

	fmt.Println("Maaş : " + strconv.FormatFloat(user1.GetPayment(), 'g', -1, 64))

	//v2
	fmt.Println("\n")

	fmt.Println("Kullanıcı oluşturma v2")

	user2 := NewUser()

	user2.FirstName = "Furkan"
	user2.LastName = "SARIKAYA"
	user2.Age = 24
	user2.UserName = "pluton"

	user2.Pay = &Payment{Salary: 4000, Bonus: 202}

	fmt.Println(user2.GetFullName())
	fmt.Println(user2.GetUserName())
	fmt.Println(user2.GetPayment())

	fmt.Println("Maaş : " + strconv.FormatFloat(user2.GetPayment(), 'g', -1, 64))

}

// kullanıcı yapısı
type User struct {
	Id        int
	FirstName string
	LastName  string
	UserName  string
	Age       int
	Pay       *Payment
}

//kullanıcının yapıcı metodu
func NewUser() *User {
	u := new(User)
	u.Pay = NewPayment()
	return u
}

//ödeme yapısı
type Payment struct {
	Salary float64
	Bonus  float64
}

// Ödemenin yapıcı metodu
func NewPayment() *Payment {
	p := new(Payment)
	return p
}

//Metotlar
func (u User) GetFullName() string {
	return u.FirstName + " " + u.LastName
}

func (u *User) GetUserName() string {
	return u.UserName
}

func (u *User) GetPayment() float64 {
	pay := u.Pay.Salary + u.Pay.Bonus
	return pay
}
