package handler

import (
	"time"

	"github.com/bwmarrin/discordgo"
	"github.com/dustyRAIN/leetcode-api-go/leetcodeapi"
)

func RoundAlertAC(s *discordgo.Session, timeSlot time.Duration) error {
	channelID := "1153931585230491658"

	tracedUsers, err := readTracedList()
	if err != nil {
		return err
	}

	message := ""
	for _, username := range tracedUsers {
		userLatestAc, err := leetcodeapi.GetUserRecentAcSubmissions(username, 1)
		if err != nil {
			return err
		}
		if len(userLatestAc) == 0 {
			continue
		}
		for _, ac := range userLatestAc {
			acTimestamp := ConvertTimestampToInt64(ac.Timestamp)
			if acTimestamp >= time.Now().Unix()-60 {
				message += username + " solved " + ac.Title + "\n"
			}
		}
	}

	s.ChannelMessageSend(channelID, message)
	time.Sleep(timeSlot)
	return nil
}
