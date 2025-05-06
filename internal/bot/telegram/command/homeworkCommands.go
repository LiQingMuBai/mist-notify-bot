package command

import (
	"fmt"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"homework_bot/internal/bot"
	"homework_bot/internal/domain"
	"homework_bot/pkg/tron"
	"log"
	"strconv"
	"strings"
	"time"
)

type StartCommand struct{}

func NewStartCommand() *StartCommand {
	return &StartCommand{}
}

func (c *StartCommand) Exec(b bot.IBot, message *tgbotapi.Message) error {

	userName := message.From.UserName
	user, err := b.GetServices().IUserService.GetByUsername(userName)

	if user.Username == "" {
		//	log.Println(">>>>>>>>>>>>>>>>>>>>>>>>>>>>空的，需要创建<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<")

		user = *domain.NewUser(message.From.UserName, "", fmt.Sprintf("%d", message.Chat.ID), "", "", "", "", "")

		err = b.GetServices().IUserService.Create(user)

		pk, _address, _ := tron.GetTronAddress(int(user.Id))

		updateUser := domain.User{
			Id:      user.Id,
			Key:     pk,
			Address: _address,
		}
		b.GetServices().IUserService.Update(updateUser)

	} else {
		log.Println("username", userName)
	}

	textStart := "\n\n\n💖您好" + userName + ",🛡️U盾在手，链上无忧！\n" +
		"歡迎使用U盾鏈上風控助手\n" +
		"🚀功能介紹：\n" +
		"✅USDT地址風險查詢\n" +
		"✅地址行爲分析報告\n" +
		"✅地址風險等級變動提醒\n" +
		"✅USDT凍結警報提醒（秒級響應，讓你的U永不被凍結）\n" +
		"🎁 新用户福利：\n🎉 免费绑定 1 个地址，开启实时风险监控\n🎉 每日赠送 1 次地址风险查询\n\n" +
		"💡常用指令：\n" +
		"/check 地址 ➜ 查詢地址風險\n" +
		"/monitor_address 地址 ➜ 開啓地址實時監控\n" +
		"/upgrade_vip ➜ 升級會員，解鎖更多權益\n" +
		"📞聯繫客服：@Ushield001\n"

	//"🚀用戶標識:" + user.UserID + "\n🏆推廣人數:0\n🔎查詢積分:0\n🕙註冊時間:+" + "\n+-----------------------+\n/query – 地址查詢\n/help –   幫助\n – 更多功能請聯繫我們的客服\n+--------------------+\n🔍@Ushield001"
	msg := domain.MessageToSend{
		ChatId: message.Chat.ID,
		Text:   textStart,
	}
	err = b.SendMessage(msg, bot.DefaultChannel)
	return err
}

type AddCommand struct{}

func NewAddCommand() *AddCommand {
	return &AddCommand{}
}

func (c *AddCommand) Exec(b bot.IBot, message *tgbotapi.Message) error {
	b.GetSwitcher().ISwitcherAdd.Next(message.From.ID)
	msg := domain.MessageToSend{
		ChatId: message.Chat.ID,
		Text:   "輸入您的波場地址",
	}

	err := b.SendMessage(msg, bot.DefaultChannel)
	return err
}

type UpdateCommand struct{}

func NewUpdateCommand() *UpdateCommand {
	return &UpdateCommand{}
}

func (c *UpdateCommand) Exec(b bot.IBot, message *tgbotapi.Message) error {
	b.GetSwitcher().ISwitcherUpdate.Next(message.From.ID)
	msg := domain.MessageToSend{
		ChatId: message.Chat.ID,
		Text:   "填寫您的條目 ID",
	}

	err := b.SendMessage(msg, bot.DefaultChannel)
	return err
}

type DeleteCommand struct{}

func NewDeleteCommand() *DeleteCommand {
	return &DeleteCommand{}
}

func (c *DeleteCommand) Exec(b bot.IBot, message *tgbotapi.Message) error {
	words := strings.Split(message.Text, " ")
	if len(words) != 2 {
		return b.SendInputError(message)
	}

	id, err := strconv.Atoi(words[1])
	if err != nil {
		return b.SendInputError(message)
	}

	err = b.GetServices().Delete(id)
	if err != nil {
		msg := domain.MessageToSend{
			ChatId: message.Chat.ID,
			Text:   "Ошибка удаления",
		}
		_ = b.SendMessage(msg, bot.DefaultChannel)
		return err
	}

	msg := domain.MessageToSend{
		ChatId: message.Chat.ID,
		Text:   "Запись успешно удалена",
	}
	err = b.SendMessage(msg, bot.DefaultChannel)
	return err
}

type GetAllCommand struct{}

func NewGetAllCommand() *GetAllCommand {
	return &GetAllCommand{}
}

func (c *GetAllCommand) Exec(b bot.IBot, message *tgbotapi.Message) error {
	homeworks, err := b.GetServices().GetAll()

	if err != nil {
		return err
	}

	for _, homework := range homeworks {
		err = b.SendHomework(homework, message.Chat.ID, bot.DefaultChannel)
		if err != nil {
			return err
		}
	}
	return nil
}

type GetOnWeekCommand struct{}

func NewGetOnWeekCommand() *GetOnWeekCommand {
	return &GetOnWeekCommand{}
}

func (c *GetOnWeekCommand) Exec(b bot.IBot, message *tgbotapi.Message) error {
	homeworks, err := b.GetServices().GetByWeek()
	if err != nil {
		return err
	}

	for _, homework := range homeworks {
		err = b.SendHomework(homework, message.Chat.ID, bot.DefaultChannel)
		if err != nil {
			return err
		}
	}

	return nil
}

type GetOnIdCommand struct{}

func NewGetOnIdCommand() *GetOnIdCommand {
	return &GetOnIdCommand{}
}

func (c *GetOnIdCommand) Exec(b bot.IBot, message *tgbotapi.Message) error {
	words := strings.Split(message.Text, " ")
	if len(words) != 2 {
		return b.SendInputError(message)
	}

	id, err := strconv.Atoi(words[1])
	if err != nil {
		return err
	}

	homework, err := b.GetServices().GetById(id)
	if err != nil {
		return err
	}

	err = b.SendHomework(homework, message.Chat.ID, bot.DefaultChannel)
	return err
}

type GetOnTodayCommand struct{}

func NewGetOnTodayCommand() *GetOnTodayCommand {
	return &GetOnTodayCommand{}
}

func (c *GetOnTodayCommand) Exec(b bot.IBot, message *tgbotapi.Message) error {
	homeworks, err := b.GetServices().GetByToday()
	if err != nil {
		return err
	}

	for _, homework := range homeworks {
		err = b.SendHomework(homework, message.Chat.ID, bot.DefaultChannel)
		if err != nil {
			return err
		}
	}
	return nil
}

type GetOnTomorrowCommand struct{}

func NewGetOnTomorrowCommand() *GetOnTomorrowCommand {
	return &GetOnTomorrowCommand{}
}

func (c *GetOnTomorrowCommand) Exec(b bot.IBot, message *tgbotapi.Message) error {
	homeworks, err := b.GetServices().GetByTomorrow()
	if err != nil {
		return err
	}

	for _, homework := range homeworks {
		err = b.SendHomework(homework, message.Chat.ID, bot.DefaultChannel)
		if err != nil {
			return err
		}
	}
	return nil
}

type GetOnDateCommand struct{}

func NewGetOnDateCommand() *GetOnDateCommand {
	return &GetOnDateCommand{}
}

func (c *GetOnDateCommand) Exec(b bot.IBot, message *tgbotapi.Message) error {
	words := strings.Split(message.Text, " ")
	if len(words) != 2 {
		return b.SendInputError(message)
	}

	date, err := time.Parse(time.DateOnly, words[1])
	if err != nil {
		return err
	}

	homeworks, err := b.GetServices().GetByDate(date)
	if err != nil {
		return err
	}

	for _, homework := range homeworks {
		err = b.SendHomework(homework, message.Chat.ID, bot.DefaultChannel)
		if err != nil {
			return err
		}
	}
	return nil
}

type HelpCommand struct{}

func NewHelpCommand() *HelpCommand {
	return &HelpCommand{}
}

func (c *HelpCommand) Exec(b bot.IBot, message *tgbotapi.Message) error {
	textHelp := "聯係我們的客服"
	msg := domain.MessageToSend{
		ChatId: message.Chat.ID,
		Text:   textHelp,
	}
	err := b.SendMessage(msg, bot.DefaultChannel)
	return err
}

type DefaultCommand struct{}

func NewDefaultCommand() *DefaultCommand {
	return &DefaultCommand{}
}

func (c *DefaultCommand) Exec(b bot.IBot, message *tgbotapi.Message) error {
	msg := domain.MessageToSend{
		ChatId: message.Chat.ID,
		Text:   "我是默认命令，不熟悉",
	}
	err := b.SendMessage(msg, bot.DefaultChannel)
	return err
}
