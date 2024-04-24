package handler

import (
	"bufio"
	"errors"
	"fmt"
	"os"

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

	message := ""
	for _, username := range tracedUsers {
		message += "👉 " + username + "\n"
	}
	return message, nil
}

func checkUserExist(username string) (bool, error) {
	file, err := os.Open("tracedList.txt")
	if err != nil {
		fmt.Println("Error opening file:", err)
		return false, err
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text() // Get the current line as a string
		if line == username {
			return true, nil
		}
	}
	return false, nil
}

func AddNewTraced(username string) (string, error) {
	if username == "" {
		return "Username is empty.", nil
	}

	filename := "tracedList.txt"
	exist, err := checkUserExist(username)
	if err != nil {
		return "Add user error when checking user duplicated!", err
	}

	if exist {
		return "User already exist.", nil
	}

	// Open the file in append mode, create if it doesn't exist
	file, err := os.OpenFile(filename, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		return "Add user error when opening file TracedList.txt!", err
	}
	defer file.Close()

	if _, err := file.WriteString("\n" + username); err != nil {
		return "Add user error when writing to the TracedList.txt!", err
	}
	return "Add new traced user successfully.", nil
}
