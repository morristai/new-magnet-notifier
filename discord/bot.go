package discord

import (
	"fmt"
	"github.com/bwmarrin/discordgo"
	c "github.com/morristai/rarbg-notifier/common"
	"github.com/morristai/rarbg-notifier/parser"
	"log"
	"os"
	"os/signal"
	"syscall"
	"time"
)

func Run() {
	Token := genToken()
	// Create a new Discord session using the provided bot token.
	dg, err := discordgo.New("Bot " + Token)
	if err != nil {
		fmt.Println("error creating Discord session,", err)
		return
	}

	// Register the messageCreate func as a callback for MessageCreate events.
	dg.AddHandler(messageCreate)

	// In this example, we only care about receiving message events.
	dg.Identify.Intents = discordgo.IntentsGuildMessages

	// Open a websocket connection to Discord and begin listening.
	err = dg.Open()
	if err != nil {
		fmt.Println("error opening connection,", err)
		return
	}

	// for testing
	channelID := "922537842083762237"
	message := "```go\nimport \"fmt\"```"
	time.Sleep(5)
	dg.ChannelMessageSend(channelID, message)

	// Wait here until CTRL-C or other term signal is received.
	fmt.Println("Bot is now running.  Press CTRL-C to exit.")
	sc := make(chan os.Signal, 1) // channel
	signal.Notify(sc, syscall.SIGINT, syscall.SIGTERM, os.Interrupt, os.Kill)
	<-sc

	// Cleanly close down the Discord session.
	dg.Close()
}

func genContent(cache *c.LeaderboardCache) discordgo.MessageEmbed {
	oneMovie := cache.VideoList["Dune"]
	return oneMovie.GenDiscordMessage()
}

func messageCreate(s *discordgo.Session, m *discordgo.MessageCreate) {
	var tmpMovie c.LeaderboardCache
	parser.Load("./cache/1080p.json", &tmpMovie)
	// fmt.Println(m.Content)
	if m.Author.ID == s.State.User.ID {
		return
	}

	if m.Content == "movie" {
		channelID := "922537842083762240"
		data := genContent(&tmpMovie)
		_, err := s.ChannelMessageSendEmbed(channelID, &data)
		if err != nil {
			log.Panicln(err)
		}
	}

	// check health status
	if m.Content == "health" {
		_, err := s.ChannelMessageSend(m.ChannelID, fmt.Sprintf("✅ UP (Latency: %s)", s.HeartbeatLatency()))
		if err != nil {
			log.Fatal(err)
		}
	}
}
