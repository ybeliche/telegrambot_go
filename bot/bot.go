// Package bot provides telegram utilities
package bot

import (
	"log"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

func Bot(bot_api_token, txt string, chat_id int64) {
	bot, err := tgbotapi.NewBotAPI(bot_api_token)
	if err != nil {
		log.Panic(err)
	}

	log.Printf("Authorized as: %s", bot.Self.UserName)
	log.Printf("Will send: %s", txt)

	msg := tgbotapi.NewMessage(chat_id, txt)

	if _, err := bot.Send(msg); err != nil {
		log.Printf("Failed to send message: %v", err)
	}
}
