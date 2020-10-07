package main

import (
	"encoding/json"
	"fmt"
	"os"
)

/*
func main() {
	jsonStr := `
		{
			"data" : {
				"object":"card",
				"id":"card_467443733436",
				"first_name":"Furkan",
				"last_name":"SARIKAYA",
				"balance":"54.950"
			}
		}
	`
	var m map[string]map[string]interface{}

	if err := json.Unmarshal([]byte(jsonStr), &m); err != nil {
		panic(err)
	}

	fmt.Println(m)

	fmt.Println("----------------------")

	b, err := json.Marshal(m)
	if err != nil {
		panic(err)
	}

	fmt.Println(string(b))

}
*/

//JSON İşlemler

// Type Nesneleri

// Name sturct'ının tanımlıyoruz
type Name struct {
	Family   string
	Personal string
}

// Email sturct'ının tanımlıyoruz
type Email struct {
	ID      int
	Kind    string
	Address string
}

//Interest struct
type Interest struct {
	ID   int
	Name string
}

type Person struct {
	ID       int
	FistName string
	LastName string
	UserName string
	Gender   string
	Name     Name
	Email    []Email
	Interest []Interest
}

func GetPerson(p *Person) string {
	return p.FistName + " " + p.LastName
}

func GetPersonEmailAddress(p *Person, i int) string {
	return p.Email[i].Address
}

func GetPersonEmail(p *Person, i int) Email {
	return p.Email[i]
}

func WriteMessage(msg string) {
	fmt.Println(msg)
}

func WriteStarLine() {
	fmt.Println("*************")
}

func checkError(err error) {
	if err != nil {
		fmt.Println("Fatal Error : ", err.Error())
		os.Exit(1)
	}
}

func SaveJSON(fileName string, key interface{}) {
	outFile, err := os.Create(fileName)
	checkError(err)
	defer outFile.Close()
	encoder := json.NewEncoder(outFile)
	err = encoder.Encode(key)
	checkError(err)

}

func main() {
	person := Person{
		ID:       9,
		FistName: "Furkan",
		LastName: "SARIKAYA",
		UserName: "FurkanSARIKAYA",
		Gender:   "Erkek",
		Name:     Name{Family: "Prs", Personal: "Furkan"},
		Email: []Email{
			Email{ID: 1, Kind: "Work", Address: "furkan@furkansarikaya.com"},
			Email{ID: 2, Kind: "Personel", Address: "furkan@gmail.com"},
		},
		Interest: []Interest{
			Interest{ID: 1, Name: "Go"},
			Interest{ID: 2, Name: "C#"},
			Interest{ID: 3, Name: "Java"},
		},
	}

	WriteMessage("Reading Operation Started")

	WriteMessage("Personel Fullname")
	WriteStarLine()

	res := GetPerson(&person)

	WriteMessage(res)

	WriteMessage("\n")

	WriteMessage("Personel Email With Index")
	WriteStarLine()

	resEmail := GetPersonEmailAddress(&person, 1)
	WriteMessage(resEmail)

	WriteStarLine()

	WriteMessage("\n")

	WriteMessage("Personel Email Object With Index")
	WriteStarLine()

	resEmail2 := GetPersonEmail(&person, 0)
	fmt.Println(resEmail2)
	WriteStarLine()

	WriteMessage("Reading Operation Ended")

	WriteMessage("\n")

	WriteMessage("Writing Operation Started")

	SaveJSON("person.json", person)
	WriteMessage("Writing Operation Ended")
}
