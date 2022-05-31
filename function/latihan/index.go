// package main

// import "fmt"

// func main() {
// 	/*
// 		Membuat sebuah function yang akan mengecek sebuah kota yang diinputkan
// 		berasal dari provinsi dan negara mana.

// 		Jakarta Selatan -> DKI Jakarta, Indonesia.
// 		Cianjur -> Jawa Barat, Indonesia
// 	*/
// 	province, region := location("Jakarta Selatan")
// 	fmt.Printf("Saya tinggal di %s provinsi %s\n", region, province)
// 	province, region = location("Bandung")
// 	fmt.Printf("Saya tinggal di %s provinsi %s\n", region, province)
// 	province, region = location("Pesawaran")
// 	fmt.Printf("Saya tinggal di %s provinsi %s\n", region, province)
// 	province, region = location("Sragen")
// 	fmt.Printf("Saya tinggal di %s provinsi %s\n", region, province)
// 	fmt.Println(location("Sragen"))
// }

// func location(city string) (province, region string) { // Naming return
// 	// var province, region string

// 	switch city {
// 	case "Jakarta Selatan", "Jakarta Barat", "Jakarta Timur":
// 		province, region = "DKI Jakarta", "Indonesia"
// 	case "Bandar Lampung", "Lampung Selatan", "Pesawaran":
// 		province, region = "Lampung", "Indonesia"
// 	case "Bogor", "Cianjur", "Bandung":
// 		province, region = "Jawa Barat", "Indonesia"
// 	default:
// 		province, region = "tidak ada", "tidak ada"
// 	}
// 	return
// 	// return province, region
// }

// package main

// import (
// 	"fmt"
// 	"strings"
// )

// func main() {
// 	nama := "Jution"
// 	alamat := "Ragunan"
// 	pendidikan := "S1"
// 	skills := []string{"Tidur", "Berlari", "Berenang", "Menghilang"} //variadic

// 	portofolio(nama, alamat, pendidikan, skills...)

// }

// func portofolio(name, address, education string, skills ...string) {
// 	fmt.Println("Nama\t\t:", name)
// 	fmt.Println("Alamat\t\t:", address)
// 	fmt.Println("Pendidikan\t:", education)
// 	fmt.Print("Skills\t\t: ")
// 	for i, skill := range skills {
// 		fmt.Print(skill)
// 		if i<len(skills)-1 {
// 			fmt.Print(", ")
// 		} else {
// 			fmt.Print("\n")
// 		}
// 	}
// 	fmt.Print("Skills\t\t: ", strings.Join(skills, ", "), "\n")
// }

package main

import (
	"fmt"
	"strings"
)

func main() {
	/*
		Function as parameter
		-> Kita akan membuat sebuah callback function yang ditaruh di parameter
		-> Artinya function itu bertindak sebagai parameter
		-> ex: Saya mau memfilter email, '0' -> jutionck@gmail.com, jution.kirana@enigmacamp.com, dst.
	*/
	getValidEmail("jutionck@gmail.com", emailFilter)
	getValidEmail("jutionck.gmail.com", emailFilter)
	fmt.Printf("%T", emailFilter)
}

func getValidEmail(email string, callback func(string) string) {
	// emailFiltered := callback(email)
	// fmt.Println("Your email", emailFiltered)
	fmt.Println("Your email", callback(email))
}

func emailFilter(email string) string {
	if strings.Contains(email, "@") {
		return email
	} else {
		return "not valid!"
	}
}
