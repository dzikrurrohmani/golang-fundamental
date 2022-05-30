package main

import "fmt"

func main() {
	/*
		Buat program map di dalam map berdasarkan jenis mobil dan cc.
		nama map toyota,
		jenis: mobil MVP, SUV, Sedan, dll
		cc : ->
			3000 -> Innova, Fortuner
			2800 -> Raize, Avanza
			2500 -> Camry, Rush
			2000 -> dll.
	*/

	toyota := map[string]map[int][]string{
		"MVP": {3000: {"Innova", "Fortuner"}, 2800: {"Raize", "Avanza"}, 2500: {"Camry", "Rush"}},
		"SUV": {3000: {"Innova", "Fortuner"}, 2800: {"Raize", "Avanza"}, 2500: {"Camry", "Rush"}},
		"Sedan": {3000: {"Innova", "Fortuner"}, 2800: {"Raize", "Avanza"}, 2500: {"Camry", "Rush"}},
	}
	fmt.Println(toyota["MVP"][3000])
	fmt.Println(toyota["SUV"][2800])
	fmt.Println(toyota["Sedan"][2500])
}