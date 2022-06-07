package main

import (
	"fmt"
)

func main() {
	// viewer.PrintResult()
	var whatever interface{}
	fmt.Printf("Type of whatever di alamat %p adalah %T\n", &whatever, whatever)

	whatever = "Enigma"
	fmt.Printf("Type of whatever di alamat %p adalah %T\n", &whatever, whatever)

	whatever = 2000
	fmt.Printf("Type of whatever di alamat %p adalah %T\n", &whatever, whatever)

	whatever = true
	fmt.Printf("Type of whatever di alamat %p adalah %T\n", &whatever, whatever)

	whatever = [2]int{1, 2}
	fmt.Printf("Type of whatever di alamat %p adalah %T\n", &whatever, whatever)

	whatever = []string{"mangga", "apel"}
	fmt.Printf("Type of whatever di alamat %p adalah %T\n", &whatever, whatever)

	// penerapan interface kosong
	// fmt.Print()
	// fmt.Printf()
	// fmt.Println()
	// fmt.Sprint()
	// fmt.Sprintf()
	// fmt.Sprintln()

	// pembuatan interface kosong harus ada -> {}
	// jika tidak ada maka itu dianggap sebagai pembuatan interface {kontrak}
	printType(1995)
	printType("Enigma")
	printType(false)
	mustNum(1995)
	mustNum("Enigma")
	mustNum(false)
}

func printType(i interface{}) {
	switch v := i.(type) {
	case int:
		fmt.Printf("%v\n", v)
	case string:
		fmt.Printf("%q memiliki panjang %v\n", v, len(v))
	default:
		fmt.Printf("Error unknown %T\n", v)
	}
}

func mustNum(i interface{}) {
	v, ok := i.(int)
	if ok {
		fmt.Println(v)
	} else {
		fmt.Println("Tipe data harus (int)")
	}
}
