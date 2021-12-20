package parser

import (
	"fmt"
	"github.com/PuerkitoBio/goquery"
	log "github.com/sirupsen/logrus"
	"net/http"
	"strings"
)

func Parse(res *http.Response) {
	doc, err := goquery.NewDocumentFromReader(res.Body)
	if err != nil {
		log.Fatal(err)
	}

	var arr []string
	//body > table:nth-child(6) > tbody > tr > td:nth-child(2) > div > table > tbody > tr:nth-child(2) > td > div:nth-child(1) > table > tbody > tr > td:nth-child(2) > a > img
	//body > table:nth-child(6) > tbody > tr > td:nth-child(2) > div > table > tbody > tr:nth-child(2) > td > div:nth-child(1) > table > tbody > tr > td:nth-child(6) > a > img
	doc.Find("table:nth-child(6) td:nth-child(2) tr:nth-child(2) div:nth-child(1) table tbody tr td").Each(func(i int, s *goquery.Selection) {
		movieName, ok := s.Find("a").Attr("title")
		if ok {
			title := strings.TrimSpace(movieName)
			arr = append(arr, title)
		}
	})
	for _, i := range arr {
		fmt.Println(i)
	}
}
