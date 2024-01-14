# LeetCode-DC-bot
A discord bot for leetcode study group, share member's latest AC submission automatically.

## LICENSE
LeetCode-Alert-Bot © 2023 by Anthony Lin is licensed under [CC BY-NC 4.0](https://creativecommons.org/licenses/by-nc/4.0/?ref=chooser-v1)

## Requirements
- [leetcode-api-go](https://github.com/dustyRAIN/leetcode-api-go)
- [discordgo](https://github.com/bwmarrin/discordgo)

## Run the bot
1. Clone this repo.
2. Put your discord bot token in `.env` file as `DCToken`.
3. Put the channel id you want to recieve the alert ac submissions in `.env` as `AlertChannelId`.
4. Put the users you want to track in `.env` as `Users`, seperated by line.
5. Run `go run main.go` to start the bot.

## Usage
1. `! help` to get the help message.
2. `! ask status <username>` to get the user's AC status.
3. `! ask ac <username>` to get the latest AC submission of the user.
4. `! ask users` to get the list of users are being tracked.
5. `! add <username>` to add user to traced list.

## Example
### Show help message
```
> ! help
Commands:
! ask status [username]
! ask ac [username]
! ask users
! help
! add [username]
```

### Show user status
```
> ! ask status luckyanthonyan
All: 351
Easy: 120
Medium: 177
Hard: 54
```

### Show user latest AC submission
```
> ! ask ac luckyanthonyan
luckyanthonyan Latest Accepted Submission
2023-12-29 23:58:33
Minimum Difficulty of a Job Schedule
```

### Show users being tracked
```
> ! ask users
Current Folowed Users:
luckyanthonyan
apple912162
hank55663
codingiscool666
alan8585
```

### Add user to track

After add user, the bot will update the `tracedList.txt`.
```
> ! add alan8585
Add new traced user successfully.
```

