package handler

import (
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"homework_bot/internal/bot"
	"homework_bot/internal/domain"
)

type ExchangeHandler struct{}

func NewExchangeHandler() *ExchangeHandler {
	return &ExchangeHandler{}
}

func (h *ExchangeHandler) Handle(b bot.IBot, message *tgbotapi.Message) error {
	msg := domain.MessageToSend{
		ChatId: message.Chat.ID,
		Text:   "请输入能量转移目标地址\n" + "余额 100trx\n",
	}

	b.GetSwitcher().Next(message.Chat.ID)
	_ = b.SendMessage(msg, bot.DefaultChannel)

	//message.From.
	return nil
}
