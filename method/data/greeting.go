package data

import (
	"fmt"
)

type greeting struct { // Unexported (private)
	Name string
}

type Exgreeting struct { // Unexported (private)
	Name  string
	age   int
	Kelas int
}

func SayHi(name string) string {
	return name
}

func SayHello() greeting { // Exported (public)
	return greeting{
		Name: "Jack",
	}
}

func ExSayHello() Exgreeting { // Exported (public)
	return Exgreeting{
		Name: "ExJack",
		age:  23,
	}
}

func (g Exgreeting) SayHi() Exgreeting {
	fmt.Println("Halo saya", g.Name)
	fmt.Printf("g SayHi (address): %p\n", &g)
	return g
}

func (g Exgreeting) SayGoodBye(name string) {
	fmt.Println("Sampai jumpa lagi", name)
	fmt.Printf("g SayGoodBye (address): %p\n", &g)
}

func (g *Exgreeting) SayHiPointer() Exgreeting {
	fmt.Println("Halo saya", g.Name)
	fmt.Printf("g SayHiPointer (address): %p\n", g)
	return *g
}

func (g *Exgreeting) SayGoodByePointer(name string) {
	fmt.Println("Sampai jumpa lagi", name)
	fmt.Printf("g SayGoodByePointer (address): %p\n", g)
}

func (g Exgreeting) ChangeNamePBV(name string) {
	fmt.Println("---> on ChangeNamePBV, name changed to", name)
	fmt.Printf("%p\n", &g)
	g.Name = name
}

func (g *Exgreeting) ChangeNamePBR(name string) {
	fmt.Println("---> on ChangeNamePBR, name changed to", name)
	fmt.Printf("%p\n", g)
	g.Name = name
}