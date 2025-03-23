package handler

import (
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"homework_bot/internal/bot"
	"homework_bot/internal/domain"
)

type TronShieldHandler struct{}

func NewTronShieldHandler() *TronShieldHandler {
	return &TronShieldHandler{}
}

func (h *TronShieldHandler) Handle(b bot.IBot, message *tgbotapi.Message) error {
	msg := domain.MessageToSend{
		ChatId: message.Chat.ID,
		Text:   "請輸入地址，會自動匹配，默認是波場USDT地址風險評估",
	}

	b.GetSwitcher().Next(message.Chat.ID)
	_ = b.SendMessage(msg, bot.DefaultChannel)
	return nil
}
