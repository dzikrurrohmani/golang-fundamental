package main

import (
	"fmt"
	"strconv"
	"strings"
)

func main() {
	// Number : int, uint, float
	var numberOne int = -1
	var numberTwo uint = 1
	var numberThree float32 = 20
	fmt.Printf("Number One: %d\n", numberOne)
	fmt.Printf("Number Two: %d\n", numberTwo)
	fmt.Printf("Number Three: %.2f\n", numberThree)

	// Boolean
	var isValid bool = true
	fmt.Printf("isValid: %t\n", isValid)

	// String
	var message string = "Halo salam kenal"
	fmt.Printf("message: %s\n", message)

	// String literal
	getMessage := `
	Hello
	hai
	gracias
	hola
	`
	fmt.Print(getMessage)

	getData := "Hari ini adalah hari jum'at ke pasar membeli \"apel, gula, dan garam\""
	fmt.Println(getData)

	// Concat
	namaDepan := "Dzikrur"
	namaBelakang := "Hilmi"
	namaLengkap := namaDepan + " " + namaBelakang // Harus tipe data yang sama
	fmt.Println(namaLengkap)

	fmt.Println(strings.ToUpper(namaLengkap))
	fmt.Println(strings.ToLower(namaLengkap))
	fmt.Println(strings.Contains(namaLengkap, "D")) // Case sensitive
	fmt.Println(strings.Count(namaLengkap, "r"))    // Case sensitive
	fmt.Println(len(namaLengkap))                   // jumlah karakter

	// Convert Data Type
	var index int8 = 15
	otherindex := int32(index)
	fmt.Println(otherindex)

	// Cek tipe data
	fmt.Printf("tipe data index: %T\n", index)
	fmt.Printf("tipe data otherindex: %T\n", otherindex)

	// Number to string
	var a int = 12
	b := fmt.Sprintf("%d", a)
	fmt.Printf("print %T", b)

	c := strconv.Itoa(a)
	fmt.Printf("print %T", c)

	// String to number
	d, _ := strconv.Atoi(b)
	e, _ := strconv.Atoi(c)
	result := d+e
	fmt.Printf("hasil jumlah : %d\n", result)

	// Hati-hati
	var angka_int64 int64 = 128
	fmt.Println(int8(127))
	fmt.Println(int8(angka_int64))

}