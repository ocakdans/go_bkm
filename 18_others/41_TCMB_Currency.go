package others

import (
	"encoding/xml"
	"fmt"
	"net/http"
	"time"
)

type tarih_date struct {
	XMLName   xml.Name `xml: "Tarih_Date"`
	Tarih     string   `xml: "Tarih",attr`
	Date      string   `xml: "Date",attr`
	Bulten_No string   `xml: "Bulten_No",attr`
	Currency  []currency_tcmb
}

type currency_tcmb struct {
	CrossOrder      string `xml: "CrossOrder",attr`
	Kod             string `xml: "Kod",attr`
	CurrencyCode    string `xml: "CurrencyCode",attr`
	Unit            string `xml: "Unit"`
	Isim            string `xml: "Isim"`
	CurrencyName    string `xml: "CurrencyName"`
	ForexBuying     string `xml: "ForexBuying"`
	ForexSelling    string `xml: "ForexSelling"`
	BanknoteBuying  string `xml: "BanknoteBuying"`
	BanknoteSelling string `xml: "BanknoteSelling"`
	CrossRateUSD    string `xml: "CrossRateUSD"`
	CrossRateOther  string `xml: "CrossRateOther"`
}

type Currency struct {
	CrossOrder      int
	Code            string
	Unit            int
	CurrencyNameTR  string
	CurrencyName    string
	ForexBuying     float64
	ForexSelling    float64
	BanknoteBuying  float64
	BanknoteSelling float64
	CrossRateUSD    float64
	CrossRateOther  float64
}

type CurrencyDay struct {
	ID         string
	Date       time.Time
	DayNo      string
	Currencies []Currency
}

func (c *tarih_date) GetData(CurrencyDate time.Time) {
	xDate := CurrencyDate
	t := new(tarih_date)
	currOfDay := t.getData(CurrencyDate, xDate)
	for {
		if currOfDay == nil {
			CurrencyDate = CurrencyDate.AddDate(0, 0, -1)
			currOfDay = t.getData(CurrencyDate, xDate)
			if currOfDay != nil {
				break
			}
		} else {
			break
		}
	}
}

func (t *tarih_date) getData(CurrencyDate time.Time, xDate time.Time) *CurrencyDay {
	currOfDay := new(CurrencyDay)
	var res *http.Response
	var err error
	var url string
	url = "http://www.tcmb.gov.tr/kurlar/" + CurrencyDate.Format("200601") + "/" + CurrencyDate.Format("02012006") + ".xml"
	res, err = http.Get(url)
	if err != nil {
		fmt.Println("Error: ", err.Error())
	} else {
		defer res.Body.Close()
	}

	if res.StatusCode != http.StatusNotFound {
		decoder := xml.NewDecoder(res.Body)
		err = decoder.Decode(t)
		if err != nil {
			fmt.Println("Error: ", err.Error())
		}

		currOfDay.ID = t.Tarih + t.Date

	}

	return currOfDay
}

func checkError(err error) {
	if err != nil {
		fmt.Println("Fatal Error: ", err.Error())
	}
}
