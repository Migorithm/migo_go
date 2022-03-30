package main

import (
	"errors"
	"fmt"
	"log"
	"os"
	"path"
)

func main() {
	var dirs = []string{"/home", "mano", "go-projects", "..", "src"}
	path := path.Join(dirs...)
	fmt.Println(path) // /home/mano/src
	fmt.Printf("Path after join: %s\n", path)

	//Create file
	CreateEmptyFile("sample.txt")

	//File Info
	FileInfo("ss.txt")

}

func CreateEmptyFile(filename string) {
	myFile, err := os.Create(filename)
	if err != nil {
		log.Fatal("Error!", err)
	}
	log.Println("Emptry file created successfully. ", myFile)
	defer myFile.Close()

}

func FileInfo(filename string) {

	_, err := os.OpenFile(filename,...)
	if os.IsNotExist(err) {
    // handle the case where the file doesn't exist
	}else {
		fileInfo, _ := os.Stat(filename)
	

		fmt.Println("File Name:", fileInfo.Name())
		fmt.Println("Size ", fileInfo.Size(), " bytes")
		fmt.Println("Permissions:", fileInfo.Mode())
		fmt.Println("Last modified:", fileInfo.ModTime())
		fmt.Println("Is Directory:", fileInfo.IsDir())
	
	}

	

	

}
