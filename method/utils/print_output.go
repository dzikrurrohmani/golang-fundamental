package utils

import (
	"fmt"
	"golang-fundamental/method/data"
)

func PrintStudentMethod() {
	fmt.Println(data.SayHi("Jono"))
	fmt.Println(data.SayHello())
	liat := data.SayHello()
	fmt.Println(liat.Name)
	liat.Name = "Sila"
	fmt.Println(liat.Name)
	coba := new(data.Exgreeting)
	fmt.Println(coba)
	fmt.Println(coba.Kelas, coba.Name)
	coba2 := data.Exgreeting{Name: "Safira", Kelas: 2}
	fmt.Println(coba2)

	cobaMethod := data.Exgreeting{Name: "Alila"}
	fmt.Printf("cobaMethod (address) %p\n", &cobaMethod)
	fmt.Println(cobaMethod.SayHi())
	cobaMethod.SayHi().SayGoodBye("Bulan")
	fmt.Printf("cobaMethod (address) %p\n", &cobaMethod)
	fmt.Print(cobaMethod.SayHiPointer())
	cobaMethod.SayGoodByePointer("Bulan")
	fmt.Printf("cobaMethod (address) %p\n", &cobaMethod)

	var cobaPass = data.Exgreeting{Name: "john wick", Kelas: 21}
	fmt.Println("Name before", cobaPass.Name)

	cobaPass.ChangeNamePBV("jason bourne")
	fmt.Println("cobaPass.Name after changeName1", cobaPass.Name)

	cobaPass.ChangeNamePBR("ethan hunt")
	fmt.Println("cobaPass.Name after changeName2", cobaPass.Name)
}
