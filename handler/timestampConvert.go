package handler

import (
	"fmt"
	"strconv"
	"time"
)

// Converts unix timestamp string to human readable time
func ConvertTimestampToTime(timestamp string) string {
	unixTimestamp, err := strconv.ParseInt(timestamp, 10, 64)
	if err != nil {
		fmt.Println("Convert timestamp to time error.")
		return ""
	}
	dataTimestamp := time.Unix(unixTimestamp, 0)
	humanDateTime := dataTimestamp.Format("2006-01-02 15:04:05")
	return humanDateTime
}

// Converts unix timestamp string to int64
func ConvertTimestampToInt64(strTimestamp string) int64 {
	unixTimestamp, err := strconv.ParseInt(strTimestamp, 10, 64)
	if err != nil {
		fmt.Println("Error parsing Unix timestamp:", err)
		return -1
	}
	return unixTimestamp
}
