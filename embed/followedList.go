package embed

import (
	"time"

	"github.com/bwmarrin/discordgo"
)

func BuildFollowedListEmbed(msg string) *discordgo.MessageEmbed {
	embedMessage := &discordgo.MessageEmbed{
		Color:       0x58b9ff, // sky blue
		Description: msg,
		Image: &discordgo.MessageEmbedImage{
			URL: "https://i.imgur.com/mWiRuiK.jpg",
		},
		Timestamp: time.Now().Format(time.RFC3339), // Discord wants ISO8601; RFC3339 is an extension of ISO8601 and should be completely compatible.
		Title:     "(ง ◉ _ ◉)ง )) Current Followed Users",
	}
	return embedMessage
}
