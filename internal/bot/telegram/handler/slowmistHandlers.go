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

	userName := message.From.UserName
	user, err := b.GetServices().IUserService.GetByUsername(userName)
	msg := domain.MessageToSend{
		ChatId: message.Chat.ID,
		Text:   "系統錯誤，請重新輸入地址",
	}
	if err != nil {
		b.GetSwitcher().Next(message.Chat.ID)
		_ = b.SendMessage(msg, bot.DefaultChannel)
		return nil
	}
	if user.Times == 1 {
		msg = domain.MessageToSend{
			ChatId: message.Chat.ID,
			Text: "🔍普通用戶每日贈送 1 次地址風險查詢\n" +
				"📞聯繫客服@ushield001\n",
		}

	} else {
		msg = domain.MessageToSend{
			ChatId: message.Chat.ID,
			Text: "🔍風險評分:87\n" +
				"⚠️有與疑似惡意地址交互\n" +
				"⚠️️有與惡意地址交互\n" +
				"⚠️️有與高風險標籤地址交互\n" +
				"⚠️️受制裁實體\n" +
				"📢📢📢更詳細報告請聯繫客服@ushield001\n",
		}

		err := b.GetServices().IUserService.UpdateTimes(1, userName)

		if err != nil {
			msg = domain.MessageToSend{
				ChatId: message.Chat.ID,
				Text:   "系統錯誤，請重新輸入地址",
			}
		}
	}

	b.GetSwitcher().Next(message.Chat.ID)
	_ = b.SendMessage(msg, bot.DefaultChannel)
	return nil
}
