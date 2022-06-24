package main

import (
	"fmt"
	"sync"
)

/*
Contoh penggunaan channel for - range - close
*/

func main() {
	test3()
}

func test1() {
	squareChan := make(chan int)
	numbers := []int{2, 3, 4, 5, 6, 7, 8}
	go square(numbers, squareChan)
	// for - range
	for i := range squareChan {
		fmt.Println("Square result:", i)
	}

	cubedChan := make(chan int)
	go cubed(numbers, cubedChan)
	// for - range
	for i := range cubedChan {
		fmt.Println("Cubed result:", i)
	}
}

func square(numbers []int, ch chan<- int) {
	for _, v := range numbers {
		ch <- v * v
		fmt.Println("Square Berhasil mengirim data", v, "ke channel")
	}

	close(ch)
}

func cubed(numbers []int, ch chan<- int) {
	for _, v := range numbers {
		ch <- v * v * v
		fmt.Println("Cubed Berhasil mengirim data", v, "ke channel")
	}

	close(ch)
}

func test2() {
	var wg sync.WaitGroup

	squareChan := make(chan int)
	numbers := []int{2, 3, 4, 5, 6, 7, 8}
	wg.Add(1)
	go square2(numbers, squareChan, &wg)
	// for - range
	go func() {
		// defer wg.Done()
		for i := range squareChan {
			fmt.Println("Square result:", i)
		}
	}()

	cubedChan := make(chan int)
	wg.Add(1)
	go cubed2(numbers, cubedChan, &wg)
	// for - range
	go func() {
		// wg.Done()
		for i := range cubedChan {
			fmt.Println("Cubed result:", i)
		}
	}()

	wg.Wait()
}

// channel di set hanya untuk mengirim saja (channel direction): ch chan <- int
func square2(numbers []int, ch chan<- int, wg *sync.WaitGroup) {
	defer wg.Done()
	for _, v := range numbers {
		ch <- v * v
		fmt.Println("Square Berhasil mengirim data", v, "ke channel")
	}

	close(ch)
}

func cubed2(numbers []int, ch chan<- int, wg *sync.WaitGroup) {
	defer wg.Done()
	for _, v := range numbers {
		ch <- v * v * v
		fmt.Println("Cubed Berhasil mengirim data", v, "ke channel")
	}

	close(ch)
}

func test3() {
	var wg sync.WaitGroup

	squareChan := make(chan int)
	cubedChan := make(chan int)
	doneChan := make(chan bool)
	numbers := []int{2, 3, 4, 5, 6, 7, 8}
	wg.Add(2)
	go square3(numbers, squareChan, &wg)
	go cubed3(numbers, squareChan, &wg)

	go func() {
		wg.Wait()
		doneChan <- true
	}()

	for {
		select {
		case msg1 := <-squareChan:
			fmt.Println("Square result:", msg1)
		case msg2 := <-cubedChan:
			fmt.Println("Square result:", msg2)
		case <-doneChan:
			fmt.Println("Done")
			return
		}
	}
}

// channel di set hanya untuk mengirim saja (channel direction): ch chan <- int
func square3(numbers []int, ch chan<- int, wg *sync.WaitGroup) {
	defer wg.Done()
	for _, v := range numbers {
		ch <- v * v
		fmt.Println("Square Berhasil mengirim data", v, "ke channel")
	}
}

func cubed3(numbers []int, ch chan<- int, wg *sync.WaitGroup) {
	defer wg.Done()
	for _, v := range numbers {
		ch <- v * v * v
		fmt.Println("Cubed Berhasil mengirim data", v, "ke channel")
	}
}
