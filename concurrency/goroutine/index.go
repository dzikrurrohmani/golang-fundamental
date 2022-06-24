// package main

// import (
// 	"fmt"
// 	"time"
// )

// func main() {
// 	test2()

// }

// func test1() {
// 	fmt.Println("Start")
// 	greeting("jution")
// 	fmt.Println("End")
// }

// func test2() {
// 	fmt.Println("Start")
// 	go greeting("jution")
// 	fmt.Println("End")
// 	time.Sleep(5 * time.Second)
// }

// func greeting(name string) {
// 	fmt.Println("Hello", name)
// }

package main

import (
	"fmt"
	"sync"
	"time"
)

func main() {

	// fmt.Println("Start")

	// goroutine
	// implementasi cukuo simpel, dengan menambahkan keyword go didepan function
	// unpredictable
	// -------------------------------------
	// var wg sync.WaitGroup

	// wg.Add(1)
	// go greeting(&wg, "Utsman")
	// wg.Add(1)
	// go sayHello(&wg)
	// fmt.Println("End")

	// time.Sleep(100 * time.Microsecond)
	// wg.Wait()

	// -------------------------------------
	var wg sync.WaitGroup

	wg.Add(1)
	go func() {
		defer wg.Done()
		greeting("Utsman")
	}()
	wg.Add(1)
	go func() {
		defer wg.Done()
		sayHello()
	}()
	fmt.Println("End")

	time.Sleep(100 * time.Microsecond)
	wg.Wait()
}

func greeting(name string) {
	fmt.Println("Start greeting func")
	fmt.Println("Hello", name)
	time.Sleep(1 * time.Second)
	fmt.Println("End greeting func")
}

func sayHello() {
	time.Sleep(1 * time.Second)
	fmt.Println("Hai...")
}

/*
	- goroutine bisa dikatakan dapat menjalankan beberapa proses secara bersamaan
	- Lantas bagaimana jika kita ingin menunggu beberapa proses tsb ?
	- kita bisa menggunakan waitGroup untuk menunggu goroutine
	- waitgroup itu sendiri sudah disediakan oleh packageGo
	- package go.Sync.WaitGroup
*/
