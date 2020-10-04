package main

import (
	"log"
	"os"
)

//Dosya Oluşturma
/*
var (
	newFile *os.File
	err     error
)

func main() {
	newFile, err = os.Create("demo.txt")
	if err != nil {
		log.Fatal(err)
	}

}
*/

//Dosya Bilgisi Almak(Get File Info)
/*
var (
	fileInfo *os.FileInfo
	err      error
)

func main() {
	//Dosya bilgisini döndürür.
	//Eğer dosya yoksa hata döner.
	fileInfo, err := os.Stat("demo.txt")
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("File Name : ", fileInfo.Name())
	fmt.Println("Size in bytes : ", fileInfo.Size())
	fmt.Println("Permission : ", fileInfo.Mode())
	fmt.Println("Last Modified : ", fileInfo.ModTime())
	fmt.Println("Is Dictionary : ", fileInfo.IsDir())
	fmt.Println("System Interface Type : ", fileInfo.Sys())
	fmt.Printf("System Info : %+v\n\n", fileInfo.Sys())

}
*/

//Yeniden İsimlendirme ve Taşıma(Rename and Move a File)
/*
func main() {
	originalPath := "demo.txt"
	newPath := "./moved/test.txt"
	err := os.Rename(originalPath, newPath)
	if err != nil {
		log.Fatal(err)
	}
}
*/

//Dosyanın var olup olmadığını kontrol etmke(Check if File Exists)
/*
var (
	fileInfo *os.FileInfo
	err      error
)

func main() {
	fileInfo, err := os.Stat("demo.txt")
	if err != nil {
		if os.IsNotExist(err) {
			log.Fatal("File does not exists")
		}
	}
	log.Println("File does exists. File Information : ")
	log.Println(fileInfo)
}
*/

//Dosyaları Açma ve Kapama(Open and Cloes Files)

//Salt okunur açma
/*
func main() {
	file, err := os.Open("demo.txt")
	defer file.Close()
	if err != nil {
		log.Fatal(err)
	}
	//üzerine veri yazdık
}
*/

/*
func main() {

	/*
		OpenFile() ikinci parametre tipleri;

		os.O_RDONLY		: Sadece okuma
		os.O_WRONLY		: Sadece yazma
		os.O_RDWR		: Okuma ve yazma yapılabilir
		os.O_APPEND		: Dosyanın sonuna ekle
		os.O_CREATE		: Dosya yoksa oluştur
		os.O_TRUNC		: Açılırken dosyayı kes

		Bu ayarlar birden fazla olarak da kullanılabilir

		-> os.O_CREATE|os.O_APPEND
		-> os.O_CREATE|os.O_TRUNC|os.O_WRONLY

	* /

	//İkinci parametre dosya açılış amacını ayarlar
	//Üçüncü parametre dosya izinlerini belirler.
	file, err := os.OpenFile("demo.txt", os.O_APPEND, 0666)
	defer file.Close()
	if err != nil {
		log.Fatal(err)
	}

}
*/

//Okuma ve Yazma İzinlerini Kontrol Etmek (Check Read and Write Permissions)
/*
func main() {
	//Yazma izinlerini test
	fileWr, errWr := os.OpenFile("demo.txt", os.O_WRONLY, 0666)
	defer fileWr.Close()
	if errWr != nil {
		if os.IsPermission(errWr) {
			log.Println("Hata : Yazma izni reddedildi.")
		} else {
			log.Fatal(errWr)
		}
	}

	//*************

	//Okuma izinlerini test
	fileR, errR := os.OpenFile("demo.txt", os.O_RDONLY, 0666)
	defer fileR.Close()
	if errR != nil {
		if os.IsPermission(errR) {
			log.Println("Hata : Okuma izni reddedildi.")
		} else {
			log.Fatal(errR)
		}
	}

}
*/

//Dosya Kopyalama (Copy a file)
/*
func main() {
	originalFile, err := os.Open("demo.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer originalFile.Close()

	//Yeni bir dosya oluştur

	newFile, errC := os.Create("./folder/demo_copy.txt")
	if errC != nil {
		log.Fatal(errC)
	}
	defer newFile.Close()

	//Byte'ları kaynaktan hedefe koppyala
	byteWritten, errI := io.Copy(newFile, originalFile)
	if errI != nil {
		log.Fatal(errI)
	}
	log.Printf("Copied %d bytes.", byteWritten)

	//Dosya içeriğini işle
	//Belleği diske boşalt
	errI = newFile.Sync()
	if errI != nil {
		log.Fatal(errI)
	}
}
*/

//Byte'ları Bir Dosyaya Yazın(Write Bytes to a File)
/*
func main() {
	//Demo.txt dosyasını sadece yazılabilir bir dosya olarak aç
	file, err := os.OpenFile(
		"demo.txt",
		os.O_WRONLY,
		0666)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	byteSlice := []byte("Bu dosyaya yazdık!\n")
	bytesWritten, errC := file.Write(byteSlice)
	if errC != nil {
		log.Fatal(errC)
	}
	log.Printf("Wrote %d bytes.\n", bytesWritten)
}
*/

//Hızlı Dosya Yazma (Quick Write to File)
/*
func main() {
	err := ioutil.WriteFile("demo.txt", []byte("Merhaba!"), 0666)
	if err != nil {
		log.Fatal(err)
	}
}
*/

//Geçici Dosya ve Klasörler ile Çalışmak
/*
func main() {
	//gecici klasor olustur
	tempDirPath, err := ioutil.TempDir("", "geciciKlasor")
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Geçici klasör oluşturuldu: ", tempDirPath)

	// gecici dosya olustur
	tempFile, errF := ioutil.TempFile(tempDirPath, "geciciDosya.txt")
	if errF != nil {
		log.Fatal(errF)
	}

	fmt.Println("Geçici dosya oluşturuldu: ", tempFile.Name())

	// close file
	errF = tempFile.Close()
	if errF != nil {
		log.Fatal(errF)
	}

	// Silme

	errF = os.Remove(tempFile.Name())
	if errF != nil {
		log.Fatal(errF)
	} else {
		fmt.Printf("%s dosyası silindi.", tempFile.Name())
	}

	err = os.Remove(tempDirPath)
	if err != nil {
		log.Fatal(err)
	} else {
		fmt.Printf("%s klasörü silindi.", tempDirPath)
	}
}
*/

//Dosya Silmek (Delete a File)
func main() {
	err := os.Remove("demo.txt")
	if err != nil {
		log.Fatal(err)
	}
}
