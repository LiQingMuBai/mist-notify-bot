package command

import (
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"homework_bot/internal/bot"
)

type ICommand interface {
	Exec(bot bot.IBot, message *tgbotapi.Message) error
}
