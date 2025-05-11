package handler

import (
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"ushield_bot/internal/bot"
	"ushield_bot/internal/domain"
)

type AccountHandler struct{}

func NewAccountHandler() *AccountHandler {
	return &AccountHandler{}
}

func (h *AccountHandler) Handle(b bot.IBot, message *tgbotapi.Message) error {
	msg := domain.MessageToSend{
		ChatId: message.Chat.ID,
		Text:   "您的波场地址\n" + "余额 100trx\n",
	}
	_ = b.SendMessage(msg, bot.DefaultChannel)

	//message.From.
	return nil
}
