package handler

import (
	"fmt"
	"github.com/go-playground/validator/v10"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"github.com/google/uuid"
	"homework_bot/internal/bot"
	"homework_bot/internal/bot/telegram/command"
	"homework_bot/internal/domain"
	"io"
	"net/http"
	"os"
	"path/filepath"
	"regexp"
	"strconv"
	"strings"
	"time"
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

type WaitingNameHandler struct{}

func NewWaitingNameHandler() *WaitingNameHandler {
	return &WaitingNameHandler{}
}

func (h *WaitingNameHandler) Handle(b bot.IBot, message *tgbotapi.Message) error {
	userId := message.From.ID
	userData := b.GetUserData()
	data := userData[userId]
	data.Name = message.Text

	userData[userId] = data
	b.SetUserData(userData)
	b.GetSwitcher().Next(userId)

	msg := domain.MessageToSend{
		ChatId: message.Chat.ID,
		Text:   "Название успешно добавлено! Теперь отправте описание к записи, или команду /done",
	}

	err := b.SendMessage(msg, bot.DefaultChannel)
	return err
}

type WaitingDescriptionHandler struct{}

func NewWaitingDescriptionHandler() *WaitingDescriptionHandler {
	return &WaitingDescriptionHandler{}
}

func (h *WaitingDescriptionHandler) Handle(b bot.IBot, message *tgbotapi.Message) error {
	userId := message.From.ID
	userData := b.GetUserData()
	data := userData[userId]
	data.Description = message.Text

	userData[userId] = data
	b.SetUserData(userData)
	b.GetSwitcher().Next(userId)

	msg := domain.MessageToSend{
		ChatId: message.Chat.ID,
		Text:   "Описание успешно добавлено! Теперь отправте фотографии к записи, или команду /done",
	}

	err := b.SendMessage(msg, bot.DefaultChannel)
	return err
}

type WaitingImageHandler struct{}

func NewWaitingImageHandler() *WaitingImageHandler {
	return &WaitingImageHandler{}
}

func (h *WaitingImageHandler) saveImage(bot *tgbotapi.BotAPI, fileId string) (string, error) {
	file, err := bot.GetFile(tgbotapi.FileConfig{FileID: fileId})
	if err != nil {
		return "", err
	}

	uniqueFileName := uuid.New().String() + filepath.Ext(file.FilePath)
	savePath := filepath.Join("/home/nikita/GolandProjects/homework_bot/media", uniqueFileName)

	fileURL := file.Link(bot.Token)
	response, err := http.Get(fileURL)
	if err != nil {
		return "", err
	}

	defer func() {
		_ = response.Body.Close()
	}()

	out, err := os.Create(savePath)
	if err != nil {
		return "", err
	}
	defer func() {
		_ = out.Close()
	}()

	_, err = io.Copy(out, response.Body)
	if err != nil {
		return "", err
	}

	return savePath, nil
}

func (h *WaitingImageHandler) Handle(b bot.IBot, message *tgbotapi.Message) error {
	userId := message.From.ID
	userData := b.GetUserData()
	data := userData[userId]

	if len(message.Photo) > 0 {
		image := message.Photo[len(message.Photo)-1]
		path, err := h.saveImage(b.GetBot(), image.FileID)
		if err != nil {
			return err
		}

		data.Images = append(data.Images, path)
		userData[userId] = data
		b.SetUserData(userData)

		msg := domain.MessageToSend{
			ChatId: message.Chat.ID,
			Text:   "Отправте изображение, или вызовите команду /done",
		}

		err = b.SendMessage(msg, bot.DefaultChannel)
		if err != nil {
			return err
		}
	} else if message.Text == "/done" {
		msg := domain.MessageToSend{
			ChatId: message.Chat.ID,
			Text:   "Фотографии успешно загружены\nОтправте мне теги к записи одной строкой разделяя слова запятой",
		}

		err := b.SendMessage(msg, bot.DefaultChannel)
		if err != nil {
			return err
		}
		b.GetSwitcher().Next(userId)
	} else {
		msg := domain.MessageToSend{
			ChatId: message.Chat.ID,
			Text:   "НЕВЕРНОЕ СООБЩЕНИЕ!\nНужно, то отправте изображение, или вызвать команду /done",
		}
		err := b.SendMessage(msg, bot.DefaultChannel)
		if err != nil {
			return err
		}
	}
	return nil
}

type Tags struct {
	tags string `validator:"tags"`
}

const tagsFormat = `^[a-zA-Z0-9]+(,[a-zA-Z0-9]+)*$`

type WaitingTagsHandler struct {
	Tags
}

func (h *WaitingTagsHandler) validateTags(fl validator.FieldLevel) bool {
	tags := fl.Field().String()
	matched, _ := regexp.MatchString(tagsFormat, tags)
	return matched
}

func NewWaitingTagsHandler() *WaitingTagsHandler {
	return &WaitingTagsHandler{}
}

func (h *WaitingTagsHandler) Handle(b bot.IBot, message *tgbotapi.Message) error {
	validate := validator.New()
	err := validate.RegisterValidation("tags", h.validateTags)
	if err != nil {
		return err
	}

	tags := Tags{
		tags: message.Text,
	}

	if err = validate.Struct(tags); err != nil {
		msg := domain.MessageToSend{
			ChatId: message.Chat.ID,
			Text:   "НЕВЕРНОЕ СООБЩЕНИЕ",
		}
		err := b.SendMessage(msg, bot.DefaultChannel)
		if err != nil {
			return err
		}

	}

	userId := message.From.ID
	userData := b.GetUserData()
	data := userData[userId]

	tagsString := strings.Split(message.Text, ",")
	data.Tags = tagsString

	userData[userId] = data
	b.SetUserData(userData)
	b.GetSwitcher().Next(userId)

	msg := domain.MessageToSend{
		ChatId: message.Chat.ID,
		Text:   "Теги успешно записаны!\nОтправте дату дедлайна записи. Формат:yyyy-mm-dd",
	}

	err = b.SendMessage(msg, bot.DefaultChannel)
	if err != nil {
		return err
	}
	return nil
}

type WaitingDeadlineHandler struct{}

func NewWaitingDeadlineHandler() *WaitingDeadlineHandler {
	return &WaitingDeadlineHandler{}
}

func (h *WaitingDeadlineHandler) Handle(b bot.IBot, message *tgbotapi.Message) error {
	validate := validator.New()

	err := validate.RegisterValidation("customDate", func(fl validator.FieldLevel) bool {
		dateStr := fl.Field().String()
		_, err := time.Parse("2006-01-02", dateStr)
		return err == nil
	})
	if err != nil {
		return err
	}

	err = validate.Var(message.Text, "required,customDate")
	if err != nil {
		msg := domain.MessageToSend{
			ChatId: message.Chat.ID,
			Text:   "НЕВЕРНОЕ СООБЩЕНИЕ\nВведите ещё раз",
		}
		err := b.SendMessage(msg, bot.DefaultChannel)
		if err != nil {
			return err
		}
	}

	userId := message.From.ID

	userData := b.GetUserData()
	data := userData[userId]

	layout := "2006-01-02"
	parsed, err := time.Parse(layout, message.Text)
	if err != nil {
		return err
	}

	data.Deadline = parsed
	userData[userId] = data
	b.SetUserData(userData)
	b.GetSwitcher().Next(userId)
	return nil
}

type WaitingIdHandler struct{}

func NewWaitingIdHandler() *WaitingIdHandler {
	return &WaitingIdHandler{}
}

func (h *WaitingIdHandler) Handle(b bot.IBot, message *tgbotapi.Message) error {
	validate := validator.New()

	err := validate.Var(message.Text, "required,number")
	if err != nil {
		err = fmt.Errorf("failed to validate text: %v", err)
		return err
	}

	userData := b.GetUserData()
	data := userData[message.From.ID]
	data.Id, err = strconv.Atoi(message.Text)
	if err != nil {
		err = fmt.Errorf("failed to parse id: %v", err)
		return err
	}

	msg := domain.MessageToSend{
		ChatId: message.Chat.ID,
		Text:   "Напишите новое название вашего дз/записи или напишите /done",
	}

	err = b.SendMessage(msg, bot.DefaultChannel)
	if err != nil {
		err = fmt.Errorf("failed to send message: %v", err)
		return err
	}

	userId := message.From.ID
	userData[userId] = data
	b.SetUserData(userData)
	b.GetSwitcher().Next(userId)
	return nil
}
