package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	// Konsol : Veri Çıkışı
	str1 := "The quick red fox"
	str2 := "jumped over"
	str3 := "the lazy brown dog."
	aNumber := 42
	isTrue := true

	stringLength, _ := fmt.Println(str1, str2, str3)

	fmt.Println("String Length", stringLength)

	fmt.Printf("Value of aNumber: %v\n", aNumber) //%v kullanarak o alana "aNumber" değişkeni yazıldı.%v yer tutucu
	fmt.Printf("Value of isTrue: %v\n", isTrue)
	fmt.Printf("Value of aNumber as float: %.2f\n", float64(aNumber)) //%.2f kullanarak o alana "aNumber" değişkeni yazıldı.%.2f float yer tutucu

	fmt.Printf("Data types: %T, %T, %T, %T, and %T\n", str1, str2, str3, aNumber, isTrue) //%T yer tutucu = değişkenin tipini gösterir

	//golang placeholders

	myString := fmt.Sprintf("Data types as var: %T, %T, %T, %T, and %T", str1, str2, str3, aNumber, isTrue)
	fmt.Print(myString)

	fmt.Println("")
	//Konsol : Veri Girişi
	reader := bufio.NewReader(os.Stdin)
	fmt.Print("Enter text: ")
	str, _ := reader.ReadString('\n')
	fmt.Println(str)

	fmt.Print("Enter a number: ")
	str, _ = reader.ReadString('\n')
	f, err := strconv.ParseFloat(strings.TrimSpace(str), 64)
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println("Value of number:", f)
	}
}
