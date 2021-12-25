package test

import (
	"github.com/morristai/rarbg-notifier/client"
	c "github.com/morristai/rarbg-notifier/common"
	"github.com/morristai/rarbg-notifier/parser"
	"os"
	"testing"
)

func TestParserHighest(t *testing.T) {
	url := "https://rarbg.to/torrents.php?category=14;48;17;44;45;47;50;51;52;42;46;54&search=1080p&order=seeders&by=DESC"
	content := client.Request(url)
	//contentBytes, _ := io.ReadAll(content.Body)
	//os.WriteFile("./cache/rarbg_response.html", contentBytes, 0644)
	leaderBoard := parser.ParseHighest(content.Body)
	if len(leaderBoard.VideoList) != 15 {
		t.Error("movie list should be 15")
	}
}

func TestMarshal(t *testing.T) {
	var leaderBoard c.LeaderboardCache
	file, _ := os.Open("./cache/rarbg_response.html")
	res := parser.ParseHighest(file)
	err := parser.Save("./cache/1080p.json", res)
	if err != nil {
		t.Error(err)
	}
	err = parser.Load("./cache/1080p.json", &leaderBoard)
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

func TestAll(t *testing.T) {
	TestParserHighest(t)
	TestEquality(t)
	TestEquality(t)
}
