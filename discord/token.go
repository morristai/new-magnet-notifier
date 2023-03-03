package discord

import c "github.com/morristai/rarbg-notifier/common"

var (
	token string
)

func genToken() string {
	config := c.ReadConfig("./resources/application.yml")
	token = config.Discord.Token
	return token
}
