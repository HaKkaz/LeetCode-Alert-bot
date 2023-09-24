package main

import (
	"LeetCode-Alert-bot/handler"
	"fmt"
	"os"
	"os/signal"
	"strings"
	"syscall"

	"github.com/bwmarrin/discordgo"
	_ "github.com/joho/godotenv/autoload"
)

const prefix string = "!" // prefix of command

func main() {
	token := os.Getenv("DCToken")

	// creates a new Discord session
	dg, err := discordgo.New("Bot " + token)
	fmt.Println("Bot " + token)
	if err != nil {
		fmt.Println("error creating Discord session,", err)
		return
	}

	// Register the messageCreate func as a callback for MessageCreate events.
	dg.AddHandler(messageCreate)

	// 只監聽訊息
	dg.Identify.Intents = discordgo.IntentsGuildMessages

	// 開啟連線
	err = dg.Open()
	if err != nil {
		fmt.Println("error opening connection,", err)
		return
	}

	go func() {
		for {
			err := handler.RoundAlertAC(dg)
			if err != nil {
				fmt.Println("error sending message,", err)
			}
		}
	}()

	// Wait here until CTRL-C or other term signal is received.
	fmt.Println("Bot is now running.  Press CTRL-C to exit.")
	sc := make(chan os.Signal, 1)
	signal.Notify(sc, syscall.SIGINT, syscall.SIGTERM, os.Interrupt)
	<-sc

	// Cleanly close down the Discord session.
	dg.Close()
}

func messageCreate(s *discordgo.Session, m *discordgo.MessageCreate) {
	// Ignore all messages created by the bot itself
	// This isn't required in this specific example but it's a good practice.
	if m.Author.ID == s.State.User.ID {
		return
	}
	fmt.Println(m.Content)

	if isCommand(m.Content) {
		channelID := m.ChannelID
		command := getCommand(m.Content)
		command_args := strings.Split(command, " ")
		fmt.Println(command_args)

		// Ask user solved count by difficulty.
		if command_args[0] == "status" && len(command_args) == 2 {
			username := command_args[1]
			message, err := handler.AskUserAcStatus(username)
			if err != nil {
				message = "No this user."
			}
			s.ChannelMessageSend(channelID, message)
		}

		if command_args[0] == "ac" && len(command_args) == 2 {
			username := command_args[1]
			message, err := handler.AskLatestAc(username)
			if err != nil {
				message = "No ac submission."
			}
			s.ChannelMessageSend(channelID, message)
		}

		if command_args[0] == "users" && len(command_args) == 2 {
			message, err := handler.AskTracedUsers()
			if err != nil {
				message = "Read traced list error."
			}
			s.ChannelMessageSend(channelID, message)
		}
	}
}

func isCommand(s string) bool {
	return len(s) > len(prefix) && s[:len(prefix)] == prefix
}

func getCommand(s string) string {
	// Command have to seperate by space.
	return s[len(prefix)+1:]
}
