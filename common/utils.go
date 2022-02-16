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
		Type:        "rich",
		URL:         v.Url,
		Description: v.GenDiscordDescription(),
		Image:       &poster,
		Fields:      v.GenDiscordFields(map[string]bool{"Resolution": true, "Encoding": true, "Rating": true}),
	}
	return data
}

func (v *VideoInfo) GenDiscordDescription() string {
	// TODO: grab from video main page
	description := ""
	description += fmt.Sprintf(" **%s**\n", fmt.Sprintf("[IMDB](%s)", v.ImdbUrl))
	return description
}

func (v *VideoInfo) GenDiscordFields(optionalFields map[string]bool) []*discordgo.MessageEmbedField {
	// If struct field not nil, then append to fields
	var result []*discordgo.MessageEmbedField

	if v.ProlificReview.Mean != 0 && v.ProlificReview.Std != 0 {
		result = append(result, &discordgo.MessageEmbedField{Name: "PReview", Value: fmt.Sprintf("**%.2f/%.2f**", v.ProlificReview.Mean, v.ProlificReview.Std)})
	}

	// mandatory fields
	result = append(result, &discordgo.MessageEmbedField{Name: "Year", Value: fmt.Sprintf("%v", v.Year), Inline: true})
	result = append(result, &discordgo.MessageEmbedField{Name: "Genre", Value: fmt.Sprintf("%v", v.Genre), Inline: true})
	result = append(result, &discordgo.MessageEmbedField{Name: "Size", Value: fmt.Sprintf("%v", v.Size), Inline: true})

	// optional fields
	s := reflect.ValueOf(v).Elem()
	t := s.Type()
	for i := 0; i < t.NumField(); i++ {
		name := t.Field(i).Name
		field := s.FieldByName(name)
		_, ok := optionalFields[name]
		if ok && !field.IsZero() {
			value := s.FieldByName(name).Interface()
			result = append(result, &discordgo.MessageEmbedField{Name: name, Value: fmt.Sprintf("%v", value), Inline: true})
		}
	}
	return result
}

func ReadConfig(path string) *Config {
	//viper.AddConfigPath(path)
	buf, err := ioutil.ReadFile(path)
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
