package main

import "fmt"

func main() {
	point := 30
	switch point {
	case 100:
		fmt.Println("Perfect!")
	case 80:
		fmt.Println("Good!")
	default:
		fmt.Println("Not Bad!")
	}

	//multiple case
	switch point {
	case 80, 90, 100:
		fmt.Println("Perfect!")
	case 60, 70:
		fmt.Println("Good!")
	case 50:
		fmt.Println("Awesome!")
	default:
		fmt.Println("Not Bad!")
	}

	//Switch Like If Else
	switch {
	case point == 100:
		fmt.Println("Perfect!")
	case point>=80 && point<100:
		fmt.Println("Good!")
	case point>=60 && point<80:
		fmt.Println("Awesome!")
	// default:
	// 	fmt.Println("Not Bad!")
	}

	// keyword fallthrough -> memaksa pengecekan 1 tingkat di bawahnya
	switch {
	case point == 100:
		fmt.Println("Perfect!")
		fallthrough
	case point>=80 && point<100:
		fmt.Println("Good!")
		fallthrough
	case point>=60 && point<80:
		fmt.Println("Awesome!")
	default:
		fmt.Println("Not Bad!")
	}
}
