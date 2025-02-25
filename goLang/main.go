package main

import (
	"errors"
	"fmt"
	"os"
)

const (
	filesFolder = "filesArea/"
)

func createFile(givenName string) {
	newPath := filesFolder + givenName
	//fmt.Println(fileAlreadyThere(newPath))
	if !fileAlreadyThere(newPath) {
		file, err := os.Create(newPath)
		if err != nil {
			fmt.Println(err)
			fmt.Println("Unable to create file due to above error...")
			return
		}
		defer file.Close()
		fmt.Println("File created successfully: ", newPath)
	} else {
		fmt.Println("File already there... <createFile> will take no further action")
	}

}

func fileAlreadyThere(filePath string) bool {
	_, error := os.Stat(filePath)
	//return !os.IsNotExist(err)
	return !errors.Is(error, os.ErrNotExist)
}

func main() {
	createFile("test_5.txt")

}
