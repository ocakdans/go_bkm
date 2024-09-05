package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"os"
)

type API struct {
	Message string "json:message"
}

type User struct {
	ID        int    "json:id"
	FirstName string "json:firstname"
	LastName  string "json:lastname"
	Age       int    "json:age"
}

func main() {
	apiRoot := "/api"
	http.HandleFunc(apiRoot, func(w http.ResponseWriter, r *http.Request) {
		message := API{"API Home"}
		output, err := json.Marshal(message)
		checkError2(err)
		//w.Header().Set("Content-Type", "application/json")
		fmt.Fprint(w, string(output))

	})

	// .../api/users
	http.HandleFunc(apiRoot+"/users", func(w http.ResponseWriter, r *http.Request) {
		users := []User{
			User{ID: 1, FirstName: "Selim", LastName: "Feyza", Age: 30},
			User{ID: 2, FirstName: "Feyza", LastName: "Selim", Age: 25},
			User{ID: 2, FirstName: "Beyza", LastName: "Selim", Age: 25},
		}
		message := users
		output, err := json.Marshal(message)
		checkError2(err)
		//w.Header().Set("Content-Type", "application/json")
		fmt.Fprint(w, string(output))
	})

	// .../api/me
	http.HandleFunc(apiRoot+"/me", func(w http.ResponseWriter, r *http.Request) {
		user := User{ID: 1, FirstName: "Selim", LastName: "Feyza", Age: 30}
		message := user
		output, err := json.Marshal(message)
		checkError2(err)
		//w.Header().Set("Content-Type", "application/json")
		fmt.Fprint(w, string(output))
	})

	http.ListenAndServe(":9090", nil)
}

func checkError2(err error) {
	if err != nil {
		fmt.Println("Fatal Error", err.Error())
		os.Exit(1)
	}
}
