package main

import "fmt"

func add(x int,y int) int{
	return x+y
}

//gelen veriyi değiştirebilmek için * kullandık.
func sayHello(message *string){
	fmt.Println(*message)
	*message="Hi Go!"
}
func main(){
	fmt.Println(add(42,13))
	message :="Hi"
	sayHello(&message)
	fmt.Println(message)
}

// Fonksiyonlar

/*
//pointer ile
func main() {
	message := "Hi"
	sayHello(&message)
	println(message)
}

func sayHello(message *string) {
	println(*message)
	*message = "Hi Go!"
}*/

/*
func main() {
	println(add(3, 4, ""))
}

func add(x, y int, z string) int {
	return x + y
}*/

//------------------------------------------------------------------------------------------------------------//

// Fonksiyonlar : Çoklu Sonuç Dönmek
/*
func main() {
	a, b := swap("Furkan", "SARIKAYA")
	println(a, b)
	numTerms, sum := add(3, 4, 5)
	println("Added :", numTerms, "terms for a total of", sum)
}

func swap(x, y string) (string, string) {
	return y, x
}

func add(terms ...int) (int, int) {
	result := 0
	for _, term := range terms {
		result += term
	}
	return len(terms), result
}*/

//------------------------------------------------------------------------------------------------------------//

// Fonksiyonlar : İsimlendirilmiş Geri Dönüş Yapan Fonksiyonlar

/*func main() {
	fmt.Println(split(17))

	numTerms, sum := add(1, 5)
	println("Added :", numTerms, "terms for a total of", sum)
}
func split(sum int) (x, y int) {
	x = sum * 4 / 9
	y = sum - x
	return
}

func add(terms ...int) (numTerms int, sum int) {
	for _, term := range terms {
		sum += term
	}
	numTerms = len(terms)
	return
}*/

//------------------------------------------------------------------------------------------------------------//

// Fonksiyonlar : Değişken Sayıda Parametre Alan Fonksiyonlar(Variadic Functions)

/*func main() {
	sayHello("Merhaba", "Go", "Developer", ".")
	result := add(3, 5, 87, 89, 4, 115555884)
	println(result)
}

func sayHello(messages ...string) {
	for _, message := range messages {
		println(message)
	}
}

func add(terms ...int) int {
	result := 0
	for _, term := range terms {
		result += term
	}
	return result
}*/

//------------------------------------------------------------------------------------------------------------//

// Fonksiyonlar : Anonim Fonksiyonlar
/*
func main() {

	addFunc := func(terms ...int) (numTerms int, sum int) {
		for _, term := range terms {
			sum += term
		}
		numTerms = len(terms)
		return
	}
	numTerms, sum := addFunc(3, 50)

	println("Added :", numTerms, "terms for a total of", sum)

}
*/