package main

import "fmt"

func main() {
	var customerName string = "James"
	var customerAge uint8 = 26
	var customerStatus bool = true

	/*
		%s -> string
	  	%d -> bilangan bulat
		%f -> bilangan desimal
		%v -> semua
	*/
	fmt.Printf("Nama\t\t: %v\nUmur\t\t: %v\nMarriage\t: %v\n", customerName,
	customerAge, customerStatus)
	fmt.Println("Halo")

	var (
		firstName, lastName string
		age int
	)

	firstName = "Cristiano"
	lastName = "Ronaldo"
	age = 37

	fmt.Println(firstName, lastName, age)

	var (bootcampName, bootcampAddress = "Enigma Camp", "Jakarta Selatan")
	fmt.Println(bootcampName, bootcampAddress)

	// Inferred Type
	days, month, year := "Monday", "May", 2022
	fmt.Println(days, month, year)

	const phi = 3.14
	fmt.Println(phi)
	const (
		StatusOK		= 200
		StatusCreated	= 201
	)
	fmt.Println(StatusOK, StatusCreated)
	// Imutable & Mutable const

	_ = 99
	_ = "ignore"
	a:= 0
	fmt.Println(a)

}
