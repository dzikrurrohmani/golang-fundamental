package main

import "fmt"

func main() {
	// var num int
	// var ptr *int

	// num = 10
	// ptr = &num
	// fmt.Println(ptr)

	// *ptr /= 2
	// fmt.Println(num)

	// variable biasa
	var num int
	fmt.Println(num) // -> 0
	num = 10
	fmt.Println(num) // -> 10

	// variable dengan pointer
	var ptr *int
	fmt.Println(ptr) // -> <nil>

	// mendapatkan alamat memori
	fmt.Println("num (address)", &num)
	fmt.Println("ptr (address)", &ptr)

	// set address dari num ke ptr
	ptr = &num
	fmt.Println("num (address)", &num)
	fmt.Println("ptr (value)", ptr)

	// get value
	fmt.Println("num (value)", num)
	fmt.Println("ptr (value)(value)", *ptr)

	// experiment
	ptradd := &ptr
	fmt.Println("ptr (address)(value)(value)(value)", **ptradd)

	// simulasi perubahan
	// dirubah pada ptr
	*ptr /= 2
	fmt.Println("num (value)", num)
	fmt.Println("ptr (value)(value)", *ptr)
	// dirubah pada num
	num /= 2
	fmt.Println("num (value)", num)
	fmt.Println("ptr (value)(value)", *ptr)

	// Membuat sebuah variable pointer menggunakan keyword new() -> function
	// Ketika membuat pointer dengan keyword new(), sebenarnya akan diperoleh sebuah pointer ke data kosong.
	// ex: var name = new(string) == var name *string

	var name1 = new(int)
	var name2 *int

	fmt.Println(name1, *name1)
	fmt.Println(name2)

	// Pass by reference not by value
	number1 := 20
	fmt.Println("Before assign to function", number1) // 20
	passByReference(&number1)
	fmt.Println("After assign to function", number1) // 10

	number2 := 20
	fmt.Println("Before assign to function", &number2) // 20
	passByValue(number2)
	fmt.Println("After assign to function", &number2) // 10
}

// Menerapkan pointer pada function
func passByReference(number *int) {
	*number = 10
}

func passByValue(number int) {
	number = 10
	fmt.Println(&number)
}
