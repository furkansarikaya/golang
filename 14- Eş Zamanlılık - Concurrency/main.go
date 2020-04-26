package main

import (
	"runtime"
	"time"
)

// description : https://go-tour-turkish.appspot.com/concurrency/1
// blog : https://go-talks.appspot.com/github.com/cihanozhan/talks-golang/go-giris/giris.slide#50

// func main() {
// 	//basit bir gorouting kullanımı
// 	// go xFunc()
// 	// time.Sleep(100 * time.Millisecond)

// 	runtime.GOMAXPROCS(1)
// 	go xFunc()
// 	time.Sleep(100 * time.Millisecond)
// }

// // func xFunc() {
// // 	for l := byte('a'); l <= byte('z'); l++ {
// // 		println(string(l))
// // 	}
// // }

// func xFunc() {
// 	for l := byte('a'); l <= byte('z'); l++ {
// 		go println(string(l))
// 	}
// }

// Kanallar(channels)

// func sum(s []int, c chan int) {
// 	sum := 0
// 	for _, v := range s {
// 		sum += v
// 	}
// 	c <- sum
// }

// func main() {
// 	s := []int{1, 2, 3, 4, 5, 6}

// 	c := make(chan int)
// 	go sum(s[:len(s)], c)
// 	go sum(s[:len(s)/2], c)
// 	x, y := <-c, <-c

// 	fmt.Println(x, y)
// }

func main() {
	runtime.GOMAXPROCS(8)
	ch := make(chan string)
	go xFunc(ch)
	go printer(ch)
	time.Sleep(100 * time.Millisecond)
}

func xFunc(ch chan string) {
	for l := byte('a'); l <= byte('z'); l++ {
		ch <- string(l)
	}
}

func printer(ch chan string) {
	for {
		println(<-ch)
	}
}
