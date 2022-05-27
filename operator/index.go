package main

import "fmt"

func main() {
	/*
		Jenis2 operator>
		1. Aritmatika
		2. Perbandingan
		3. Logika
	*/

	fmt.Println("Operator Aritmatika")
	a := 10
	b := 5
	result := a + b
	fmt.Println(result)
	result = a - b
	fmt.Println(result)
	result = a * b
	fmt.Println(result)
	result = a / b
	fmt.Println(result)
	result = a % b
	fmt.Println(result)

	fmt.Println("Operator Perbandingan")
	x := 2
	y := 2
	fmt.Println(x == y)
	fmt.Println(x != y)
	fmt.Println(x > y)
	fmt.Println(x < y)
	fmt.Println(x >= y)
	fmt.Println(x <= y)

	fmt.Println("Operator Logika")
	p := true
	q := false
	fmt.Println(p && q)
	fmt.Println(p || q)
	fmt.Println(!q)

}
