package main

import (
	"errors"
	"fmt"
	"strconv"
)

func main() {
	/*
		Go tidak memiliki error handling seperti try/catch/finally
		Go mempunyai built-in error {interface}
		Go sendiri ada namanya defer, panic & recover
	*/
	test2()
}

// biasanya recovery dibuat function
func DeferRecovery() {
	if r := recover(); r!= nil {
		fmt.Println("Terjadi error: ", r)
	}
}

func test2() {
	var (
		input string
	)
	fmt.Print("Masukkan angka: ")
	fmt.Scanln(&input)

	defer DeferRecovery()

	if valid, err := validateNumber(input); valid {
		fmt.Println(input)
	} else {
		// fmt.Println(err.Error())
		panic(err.Error())
	}
}

func validateNumber(input string) (bool, error) {
	_, err := strconv.Atoi(input)
	if err != nil {
		return false, errors.New("bukan angka")
	}
	return true, nil
}

func test1() {
	var (
		input  string
		number int
		err    error
	)
	fmt.Print("Masukkan angka: ")
	fmt.Scanln(&input)

	number, err = strconv.Atoi(input)
	if err == nil {
		fmt.Println(number)
	} else {
		fmt.Println(input, "bukan sebuah angka")
		fmt.Println(err.Error())
	}
}
