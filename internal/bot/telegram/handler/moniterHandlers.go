package handler

import (
	"fmt"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"strings"
	"ushield_bot/internal/bot"
	"ushield_bot/internal/domain"
)

type MoniterHandler struct{}

func NewMoniterHandler() *MoniterHandler {
	return &MoniterHandler{}
}

func (h *MoniterHandler) Handle(b bot.IBot, message *tgbotapi.Message) error {
	msg := domain.MessageToSend{
		ChatId: message.Chat.ID,
		Text:   "🎉請輸入您想要監控的地址？\n",
	}
	meta := message.Text

	user, _ := b.GetServices().IUserService.GetByUsername(message.Chat.UserName)

	if len(user.TronAddress) != 0 || len(user.EthAddress) != 0 {

		msg := domain.MessageToSend{
			ChatId: message.Chat.ID,
			Text: "🎉您已經綁定過地址\n" +
				"\n\n🎁 當前爲試用期，1個地址免費監控30天" +
				"\n\n昇級會員 ➜ 綁定2個地址 ➕ 實時凍結提醒" +
				"\n\n👉 /vip 查看會員權益" +
				"\n\n🛡️ U盾在手，鏈上無憂！" +
				"\n",
		}
		_ = b.SendMessage(msg, bot.DefaultChannel)
		return nil
	}
	b.GetServices().IUserService.BindChat(fmt.Sprintf("%d", message.Chat.ID), message.From.UserName)
	userStateMap := b.GetUserStates()
	userStateMap[message.Chat.ID] = "pre-monitor"
	if len(meta) == 34 && strings.HasPrefix(meta, "T") {
		userStateMap := b.GetUserStates()
		userStateMap[message.Chat.ID] = "monitor-done"

		b.GetServices().IUserService.BindTronAddress(meta, message.From.UserName)
		msg = domain.MessageToSend{
			ChatId: message.Chat.ID,
			Text: "✅ 地址監控已開啟！\n\n已绑定地址：" + meta + "\n\n繫統將每天自動檢測該地址的風險等級\n\n" +
				"⚠️ 若风险等级发生变化或有列入黑名单风险，您将第一时间收到通知！" +
				"\n\n🎁當前爲試用期，1個地址免費監控30天" +
				"\n\n昇級會員 ➜ 綁定2個地址 ➕ 實時凍結提醒" +
				"\n\n👉 /vip 查看會員權益" +
				"\n\n🛡️ U盾在手，鏈上無憂！" +
				"\n",
		}

	}
	if len(meta) == 42 && strings.HasPrefix(meta, "0x") {
		userStateMap := b.GetUserStates()
		userStateMap[message.Chat.ID] = "monitor-done"
		b.GetServices().IUserService.BindEthereumAddress(meta, message.From.UserName)
		msg = domain.MessageToSend{
			ChatId: message.Chat.ID,
			Text: "✅ 地址監控已開啟！\n\n已绑定地址：" + meta + "\n\n繫統將每天自動檢測該地址的風險等級\n\n" +
				"⚠️ 若风险等级发生变化或有列入黑名单风险，您将第一时间收到通知！" +
				"\n\n🎁當前爲試用期，1個地址免費監控30天" +
				"\n\n昇級會員 ➜ 綁定2個地址 ➕ 實時凍結提醒" +
				"\n\n👉 /vip 查看會員權益" +
				"\n\n🛡️ U盾在手，鏈上無憂！" +
				"\n",
		}
	}

	b.GetSwitcher().Next(message.Chat.ID)

	_ = b.SendMessage(msg, bot.DefaultChannel)
	return nil
}
