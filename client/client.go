package client

import (
	"log"
	"net/http"
)

func Request(url string) *http.Response {
	header := NewHeader()
	client := &http.Client{}
	req, err := http.NewRequest(http.MethodGet, url, nil)
	req.Header = header
	resp, err := client.Do(req)
	if err != nil {
		log.Fatalln(err)
	} else if resp.StatusCode != 200 {
		log.Fatalln("Request Error, Response StatusCode: ", resp.StatusCode)
	} else {
		log.Println("Request Successful")
	}
	return resp
}
