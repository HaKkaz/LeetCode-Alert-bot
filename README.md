# LeetCode-DC-bot
A discord bot for leetcode study group, share member's latest AC submission automatically.

## Run the bot
1. Clone this repo.
2. Put your discord bot token in `.env` file as `DCToken`.
3. Put the channel id you want to recieve the alert ac submissions in `.env` as `AlertChannelId`.
4. Put the users you want to track in `.env` as `Users`, seperated by line.
5. Run `go run main.go` to start the bot.

## Usage
1. `!help` to get the help message.
2. `!status <username>` to get the user's AC status.
3. `!ac <username>` to get the latest AC submission of the user.
4. `!users` to get the list of users are being tracked.

## Example
```
> ! status luckyanthonyan
All: 351
Easy: 120
Medium: 177
Hard: 54
```