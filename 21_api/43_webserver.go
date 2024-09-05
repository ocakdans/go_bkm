package main

import (
	"fmt"
	"net/http"
	"strconv"
)

type Human struct {
	Fname string
	Lname string
	Age   int
}

func (human Human) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	human.Fname = "Selim"
	human.Lname = "Feyza"
	human.Age = 30

	// Formu parse etmek için
	r.ParseForm()

	// Sunucudan form bilgisini almak için
	fmt.Println(r.Form)

	// URL path bilgisini almak için
	fmt.Println("path", r.URL.Path)

	fmt.Fprintf(w, human.Fname+human.Lname+strconv.Itoa(human.Age))
}

func main() {

	var human Human
	err := http.ListenAndServe(":9000", human)
	checkError(err)

}

func checkError(err error) {
	if err != nil {
		panic(err)
	}
}
