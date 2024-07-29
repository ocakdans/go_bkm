package maps

import "fmt"

func Map() {

	dictionary := make(map[string]string)
	dictionary["table"] = "masa"
	dictionary["book"] = "kitap"
	dictionary["pencil"] = "kalem"

	fmt.Println(dictionary["table"])
	fmt.Println("number of keys", len(dictionary))
	fmt.Println("dictionary:", dictionary)

	// silme işlemi
	delete(dictionary, "book")
	fmt.Println("number of keys", len(dictionary))
	fmt.Println("dictionary:", dictionary)

	value, isThere := dictionary["table"]
	fmt.Println("value: ", value)
	fmt.Println("is there: ", isThere) // to check if it exists in the map.

	dictionary2 := map[string]string{"glass": "cam", "hand": "el"}
	fmt.Println(dictionary2)

}
