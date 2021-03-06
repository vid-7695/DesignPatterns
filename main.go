package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

func main() {
	endpoint := "numbersapi.p.rapidapi.com/6/21/date"
	req, _ := http.NewRequest(http.MethodGet, fmt.Sprintf("https://%s", endpoint), nil)
	addHeader(req)
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		log.Fatalln(err)
	}
	//We Read the response body on the line below.
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Fatalln(err)
	}
	//Convert the body to type string
	sb := string(body)
	log.Printf(sb)
}

func addHeader(request *http.Request) {
	request.Header.Add("X-RapidAPI-Key", "f0b69255acmsh832c370eaa4f68ap17e30bjsndaac3e61aa04")
	request.Header.Add("X-RapidAPI-Host", "numbersapi.p.rapidapi.com")
}
