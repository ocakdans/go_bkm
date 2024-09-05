package main

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
)

type API2 struct {
	Message string `"json:message"`
}

func Hello(w http.ResponseWriter, r *http.Request) {
	urlParams := mux.Vars(r)
	id := urlParams["id"]
	messageContent := "Kullanıcı ID: " + id

	message := API2{Message: messageContent}
	output, err := json.Marshal(message)
	if err != nil {
		http.Error(w, err.Error(), 500)
		return
	}
	fmt.Fprintf(w, string(output))

}

func main() {
	gorillaRoute := mux.NewRouter()
	gorillaRoute.HandleFunc("/api/user/{id:[0-9]}", Hello)
	http.Handle("/", gorillaRoute)
	http.ListenAndServe(":9090", nil)
}
