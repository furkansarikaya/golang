package main

import (
	"fmt"
	"net/http"
	"strconv"
)

type Human struct {
	FName string
	LName string
	Age   int
}

func (hum Human) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	hum.FName = "Furkan"
	hum.LName = "SARIKAYA"
	hum.Age = 25

	// Formu parse etmek için
	r.ParseForm()

	//Sunucudan form bilgisini almak için
	fmt.Println(r.Form)

	//URL'nin path bilgisini almak için
	fmt.Println("path", r.URL.Path)

	fmt.Fprint(w, "<table><tr><td><b>Ad</b></td><td><b>Soyad</b></td><td><b>Yaş</b></td></tr><tr><td>"+hum.FName+"</td><td>"+hum.LName+"</td><td>"+strconv.Itoa(hum.Age)+"</td></tr><tr><tr></tr><tr></tr><tr><td><b>Path</b></td><td>"+r.URL.Path+"</td></tr></table>")

}

func main() {
	var hum Human
	err := http.ListenAndServe("localhost:9000", hum)
	checkError(err)
}
func checkError(err error) {
	if err != nil {
		panic(err)
	}
}
