package handler

import (
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"homework_bot/internal/bot"
	"homework_bot/internal/domain"
)

type VIPHandler struct{}

func NewVIPHandler() *VIPHandler {
	return &VIPHandler{}
}

func (h *VIPHandler) Handle(b bot.IBot, message *tgbotapi.Message) error {

	msg := domain.MessageToSend{
		ChatId: message.Chat.ID,
		Text:   "📞VIP請聯繫客服@Ushield001\n",
	}

	b.GetSwitcher().Next(message.Chat.ID)
	_ = b.SendMessage(msg, bot.DefaultChannel)
	return nil
}
