package selfapi

import (
	"fmt"

	"github.com/dustyRAIN/leetcode-api-go/leetcodeapi"
)

type QuestionResponseBody struct {
	Data struct {
		Question struct {
			QuestionId    string `json:"questionId"`
			FrontendId    string `json:"questionFrontendId"`
			Title         string `json:"title"`
			TitleSlug     string `json:"titleSlug"`
			IsPaidOnly    bool   `json:"isPaidOnly"`
			Difficulty    string `json:"difficulty"`
			Likes         int    `json:"likes"`
			Dislikes      int    `json:"dislikes"`
			CategoryTitle string `json:"categoryTitle"`
		} `json:"question"`
	} `json:"data"`
}

func GetProblemDetails(problemName string) (*QuestionResponseBody, error) {
	var responseBody QuestionResponseBody
	payload := fmt.Sprintf(`{
		"query": "\n    query questionTitle($titleSlug: String!) {\n  question(titleSlug: $titleSlug) {\n    questionId\n    questionFrontendId\n    title\n    titleSlug\n    isPaidOnly\n    difficulty\n    likes\n    dislikes\n    categoryTitle\n  }\n}\n    ",
    	"variables": {
			"titleSlug": "%v"
		}
	}`, problemName)
	fmt.Println(payload)
	err := (&leetcodeapi.Util{}).MakeGraphQLRequest(payload, &responseBody)

	if err != nil {
		return nil, err
	}
	return &responseBody, nil
}

func GetProblemDifficultyAndId(problemName string) (string, string) {
	responseBody, _ := GetProblemDetails(problemName)

	fmt.Println(responseBody.Data.Question.Difficulty)
	return responseBody.Data.Question.Difficulty, responseBody.Data.Question.FrontendId
}
