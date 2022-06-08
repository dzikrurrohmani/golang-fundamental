package data

import (
	"fmt"
	"golang-fundamental/file/utils"
	"os"
)

func WriteToFile(fileLocation string, input string) {
	// buka file bisa pakai os atau ioutil
	// file permission 0644 777

	outputFile, outputError := os.OpenFile(fileLocation, os.O_RDWR|os.O_APPEND, 0644)
	if outputError != nil {
		fmt.Println(outputError.Error())
		fmt.Println("An error occured with file creation")
		return
	}
	defer outputFile.Close()
	_, err := outputFile.WriteString(input + "\n")
	if utils.IsError(err){
		return
	}

	err = outputFile.Sync()
	if utils.IsError(err) {
		return
	}
}
