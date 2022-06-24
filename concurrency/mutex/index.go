package main

import (
	"fmt"
	"sync"
)

/*
Program kita mengalami reace condition
race condition = kondisi dimana 2 atau lebih goroutine mengakses data yang sama secara bersamaan
efek: perhitungan kacau, or salah

Go menyediakan solusi yaitu "mutex"
dengan mutex: ketika terjadi rc, hanya goroutine yang beruntung yang dapat mengakses data tsb.
kemudian goroutine yang lain menunggu (waiting status) sampai mendapatkan data tersebut.

Go: package yang sama seperti waitGroup yaitu sync.Mutex
*/

type counter struct {
	val int
}

func (c *counter) Add(x int) {
	c.val = c.val + x
}

func (c *counter) Value() int {
	return c.val
}

func print(till int, message *counter, wg *sync.WaitGroup, mtx *sync.Mutex) {
	defer wg.Done()
	for i := 0; i < till; i++ {
		mtx.Lock()
		message.Add(1)
		mtx.Unlock()
	}
}

func main() {
	var wg sync.WaitGroup
	var mtx sync.Mutex
	var meter counter

	till := 100
	for i := 0; i < till; i++ {
		wg.Add(1)
		go print(till, &meter, &wg, &mtx)
	}

	wg.Wait()
	fmt.Println(meter.Value())
}
