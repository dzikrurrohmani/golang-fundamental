package main

import "fmt"

func main() {
	// // for i:=0; i<10; {
	// // 	fmt.Println("Halo")
	// // 	i++
	// // }
	// // basic for
	// for i:=8; i<10; i++{
	// 	fmt.Println("pertama")
	// }
	// // for a while
	// j:=8
	// for j<10 {
	// 	fmt.Println("kedua")
	// 	j++
	// }

	// // for ever -> tanpa argumen
	// for {
	// 	fmt.Println("ketiga")

	// 	// untuk menghentikan perulangan
	// 	j++
	// 	if j == 20 {
	// 		break
	// 	}
	// }

	// Simulasi continue - even dan odd
	for i:=1; i<12; i++ {
		if i%2 == 1 {
			continue
		}
		fmt.Println(i)
	}
}