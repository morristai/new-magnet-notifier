package main

import (
	"fmt"
	"github.com/morristai/rarbg-notifier/parser"
	"io"
	"os"
)

func main() {
	//url := "https://rarbg.to/torrents.php?category=14;17;42;44;45;46;47;48;50;51;52;54&search=&order=seeders&by=DESC"
	//content := client.Request(url)
	// local test
	file, _ := os.Open("./resources/rarbg_response.html")
	defer file.Close()
	//parser.ParseBillboard(file)
	file.Seek(0, io.SeekStart) // io.Reader need to reset since the pointer is moved by previous parser
	res := parser.ParseHighest(file)
	for idx, i := range res {
		fmt.Println(idx, "========================================")
		fmt.Println("Title:", i.Title)
		fmt.Println("Year:", i.Year)
		fmt.Println("Size:", i.Size)
		fmt.Println("Poster:", i.Poster)
		fmt.Println("Genre:", i.Genre)
		fmt.Println("Resolution:", i.Resolution)
		fmt.Println("IMDB:", i.Imdb)
		fmt.Println("Rating:", i.Rating)
	}

	//parser.ParseTest(file)
	//discord.Run()
	//test := "Christmas.with.the.Chosen.The.Messengers.2021.1080p.WEBRip.x264-RARBG"
	//fmt.Printf("%+v\n", parser.MatchBasic(test))
}
