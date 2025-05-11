package handler

import (
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"ushield_bot/internal/bot"
	"ushield_bot/internal/domain"
)

type UserRelationHandler struct{}

func NewUserRelationHandler() *UserRelationHandler {
	return &UserRelationHandler{}
}

func (h *UserRelationHandler) Handle(b bot.IBot, message *tgbotapi.Message) error {
	msg := domain.MessageToSend{
		ChatId: message.Chat.ID,
		Text:   "🔍请输入你的推荐关系的用户名，不需要加@符号\n",
	}
	_ = b.SendMessage(msg, bot.DefaultChannel)

	//message.From.
	return nil
}
