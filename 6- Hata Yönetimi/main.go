package main

import (
	"errors"
	"fmt"
	"os"
)

//Hata Yönetimi : Error Handling / Exception Handling

func main() {
	// file, err := os.Open("dosyam.txt")
	// if err != nil {
	// 	fmt.Println(err.Error)
	// }
	// fmt.Println(file.Name)

	myError := errors.New("Bu bir hatadır!")
	fmt.Println(myError)

	_, err := os.Open("abc.rar")
	checkError(err)
}

func checkError(err error) {
	if err != nil {
		fmt.Println("ERROR : ", err.Error())
		//loglama
		os.Exit(1) //programı kapat
	}
}
