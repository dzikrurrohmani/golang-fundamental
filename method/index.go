package main

import (
	"fmt"
	"golang-fundamental/method/data"
)

func main() {
	// utils.PrintStudentMethod()
	// utils.PrintShaper()

	// Embedded interface
	var cube data.ShapeGeometry = &data.Cube{
		Side: 4,
	}
	fmt.Println("Volume kubus:", cube.Volume())
	fmt.Println("Luas Permukaan kubus:", cube.Area())
}
