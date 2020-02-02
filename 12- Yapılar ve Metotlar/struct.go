package main

import (
	"fmt"
	"strconv"
)

type Human struct {
	FirstName string
	LastName  string
	Age       int
}

// Varsayılan ve boş yapıcı metot
func NewHuman() *Human {
	h := new(Human)
	return h
}

func NewHumanWithParam(firstName, lastName string, age int) *Human {
	h := new(Human)
	h.FirstName = firstName
	h.LastName = lastName
	h.Age = age
	return h
}

func main() {
	//Nesne örneği oluşturma

	//v1

	/*
		x := Human{FirstName: "Furkan"}
		fmt.Println(x.FirstName)*/

	//v2
	/*
		x := new(Human)
		x.FirstName = "Furkan"

		fmt.Println(x.FirstName)
	*/

	//Yapıcı metot kullanımı
	/*
		x := NewHuman()
		x.Age = 24
		fmt.Println(x.Age)
	*/
	//Parametreli yapıcı metot kullanımı
	/*
		x := NewHumanWithParam("Furkan", "SARIKAYA", 24)
		fmt.Println(x.FirstName)
	*/

	//Veriyi okuyalım
	x := NewHumanWithParam("Furkan", "SARIKAYA", 24)
	var data = x.FirstName + " " + x.LastName + " / " + strconv.Itoa(x.Age)
	fmt.Println(data)
}
