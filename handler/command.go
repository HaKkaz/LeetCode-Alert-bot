package handler

import (
	"errors"
	"fmt"

	"github.com/dustyRAIN/leetcode-api-go/leetcodeapi"
)

func AskUserAcStatus(username string) (string, error) {
	// var userSolved leetcodeapi.UserSolveCountByDifficultyDetails
	userSolved, err := leetcodeapi.GetUserSolveCountByDifficulty(username)
	if err != nil {
		fmt.Println("api error")
		return "", err
	}
	acs := userSolved.SolveCount.SubmitStatsGlobal.AcSubmissionNum
	if len(acs) == 0 {
		return "", errors.New("User Didn't Exist.")
	}

	message := fmt.Sprintf("All: %d\nEasy: %d\nMedium: %d\nHard: %d\n",
		acs[0].Count, acs[1].Count, acs[2].Count, acs[3].Count)
	return message, nil
}

func AskLatestAc(username string) (string, error) {
	userLatestAc, err := leetcodeapi.GetUserRecentAcSubmissions(username, 1)
	if err != nil {
		fmt.Println("api error")
		return "", err
	}
	if len(userLatestAc) == 0 {
		fmt.Println("There are no any ac submission.")
		return "", err
	}
	latestAc := userLatestAc[0]
	message := fmt.Sprintf("%s Latest Accepted Submission\n%s\n%s",
		username, ConvertTimestampToTime(latestAc.Timestamp), latestAc.Title)
	return message, nil
}

func AskTracedUsers() (string, error) {
	tracedUsers, err := readTracedList()
	if err != nil {
		fmt.Println("Read traced list error.")
		return "Read traced list error.", err
	}

	message := "Current Folowed Users:"
	for _, username := range tracedUsers {
		message += "\n" + username
	}
	return message, nil
}
