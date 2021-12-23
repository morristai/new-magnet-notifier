package parser

import (
	"fmt"
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

func MatchInfo(origin string) c.VideoInfo {
	//Christmas.with.the.Chosen.The.Messengers.2021.1080p.WEBRip.x264-RARBG
	r, _ := regexp.Compile("(^\\S*).((?:19|20)[0-9]{2})(.*)")
	titleYear := r.FindStringSubmatch(origin)
	title = strings.Replace(titleYear[1], ".", " ", -1)
	fmt.Println(titleYear[2])
	year, _ = strconv.Atoi(titleYear[2])
	others := titleYear[3]
	// Resolution
	r, _ = regexp.Compile(".*((?:720|1080|2160)p).*")
	ifResolution := r.FindStringSubmatch(others)
	if len(ifResolution) != 0 {
		resolution = ifResolution[1]
	}
	// Encoding

	fmt.Println(resolution)

	video := c.VideoInfo{
		Title:      title,
		Year:       year,
		Resolution: resolution,
	}
	return video
}
