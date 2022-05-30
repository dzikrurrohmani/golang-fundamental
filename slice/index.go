package main

import "fmt"

func main() {
	/*
		Slice -> sama seperti array, bedanya tidak perlu definisikan jumlah elemen
		Slice -> merupakan reference
	*/

	var person1 = [2]string{"Antonio", "Christiano"}
	var person2 = [...]string{"Antonio", "Christiano"}
	var person3 = []string{"Antonio", "Christiano"}
	// person3 = append(person3, "saya")

	fmt.Printf("[person1: %T dengan jumlah elemen %d\n", person1, len(person1))
	fmt.Printf("[person2: %T dengan jumlah elemen %d\n", person2, len(person2))
	fmt.Printf("[person3: %T dengan jumlah elemen %d\n", person3, len(person3))

	// Variasi untuk mengambil nilai
	var personArray = [2]string{"Antonio", "Christiano"}
	var personSlice = []string{"Antonio", "Christiano"}

	fmt.Println(personArray[0:2])
	fmt.Println(personSlice)
	fmt.Println(personSlice[:1])

	// len&cap
	// len -> menghitung jumlah elemen slice
	// cap -> menghitung lebar atau kapasitas slice
	var fruits = []string{"Mangga", "Apel", "Jeruk", "Alpukat"}
	fmt.Println(fruits)
	fmt.Println("len: ", len(fruits)) //4
	fmt.Println("cap: ", cap(fruits)) //4

	getFruits1 := fruits[0:2]
	fmt.Println(getFruits1)
	fmt.Println("len: ", len(getFruits1)) //2
	fmt.Println("cap: ", cap(getFruits1)) //4
	fmt.Printf("%p\n", &getFruits1[0])

	// append() menambahkan elemen pada slice
	// elemen yang ditambahkan, posisinya di akhir
	// jumlah kapasitas memenuhi, pasti masuk ke array yang mempunyai referensi yang sama (lama)
	// jumlah kapasitas melebihi, otomatis yang masuk akan dijadikan sebuah referensi baru
	// ex: cap: 4, append sampai 5 sehingga dianggap sebuah array baru

	getFruits1 = append(getFruits1, "haha", "hihi", "huhu")
	fmt.Println(getFruits1)
	fmt.Println("len: ", len(getFruits1)) //2
	fmt.Println("cap: ", cap(getFruits1)) //4
	fmt.Printf("%p\n", &getFruits1[0])

	getFruits2 := fruits[1:3]
	fmt.Println(getFruits2)
	fmt.Println("len: ", len(getFruits2)) //2
	fmt.Println("cap: ", cap(getFruits2)) //3	
	fmt.Printf("%p\n", &getFruits2[0])

	getFruits2 = append(getFruits2, "haha")
	fmt.Println(getFruits2)
	fmt.Println("len: ", len(getFruits2)) //3
	fmt.Println("cap: ", cap(getFruits2)) //3
	fmt.Printf("%p\n", &getFruits2[0])

	getFruits2 = append(getFruits2, "haha", "hihi")
	fmt.Println(getFruits2)
	fmt.Println("len: ", len(getFruits2)) //6
	fmt.Println("cap: ", cap(getFruits2)) //6
	fmt.Printf("%p\n", &getFruits2[0])

	getFruits2 = append(getFruits2, "haha", "hihi", "haha", "hihi", "haha", "hihi", "haha", "hihi")
	fmt.Println(getFruits2)
	fmt.Println("len: ", len(getFruits2)) //6
	fmt.Println("cap: ", cap(getFruits2)) //6
	fmt.Printf("%p\n", &getFruits2[0])

	// make -> membuat slice kosong dengan spesifik length
	cities := make([]string, 3)
	fmt.Printf("cities: %T\n", cities)

	cities[0] = "Jakarta Selatan"
	cities[1] = "Jakarta Pusat"
	cities[2] = "Jakarta Barat"

	fmt.Printf("%q\n", cities)
}