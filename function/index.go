package main

import "fmt"

func main() {
	// Bisa digunakan berkali2 agar tidak mengulang code
	fmt.Println(add(3,5))
	fmt.Println(add(4,5))
	fmt.Println(add(5,5))

	printMessage()
	greeting()
	fmt.Println("")

	message, integer := Message("Dzikrur", "Sragen")
	fmt.Println(message)
	fmt.Println("saya", "mudik", integer, "kali")
}

func printMessage() {
	fmt.Println("Halo dunia...")
}

func greeting() {
	fmt.Println("Salam kenal...")
}

func add(param1, param2 int) int {
	return param1+param2-7
}

func Message(name, address string) (string, int) {
	return fmt.Sprintf("Halo nama saya %s, saya tinggal di %s", name, address), 1
}