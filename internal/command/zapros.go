package command

import (
	"io/ioutil"
	"log"
	"net/http"
)

func Zapros(url string) []byte {
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		log.Println("Error on response.\n[ERROR] -", err)
	}

	// add authorization header to the req
	req.Header.Add("X-Yandex-API-Key", "---")

	// Send req using http Client
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		log.Println("Error on response.\n[ERROR] -", err)
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Println("Error while reading the response bytes:", err)
	}
	return body

}
