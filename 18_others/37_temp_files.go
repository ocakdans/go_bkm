package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"os"
)

func main() {
	TempDirectory()
}

func TempDirectory() {
	tempDirPath, err := ioutil.TempDir("", "myTempDir")
	if err != nil {
		log.Fatal(err)

	}
	fmt.Println("Temp dir created:", tempDirPath)

	tempFile, err := ioutil.TempFile(tempDirPath, "myTempFile.txt")
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Temp file created:", tempFile.Name())

	err = tempFile.Close()
	if err != nil {
		log.Fatal(err)
	}

	// Delete the temp file and directory
	err = os.Remove(tempFile.Name())
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Temp file deleted", tempFile.Name())

	err = os.Remove(tempDirPath)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Temp directory deleted", tempDirPath)

}
