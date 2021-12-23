package parser

import (
	"fmt"
	"github.com/PuerkitoBio/goquery"
	c "github.com/morristai/rarbg-notifier/common"
	log "github.com/sirupsen/logrus"
	"io"
	"regexp"
	"strings"
)

var (
	jpgUrl        string
	highestList   []c.VideoInfo
	billboardList []string
)

func ParseHighest(res io.Reader) []c.VideoInfo {
	doc, err := goquery.NewDocumentFromReader(res)
	if err != nil {
		log.Fatal(err)
	}

	//body > table:nth-child(6) > tbody > tr > td:nth-child(2) > div > table > tbody > tr:nth-child(2) > td > table.lista2t > tbody > tr:nth-child(2) > td:nth-child(2) > a:nth-child(1)
	selectors := "table:nth-child(6) td:nth-child(2) table.lista2t tr.lista2 td:nth-child(2) a:nth-child(1)"

	doc.Find(selectors).Each(func(i int, s *goquery.Selection) {
		var video c.VideoInfo
		r, _ := regexp.Compile("^.*(https.*jpg).*$")
		originUrl, ok := s.Attr("onmouseover")
		if ok {
			jpgUrl = r.FindStringSubmatch(originUrl)[1]
		}
		title := s.Contents().Text()
		if title != "" {
			video.Title = title
			video.Year = 2021
			video.Poster = jpgUrl
		}
		highestList = append(highestList, video)
	})
	if len(highestList) == 0 {
		log.Error("Highest list Not found!")

	} else {
		for idx, i := range highestList {
			fmt.Printf("%d: %s %s %s\n", idx, i.Title, i.Poster, i.Year)
		}
	}
	return highestList
}

func ParseBillboard(res io.Reader) []string {
	doc, err := goquery.NewDocumentFromReader(res)
	if err != nil {
		log.Fatal(err)
	}

	// body > table:nth-child(6) > tbody > tr > td:nth-child(2) > div > table > tbody > tr:nth-child(2) > td > div:nth-child(1) > table > tbody > tr > td:nth-child(6) > a > img
	selectors := "table:nth-child(6) td:nth-child(2) tr:nth-child(2) div:nth-child(1) table tbody tr td a" // Billboard

	doc.Find(selectors).Each(func(i int, s *goquery.Selection) {
		movieName, ok := s.Attr("title")
		if ok {
			title := strings.TrimSpace(movieName)
			billboardList = append(billboardList, title)
		}
	})
	if len(billboardList) == 0 {
		log.Error("Rarbg Billboard Not found!")
	} else {
		//for idx, i := range highestList {
		//	fmt.Printf("%d: %s\n", idx, i)
		//}
	}
	return billboardList
}
