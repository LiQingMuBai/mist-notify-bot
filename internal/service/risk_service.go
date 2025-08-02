package service

import (
	"context"
	"fmt"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"gorm.io/gorm"
	"ushield_bot/internal/domain"
	"ushield_bot/internal/infrastructure/repositories"
	. "ushield_bot/internal/infrastructure/tools"
)

func START_FREEZE_RISK_1(db *gorm.DB, callbackQuery *tgbotapi.CallbackQuery, bot *tgbotapi.BotAPI) {
	userRepo := repositories.NewUserRepository(db)
	user, _ := userRepo.GetByUserID(callbackQuery.Message.Chat.ID)
	if IsEmpty(user.Amount) {
		user.Amount = "0.00"
	}

	if IsEmpty(user.TronAmount) {
		user.TronAmount = "0.00"
	}

	userAddressRepo := repositories.NewUserAddressMonitorRepo(db)

	addresses, _ := userAddressRepo.Query(context.Background(), callbackQuery.Message.Chat.ID)

	nums := len(addresses)
	//扣trx
	var COST_FROM_TRX bool
	var COST_FROM_USDT bool

	if CompareStringsWithFloat(user.TronAmount, "2800", float64(nums)) || CompareStringsWithFloat(user.Amount, "800", float64(nums)) {
		//扣减

		if CompareStringsWithFloat(user.TronAmount, "2800", float64(nums)) {
			rest, _ := SubtractStringNumbers(user.TronAmount, "2800", float64(nums))

			user.TronAmount = rest
			userRepo.Update2(context.Background(), &user)
			fmt.Printf("rest: %s", rest)
			COST_FROM_TRX = true
			//扣usdt
		} else if CompareStringsWithFloat(user.Amount, "800", float64(nums)) {
			rest, _ := SubtractStringNumbers(user.Amount, "800", float64(nums))
			fmt.Printf("rest: %s", rest)
			user.Amount = rest
			userRepo.Update2(context.Background(), &user)
			COST_FROM_USDT = true
		}

		//添加记录
		userAddressEventRepo := repositories.NewUserAddressMonitorEventRepo(db)

		for _, address := range addresses {
			var event domain.UserAddressMonitorEvent
			event.ChatID = callbackQuery.Message.Chat.ID
			event.Status = 1
			event.Address = address.Address
			event.Network = address.Network
			event.Days = 1
			if COST_FROM_TRX {
				event.Amount = "2800 TRX"
			}
			if COST_FROM_USDT {
				event.Amount = "800 USDT"
			}
			userAddressEventRepo.Create(context.Background(), &event)
		}
		//后台跟踪起来
		user, _ := userRepo.GetByUserID(callbackQuery.Message.Chat.ID)
		msg := tgbotapi.NewMessage(callbackQuery.Message.Chat.ID,
			"💬"+"<b>"+"用户姓名: "+"</b>"+user.Username+"\n"+
				"👤"+"<b>"+"用户电报ID: "+"</b>"+user.Associates+"\n"+
				"💵"+"<b>"+"当前TRX余额:  "+"</b>"+user.TronAmount+" TRX"+"\n"+
				"💴"+"<b>"+"当前USDT余额:  "+"</b>"+user.Amount+" USDT")
		msg.ParseMode = "HTML"
		inlineKeyboard := tgbotapi.NewInlineKeyboardMarkup(
			tgbotapi.NewInlineKeyboardRow(
				tgbotapi.NewInlineKeyboardButtonData("⬅️返回", "address_manager_return"),
			),
		)

		msg.ReplyMarkup = inlineKeyboard
		bot.Send(msg)
	} else {

		//余额不足，需充值
		msg := tgbotapi.NewMessage(callbackQuery.Message.Chat.ID,
			"💬"+"<b>"+"用户姓名: "+"</b>"+user.Username+"\n"+
				"👤"+"<b>"+"用户电报ID: "+"</b>"+user.Associates+"\n"+
				"💵"+"<b>"+"当前TRX余额:  "+"</b>"+user.TronAmount+" TRX"+"\n"+
				"💴"+"<b>"+"当前USDT余额:  "+"</b>"+user.Amount+" USDT")
		msg.ParseMode = "HTML"
		inlineKeyboard := tgbotapi.NewInlineKeyboardMarkup(
			tgbotapi.NewInlineKeyboardRow(
				tgbotapi.NewInlineKeyboardButtonData("💵充值", "deposit_amount"),
			),
		)

		msg.ReplyMarkup = inlineKeyboard
		bot.Send(msg)
	}
}
