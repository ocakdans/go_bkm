package restful

import (
	"bytes"
	"encoding/json"
	"io/ioutil"
	"net/http"
)

type Product struct {
	Id          string  `json:"id"`
	ProductName string  `json:"productName"`
	CategoryId  int     `json:"categoryId"`
	UnitPrice   float64 `json:"unitPrice"`
}

type Category struct {
	Id           int    `json:"id"`
	CategoryName string `json:"categoryName"`
}

func GetAllProducts() ([]Product, error) {
	res, err := http.Get("http://localhost:3000/products")
	if err != nil {
		return nil, err
	}
	defer res.Body.Close()
	bodyBytes, _ := ioutil.ReadAll(res.Body)
	var products []Product
	json.Unmarshal(bodyBytes, &products)
	return products, nil
}

func AddProduct() (Product, error) {
	product := Product{"6", "Microphone", 1, 10000}
	jsonProduct, _ := json.Marshal(product)
	res, err := http.Post("http://localhost:3000/products",
		"application/json;charset:utf-8",
		bytes.NewBuffer(jsonProduct))
	if err != nil {
		return Product{}, err
	}

	defer res.Body.Close()
	bodyBytes, _ := ioutil.ReadAll(res.Body)
	var productResponse Product
	json.Unmarshal(bodyBytes, &productResponse)
	//fmt.Println("saved to db", productResponse)
	return productResponse, nil

}
