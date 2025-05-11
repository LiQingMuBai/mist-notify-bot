package command

import (
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"log"
	"ushield_bot/internal/bot"
	"ushield_bot/internal/domain"
	"ushield_bot/pkg/switcher"
	"ushield_bot/pkg/tron"
)

type ExchangeEnergyCommand struct{}

func NewExchangeEnergyCommand() *ExchangeEnergyCommand {
	return &ExchangeEnergyCommand{}
}

func (c *ExchangeEnergyCommand) Exec(b bot.IBot, message *tgbotapi.Message) error {
	userId := message.From.ID
	userName := message.From.UserName

	textStart := "\n\n\n💖您好" + userName + ",🛡️U盾在手，链上无忧！\n" +
		"歡迎使用U盾鏈上風控助手\n" +
		" 📢請輸入兌換能量筆數，格式如下：\n\n" +
		"地址" + "英文下劃綫" + "筆數" + "\n\n" +
		"案例TJCo98saj6WND61g1uuKwJ9GMWMT9WkJFo轉賬一筆能量" + "\n" +
		"TJCo98saj6WND61g1uuKwJ9GMWMT9WkJFo_1" + "\n" +
		"📞聯繫客服：@Ushield001\n"

	msg := domain.MessageToSend{
		ChatId: message.Chat.ID,
		Text:   textStart,
	}
	//b.GetSwitcher().ISwitcherUser.Next(userId)
	b.GetTaskManager().SetTaskStatus(userId, "exchange", switcher.StatusBefore)
	err := b.SendMessage(msg, bot.DefaultChannel)
	return err
}

type GetAccountCommand struct{}

func NewGetAccountCommand() *GetAccountCommand {
	return &GetAccountCommand{}
}

func (c *GetAccountCommand) Exec(b bot.IBot, message *tgbotapi.Message) error {
	userId := message.From.ID
	userName := message.From.UserName

	log.Println("userid>>", userId)
	user, errmsg := b.GetServices().IUserService.GetByUsername(userName)

	if errmsg != nil {

		log.Println("error", errmsg)

	}
	log.Println("user>>", user)
	textStart := "\n\n\n💖您好" + userName + ",🛡️U盾在手，链上无忧！\n" +
		"歡迎使用U盾鏈上風控助手\n\n" +
		"🚀您的地址，請充值：\n\n" +
		user.Address + "\n" +
		"✅您的餘額\n" +
		" 📢" + user.Amount + "\n\n" +
		"📞聯繫客服：@Ushield001\n"

	if len(user.Username) > 0 && len(user.Address) == 0 {

		log.Println("新增地址")
		pk, _address, _ := tron.GetTronAddress(int(user.Id))
		updateUser := domain.User{
			Username: userName,
			Key:      pk,
			Address:  _address,
		}
		b.GetServices().IUserService.UpdateAddress(updateUser)
		textStart = "\n\n\n💖您好" + userName + ",🛡️U盾在手，链上无忧！\n" +
			"歡迎使用U盾鏈上風控助手\n" +
			"🚀您的地址，請充值：\n" +
			_address + "\n" +
			"✅您的餘額\n" +
			"📢0.0" + "\n" +
			"📞聯繫客服：@Ushield001\n"
	}

	msg := domain.MessageToSend{
		ChatId: message.Chat.ID,
		Text:   textStart,
	}
	err := b.SendMessage(msg, bot.DefaultChannel)
	return err
}

type UserRelationCommand struct{}

func NewUserRelationCommand() *UserRelationCommand {
	return &UserRelationCommand{}
}

func (c *UserRelationCommand) Exec(b bot.IBot, message *tgbotapi.Message) error {
	//userId := message.From.ID
	msg := domain.MessageToSend{
		ChatId: message.Chat.ID,
		Text:   "绑定上级关系成功",
	}
	err := b.SendMessage(msg, bot.DefaultChannel)
	return err
}
