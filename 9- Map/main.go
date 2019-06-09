package main

import (
	"fmt"
)

//Map

/*
Bir eşlem ("map"), anahtarları ("keys") değerlere ("values") eşler.

Bir eşlemin "sıfır değeri" `nil`'dir. Bir nil eşleminde anahtar bulunmadığı gibi anahtar da oluşturulamaz.

make fonksiyonu verilen türde, ilklendirilmiş ve kullanıma hazır durumda bir eşlem döner.
*/

type Vertex struct {
	Lat, Long float64
}

var m map[string]Vertex

func main() {
	m = make(map[string]Vertex)
	m["Bell Labs"] = Vertex{
		40.68433, -74.39967,
	}
	fmt.Println(m["Bell Labs"])

	fmt.Println("---------------------------")
	//1.
	myMap := make(map[int]string)
	fmt.Println(myMap)
	myMap[43] = "foo"
	myMap[12] = "bar"
	fmt.Println(myMap)
	//KeyValuePair

	fmt.Println("---------------------------")
	//2.
	states := make(map[string]string)
	states["IST"] = "İstanbul"
	states["ANK"] = "Ankara"
	states["ANT"] = "Antalya"
	states["ZNG"] = "Zonguldak"
	fmt.Println(states)

	//şehir listesinde ANT anahtar adına sahip veriyi elde et.
	zonguldak := states["ZNG"]
	fmt.Println("Seçili şehir :", zonguldak)

	//ANK anahtar adına sahip veriyi sil
	delete(states, "ANK")
	fmt.Println(states)

	states["ERZ"] = "Erzurum"

	for k, v := range states {
		fmt.Printf("%v: %v\n", k, v)
	}

	//states map nesnesinin elemanı adedince kapasiteli bir key listesi oluştur
	keys := make([]string, len(states))
	i := 0
	for k := range states {
		keys[i] = k
		i++
	}
	fmt.Println(keys)

	// key listesindeki index değerlerine göre states nesnesinde bulunan şehileri yazdır
	for i := range keys {
		fmt.Println(states[keys[i]])
	}
}
