package main

import "fmt"

func main() {
	/*
		Membuat sebuah function yang akan mengecek sebuah kota yang diinputkan
		berasal dari provinsi dan negara mana.

		Jakarta Selatan -> DKI Jakarta, Indonesia.
		Cianjur -> Jawa Barat, Indonesia
	*/
	province, region := location("Jakarta Selatan")
	fmt.Printf("Saya tinggal di %s provinsi %s\n", region, province)
	province, region = location("Bandung")
	fmt.Printf("Saya tinggal di %s provinsi %s\n", region, province)
	province, region = location("Pesawaran")
	fmt.Printf("Saya tinggal di %s provinsi %s\n", region, province)
	province, region = location("Sragen")
	fmt.Printf("Saya tinggal di %s provinsi %s\n", region, province)
	fmt.Println(location("Sragen"))
}

func location(city string) (string, string) {
	var province, region string

	switch city {
	case "Jakarta Selatan", "Jakarta Barat", "Jakarta Timur":
		province, region = "DKI Jakarta", "Indonesia"
	case "Bandar Lampung", "Lampung Selatan", "Pesawaran":
		province, region = "Lampung", "Indonesia"
	case "Bogor", "Cianjur", "Bandung":
		province, region = "Jawa Barat", "Indonesia"
	default:
		province, region = "tidak ada", "tidak ada"
	}
	return province, region
}
