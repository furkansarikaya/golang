package main

import (
	"fmt"
)

//Diziler

func main() {

	//basit bir dizi
	myArray1 := [3]int{}
	myArray1[0] = 32
	myArray1[1] = 23
	myArray1[2] = 54
	fmt.Println(myArray1)

	fmt.Println("-----------------------------------------------")

	//renk dizisi
	var colors [3]string
	colors[0] = "Red"
	colors[1] = "Green"
	colors[2] = "Blue"

	fmt.Println(colors)
	fmt.Println(colors[1]) //get
	colors[1] = "Yellow"   //set
	fmt.Println(colors[1])

	fmt.Println("-----------------------------------------------")

	var numbers = [5]int{4, 6, 7, 3, 66}
	fmt.Println(numbers)
	fmt.Println("Number of numbers", len(numbers))

	//eleman sayısını belirtmeden dizi oluşturma(otomatik size değerini alıyor)
	myArray2 := [...]int{4, 3, 5, 6, 7, 332}
	myArray2[3] = 5555
	fmt.Println(myArray2)

	fmt.Println("-----------------------------------------------")

	//dizilerde dögüler
	var cars [3]string
	cars[0] = "Tesla"
	cars[1] = "Mercedes"
	cars[2] = "Jaguar"
	i := 0
	for i <= len(cars)-1 {
		fmt.Println(cars[i])
		i++
	}
	fmt.Println("-----------------------------------------------")
	for i := 0; i < len(cars); i++ {
		fmt.Println(cars[i])
	}
	fmt.Println("-----------------------------------------------")
	for index, value := range cars {
		fmt.Println("i =", index, "&value =", value)
	}
	fmt.Println("-----------------------------------------------")
	for index := range cars {
		fmt.Println("i =", index, "&value =", cars[index])
	}
	fmt.Println("-----------------------------------------------")
	for index, _ := range cars {
		fmt.Println("i =", index, "&value =", cars[index])
	}
	fmt.Println("-----------------------------------------------")
	for _, value := range cars {
		fmt.Println("value =", value)
	}
}
