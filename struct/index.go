package main

import (
	"fmt"
	"os"
)

// import (
// 	"fmt"
// 	"strings"
// )

// type student struct { // private access modifier -> hanya bisa diakses di 1 package
// 	npm, name string
// 	grade     int
// 	isDone    bool
// }

func main() {
	// // fmt.Println(student) // tidak bisa langsng memanggil struct tanpa diinisialisasi
	// var student1 student  // membuat sebuah objek atau variavle dari struct dengan nama student1
	// fmt.Println(student1) // akan mencetak zero value dari masing masing properti/variable

	// // isi value dari variable/properti
	// student1.npm = "12345678"
	// student1.name = "Budi"
	// student1.grade = 100
	// student1.isDone = false

	// fmt.Println("student1:", student1)

	// // menggunakan struct literal
	// var student2 = student{"12345679", "Joko", 98, true}
	// fmt.Println("student2:", student2.npm, student2.name, student2.grade, student2.isDone)

	// student3 := student{"12345680", "Sari", 80, true}
	// fmt.Println("student3:", student3)

	// // deklarasi struct tanpa type alias
	// student4 := struct {
	// 	nama string
	// }{
	// 	"Joko",
	// }

	// fmt.Println("student4:", student4)

	// var student5 struct {
	// 	alamat string
	// }

	// student5.alamat = "Ragunan"
	// fmt.Println("student5:", student5.alamat)

	// // pointer juga dapat diterapkan pada struct
	// fmt.Println(&student1)
	// fmt.Println(&student1.name)

	// // mengambil referensi dari student2
	// var studentPointer *student = &student2
	// fmt.Println("&student2:", &student2)
	// fmt.Println("studentPointer:", studentPointer)
	// fmt.Println("*studentPointer:", *studentPointer)
	// fmt.Printf("student2 address: %p\n", &student2)
	// fmt.Printf("studentPointer value: %p\n", studentPointer)
	// fmt.Println("studentPointer.npm", studentPointer.npm)
	// fmt.Println("(*studentPointer).npm", (*studentPointer).npm)
	// // mengakses value yang bukan tipe data pointer, karena property name adalah string
	// // fmt.Println(*studentPointer.name) // cara salah
	// fmt.Println("(*studentPointer).name", (*studentPointer).name)

	// // membuat sebuah variable dengan keyword new(), di struct juga bisa
	// // kita menggunakan student struct
	// var student6 student
	// var student6New = new(student) // mendefinisikan sebuah object struct menggunakan keyword new()
	// fmt.Println("student6:", student6)
	// fmt.Println("student6 new:", student6New)

	// // mencetak alamat untuk setiap struct
	// fmt.Println("student6 new:", student6New)

	// // tentang keyword make() yang sebelumnya digunakan untuk membuat slice, map, channels
	// // kali ini juga bisa digunakan untuk membuat struct

	// // embedded struct
	// fmt.Println(strings.Repeat("-", 40))
	// mahasiswa1 := mahasiswa{
	// 	npm:  "13753030",
	// 	name: "Jution Candra",
	// 	ins: institut{
	// 		institutName:    "Politeknik Bandar Lampung",
	// 		institutAddress: "Bandar Lampung",
	// 		institutAccred:  "B",
	// 		institutSince:   1962,
	// 	},
	// 	course: []course{
	// 		{
	// 			name: "Matematika Diskrit",
	// 			sks:  3,
	// 		},
	// 		{
	// 			name: "Algoritma Pemrograman",
	// 			sks:  2,
	// 		},
	// 	},
	// 	grade:  78,
	// 	isDone: false,
	// }
	// fmt.Println("mahasiswa1")
	// fmt.Println(mahasiswa1)
	// fmt.Printf("%+v\n", mahasiswa1)

	// siswasma1 := siswasma{
	// 	nis:  "13753031",
	// 	name: "Jution Kirana",
	// }
	// fmt.Println("siswasma1")
	// fmt.Println(siswasma1)
	// siswasma1.institut.institutName = "ITS"
	// siswasma1.institut.institutAddress = "Surabaya"
	// siswasma1.institut.institutAccred = "A"
	// siswasma1.institut.institutSince = 1960

	// fmt.Println(siswasma1)
	// fmt.Printf("%+v\n", siswasma1)

	// fmt.Println(strings.Repeat("-", 40))
	// mahasiswaMahasiswa := []mahasiswa{
	// 	{
	// 		npm:  "22111998",
	// 		name: "Dzikrur",
	// 		ins: institut{
	// 			institutName:    "UNS",
	// 			institutAddress: "Bandar Lampung",
	// 			institutAccred:  "B",
	// 			institutSince:   1962,
	// 		},
	// 		course: []course{
	// 			{
	// 				name: "Matematika Diskrit",
	// 				sks:  3,
	// 			},
	// 			{
	// 				name: "Algoritma Pemrograman",
	// 				sks:  2,
	// 			},
	// 		},
	// 		grade:  78,
	// 		isDone: false,
	// 	},
	// 	{
	// 		npm:  "19031999",
	// 		name: "Rohmani",
	// 		ins: institut{
	// 			institutName:    "UGM",
	// 			institutAddress: "Bandar Lampung",
	// 			institutAccred:  "B",
	// 			institutSince:   1962,
	// 		},
	// 		course: []course{
	// 			{
	// 				name: "Metode Numerik",
	// 				sks:  3,
	// 			},
	// 			{
	// 				name: "Dasar Pengolahan Sinyal",
	// 				sks:  2,
	// 			},
	// 		},
	// 		grade:  78,
	// 		isDone: false,
	// 	},
	// }
	// fmt.Println("mahasiswaMahasiswa")
	// for _, mhsiswa := range mahasiswaMahasiswa {
	// 	fmt.Println(mhsiswa.npm, mhsiswa.name, mhsiswa.ins.institutName)
	// 	for _, matkul := range mhsiswa.course {
	// 		fmt.Println(matkul.name)
	// 	}
	// }

	// temp := mahasiswa{
	// 	npm:  "19031999",
	// 	name: "Rohmani",
	// 	ins: institut{
	// 		institutName:    "UGM",
	// 		institutAddress: "Bandar Lampung",
	// 		institutAccred:  "B",
	// 		institutSince:   1962,
	// 	},
	// 	course: []course{
	// 		{
	// 			name: "Metode Numerik",
	// 			sks:  3,
	// 		},
	// 		{
	// 			name: "Dasar Pengolahan Sinyal",
	// 			sks:  2,
	// 		},
	// 	},
	// 	grade:  78,
	// 	isDone: false,
	// }
	// tambahMahasiswa(&mahasiswaMahasiswa, temp)
	// fmt.Println("mahasiswaMahasiswa")
	// for _, mhsiswa := range mahasiswaMahasiswa {
	// 	fmt.Println(mhsiswa.npm, mhsiswa.name, mhsiswa.ins.institutName)
	// 	for _, matkul := range mhsiswa.course {
	// 		fmt.Println(matkul.name)
	// 	}
	// }
	// fmt.Println("mahasiswaMahasiswa")
	// for _, mhsiswa := range mahasiswaMahasiswa {
	// 	fmt.Println(mhsiswa.npm, mhsiswa.name, mhsiswa.ins.institutName)
	// 	for _, matkul := range mhsiswa.course {
	// 		fmt.Println(matkul.name)
	// 	}
	// }

	// emp := employee{
	// 	FirstName: "Jack",
	// 	LastName:  "Ma",
	// }
	// output, err := json.MarshalIndent(emp, "", " ")

	// if err != nil {
	// 	fmt.Println(err)
	// 	return
	// }
	// fmt.Println(string(output))

	// env? environment -> KEY-VALUE
	// Tiap OS berbeda untuk cara set (mengisi si env secara dinamis) dan get(mengambil yang sudah diset)
	// GET -> echo %ENV_NAME% || %ENV_NAME (windows)| Terminal cmd, powershell (x)
	// GET -> echo $ENV_NAME (OsX, Linux)

	// SET -> set FULL
	// Golang Inbuilt OS Package

	os.Setenv("FULL_NAME", "jution") // SET Env
	fullName := os.Getenv("FULL_NAME")
	fmt.Println(fullName)
}

// // sebagai alternatif inheritance pada OOP bisa dilakukan spt di bawah ini
// type institut struct {
// 	institutName    string
// 	institutAddress string
// 	institutAccred  string
// 	institutSince   int
// }

// type course struct {
// 	name string
// 	sks  int
// }

// type mahasiswa struct {
// 	npm, name string
// 	grade     int
// 	isDone    bool
// 	ins       institut // embedded struct
// 	course    []course
// }

// type siswasma struct {
// 	nis, name string
// 	institut  // embedded struct with syntactic sugar (nama variable sama dg structnya bisa ditulis sekali)
// }

// Reflection is the ability of a program to inspect its variables and values at run time and find their type
// Struct tag bisa pada JSON, XML
// type employee struct {
// 	FirstName string `json:"firstname"` // metadata information
// 	LastName  string `json:"lastname"`  //`tagName`
// 	IsStatus  bool   `json:"is_status"`
// 	// isStatus  bool   `json:"is_status, omitempty"`
// }

// func tambahMahasiswa(sliceNya *[]mahasiswa, mahaNya ...mahasiswa) {
// 	*sliceNya = append(*sliceNya, mahaNya...)
// }
