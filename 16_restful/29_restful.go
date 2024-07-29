package restful

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

type Todo struct {
	UserId    int    `json: "userId"`
	Id        int    `json: "id"`
	Title     string `json: "title"`
	Completed bool   `json: "completed"`
}

func GetRestful() {
	response, err := http.Get("https://jsonplaceholder.typicode.com/todos/1")
	if err != nil {
		fmt.Println(err)
	}
	defer response.Body.Close()
	bodyBytes, _ := ioutil.ReadAll(response.Body) // byte formatında getiriyor.
	bodyString := string(bodyBytes)               // byte ı stringe çevir.
	fmt.Println(bodyString)

	var todo Todo
	json.Unmarshal(bodyBytes, &todo)
	fmt.Println(todo)
}

func PostRestful() {
	postTodo := Todo{1, 2, "Alışveriş", true}
	jsonTodo, err := json.Marshal(postTodo) //json a çeviricez
	if err != nil {
		fmt.Println(err)
	}
	response, err := http.Post("https://jsonplaceholder.typicode.com/todos",
		"application/json;charset=utf-8",
		bytes.NewBuffer(jsonTodo))

	if err != nil {
		fmt.Println(err)
	}

	defer response.Body.Close()

	bodyBytes, _ := ioutil.ReadAll(response.Body) // byte formatında getiriyor.
	bodyString := string(bodyBytes)               // byte ı stringe çevir.
	fmt.Println(bodyString)

	var todoResponse Todo
	json.Unmarshal(bodyBytes, &todoResponse)
	fmt.Println(todoResponse)
}
