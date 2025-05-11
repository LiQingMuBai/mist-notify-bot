package handler

import (
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"log"
	"strconv"
	"strings"
	"ushield_bot/internal/bot"
	"ushield_bot/internal/domain"
)

type ExchangeHandler struct{}

func NewExchangeHandler() *ExchangeHandler {
	return &ExchangeHandler{}
}

func (h *ExchangeHandler) Handle(b bot.IBot, message *tgbotapi.Message) error {
	msg := domain.MessageToSend{
		ChatId: message.Chat.ID,
		Text:   "请输入能量转移目标地址\n" + "余额 100trx\n",
	}
	_ = b.SendMessage(msg, bot.DefaultChannel)

	//message.From.
	return nil
}

type ExchangeExecHandler struct{}

func NewExchangeExecHandler() *ExchangeExecHandler {
	return &ExchangeExecHandler{}
}

func (h *ExchangeExecHandler) Handle(b bot.IBot, message *tgbotapi.Message) error {
	text := message.Text
	username := message.From.UserName

	log.Println(username)

	if !strings.Contains(text, "_") {
		msg := domain.MessageToSend{
			ChatId: message.Chat.ID,
			Text:   "请输入正确的转账格式，地址_笔数\n",
		}
		_ = b.SendMessage(msg, bot.DefaultChannel)
	} else {

		target := strings.Split(text, "_")[0]
		num := strings.Split(text, "_")[1]
		userName := message.From.UserName
		user, err := b.GetServices().IUserService.GetByUsername(userName)
		if err != nil {
			msg := domain.MessageToSend{
				ChatId: message.Chat.ID,
				Text:   "请输入能量转移目标地址\n" + "余额 100trx\n",
			}
			_ = b.SendMessage(msg, bot.DefaultChannel)
		} else {
			userAmount, _ := strconv.ParseFloat(user.Amount, 64)
			num, _ := strconv.ParseFloat(num, 64)
			if userAmount < num*2.5 {
				msg := domain.MessageToSend{
					ChatId: message.Chat.ID,
					Text:   "对不起你的资金不够，请充值\n",
				}

				_ = b.SendMessage(msg, bot.DefaultChannel)
			}

			log.Println(target)
			//user.Amount
			//判断用户余额>笔数*trx，小于的话就报错，大于就执行转账

			//判断用户的地址输入是否正确

			//如果都正确补充能量

			//划扣资金

			msg := domain.MessageToSend{
				ChatId: message.Chat.ID,
				Text:   "请输入能量转移目标地址\n" + "余额 100trx\n",
			}

			_ = b.SendMessage(msg, bot.DefaultChannel)
		}
	}
	//message.From.
	return nil
}
