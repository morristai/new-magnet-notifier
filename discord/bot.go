package discord

import (
	"fmt"
	"github.com/bwmarrin/discordgo"
	"github.com/morristai/rarbg-notifier/task"
	"log"
	"os"
	"os/signal"
	"syscall"
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

	// Trigger cron job
	task.CronJobs(dg)

	// Wait here until CTRL-C or other term signal is received.
	fmt.Println("Bot is now running.  Press CTRL-C to exit.")
	sc := make(chan os.Signal, 1) // channel
	signal.Notify(sc, syscall.SIGINT, syscall.SIGTERM, os.Interrupt, os.Kill)
	<-sc

	// Cleanly close down the Discord session.
	dg.Close()
}

func messageCreate(s *discordgo.Session, m *discordgo.MessageCreate) {
	// check health status
	if m.Content == "health" {
		_, err := s.ChannelMessageSend(m.ChannelID, fmt.Sprintf("✅ UP (Latency: %s)", s.HeartbeatLatency()))
		if err != nil {
			log.Fatalln(err)
		}
	}
}
