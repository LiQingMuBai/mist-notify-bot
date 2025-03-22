package handler

import (
	"fmt"
	"github.com/go-playground/validator/v10"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"homework_bot/internal/bot"
	"strings"
)

type AskGroupHandler struct{}

func NewAskGroupHandler() *AskGroupHandler {
	return &AskGroupHandler{}
}

func validateGroup(message *tgbotapi.Message) error {
	group := message.Text
	fields := strings.Split(group, " ")
	if len(fields) != 2 {
		return fmt.Errorf("invalid group format")
	}

	if len(fields[0]) != 15 {
		fmt.Println(len(fields[0]))
		return fmt.Errorf("invalid group format")
	}

	valid := validator.New()
	if err := valid.Var(fields[1], "numeric"); err != nil {
		return err
	}
	return nil
}

func (h *AskGroupHandler) Handle(b bot.IBot, message *tgbotapi.Message) error {
	//if err := validateGroup(message); err != nil {
	//	msg := domain.MessageToSend{
	//		ChatId: message.Chat.ID,
	//		Text:   "資訊不正確！請重新輸入",
	//	}
	//	err = b.SendMessage(msg, bot.DefaultChannel)
	//	if err != nil {
	//		return err
	//	}
	//	return nil
	//}
	//
	//fields := strings.Split(message.Text, " ")
	//user, err := b.GetServices().IUserService.GetByUsername(message.From.UserName)
	//studyGroup, _ := strconv.Atoi(fields[1])
	//
	//if err != nil {
	//	user = *domain.NewUser(message.From.UserName, fields[0], studyGroup)
	//	err = b.GetServices().IUserService.Create(user)
	//	if err != nil {
	//		msg := domain.MessageToSend{
	//			ChatId: message.Chat.ID,
	//			Text:   "Server error\n",
	//		}
	//		_ = b.SendMessage(msg, bot.DefaultChannel)
	//		return err
	//	}
	//	userId := message.From.ID
	//	b.GetSwitcher().Next(userId)
	//
	//	msg := domain.MessageToSend{
	//		ChatId: message.Chat.ID,
	//		Text:   "Группа успешно задана",
	//	}
	//	_ = b.SendMessage(msg, bot.DefaultChannel)
	//	return nil
	//}
	//
	//user = domain.User{
	//	Username:      message.From.UserName,
	//	CodeDirection: fields[0],
	//	StudyGroup:    studyGroup,
	//}
	//err = b.GetServices().IUserService.Update(user)
	//if err != nil {
	//	return err
	//}
	//
	//msg := domain.MessageToSend{
	//	ChatId: message.Chat.ID,
	//	Text:   "Группа успешно задана",
	//}
	//
	//b.GetSwitcher().Next(message.Chat.ID)
	//_ = b.SendMessage(msg, bot.DefaultChannel)
	//return err
	return validateGroup(message)
}
