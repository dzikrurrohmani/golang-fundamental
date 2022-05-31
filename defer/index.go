package main

import "fmt"

func main() {
	/*
		Defer
		-> Sebuah baris eksekusi yang terakhir
	*/

	fmt.Println(1)
	defer fmt.Println(2)
	fmt.Println(3)

	defer a()
	b()
	c()
}

func a() {fmt.Println("a Function")}
func b() {fmt.Println("b Function")}
func c() {fmt.Println("c Function")}
