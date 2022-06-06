package viewer

import (
	"fmt"
	"golang-fundamental/methodInterfacelatihan/procedure"
)

func PrintResult() {
	fmt.Println("PEMAHAMAN 1")
	var calculation1 procedure.Calculator

	calculation1.Tampil = procedure.Substraction{Num1: 5, Num2: 4}
	calculation1.PrintResult()

	calculation1.Tampil = procedure.Addition{Num1: 5, Num2: 4}
	calculation1.PrintResult()

	// fmt.Println("PEMAHAMAN 2")
	// var calculation2 procedure.CalcOps

	// calculation2 = procedure.Substraction{Num1: 30, Num2: 14}
	// // fmt.Println(calculation2.DoCalc())
	// fmt.Printf("Hasil pengurangan (Num1: %.2f dan Num2: %.2f) adalah:%.2f\n", calculation2.(procedure.Substraction).Num1, calculation2.(procedure.Substraction).Num2 , calculation2.DoCalc())

	// calculation2 = procedure.Addition{Num1: 12, Num2: 7}
	// // fmt.Println(calculation2.DoCalc())
	// fmt.Printf("Hasil penjumlahan (Num1: %.2f dan Num2: %.2f) adalah:%.2f\n", calculation2.(procedure.Addition).Num1, calculation2.(procedure.Addition).Num2 , calculation2.DoCalc())
}
