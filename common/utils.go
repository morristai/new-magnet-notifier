package common

import (
	"fmt"
	"github.com/bwmarrin/discordgo"
	"gopkg.in/yaml.v3"
	"io/ioutil"
	"log"
	"reflect"
)

func (v *VideoInfo) GenDiscordMessage() discordgo.MessageEmbed {
	poster := discordgo.MessageEmbedImage{
		URL:    v.Poster,
		Height: 300,
		Width:  200,
	}

	data := discordgo.MessageEmbed{
		Title:       fmt.Sprintf("🍿 %s", v.Title),
		URL:         v.Url,
		Description: v.GenDiscordDescription(),
		Image:       &poster,
	}
	return data
}

func (v *VideoInfo) GenDiscordDescription() string {
	description := fmt.Sprintf(
		// TODO: 1. IMDB emoji 2. Table view
		"Year: **%d**\nGenre: **%s**\n",
		v.Year, v.Genre,
	)
	if v.ProlificReview.Mean != 0 {
		description += fmt.Sprintf("PReview: **%.2f/%.2f**", v.ProlificReview.Mean, v.ProlificReview.Std)
	}

	optionalFields := map[string]bool{"Resolution": true, "Encoding": true, "Rating": true}
	description += v.validFields(optionalFields)
	description += fmt.Sprintf(" **%s**\n", fmt.Sprintf("[IMDB](%s)", v.ImdbUrl))
	description += fmt.Sprintf("Size: **%s**\n", v.Size)

	return description
}

func (v *VideoInfo) validFields(optionalFields map[string]bool) string {
	// If struct field not nil, then append to description
	result := ""
	s := reflect.ValueOf(v).Elem()
	t := s.Type()
	for i := 0; i < t.NumField(); i++ {
		name := t.Field(i).Name
		field := s.FieldByName(name)
		_, ok := optionalFields[name]
		if ok && !field.IsZero() {
			value := s.FieldByName(name).Interface()
			result += fmt.Sprintf("\n%s: **%s**", name, value)
		}
	}
	return result
}

func ReadConfig(filePath string) *Config {
	buf, err := ioutil.ReadFile(filePath)
	if err != nil {
		log.Fatalln(err)
	}

	config := &Config{}
	err = yaml.Unmarshal(buf, config)
	if err != nil {
		log.Fatalln(err)
	}
	return config
}

func Contains(s []string, str string) bool {
	for _, v := range s {
		if v == str {
			return true
		}
	}
	return false
}
