package main

import (
	"bytes"
	"html/template"
	"io/ioutil"
	"net/http"
)

type Page struct {
	Title           string
	Author          string
	Header          string
	PageDescription string
	Content         string
	URI             string
}

func loadFile(filename string) (string, error) {
	bytes, err := ioutil.ReadFile(filename)
	if err != nil {
		return "", err

	}
	return string(bytes), nil
}

// for 46_page.html

// func handler(w http.ResponseWriter, r *http.Request) {
// 	var body, _ = loadFile("46_page.html")
// 	fmt.Fprintf(w, body)
// }

// for 46_template.html

func handler(w http.ResponseWriter, r *http.Request) {
	// String birleştirme işlemi
	var builder bytes.Buffer
	builder.WriteString("KodLab yayinevinden çikardiğimiz Yazilimcilar İçin İleri Seviye T-SQL kitabimin özellikleri aşağidaki gibidir;\n")
	builder.WriteString("704 Sayfa\n")
	builder.WriteString("ISBN: 9.786.055.201.142\n")
	builder.WriteString("Fiyat: 37 TL\n")
	builder.WriteString("Boyut: 15 x 21\n")
	builder.WriteString("2. Baski\n")

	uri := "https://www.kodlab.com.tr/kitaplar/yazilimcilar-icin-ileri-seviye-t-sql"

	page := Page{
		Title:           "Yazilimcilar İçin İleri Seviye T-SQL",
		Author:          "Mehmet Salih Deveci",
		Header:          "Yazilimcilar İçin İleri Seviye T-SQL",
		PageDescription: "Yazilimcilar İçin İleri Seviye T-SQL kitabi hakkinda detayli bilgiler",
		Content:         builder.String(),
		URI:             uri,
	}
	t, _ := template.ParseFiles("46_template.html")
	t.Execute(w, page)
}

func main() {
	http.HandleFunc("/", handler)
	http.ListenAndServe(":9000", nil)
}
