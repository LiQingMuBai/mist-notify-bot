package handler

import (
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"homework_bot/internal/bot"
)

type IHandler interface {
	Handle(bot bot.IBot, message *tgbotapi.Message) error
}
