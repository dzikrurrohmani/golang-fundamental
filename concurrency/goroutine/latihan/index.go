package main

import (
	"fmt"
	"sync"
	"time"
)

type PenjualanToko struct {
	KodeToko  string
	Penjualan []float64
}

func main() {
	var wg sync.WaitGroup
	// sample data
	datatoko := []PenjualanToko{
		{"123", []float64{1, 2, 3, 4, 5}},
		{"124", []float64{13, 21, 33}},
		{"125", []float64{4, 3}},
		{"126", []float64{21, 22, 23, 9, 1, 22}},
		{"127", []float64{8, 4, 10}},
	}
	for _, data := range datatoko {
		wg.Add(1)
		go func(dataNya PenjualanToko) {
			defer wg.Done()
			HitungPenjualan(dataNya)
		}(data)
	}
	wg.Wait()
}

func HitungPenjualan(n PenjualanToko) {
	var penjualan float64
	for _, penjualanke := range n.Penjualan {
		penjualan += penjualanke
		time.Sleep(time.Second)
	}
	fmt.Printf("Penjualan Toko-%s sebesar %.2f\n", n.KodeToko, penjualan)
}
