package handler

import (
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"log"
	"strings"
	"ushield_bot/internal/bot"
	"ushield_bot/pkg/switcher"
)

type Factory struct {
}

func NewFactory() *Factory {
	return &Factory{}
}

func (f *Factory) GetHandler(b bot.IBot, message *tgbotapi.Message) IHandler {
	userId := message.From.ID

	log.Println(message.Text)

	status, flag := b.GetTaskManager().GetTaskStatus(userId, "exchange")

	meta := message.Text
	state := b.GetUserStates()[message.Chat.ID]
	if strings.Contains(message.Text, "monitor") {
		return NewMoniterHandler()
	}
	if len(meta) == 34 && strings.HasPrefix(meta, "T") && strings.Contains(state, "pre-monitor") {
		return NewMoniterHandler()
	}
	if len(meta) == 42 && strings.HasPrefix(meta, "0x") && strings.Contains(state, "pre-monitor") {
		return NewMoniterHandler()
	}
	if strings.Contains(message.Text, "vip") {
		return NewVIPHandler()
	}
	if strings.Contains(message.Text, "address") {
		return NewStatsHandler()
	}
	if strings.Contains(message.Text, "help") {
		return NewHelpHandler()
	}
	if strings.Contains(message.Text, "check") {
		return NewTronShieldHandler()
	}
	if strings.Contains(message.Text, "relation") {
		return NewUserRelationHandler()
	}

	if len(meta) == 34 && strings.HasPrefix(meta, "T") {
		return NewMisttrackHandler()
	}
	if len(meta) == 42 && strings.HasPrefix(meta, "0x") {
		return NewMisttrackHandler()
	}
	if strings.Contains(message.Text, "get_account") {
		return NewCommandHandler()
	}
	if strings.Contains(message.Text, "exchange_energy") {
		return NewCommandHandler()
	}
	if message.IsCommand() {
		return NewCommandHandler()
	}
	if flag && status == switcher.StatusBefore {
		return NewExchangeExecHandler()
	}

	return NewMessageHandler()
}

//}
