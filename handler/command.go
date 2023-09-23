package handler

import (
	"errors"
	"fmt"

	"github.com/dustyRAIN/leetcode-api-go/leetcodeapi"
)

func AskUserAcStatus(username string) ([]int, error) {
	// var userSolved leetcodeapi.UserSolveCountByDifficultyDetails
	userSolved, err := leetcodeapi.GetUserSolveCountByDifficulty(username)
	if err != nil {
		fmt.Println("api error")
		return nil, err
	}
	acs := userSolved.SolveCount.SubmitStatsGlobal.AcSubmissionNum
	if len(acs) == 0 {
		return nil, errors.New("User Didn't Exist.")
	}
	return []int{acs[0].Count, acs[1].Count, acs[2].Count, acs[3].Count}, nil
}

func AskLatestAc(username string) ([]string, error) {
	userLatestAc, err := leetcodeapi.GetUserRecentAcSubmissions(username, 1)
	if err != nil {
		fmt.Println("api error")
		return nil, err
	}
	if len(userLatestAc) == 0 {
		fmt.Println("There are no any ac submission.")
		return nil, err
	}
	ac := userLatestAc[0]
	return []string{ac.Timestamp, ac.Title}, nil
}

func skTracedUsers() (string, error) {
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
