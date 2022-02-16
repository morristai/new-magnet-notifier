package test

import (
	"github.com/morristai/rarbg-notifier/client"
	c "github.com/morristai/rarbg-notifier/common"
	"github.com/morristai/rarbg-notifier/parser"
	"os"
	"testing"
)

func TestParserHighest(t *testing.T) {
	config := c.ReadConfig("./resources/application.yml")
	content := client.RequestRarbg(config.Rarbg.Movie.Url1080p, config.Cache.Json1080p)
	var preLeaderBoard *c.LeaderboardCache
	err := parser.Load(config.Cache.Json1080p, &preLeaderBoard)
	// contentBytes, _ := io.ReadAll(content.Body)
	// os.WriteFile("./cache/rarbg_response.html", contentBytes, 0644)
	leaderBoard, _ := parser.ParseHomePage(preLeaderBoard, content.Body, config.Rarbg.Headers.Cookie, config.Imdb.Cookie)
	if len(leaderBoard.VideoList) != 15 {
		t.Error("movie list should be 15")
	}
	err = parser.Save("./cache/1080p.json", leaderBoard)
	if err != nil {
		t.Error("Save to local cache error", err)
	}
}

func TestMarshal(t *testing.T) {
	config := c.ReadConfig("./resources/application.yml")
	var leaderBoard *c.LeaderboardCache
	var preLeaderBoard *c.LeaderboardCache
	file, _ := os.Open("./cache/rarbg_response.html")
	res, _ := parser.ParseHomePage(preLeaderBoard, file, config.Rarbg.Headers.Cookie, config.Imdb.Cookie)
	err := parser.Save("./cache/1080p.json", res)
	if err != nil {
		t.Error(err)
	}
	err = parser.Load("./cache/1080p.json", leaderBoard)
	if err != nil {
		t.Error(err)
	}
	if len(leaderBoard.VideoList) != 14 {
		t.Error("movie list should be 14")
	}
}

func TestEquality(t *testing.T) {
	var leaderBoard c.LeaderboardCache
	var leaderBoardNext c.LeaderboardCache
	parser.Load("./cache/1080p.json", &leaderBoard)
	parser.Load("./cache/1080p_next.json", &leaderBoardNext)
	for k := range leaderBoardNext.VideoList {
		_, ok := leaderBoard.VideoList[k]
		if !ok {
			if k != "Dune" {
				t.Error("New movie should be Dune")
			}
		}
	}
}
