package main

import (
	"fmt"
	"sort"
)

func main() {
	mapScore := make(map[string]int)

	// isi vallue
	// mapScore["key"] = value

	mapScore["Anggar"] = 75
	mapScore["Adise"] = 70
	mapScore["Roi"] = 85
	fmt.Println(mapScore)
	fmt.Println(mapScore["Anggar"]) // 75
	fmt.Println(mapScore["anggar"]) // 0 karena case sensitive
	fmt.Printf("maps: %T\n", mapScore)
	fmt.Printf("maps: %v\n", mapScore)
	fmt.Printf("maps: %#v\n", mapScore)

	// pembuatan maps menggunakan var
	var month = map[int]string{
		1: "Januari",
		2: "Februari",
		3: "Maret",
		4: "April",
		5: "Mei",
	}

	for key, value := range month {
		fmt.Println("Bulan ke-", key, " adalah ", value)
	}

	// SORTING USE:
	/*
	m := map[string]int{"Alice": 23, "Eve": 2, "Bob": 25}

	keys := make([]string, 0, len(m))
	for k := range m {
		keys = append(keys, k)
	}
	sort.Strings(keys)

	for _, k := range keys {
		fmt.Println(k, m[k])
	}
	source: https://yourbasic.org/golang/sort-map-keys-values/
	*/
	keyList := make([]int, 0, len(month))
	for k := range month {
		keyList = append(keyList, k)
	}
	sort.Ints(keyList)

	for _, k := range keyList {
		fmt.Println(k, month[k])
	}

	// Menghapus elemen dalam map menggunakan key
	// Menggunakan keyword delete()
	fmt.Println("before deleted: ", month)
	delete(month, 5)
	fmt.Println("after deleted: ", month)
	fmt.Printf("%T\n",month[5])

	// Cek keberadaan kkey apakah ada atau tidak dalam satu map
	// Ketika memanggil sebuah map diperoleh return value (value, dan optional value (bool))
	value, isExist := month[5]

	if isExist {
		fmt.Println(value)
	} else {
		fmt.Println("Item is not exist")
	}
	// bisa dipersingkat menjadi >

	if value, isExist := month[4]; isExist {
		fmt.Println(value)
	} else {
		fmt.Println("Item is not exist")
	}

	// Kombinasi map dan slice
	var students = []map[string]string{
		{"ID": "S001", "Name": "Baim", "Address": "Jakarta"},
		{"ID": "S002", "Name": "Paula", "Address": "Jakarta"},
		{"ID": "S003", "Name": "Kiano", "Address": "Tangerang"},
	}

	for _, student := range students {
		fmt.Println(student["Name"])
	}
}
