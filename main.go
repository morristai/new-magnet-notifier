package main

import (
	"github.com/morristai/rarbg-notifier/client"
	"github.com/morristai/rarbg-notifier/parser"
)

func main() {
	url := "https://rarbg.to/torrents.php?category=movies"
	//url := "https://github.com/trendmicro"
	content := client.Request(url)
	parser.Parse(content)
}
