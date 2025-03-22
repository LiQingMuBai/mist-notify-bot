package command

import (
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"homework_bot/internal/bot"
	"homework_bot/internal/domain"
)

type AskGroupCommand struct{}

func NewAskGroupCommand() *AskGroupCommand {
	return &AskGroupCommand{}
}

func (c *AskGroupCommand) Exec(b bot.IBot, message *tgbotapi.Message) error {
	userId := message.From.ID
	msg := domain.MessageToSend{
		ChatId: message.Chat.ID,
		Text:   "Напишите номер группы, которая вам нужна. Формат: Б9122-09.03.04 4, где 4 - это 4я подгруппа",
	}
	b.GetSwitcher().ISwitcherUser.Next(userId)
	err := b.SendMessage(msg, bot.DefaultChannel)
	return err
}
