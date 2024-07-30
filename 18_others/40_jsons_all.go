package main

import (
	"encoding/json"
	"fmt"
	"os"
	//"github.com/aws/smithy-go/document/json"
)

func main() {
	p := Person{
		ID:        1,
		FirstName: "Selim",
		LastName:  "Ock",
		UserName:  "ocakdans",
		Gender:    "male",
		Name: Name{
			Family:   "Ock",
			Personal: "Selim",
		},
		Email: []Email{
			Email{ID: 1, Kind: "home", Address: "home@gmail.com"},
			Email{ID: 2, Kind: "work", Address: "work@gmail.com"},
		},
		Interest: []Interest{
			Interest{ID: 1, Name: "Golang"},
			Interest{ID: 2, Name: "Docker"},
			Interest{ID: 3, Name: "Kubernetes"},
		},
	}
	WriteMessage("Reading operariton started")
	WriteMessage("Personal Fullname")
	WriteStarLine()
	res := GetPerson(&p)
	WriteMessage(res)
	WriteStarLine()
	WriteMessage("\n")

	WriteMessage("Personal Email with Index")
	WriteStarLine()
	resEmail := GetPersonEmailAddress(&p, 1)
	WriteMessage(resEmail)
	WriteStarLine()
	WriteMessage("\n")

	WriteMessage("Personal Email Object with Index")
	WriteStarLine()
	resEmailObj := GetPersonEmail(&p, 1)
	fmt.Println(resEmailObj)
	//WriteMessage(resEmailObj)
	WriteStarLine()

	WriteMessage("Reading operariton completed")
	WriteMessage("\n")

	WriteMessage("Writing operariton started")
	SaveJson("person.json", p)
	WriteMessage("Writing operariton completed")

}

type Name struct {
	Family   string
	Personal string
}

type Email struct {
	ID      int
	Kind    string
	Address string
}

type Interest struct {
	ID   int
	Name string
}

type Person struct {
	ID        int
	FirstName string
	LastName  string
	UserName  string
	Gender    string
	Name      Name
	Email     []Email
	Interest  []Interest
}

func GetPerson(p *Person) string {
	return p.FirstName + " " + p.LastName
}

func GetPersonEmailAddress(p *Person, i int) string {
	return p.Email[i].Address
}

func GetPersonEmail(p *Person, i int) Email {
	return p.Email[i]
}

func WriteMessage(msg string) {
	fmt.Println(msg)
}

func WriteStarLine() {
	fmt.Println("*********")
}

func checkError(err error) {
	if err != nil {
		fmt.Println("Fatal Error: ", err.Error())
		os.Exit(1)
	}
}

func SaveJson(fileName string, key interface{}) {
	outFile, err := os.Create(fileName)
	checkError(err)
	encoder := json.NewEncoder(outFile)
	err = encoder.Encode(key)
	checkError(err)
	outFile.Close()
}
