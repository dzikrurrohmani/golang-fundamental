package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	// var (
	// 	fullName, address string
	// 	age               int
	// )

	fmt.Println("Data Diri Trainee PT Enigma Cipta Humanika")
	fmt.Println(strings.Repeat("-", 48))

	// fmt.Println("Masukkan nama anda: ")
	// fmt.Scanln(&fullName)

	// fmt.Println("Masukkan alamat anda: ")
	// fmt.Scanln(&address)

	// fmt.Println("Masukkan usia anda: ")
	// fmt.Scanln(&age)

	// fmt.Println(strings.Repeat("-", 48))
	// fmt.Print("Nama \t Alamat \t Usia\n")
	// fmt.Println(strings.Repeat("-", 48))
	// fmt.Printf("%s \t %s \t %d\n", fullName, address, age)
	// fmt.Println(strings.Repeat("-", 48))

	// bufio.NewReader(os.Stdin)
	inputReader := bufio.NewReader(os.Stdin)
	fmt.Print("Masukkan nama anda: ")

	name, _ := inputReader.ReadString('\n')
	fmt.Println("Nama anda", name)

}
