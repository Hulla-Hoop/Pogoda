package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"pogoda/internal/command"
	"strconv"
)

func main() {
	fmt.Println("Введите название города")
	var gorod string
	fmt.Scanln(&gorod)
	lon, lat := command.Getlonlat(gorod)
	lon2, _ := strconv.ParseFloat(lon, 32)
	lat2, _ := strconv.ParseFloat(lat, 32)
	m := command.GetWeather(lon2, lat2)
	fmt.Println("Температура в", gorod, "сейчас ", m.Fact.Temp)
	/* fmt.Printf("struct:\n\t%#v\n\n", m) */
	bodys, err := json.MarshalIndent(m, "", " ")
	if err != nil {
		log.Fatal("Json marshaling failed:", err)
	}
	err = ioutil.WriteFile("dannie.json", bodys, 0)
	if err != nil {
		log.Fatal("cannot write updated settings file:", err)
	}
}
