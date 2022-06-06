package utils

import (
	"fmt"
	"golang-fundamental/method/data"
)

func PrintShaper() {
	rectangle := new(data.Rectangle)
	rectangle.Length = 5
	rectangle.Width = 5

	getRectangleArea := rectangle.Area()
	getRectanglePerimeter := rectangle.Perimeter()

	fmt.Printf("Luas Persegi Panjang (P = %f dan L=%f) adalah: %f\n", rectangle.Length, rectangle.Width, getRectangleArea)
	fmt.Printf("Keliling Persegi Panjang (P = %f dan L=%f) adalah: %f\n", rectangle.Length, rectangle.Width, getRectanglePerimeter)
}
