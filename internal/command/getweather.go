package command

import (
	"encoding/json"
	"fmt"
	"pogoda/internal/product"
)

func GetWeather(lat float64, lon float64) product.Jandex {
	m := &product.Jandex{}
	url := fmt.Sprintf("https://api.weather.yandex.ru/v2/informers/?lat=%f&lon=%f&lang=ru_RU", lat, lon)
	body := Zapros(url)
	json.Unmarshal(body, m)
	return *m
}
