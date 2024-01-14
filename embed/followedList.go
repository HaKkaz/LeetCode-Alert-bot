package embed

import (
	"time"

	"github.com/bwmarrin/discordgo"
)

func BuildFollowedListEmbed(msg string) *discordgo.MessageEmbed {
	embedMessage := &discordgo.MessageEmbed{
		// Author: &discordgo.MessageEmbedAuthor{},
		Color:       0x58b9ff, // sky blue
		Description: msg,
		Image: &discordgo.MessageEmbedImage{
			URL: "https://i.imgur.com/mWiRuiK.jpg",
		},
		// Thumbnail: &discordgo.MessageEmbedThumbnail{
		// 	URL: "https://cdn.discordapp.com/avatars/119249192806776836/cc32c5c3ee602e1fe252f9f595f9010e.jpg?size=2048",
		// },
		Timestamp: time.Now().Format(time.RFC3339), // Discord wants ISO8601; RFC3339 is an extension of ISO8601 and should be completely compatible.
		Title:     "(ง ◉ _ ◉)ง )) Current Followed Users",
	}
	return embedMessage
}
