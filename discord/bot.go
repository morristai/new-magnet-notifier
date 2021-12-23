package discord

import (
	"fmt"
	"github.com/bwmarrin/discordgo"
	log "github.com/sirupsen/logrus"
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

func genContent() discordgo.MessageEmbed {
	posterUrl := "https://dyncdn2.com/mimages/365221/over_opt.jpg"
	title := "Swan.Song.2021.1080p.WEB.H264-NAISU"
	movieUrl := "https://rarbg.to/torrents.php?category=movies"

	poster := discordgo.MessageEmbedImage{
		URL:    posterUrl,
		Height: 300,
		Width:  200,
	}
	data := discordgo.MessageEmbed{
		Title:       title,
		URL:         movieUrl,
		Description: "Just some random test",
		Image:       &poster,
	}
	return data
}

func messageCreate(s *discordgo.Session, m *discordgo.MessageCreate) {
	// fmt.Println(m.Content)
	// Ignore all messages created by the bot itself
	if m.Author.ID == s.State.User.ID {
		return
	}
	// If the message is "ping" reply with "Pong!"
	if m.Content == "movie" {
		channelID := "922537842083762240"
		data := genContent()
		s.ChannelMessageSendEmbed(channelID, &data)
		fmt.Println(m.ChannelID)
	}

	// check health status
	if m.Content == "health" {
		_, err := s.ChannelMessageSend(m.ChannelID, fmt.Sprintf("✅ UP (Latency: %s)", s.HeartbeatLatency()))
		if err != nil {
			log.Fatal(err)
		}
	}
}
