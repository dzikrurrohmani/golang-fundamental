package main

import "golang-fundamental/file/data"

/*
Buat file
Baca file
Menulis file
Delete file
*/

var fileLocation = "/home/ubuntu/ITDP/golang-fundamental/file/db.dat"

func main() {
	// data.CreateFile(fileLocation)
	// data.WriteToFile(fileLocation, "Joko")
	data.OpenWithOsOpenFile(fileLocation)
}
