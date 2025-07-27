// Package main
package main

import (
	"log"
	"os"
	"strconv"

	bot "github.com/ybeliche/telegrambot_go/bot"
	msg "github.com/ybeliche/telegrambot_go/msg"
)

var (
	ChatID      = os.Getenv("TG_CHAT_ID")
	BotAPIToken = os.Getenv("BOT_API_TOKEN")
	Message     = os.Getenv("MESSAGE")
)

func main() {
	log.Println("Sending message...")
	txt := msg.Msg(Message)
	chatid, err := strconv.ParseInt(ChatID, 10, 64)
	if err != nil {
		log.Fatal("Failed to convert chat_id to int")
	}
	bot.Bot(BotAPIToken, txt, chatid)
}
