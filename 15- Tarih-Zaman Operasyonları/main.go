package main

import (
	"fmt"
	"time"
)

// func main() {
// 	fmt.Printf("Şuanki Unix Zamanı: %v\n", time.Now().Unix())
// 	time.Sleep(2 * time.Second)
// 	fmt.Printf("Şuanki Unix Zamanı: %v\n", time.Now().Unix())
// }

// Tarih ve Zaman Operasyonları

// func main() {
// 	t := time.Date(2018, time.November, 10, 20, 0, 0, 0, time.UTC)
// 	fmt.Printf("Çalışıyor: %s\n", t)

// 	fmt.Println("*****************")

// 	now := time.Now()
// 	fmt.Printf("Mevcut zaman: %s\n", now)

// 	fmt.Println("*****************")

// 	fmt.Println("Ay: ", now.Month())
// 	fmt.Println("Gün: ", now.Day())
// 	fmt.Println("Haftanın Günü: ", now.Weekday())

// 	fmt.Println("*****************")

// 	tomorrow := now.AddDate(0, 0, 1)
// 	fmt.Printf("Tomorrow is %v, %v,%v,%v\n", tomorrow.Weekday(), tomorrow.Month(), tomorrow.Day(), tomorrow.Year())

// 	fmt.Println("*****************")

// 	longFormat := "Monday, January 2,2018"
// 	fmt.Println("Tomorrow is", tomorrow.Format(longFormat))

// 	fmt.Println("*****************")

// 	shortFormat := "1/2/2020"
// 	fmt.Println("Tomorrow is", tomorrow.Format(shortFormat))

// 	fmt.Println("*****************")
// }

func main() {
	x := fmt.Println

	xTime := time.Date(1071, 10, 11, 20, 23, 0, 0, time.UTC)

	now := time.Now()
	x(now)

	x("------------------")

	x(now.Year())
	x(now.Month())
	x(now.Day())
	x(now.Hour())
	x(now.Minute())
	x(now.Second())
	x(now.Nanosecond())
	x(now.Location())

	x(now.Weekday())

	// Tarih karşılaştırma ya da kontrol ya da test

	x(xTime.Before(now))
	x(xTime.After(now))
	x(xTime.Equal(now))

	diff := now.Sub(xTime)
	x(diff)

	x(diff.Hours())
	x(diff.Minutes())
	x(diff.Seconds())
	x(diff.Nanoseconds())
}
