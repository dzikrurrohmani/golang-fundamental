package main

import "fmt"

func main() {
	/*
	Array -> kumpulan data dengan tipe data yang sama
	Array -> harus didefinisikan alokasinya
	*/

	// var fruits [4]string

	// fruits[0] = "Mangga"
	// fruits[1] = "Apel"
	// fruits[2] = "Jeruk"
	// fruits[3] = "Alpukat"

	// // Cara memanggilnya / cetak
	// fmt.Println(fruits)
	// fmt.Printf("%s %s %s %s\n", fruits[0], fruits[1], fruits[2], fruits[3])

	// Inisialisasi array di awal
	// fruits := [4]string{"Mangga", "Apel", "Jeruk", "Alpukat"}
	// atau
	// var fruits = [4]string{"Mangga", "Apel", "Jeruk", "Alpukat"}
	var fruits = [4]string{"Mangga", "Apel", "Jeruk", "Alpukat"}
	fmt.Println(fruits)
	fmt.Printf("%q\n", fruits)
	fmt.Printf("%s %s %s %s\n", fruits[0], fruits[1], fruits[2], fruits[3])

	/*
	Inisialisasi value:
	- Horizontal -> ex: fruits := [4]string{"Mangga", "Apel", "Jeruk", "Alpukat"}
	- Vertical -> ex: fruits := [4]string{
										"Mangga",
										"Apel",
										"Jeruk",
										"Alpukat",
									}
	*/

	// Contoh deklarasi horizontal dan vertical
	fruitshor := [4]string{"Mangga", "Apel", "Jeruk", "Alpukat"}
	fruitsver := [4]string{
		"Mangga",
		"Apel",
		"Jeruk",
		"Alpukat",
	}
	fmt.Printf("%q\n%q\n", fruitshor, fruitsver)

	// Inisialisasi tanpa diketahui (tak terbatas) jumlah elemen -> [...]
	var days = [...]string{"Senin", "Selasa", "Rabu", "Kamis"}
	fmt.Println(len(days))
	fmt.Printf("%q\n", days)

	var months = [...]string{"Januari", "Februari",}
	fmt.Println(len(months))
	fmt.Printf("%q\n", months)

	// Multidimentional Array -> [baris][kolom]
	var numbers = [2][3]int{
		{1,2,3},
		{4,5,6},
	}
	fmt.Println(numbers)
	fmt.Println(len(numbers))
	fmt.Println(len(numbers[0]))

	// Looping Array
	var students = [...]string{
		"Jack",
		"Milla",
		"Thomas",
		"Anna",
	}
	// Basic for
	for i := 0; i < len(students); i++ {
		fmt.Printf("Mahasiswa ke - %d: %s \n", i, students[i])
	}

	// For range
	for i, student := range students {
		fmt.Printf("Mahasiswa ke - %d: %s \n", i, student)
	}
	fmt.Printf("%T %T %T %T\n", fruits, days, months, students)
	fmt.Printf("%T %T %T %T\n", fruits[0], days[0], months[0], students[0])
	
	for i, number := range numbers {
		for j, num := range number {
			fmt.Printf("Number ke - [%d][%d]: %d \n", i, j, num)
		}
	}

	/*
	Bisa ngga seperti ini?
	For i, {a,b,c} := 
	*/
}