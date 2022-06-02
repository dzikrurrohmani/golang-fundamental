// package main

// import (
// 	"fmt"
// )

// func main() {
// 	// Bisa digunakan berkali2 agar tidak mengulang code
// 	fmt.Println(add(3,5))
// 	fmt.Println(add(4,5))
// 	fmt.Println(add(5,5))

// 	printMessage()
// 	greeting()
// 	fmt.Println("")

// 	message, integer := Message("Dzikrur", "Sragen")
// 	fmt.Println(message)
// 	fmt.Println("saya", "mudik", integer, "kali")

// 	oddEven(20) // GENAP
// 	oddEven(0) // Invalid
// 	oddEven(11) // GANJIL

// 	var animals = []string{"Ayam", "Bebek", "Kucing", "Merpati"}
// 	enigmaZoo(1, animals, 0)
// 	enigmaZooVariadic(2, animals...)
// 	enigmaZooVariadic(3, "Ayam", "Bebek", "Kucing", "Merpati")

// 	//CLOSURE
// 	printGreetingFunc := greeting2
// 	fmt.Println(printGreetingFunc("Jution"))

// 	/*
// 	Anonynmous function > Closure too
// 	-> Function yang tidak didefinisikan nama functionnya
// 	-> Disimpan dalam sebuah variable
// 	-> Pembuatannya bisa dilakukan dalam function main
// 	*/
// 	var greeting3 = func ()  {
// 		fmt.Println("Hai salam kenal...!")
// 	}
// 	greeting3()
// }

// func printMessage() {
// 	fmt.Println("Halo dunia...")
// }

// func greeting() {
// 	fmt.Println("Salam kenal...")
// }

// func add(param1, param2 int) int {
// 	return param1+param2-7
// }

// func Message(name, address string) (string, int) {
// 	return fmt.Sprintf("Halo nama saya %s, saya tinggal di %s", name, address), 1
// }

// func oddEven(number int) {
// 	if number == 0 {
// 		fmt.Println("Invalid number, can't zero number")
// 		return
// 	}
// 	if number%2 == 0 {
// 		fmt.Printf("Angka %d merupakan bilangan GENAP\n", number)
// 	} else {
// 		fmt.Printf("Angka %d merupakan bilangan GANJIL\n", number)
// 	}
// }
// /*
// 	Variadic function
// 	-> Pembuatan function dengan parameter seperti biasa
// 	-> Sama halnya dengan membuat slice, []string, []int, dll
// 	-> Pembuatannya cukup dg (...tipeData) ex: ...string, ...int dst
// 	-> Mempunyai sebuah index
// 	-> Catatan: Penerapan variadic function harus di akhir parameter yg dibuat
// 	   ex: func numbers(bumber ...int) (v) | func numbers(number ...int, x, y int)(x)
// */

// func enigmaZoo(count int, animal []string, last int) { //slice
// 	fmt.Println(count)
// 	for _, animal := range animal {
// 		fmt.Println(animal)
// 	}
// 	fmt.Println(last)
// }

// func enigmaZooVariadic(count int, animal ...string) { //variadic
// 	fmt.Println(count)
// 	for _, animal := range animal {
// 		fmt.Println(animal)
// 	}
// }

// /*
// 	Function value (closure)
// 	-> Sebuha function yang bisa dijadikan sebuah value dalam sebuah variable
// 	-> Sehingga variable tersebut bertindak selayaknya sebuah function
// */

// func greeting2 (name string) string {
// 	return "Salam kenal " + name
// }

package main

import "fmt"

func main() {
	var number int = 10

	req(number)
}

func req(num int) {
	if num == 0 {
		fmt.Println("Proses selesai")
		return
	}
	num = num - 1
	fmt.Println("num: ", num)
	req(num)
}
