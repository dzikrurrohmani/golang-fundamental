package data

import (
	"fmt"
	"golang-fundamental/file/utils"
	"os"
)

func CreateFile(fileLocation string) {
	// menggunakan package os

	// deteksi apakah file yang akan dibuat sudah ada ata belum
	_, err := os.Stat(fileLocation)

	// setelah itu akan dibuat
	if os.IsNotExist(err) {
		file, err := os.Create(fileLocation)
		if utils.IsError(err) {
			return
		}
		defer file.Close()
	}
	fmt.Println("File berhasil dibuat: ", fileLocation)
}
