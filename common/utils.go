package common

import (
	"fmt"
	"github.com/bwmarrin/discordgo"
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
	v.Imdb = fmt.Sprintf("[IMDB](%s)", v.Imdb)
	description := fmt.Sprintf(
		"Year: **%d**\nSize: **%s**\nGenre: **%s**\n**%s**\n",
		v.Year, v.Size, v.Genre, v.Imdb,
	)
	optionalFields := map[string]bool{"Resolution": true, "Rating": true, "Encoding": true}
	description += v.validFields(optionalFields)
	return description
}

func (v *VideoInfo) validFields(optionalFields map[string]bool) string {
	result := ""
	s := reflect.ValueOf(v).Elem()
	t := s.Type()
	for i := 0; i < t.NumField(); i++ {
		name := t.Field(i).Name
		_, ok := optionalFields[name]
		if ok {
			value := s.FieldByName(name).Interface()
			result += fmt.Sprintf("%s: **%s**\n", name, value)
		}
	}
	fmt.Println(result)
	return result
}
