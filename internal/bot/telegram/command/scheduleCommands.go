package command

import (
	"github.com/go-playground/validator/v10"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"homework_bot/internal/bot"
	"homework_bot/internal/domain"
	"time"
)

type ScheduleWeekCommand struct{}

func isUser(b bot.IBot, message *tgbotapi.Message) (domain.User, error) {
	user, err := b.GetServices().GetByUsername(message.From.UserName)
	return user, err
}

func NewScheduleWeekCommand() *ScheduleWeekCommand {
	return &ScheduleWeekCommand{}
}

func (c *ScheduleWeekCommand) Exec(b bot.IBot, message *tgbotapi.Message) error {
	user, err := isUser(b, message)
	if err != nil {
		command := NewAskGroupCommand()
		return command.Exec(b, message)
	}

	schedule := b.GetServices().GetOnThisWeek(user)
	err = b.SendSchedule(schedule, message.Chat.ID, bot.DefaultChannel)
	return err
}

type ScheduleNextWeakCommand struct{}

func NewScheduleNextWeakCommand() *ScheduleNextWeakCommand {
	return &ScheduleNextWeakCommand{}
}

func (c *ScheduleNextWeakCommand) Exec(b bot.IBot, message *tgbotapi.Message) error {
	user, err := isUser(b, message)
	if err != nil {
		command := NewAskGroupCommand()
		return command.Exec(b, message)
	}

	schedule := b.GetServices().GetOnNextWeek(user)
	err = b.SendSchedule(schedule, message.Chat.ID, bot.DefaultChannel)
	return err
}

type ScheduleDayCommand struct{}

func NewScheduleDayCommand() *ScheduleDayCommand {
	return &ScheduleDayCommand{}
}

func (c *ScheduleDayCommand) Exec(b bot.IBot, message *tgbotapi.Message) error {
	user, err := isUser(b, message)
	if err != nil {
		command := NewAskGroupCommand()
		return command.Exec(b, message)
	}

	validate := validator.New()
	err = validate.RegisterValidation("customDate", func(fl validator.FieldLevel) bool {
		dateStr := fl.Field().String()
		_, err := time.Parse("2006-01-02", dateStr)
		return err == nil
	})
	if err != nil {
		return err
	}

	date, err := time.Parse("2006-01-02", message.Text)
	if err != nil {
		return err
	}

	schedule := b.GetServices().GetOnDate(user, date)
	err = b.SendSchedule(schedule, message.Chat.ID, bot.DefaultChannel)
	return err
}

type ScheduleTodayCommand struct{}

func NewScheduleTodayCommand() *ScheduleTodayCommand {
	return &ScheduleTodayCommand{}
}

func (c *ScheduleTodayCommand) Exec(b bot.IBot, message *tgbotapi.Message) error {
	user, err := isUser(b, message)
	if err != nil {
		command := NewAskGroupCommand()
		return command.Exec(b, message)
	}

	schedule := b.GetServices().GetOnToday(user)
	err = b.SendSchedule(schedule, message.Chat.ID, bot.DefaultChannel)
	return err
}

type ScheduleTomorrowCommand struct{}

func NewScheduleTomorrowCommand() *ScheduleTomorrowCommand {
	return &ScheduleTomorrowCommand{}
}

func (c *ScheduleTomorrowCommand) Exec(b bot.IBot, message *tgbotapi.Message) error {
	user, err := isUser(b, message)
	if err != nil {
		command := NewAskGroupCommand()
		return command.Exec(b, message)
	}

	schedule := b.GetServices().GetOnTomorrow(user)

	err = b.SendSchedule(schedule, message.Chat.ID, bot.DefaultChannel)
	return err
}
