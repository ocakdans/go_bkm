package main

import (
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"os"
)

var (
	newFile *os.File
	err     error
	//fileInfo *os.FileInfo
)

func main() {
	//CreateFile()
	//CheckFile()
	//MoveFile()
	//CheckFileExists()
	//FileOpenAndClose()
	//CheckPermission()
	//CopyFile()
	//WriteBytesToFile()
	//WriteBytesToFileIOUtil()
	//DeleteFile()

}

func DeleteFile() {
	err := os.Remove("./movedFile/test_copy.txt")
	if err != nil {
		log.Fatal(err)
	}
}

func WriteBytesToFileIOUtil() {
	err := ioutil.WriteFile("./movedFile/test.txt", []byte("Hi, I am Selim"), 0666)
	if err != nil {
		log.Fatal(err)

	}
}

func WriteBytesToFile() {
	file, err := os.OpenFile("./movedFile/test.txt", os.O_WRONLY|os.O_APPEND, 0666)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()
	byteSlice := []byte("I learn Golang")
	bytesWritten, err := file.Write(byteSlice)
	if err != nil {
		log.Fatal(err)
	}
	log.Printf("Wrote %d bytes.\n", bytesWritten)
}

func CopyFile() {
	originalFile, err := os.Open("./movedFile/test.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer originalFile.Close()

	newFile, err := os.Create("./movedFile/test_copy.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer newFile.Close()

	bytesWritten, err := io.Copy(newFile, originalFile)
	if err != nil {
		log.Fatal(err)
	}
	log.Printf("Copied %d bytes.", bytesWritten)

	//aşağıdaki kısımda hafızadaki kopiyayı diske yazdırıyoruz.
	err = newFile.Sync()
	if err != nil {
		log.Fatal(err)
	}
}

func CheckPermission() {
	file, err := os.OpenFile("./movedFile/test.txt", os.O_WRONLY, 0666)
	if err != nil {
		if os.IsPermission(err) {
			log.Println("Error: Write permission denied.")
		}
	}
	defer file.Close()
}

func FileOpenAndClose() {

	// Version 1
	// file, err := os.Open("./movedFile/test.txt")
	// defer file.Close()

	// if err != nil {
	// 	log.Fatal(err)
	// }
	// Version 2
	// os.O_APPEND --> dosya açılış amacı
	// 06666 --> dosya izinleri
	file, err := os.OpenFile("./movedFile/test.txt", os.O_CREATE|os.O_APPEND, 0666)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

}

func CheckFileExists() {
	fileInfo, err := os.Stat("./movedFile/test.txt")
	if err != nil {
		if os.IsNotExist(err) {
			log.Fatal("File does not exist.")
		}
		log.Fatal(err)
	}
	log.Println("File does exist. File information: ", fileInfo)
}

func MoveFile() {
	originalPath := "test.txt"
	newPath := "./movedFile/test.txt"
	err := os.Rename(originalPath, newPath)
	if err != nil {
		log.Fatal(err)
	}

}

func CheckFile() {
	fileInfo, err := os.Stat("test.txt")
	if err != nil {
		if os.IsNotExist(err) {
			log.Fatal("File does not exist.")
		}
	}
	fmt.Println("File name: ", fileInfo.Name())
	fmt.Println("File size: ", fileInfo.Size())
	fmt.Println("Permissions: ", fileInfo.Mode())
	fmt.Println("Last Modified: ", fileInfo.ModTime())
	fmt.Println("Is dictionary: ", fileInfo.IsDir())
	fmt.Println("Sys Interface Type: ", fileInfo.Sys())
	fmt.Printf("Sys Info: %v ", fileInfo.Sys())
}

func CreateFile() {
	newFile, err = os.Create("test.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer newFile.Close()
}
