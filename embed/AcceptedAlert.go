package embed

import (
	"fmt"
	"os"
	"time"

	"github.com/bwmarrin/discordgo"
)

func SendAcceptedEmbed(
	channelID string,
	username string,
	problemId string,
	problemTitle string,
	difficulty string,
	s *discordgo.Session,
) error {
	msg := fmt.Sprintf("%s solved the problem %s. %s.", username, problemId, problemTitle)
	fmt.Println(msg)
	var duckImagePath string
	if difficulty == "Easy" {
		duckImagePath = "image/EasyDuck.png"
	} else if difficulty == "Medium" {
		duckImagePath = "image/MediumDuck.png"
	} else {
		duckImagePath = "image/HardDuck.png"
	}

	imageFile, imageErr := os.Open(duckImagePath)
	if imageErr != nil {
		fmt.Println("read file error", imageErr)
	}
	defer imageFile.Close()

	ms := &discordgo.MessageSend{
		Embed: &discordgo.MessageEmbed{
			Color:       0x58b9ff, // sky blue
			Description: msg,
			// Image: &discordgo.MessageEmbedImage{
			// 	URL: "attachment://" + duckImagePath,
			// },
			Timestamp: time.Now().Format(time.RFC3339), // Discord wants ISO8601; RFC3339 is an extension of ISO8601 and should be completely compatible.
			Title:     "🎉🎉 New Accepted Submission 🎉🎉",
		},
		Files: []*discordgo.File{
			{
				Name:        duckImagePath,
				Reader:      imageFile,
				ContentType: "image/png",
			},
		},
	}

	s.ChannelMessageSendComplex(channelID, ms)

	return nil
}
