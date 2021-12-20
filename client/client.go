package client

import (
	log "github.com/sirupsen/logrus"
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
		log.Error("Response Status: ", resp.StatusCode)
	} else {
		log.Debug("Successful get response")
	}
	return resp
}
