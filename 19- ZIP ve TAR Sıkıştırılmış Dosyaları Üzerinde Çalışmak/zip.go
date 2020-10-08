package main

import (
	"archive/zip"
	"io"
	"log"
	"os"
	"path/filepath"
	"strings"
)

// Dosyaları ZIP dosyası olarak sıkıştırmak.
// ZIP dosyası üretmek

/*
var fileFolderPath = "files\\"
var files = []string{fileFolderPath + "demo.go", fileFolderPath + "note1.txt"}

func addFile(fileName string, zw *zip.Writer) error {
	file, err := os.Open(fileName)
	if err != nil {
		return fmt.Errorf("Dosya açılırken bir hata meydana geldi (%s): %s", fileName, err)
	}
	defer file.Close()

	wr, err := zw.Create(fileName)
	if err != nil {
		msg := "(%s) ZIP dosyası içerisine yeni bir dosya oluşturulurken hata meydana geldi : %s"
		return fmt.Errorf(msg, fileName, err)
	}
	if _, err := io.Copy(wr, file); err != nil {
		return fmt.Errorf("%s dosyası ZIP dosyasına yazılırken bir hata meydana geldi : %s", fileName, err)
	}

	return nil
}

func createArchiveZIPFile(archiveFileName string) int {
	if len(archiveFileName) == 0 {
		return -1
	}

	flags := os.O_WRONLY | os.O_CREATE | os.O_TRUNC
	file, err := os.OpenFile(archiveFileName+".zip", flags, 0644)
	if err != nil {
		log.Fatalf("ZIP dosyasına veri yazmak için açılırken bir hata meydana geldi : %s", err)
		return -1
	}
	defer file.Close()

	zw := zip.NewWriter(file)
	defer zw.Close()

	for _, fileMame := range files {
		if err := addFile(fileMame, zw); err != nil {
			log.Fatalf("%s dosyası ZIP dosyasına eklenirken bir hata meydana geldi : %s", fileMame, err)
		}
	}
	return 1

}

func main() {
	result := createArchiveZIPFile("dosya")
	if result > 0 {
		fmt.Println("İşlem başarılı : ", result)
	} else {
		fmt.Println("İşlem başarısız oldu : ", result)
	}
}
*/

// Arşivlenmiş Dosyaları Dışarı Çıkarmak

func CreateDirIfNotExits(dir string) {
	if _, err := os.Stat(dir); os.IsNotExist(err) {
		err := os.MkdirAll(dir, 0755)
		if err != nil {
			log.Fatal(err)
		}
	}

}

func main() {
	zr, err := zip.OpenReader("dosya.zip")
	if err != nil {
		log.Fatal(err)
	}
	defer zr.Close()

	for _, file := range zr.Reader.File {
		zippedFile, err := file.Open()
		if err != nil {
			log.Fatal(err)
		}
		defer zippedFile.Close()

		targetDir := "./"
		extractedFilePath := filepath.Join(
			targetDir,
			file.Name,
		)

		dirName := strings.Split(file.Name, "\\")

		CreateDirIfNotExits(dirName[0])

		if file.FileInfo().IsDir() {
			log.Println("Klasör oluşturuluyor: ", extractedFilePath)
			os.MkdirAll(extractedFilePath, file.Mode())
		} else {
			log.Println("Dosya çıkarılıyor: ", file.Name)

			outFile, err := os.OpenFile(
				extractedFilePath,
				os.O_WRONLY|os.O_CREATE|os.O_TRUNC,
				file.Mode(),
			)
			if err != nil {
				log.Fatal(err)
			}
			defer outFile.Close()

			_, err = io.Copy(outFile, zippedFile)
			if err != nil {
				log.Fatal(err)
			}
		}
	}
}
