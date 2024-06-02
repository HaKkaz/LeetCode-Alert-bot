package handler

import (
	"LeetCode-Alert-bot/embed"
	"LeetCode-Alert-bot/selfapi"
	"fmt"
	"os"
	"time"

	"github.com/bwmarrin/discordgo"
	"github.com/dustyRAIN/leetcode-api-go/leetcodeapi"
)

func RoundAlertAC(s *discordgo.Session, timeSlot time.Duration) error {
	channelID := os.Getenv("DCChannel")

	tracedUsers, err := readTracedList()
	if err != nil {
		return err
	}

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
				var (
					problemId         string
					problemTitle      string
					problemSlug       string
					problemDifficulty string
				)

				problemTitle = ac.Title
				problemSlug = ac.TitleSlug
				problemDifficulty, problemId = selfapi.GetProblemDifficultyAndId(problemSlug)

				embed.SendAcceptedEmbed(channelID, username, problemId, problemTitle, problemDifficulty, s)

				debugMsg := fmt.Sprintf("%s accepted %s", username, problemTitle)
				fmt.Println(debugMsg)
			}
		}
	}

	time.Sleep(timeSlot)
	return nil
}
