package parser

import (
	c "github.com/morristai/rarbg-notifier/common"
	"regexp"
	"strconv"
	"strings"
)

var (
	title      string
	year       int
	resolution string
)

func MatchBasic(origin string, video *c.VideoInfo) {
	//Christmas.with.the.Chosen.The.Messengers.2021.1080p.WEBRip.x264-RARBG
	r, _ := regexp.Compile("(^\\S*).((?:19|20)[0-9]{2})(.*)")
	titleYear := r.FindStringSubmatch(origin)
	// TODO: index error handling
	title = strings.Replace(titleYear[1], ".", " ", -1)
	year, _ = strconv.Atoi(titleYear[2])
	others := titleYear[3]
	// Resolution
	r, _ = regexp.Compile(".*((?:720|1080|2160)p).*")
	ifResolution := r.FindStringSubmatch(others)
	if len(ifResolution) != 0 {
		resolution = ifResolution[1]
	}
	// TODO: Encoding

	video.Title = title
	video.Year = year
	video.Resolution = resolution
}

func MatchGenre(origin string, video *c.VideoInfo) {
	r, _ := regexp.Compile("\\sIMDB.*")
	genre := r.ReplaceAllString(origin, "")
	video.Genre = genre
}

func MatchRating(origin string, video *c.VideoInfo) {
	r, _ := regexp.Compile("^.*IMDB:\\s(.*)$")
	ifRating := r.FindStringSubmatch(origin)
	if len(ifRating) != 0 {
		video.Rating = ifRating[1]
	}
}
