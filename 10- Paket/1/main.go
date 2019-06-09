package main

import (
	"fmt"
	"strings"

	"github.com/fatih/color"
	log "github.com/koding/logging"
	//"github.com/cihanozcan/logging"
)

//Paketler
/*
	Paket Nedir?

	Bir araya toplanmış(bundled) fonksiyonlar topluluğudur. Farklı bir çok nesneyi de kendi içinde barındırabilir.

	- Her Go programı paketlerden oluşur
	- Programlar main paketlerinde çalışmaya başlar

	Faydaları Nelerdir?

	- Namespace, yani ilgili fonkisyonları kapsayan ortak isim uzayı sağlar ve proje ya da sistemdeki diğer fonksiyonlarla karışmamayı sağlar.
	- Kod organizasyonu sağlar.
	- Derleyiciyi hızlandırır. Örneğin fmt her kullanışımızda derleyici tarafından derlenmez.

*/
func main() {

	//strings
	fmt.Println(strings.Contains("test", "es"))
	fmt.Println(strings.Count("test", "t"))
	fmt.Println(strings.HasPrefix("unit_test", "unit")) //_ öncesini prefix olarak algılar ve unit mi olarak bakıyor true döner
	fmt.Println(strings.HasSuffix("dosya.rar", "rar"))  // . sonrasını suffix olarak algılar ve rar mı olarak bakıyor true döner
	fmt.Println(strings.Index("test", "e"))

	color.Red("Error Message")

	log.Info("Uygulama sonu")
}
