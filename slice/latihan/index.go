package main

import "fmt"

func main() {
	/*
		Tolong hapus aku dari antrian
		[]int{1, 3, 5, 6, 7, 8, 10}

		-> 1
		[]int{3, 5, 6, 7, 8, 10}

		-> 3
		[]int{5, 6, 7, 8, 10}
	*/
	var delValue int
	var antrian = []int{1, 3, 5, 6, 7, 8, 10}

	fmt.Print("Tolong hapus aku dari antrian")
	fmt.Printf("%#v\n\n",antrian)
	for {
		var newAntrian []int
		fmt.Print("-> ")
		fmt.Scanln(&delValue)
		for i := 0; i < len(antrian); i++ {
			if antrian[i] != delValue {
				newAntrian = append(newAntrian, antrian[i])
			}
		}
		antrian = newAntrian
		fmt.Printf("%#v\n\n",newAntrian)
		if len(antrian) == 0 {
			break
		}
	}
}
