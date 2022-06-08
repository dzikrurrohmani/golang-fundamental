package data

import (
	"fmt"
	"golang-fundamental/file/utils"
	"io"
	"os"
)

func OpenWithOsOpenFile(fileLocation string) {
	file, err := os.OpenFile(fileLocation, os.O_RDWR, 0644)
	if utils.IsError(err) {
		return
	}
	defer file.Close()

	var outputText = make([]byte, 1024)
	for {
		n, err := file.Read(outputText)
		if err != io.EOF {
			if utils.IsError(err) {
				break
			}
		}

		if n == 0 {
			break
		}
	}
	if utils.IsError(err) {
		return
	}
	fmt.Println(string(outputText))
}
