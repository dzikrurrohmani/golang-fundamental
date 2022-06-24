package main

import (
	"fmt"
	"sync"
	"time"
)

func main() {
	/*
		Channel : jembatan komunikasi antar goroutine
		Channel dapat sebagai pengirim dan penerima (<-)
		Ilustrasinya:
		Pengirim: channel <- value
		Penerima: value := <- channel

		Cara buatnya, sama seperti pembuatan map, slice, menggunakan keyword make(chan datatype)
		Ex: message := make(chan string) -> sebuah channel dengan tipe data string.
	*/
	// test1()
	// test2()
	// test3()
	test5()
}

func test1() {
	message := make(chan string)

	sayHi := func(name string) {
		v := fmt.Sprintf("Hello %s", name)
		message <- v
	}

	go sayHi("Bulan")
	v := <-message
	fmt.Println(v)
}

func test2() {
	message := make(chan string)

	sayHi := func(name string) {
		v := <- message
		fmt.Println(v)
	}

	go sayHi("Bulan")
	message <- "pesan"
}

func test3() {
	var wg sync.WaitGroup
	message := make(chan string)

	wg.Add(1)
	go sayHi3(message, &wg)
	message <- "pesan"
	// message <- "pesan"
}

func sayHi3(n chan string, wg *sync.WaitGroup) {
	defer wg.Done()
	time.Sleep(time.Second)
	v := <- n
	fmt.Println(v)
}

func test4() {
	message := make(chan string)

	go func ()  {
		message <- "Joko"
	}()

	for {
		isDone := sayHi4(message)
		if isDone {
			break
		}
	}
}

func sayHi4(n chan string) bool {
	time.Sleep(time.Second)
	v := <- n
	fmt.Println(v)
	return true
}

// for range close
func test5() {
	message := make(chan string)

	go func ()  {
		defer close(message)
		message <- "Joko"
	}()

	for n := range message {
		sayHi5(n)
	}
}

func sayHi5(m string) {
	fmt.Println(m)
}
