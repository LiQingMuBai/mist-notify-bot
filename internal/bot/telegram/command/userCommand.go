package command

import (
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"homework_bot/internal/bot"
	"homework_bot/internal/domain"
)

type ExchangeEnergyCommand struct{}

func NewExchangeEnergyCommand() *ExchangeEnergyCommand {
	return &ExchangeEnergyCommand{}
}

func (c *ExchangeEnergyCommand) Exec(b bot.IBot, message *tgbotapi.Message) error {
	userId := message.From.ID
	msg := domain.MessageToSend{
		ChatId: message.Chat.ID,
		Text:   "兑换能量",
	}
	b.GetSwitcher().ISwitcherUser.Next(userId)
	err := b.SendMessage(msg, bot.DefaultChannel)
	return err
}

type GetAccountCommand struct{}

func NewGetAccountCommand() *GetAccountCommand {
	return &GetAccountCommand{}
}

func (c *GetAccountCommand) Exec(b bot.IBot, message *tgbotapi.Message) error {
	userId := message.From.ID
	msg := domain.MessageToSend{
		ChatId: message.Chat.ID,
		Text:   "获取账户信息",
	}
	b.GetSwitcher().ISwitcherUser.Next(userId)
	err := b.SendMessage(msg, bot.DefaultChannel)
	return err
}

type UserRelationCommand struct{}

func NewUserRelationCommand() *UserRelationCommand {
	return &UserRelationCommand{}
}

func (c *UserRelationCommand) Exec(b bot.IBot, message *tgbotapi.Message) error {
	userId := message.From.ID
	msg := domain.MessageToSend{
		ChatId: message.Chat.ID,
		Text:   "绑定上级关系成功",
	}
	b.GetSwitcher().ISwitcherUser.Next(userId)
	err := b.SendMessage(msg, bot.DefaultChannel)
	return err
}
