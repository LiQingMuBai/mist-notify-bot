package handler

import (
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"ushield_bot/internal/bot"
	"ushield_bot/internal/bot/telegram/command"
)

type CommandHandler struct{}

func NewCommandHandler() *CommandHandler {
	return &CommandHandler{}
}

func (h *CommandHandler) Handle(bot bot.IBot, message *tgbotapi.Message) error {
	factory := command.NewFactory()
	cmd := factory.GetCommand(message)

	if err := cmd.Exec(bot, message); err != nil {
		return err
	}
	return nil
}

type MessageHandler struct{}

func NewMessageHandler() *MessageHandler {
	return &MessageHandler{}
}

func (h *MessageHandler) Handle(bot bot.IBot, message *tgbotapi.Message) error {
	return nil
}
