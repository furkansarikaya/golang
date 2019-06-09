package main

import (
	"fmt"
)

func main() {
	//for

	for i := 0; i < 5; i++ {
		fmt.Println("Value : ", i)
	}

	//for(while yöntemi)

	sum := 1
	for sum < 10 {
		sum += sum
		fmt.Println(sum)
	}

	//range

	/*
		for döngüsünün range ("aralık") formu ile bir dilim ("slice") veya eşlem ("map") üzerinde dolaşılır.
		Bir dilim üzerinde dolaşılırken her seferinde iki değer dönülür. İlki indis, ikincisi o indisli elemanın bir kopyası.

		Değer dizisinin indislerini veya o indise karşılık gelen değerleri _ ataması ile atlayabilirsiniz.
		Yalnızca indisi kullanmak istiyorsanız ", value" kısmını tamamen çıkarınız.
	*/

	/*
		Furkan not : c# daki foreach gibi kullanılabilir
	*/

	var pow = []int{1, 2, 4, 8, 16, 32, 64, 128}
	for i, v := range pow {
		fmt.Printf("2**%d = %d\n", i, v)
	}

	//array
	a := [...]string{"a", "b", "c", "ç"}
	for i := range a {
		fmt.Println("Array item", i, "is", a[i])
	}

	for i, v := range a {
		fmt.Println("Array item", i, "is", v)
	}

	//map
	baskentler := map[string]string{"Turkey": "Ankara", "France": "Paris", "Italy": "Roma", "Japan": "Tokyo"}
	for key := range baskentler {
		fmt.Println("Map item: Capital of", key, "is", baskentler[key])
	}

	for key, val := range baskentler {
		fmt.Println("Map item: Capital of", key, "is", val)
	}
}
