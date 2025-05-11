package telegram

import (
	"fmt"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"github.com/sirupsen/logrus"
	"homework_bot/internal/application/services"
	"homework_bot/internal/bot"
	"homework_bot/internal/bot/telegram/handler"
	"homework_bot/internal/domain"

	"homework_bot/pkg/switcher"
)

type Bot struct {
	bot        *tgbotapi.BotAPI
	userStates map[int64]string

	services *services.Service
	switcher *switcher.Switcher

	Task *switcher.TaskFlowManager
}

func NewBot(b *tgbotapi.BotAPI, service *services.Service) *Bot {
	statusesAdd := []string{
		bot.WaitingName,
		bot.WaitingDescription,
		bot.WaitingImages,
		bot.WaitingTags,
		bot.WaitingDeadline,
	}

	statusesUpdate := []string{bot.WaitingId}
	statusesUpdate = append(statusesUpdate, statusesAdd...)
	statusesGetTags := []string{bot.WaitingTags}
	statusesAskGroup := []string{bot.WaitingGroup}

	manager := switcher.NewTaskFlowManager()

	return &Bot{
		bot:        b,
		services:   service,
		switcher:   switcher.NewSwitcher(statusesAdd, statusesUpdate, statusesGetTags, statusesAskGroup),
		userStates: make(map[int64]string),
		Task:       manager,
	}
}

func (b *Bot) GetUserStates() map[int64]string {
	return b.userStates
}

func (b *Bot) GetServices() *services.Service {
	return b.services
}

func (b *Bot) GetSwitcher() *switcher.Switcher {
	return b.switcher
}

func (b *Bot) Start() error {
	_, err := b.bot.Request(getCommandMenu())
	if err != nil {
		return err
	}
	updates := b.initUpdatesChannel()
	b.handleUpdates(updates)
	return nil
}
func (b *Bot) GetTaskManager() *switcher.TaskFlowManager {
	return b.Task
}
func (b *Bot) GetBot() *tgbotapi.BotAPI {
	return b.bot
}

func (b *Bot) handleUpdate(update tgbotapi.Update) {
	if update.Message == nil {
		return
	}

	if update.Message.Chat.Type == "supergroup" {
		return
	}

	factory := handler.NewFactory()
	h := factory.GetHandler(b, update.Message)

	//log.Println("==============", h)
	err := h.Handle(b, update.Message)
	if err != nil {
		logrus.Errorf("Error with handlers: %v", err)
	}
}

func (b *Bot) handleUpdates(updates tgbotapi.UpdatesChannel) {
	for update := range updates {
		go func(update tgbotapi.Update) {
			b.handleUpdate(update)
		}(update)
	}
}

func (b *Bot) initUpdatesChannel() tgbotapi.UpdatesChannel {
	u := tgbotapi.NewUpdate(0)
	u.Timeout = 10

	updates := b.bot.GetUpdatesChan(u)

	return updates
}

func (b *Bot) sendMediaGroup(message domain.MessageToSend, channel int) error {
	var mediaGroup []interface{}

	for i, photo := range message.Images {
		inputPhoto := tgbotapi.NewInputMediaPhoto(tgbotapi.FilePath(photo))
		if i == 0 {
			inputPhoto.Caption = message.Text
		}
		mediaGroup = append(mediaGroup, inputPhoto)
	}

	mediaGroupCfg := tgbotapi.NewMediaGroup(message.ChatId, mediaGroup)
	if channel == bot.ChannelBot {
		mediaGroupCfg.ReplyToMessageID = bot.ChannelBot
	} else if channel == bot.ChannelInformation {
		mediaGroupCfg.ReplyToMessageID = bot.ChannelInformation
	}

	_, err := b.bot.SendMediaGroup(mediaGroupCfg)
	return err
}

func (b *Bot) sendText(message domain.MessageToSend, channel int) error {
	msg := tgbotapi.NewMessage(message.ChatId, "")
	msg.Text = message.Text

	if channel == bot.ChannelBot {
		msg.ReplyToMessageID = bot.ChannelBot
	} else if channel == bot.ChannelInformation {
		msg.ReplyToMessageID = bot.ChannelInformation
	}

	_, err := b.bot.Send(msg)
	return err
}

func (b *Bot) SendMessage(message domain.MessageToSend, channel int) error {
	if len(message.Images) > 0 {
		return b.sendMediaGroup(message, channel)
	}
	return b.sendText(message, channel)
}

func (b *Bot) SendInputError(message *tgbotapi.Message) error {
	msg := domain.MessageToSend{
		ChatId: message.Chat.ID,
		Text:   "資訊不正確！我期望/命令數據",
	}
	err := b.SendMessage(msg, bot.DefaultChannel)
	if err != nil {
		return err
	}
	return fmt.Errorf("error in get on id")
}
