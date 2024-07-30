package main

import (
	"encoding/json"
	"fmt"
)

func main() {
	jsonStr := `
	{
		"data": {
			"object": "card",
			"id": "card_1H9gYz2eZvKYlo2C5Z2eZvKY",
			"first_name": "Selim",
			"last_name": "Ock",
			"balance": "100"
		}

	}`

	var m map[string]map[string]interface{}
	// Unmarshalling is the process of converting JSON data into a Go data structure.
	// The &m is a pointer to the map where the unmarshalled data will be stored.
	if err := json.Unmarshal([]byte(jsonStr), &m); err != nil {
		panic(err)
	}
	//This prints the Go map which now contains the data from the JSON string.
	fmt.Println(m)
	fmt.Println("*********")

	// Marshalling is the process of converting a Go data structure into JSON data.
	b, err := json.Marshal(m)
	if err != nil {
		panic(err)
	}
	fmt.Println(string(b))

}
