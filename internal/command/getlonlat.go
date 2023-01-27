package command

import (
	"encoding/csv"
	"io"
	"io/ioutil"
	"log"
	"pogoda/internal/product"
	"strings"
)

func Getlonlat(name string) (lon string, lat string) {
	content, err := ioutil.ReadFile("cities.csv")
	if err != nil {
		log.Fatal(err)
	}
	r := csv.NewReader(strings.NewReader(string(content)))
	var Goroda product.Citys
	for {
		record, err := r.Read()
		if err == io.EOF {
			break
		}
		if err != nil {
			log.Fatal(err)
		}
		Goroda = append(Goroda, product.City{Name: record[6], Lon: record[17], Lat: record[18]})
	}
	for _, mesto := range Goroda {
		if mesto.Name == name {
			lon := mesto.Lon
			lat := mesto.Lat
			return lon, lat
		}
	}
	return
}
