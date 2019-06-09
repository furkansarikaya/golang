package main

func main() {
	/*degiskenler
	  var message string
	  message = "Merhaba Go!"

	  var message string="Merhaba Go!"

	  var message="Merhaba Go!"

	  var a,b,c int=3,4,5

	  fmt.Println(a)
	  fmt.Println(b)
	  fmt.Println(c)

	  var message = "Hello World"
	  var a,b,c=3,true,4.5
	  var k,o, string="abc","xyz"
	  fmt.Println(message,a,b,c)
	  u := 55
	  v,n :="abc",true

	  a := "Metin"
	  b := 'M'
	  c := `Metin`

	  fmt.Println(a,b,c)

	  a := "Metin"
	  b := 'M'
	  c := `Metin`

	  fmt.Println(a, b, c)

	  var (
	  	degisken1 = "Furkan"
	  	degisken2 = "SARIKAYA"
	  	degisken3 = 1907
	  )

	  println(degisken1, degisken2, degisken3)
	*/

	/*operatorler
	  https://www.tutorialspoint.com/go/go_operators.htm
	*/

	/*Go ile Enum Kullanımı
	  	package main1

	  import "fmt"

	  type Brand string

	  const (
	      FACEBOOK  Brand = "Facebook"
	      MICROSOFT Brand = "Microsoft"
	      GOOGLE    Brand = "Google"
	      DIJIBIL   Brand = "Dijibil"
	  )

	  func PrintBrand(brand Brand) {
	      fmt.Println(brand)
	  }

	  func main1() {

	      PrintBrand(GOOGLE)
	      PrintBrand(DIJIBIL)

	      // Compile Error : undefined: VERIZON
	      // PrintBrand(VERIZON)

	      // Compile error : cannot use 25 (type int) as type Brand in function argument
	      // PrintBrand(25)

	  }
	*/

	/*Convert
	  var myString = "1"
	  var x = 10
	  var f float32 = 2.0
	  fmt.Println(myString, x, f)

	  //string'den int'e dönüştürme
	  //number hata olmazsa int hali,"_" boş tanımlayıcı eğer hata olur ise ve hatayı değişkene aktarmak istemediğimizden "_" boş tanımlayıcı kullandık
	  number, _ := strconv.Atoi(myString)
	  fmt.Println("Sonuç : " + strconv.Itoa(number))

	  result := number + 2
	  fmt.Println(result)
	*/

	/*Casting
	  var i int = 55
	  var f float64 = float64(i)
	  var u uint = uint(f)
	  fmt.Println(i, f, u)
	*/

}
