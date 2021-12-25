package parser

import (
	"fmt"
	"github.com/PuerkitoBio/goquery"
	c "github.com/morristai/rarbg-notifier/common"
	"io"
	"log"
	"regexp"
	"strings"
	"time"
)

var (
	leaderboard   c.LeaderboardCache
	billboardList []string // only Title
)

func ParseHighest(res io.Reader) *c.LeaderboardCache {
	leaderboard.VideoList = map[string]*c.VideoInfo{} // Init
	doc, err := goquery.NewDocumentFromReader(res)
	if err != nil {
		log.Fatal(err)
	}
	//body > table:nth-child(6) > tbody > tr > td:nth-child(2) > div > table > tbody > tr:nth-child(2) > td > table.lista2t > tbody > tr:nth-child(2) > td:nth-child(2) > a:nth-child(1)
	baseSelectors := "table:nth-child(6) td:nth-child(2) table.lista2t tr.lista2 td:nth-child(2) a:nth-child(1)"
	// loop each movie
	doc.Find(baseSelectors).Each(func(i int, s *goquery.Selection) {
		var video c.VideoInfo
		// title, year, resolution
		title := s.Contents().Text()
		MatchBasic(title, &video)
		// genre
		MatchGenre(s.Siblings().Text(), &video)
		// rating
		MatchRating(s.Siblings().Text(), &video)
		// size
		size := s.Parent().SiblingsFiltered("[width=\"100px\"]")
		video.Size = size.Text()
		// IMDB
		r, _ := regexp.Compile("^.*imdb=(\\S*)$")
		imdb, ok := s.Siblings().Attr("href")
		if ok {
			// https://www.imdb.com/title/tt13207508/
			imdbCode := r.FindStringSubmatch(imdb)[1]
			video.Imdb = fmt.Sprintf("https://www.imdb.com/title/%s", imdbCode)
		}
		// poster URL
		r, _ = regexp.Compile("^.*(https.*jpg).*$")
		posterUrl, ok := s.Attr("onmouseover")
		if ok {
			video.Poster = r.FindStringSubmatch(posterUrl)[1]
		}
		// Rarbg URL
		video.Url, ok = s.Attr("href")
		video.Url = fmt.Sprintf("https://rarbg.to%s", video.Url)

		leaderboard.VideoList[video.Title] = &video
	})
	leaderboard.Time = time.Now()
	return &leaderboard
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
		log.Panicln("Rarbg Billboard Not found!")
	} else {
		//for idx, i := range highestList {
		//	fmt.Printf("%d: %s\n", idx, i)
		//}
	}
	return billboardList
}
