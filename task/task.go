package task

import (
	"bytes"
	"github.com/bwmarrin/discordgo"
	"github.com/go-co-op/gocron"
	"github.com/morristai/rarbg-notifier/client"
	c "github.com/morristai/rarbg-notifier/common"
	"github.com/morristai/rarbg-notifier/parser"
	"io"
	"io/ioutil"
	"log"
	"os"
	"time"
)

func CronJobs(dg *discordgo.Session) {
	s := gocron.NewScheduler(time.UTC)
	s.Every(30).Minute().Tag("1080p", "movie").Do(highest1080p, dg)
	//s.Every(30).Minute().Tag("2160p", "movie").Do(highest2160p, dg)
	s.RunByTag("1080p")
	log.Println("Start cron job by tags: [1080p]")
	s.StartAsync() // starts the scheduler and blocks current execution path
}

func highest1080p(dg *discordgo.Session) {
	config := c.ReadConfig("./resources/application.yml")
	var preLeaderBoard *c.LeaderboardCache
	err := parser.Load(config.Cache.Json1080p, &preLeaderBoard)
	if err != nil {
		log.Panicln(err)
	}
	curLeaderBoard, leaderBoard, top9, err := processor(config, config.Rarbg.Movie.Url1080p, config.Cache.Json1080p, preLeaderBoard)
	if err != nil {
		log.Panicln(err)
	}
	curLeaderBoard.Notified = preLeaderBoard.Notified // avoid further for loop append
	// Leader Board
	if len(leaderBoard) != 0 {
		log.Printf("Found %d 1080p new in leader board", len(leaderBoard))
		for _, m := range leaderBoard {
			if _, ok := curLeaderBoard.Notified[m.Title]; !ok {
				message := m.GenDiscordMessage()
				for _, r := range config.Discord.Channels { // send to multiple channels
					_, err := dg.ChannelMessageSendEmbed(r, &message)
					if err != nil {
						log.Fatalln(err)
					}
				}
				curLeaderBoard.Notified[m.Title] = time.Now() // set Discord notified
				log.Printf("Send [%s] to Discord %v successful\n", m.Title, config.Discord.Channels)
			} else {
				log.Println(m.Title, "is already been notified")
			}
		}
	} else {
		log.Println("No new movie found in [1080p Leader Board]")
	}
	// Newest Top 9
	if len(top9) > 0 {
		log.Printf("Found %d movies in Top9", len(top9))
		for _, m := range top9 {
			// Check if notified exist already
			if _, ok := curLeaderBoard.Notified[m.Title]; !ok {
				message := m.GenDiscordMessage()
				for _, r := range config.Discord.Channels {
					_, err := dg.ChannelMessageSendEmbed(r, &message)
					if err != nil {
						log.Fatalln(err)
					}
				}
				log.Printf("Send [%s] to Discord %v successful\n", m.Title, config.Discord.Channels)
				curLeaderBoard.Notified[m.Title] = time.Now()
			} else {
				log.Println(m.Title, "already been notified")
			}
		}
	} else {
		log.Println("No new movie found in [1080p Leader Board]")
	}

	// TODO: scheduler to clear cache
	err = parser.Save(config.Cache.Json1080p, curLeaderBoard)
	if err != nil {
		log.Fatalln(err)
	}
	log.Println("Task [Highest 1080p] done; Session Time: ", curLeaderBoard.Time)
}

func highest2160p(dg *discordgo.Session) {

}

// TODO: adopt different task with interface
func processor(config *c.Config, url string, cachePath string, preLeaderBoard *c.LeaderboardCache) (*c.LeaderboardCache, []*c.VideoInfo, []*c.VideoInfo, error) {
	var curLeaderBoard *c.LeaderboardCache
	var newFoundLeaderBoard []*c.VideoInfo
	var newFoundTop9 []*c.VideoInfo

	content := client.RequestRarbg(url, config.Rarbg.Headers.Cookie)
	// Persist Response HTML to local cache
	if cachePath == "./cache/1080p.json" {
		// Since response body will be change during parsing, so we clone to another one.
		buf, _ := ioutil.ReadAll(content.Body)
		tmpIoReader := ioutil.NopCloser(bytes.NewBuffer(buf))
		originIoReader := ioutil.NopCloser(bytes.NewBuffer(buf))
		persistResponse(tmpIoReader)
		content.Body = originIoReader
	}
	curLeaderBoard, err := parser.ParseHomePage(preLeaderBoard, content.Body, config.Rarbg.Headers.Cookie, config.Imdb.Cookie)
	if err != nil {
		return nil, nil, nil, err
	}
	// Compare LeaderBoard VideoList
	for k := range curLeaderBoard.VideoList {
		_, ok := preLeaderBoard.VideoList[k]
		if !ok {
			newFoundLeaderBoard = append(newFoundLeaderBoard, curLeaderBoard.VideoList[k])
		}
	}
	// Compare The Newest Top 9 VideoList
	if cachePath == "./cache/1080p.json" {
		for k := range curLeaderBoard.Newest9 {
			_, ok := preLeaderBoard.Newest9[k]
			if ok == true {
				log.Println(k, "is already in the leader board")
			} else {
				newFoundTop9 = append(newFoundTop9, curLeaderBoard.Newest9[k])
			}
		}
	}
	return curLeaderBoard, newFoundLeaderBoard, newFoundTop9, nil
}

func persistResponse(r io.ReadCloser) {
	outFile, err := os.Create("./cache/rarbg_response.html")
	defer outFile.Close()
	_, err = io.Copy(outFile, r)
	if err != nil {
		log.Panicln(err)
	}
}

// TODO: TvShows
// func highestTvShows() {
//
// }
