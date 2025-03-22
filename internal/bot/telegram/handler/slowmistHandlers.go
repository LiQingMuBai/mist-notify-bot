package handler

import (
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"homework_bot/internal/bot"
	"homework_bot/internal/domain"
)

type MisttrackHandler struct{}

func NewMisttrackHandler() *MisttrackHandler {
	return &MisttrackHandler{}
}

func (h *MisttrackHandler) Handle(b bot.IBot, message *tgbotapi.Message) error {
	msg := domain.MessageToSend{
		ChatId: message.Chat.ID,
		Text: "🔍风险评分:87\n" +
			"⚠️与疑似恶意地址交互\n" +
			"⚠️与恶意地址交互\n" +
			"⚠️与高风险标签地址交互\n" +
			"⚠️受制裁实体\n" +
			"📢📢📢更详细报告请联系客服@vip664\n",
	}

	b.GetSwitcher().Next(message.Chat.ID)
	_ = b.SendMessage(msg, bot.DefaultChannel)
	return nil
}
