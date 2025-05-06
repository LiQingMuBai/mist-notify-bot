package command

import (
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"github.com/sirupsen/logrus"
	"os"
	"strconv"
	"strings"
)

const (
	commandStart          = "start"
	commandHelp           = "help"
	commandGetAccount     = "get_account"
	commandExchangeEnergy = "exchange_energy"
)

type Factory struct {
}

func NewFactory() *Factory {
	return &Factory{}
}

func (f *Factory) isAdmin(chatId int64) bool {
	adminsString := os.Getenv("ADMIN4")
	adminId := strings.Split(adminsString, ",")
	for _, item := range adminId {
		id, err := strconv.ParseInt(item, 10, 64)
		if err != nil {
			logrus.Errorf("Error with convert string to int, %s", err.Error())
		}
		if id == chatId {
			return true
		}
	}

	return false
}

func (f *Factory) GetCommand(message *tgbotapi.Message) ICommand {
	switch message.Command() {
	//case用户关系

	case commandStart:
		return NewStartCommand()
	case commandExchangeEnergy:
		return NewExchangeEnergyCommand()
	case commandHelp:
		return NewHelpCommand()
	case commandGetAccount:
		return NewGetAccountCommand()
	default:
		return NewDefaultCommand()
	}
}
