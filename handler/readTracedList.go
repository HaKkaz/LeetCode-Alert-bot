package handler

import (
	"bufio"
	"fmt"
	"os"
)

func readTracedList() ([]string, error) {
	filename := "tracedList.txt"
	var tracedUser []string

	file, err := os.Open(filename)
	if err != nil {
		fmt.Println("Error opening file:", err)
		return nil, err
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text() // Get the current line as a string
		tracedUser = append(tracedUser, line)
	}
	if err := scanner.Err(); err != nil {
		fmt.Println("Error scanning file:", err)
		return nil, err
	}
	return tracedUser, nil
}
