package client

import (
	"fmt"
	"log"
	"net/http"
)

func RequestRarbg(url string, cookie string) *http.Response {
	header := NewHeader(cookie)
	client := &http.Client{}
	req, err := http.NewRequest(http.MethodGet, url, nil)
	req.Header = header
	resp, err := client.Do(req)
	if err != nil {
		log.Fatalln(err)
	} else if resp.StatusCode != 200 {
		log.Fatalln("Request Rarbg Error, Response StatusCode: ", resp.StatusCode)
	} else {
		// Peak response body
		// bodyBytes, _ := io.ReadAll(resp.Body)
		// fmt.Println(string(bodyBytes))
		log.Println("Request Rarbg Successful")
	}
	return resp
}

func RequestImdb(movieCode string, cookie string) *http.Response {
	header := NewImdbHeader(cookie)
	client := &http.Client{}
	url := fmt.Sprintf("https://www.imdb.com/title/%s/reviews?sort=reviewVolume&dir=desc&ratingFilter=0", movieCode)
	req, err := http.NewRequest(http.MethodGet, url, nil)
	req.Header = header
	resp, err := client.Do(req)
	if err != nil {
		log.Fatalln(err)
	} else if resp.StatusCode != 200 {
		log.Fatalln("Request IMDB Error, Response StatusCode: ", resp.StatusCode)
	} else {
		// TODO: log control certain packages
		log.Println("Request IMDB Successful")
	}
	return resp
}
