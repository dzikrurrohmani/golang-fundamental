package main

import (
	"fmt"
	"strings"
)

func main() {
	a := 20
	if a%2 == 0 {
		fmt.Println("This is even number!")
	} else {
		fmt.Println("This is odd number!")
	}
	fmt.Println(strings.Repeat("-", 58))

	// Simulasi untuk menghitung Nilai (GPA)
	/*
		GPA: 3.5 - 4.0 Cumlaude
		GPA: 3.0 - 3.5 Memuaskan
		GPA: 2.75 - 3.0 Cukup
		GPA: <2.75 Kurang Memuaskan
	*/

	GPA := 5.0
	var status string
	if GPA >= 3.5 && GPA <= 4.0 {
		status = "Cumlaude"
	} else if GPA >= 3.0 && GPA < 3.5 {
		status = "Memuaskan"
	} else if GPA >= 2.75 && GPA < 3.0 {
		status = "Cukup"
	} else if GPA >= 0 && GPA < 2.75 {
		status = "Kurang Memuaskan"
	} else {
		status = "Diluar range GPA"
	}
	fmt.Printf("GPA %.2f : %s\n", GPA, status)
	fmt.Println(strings.Repeat("-", 58))

	// Simulasi Perhitungan Nilai
	/*
		100% -> Perfect
		>= 70 & < 100 -> Good
		< 70 -> Not bad
	*/

	point := 10000.0

	if percent := point / 100; percent >= 100 {
		fmt.Printf("%.2f -> Perfect\n", percent)
	}
	fmt.Println(strings.Repeat("-", 58))
}
