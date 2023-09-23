package handler

import (
	"time"

	"github.com/bwmarrin/discordgo"
	"github.com/dustyRAIN/leetcode-api-go/leetcodeapi"
)

func RoundAlertAC(s *discordgo.Session) error {
	channelID := "1153931585230491658"

	tracedUsers, err := readTracedList()
	if err != nil {
		return err
	}
	message := "Current Folowed Users:\n"

	for _, username := range tracedUsers {
		leetcodeapi.GetUserRecentAcSubmissions(username, 1)
		message += username + "\n"
	}

	s.ChannelMessageSend(channelID, message)
	time.Sleep(10 * time.Second)
	return nil
}
