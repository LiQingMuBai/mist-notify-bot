package service

import (
	"context"
	"fmt"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"gorm.io/gorm"
	"strconv"
	"time"
	"ushield_bot/internal/cache"
	"ushield_bot/internal/infrastructure/repositories"
	. "ushield_bot/internal/infrastructure/tools"
)

func BUNDLE_CHECK(cache cache.Cache, bot *tgbotapi.BotAPI, callbackQuery *tgbotapi.CallbackQuery, db *gorm.DB) {
	deductionAmount := callbackQuery.Data[7:len(callbackQuery.Data)]
	fmt.Printf("deductionAmount: %v\n", deductionAmount)
	userRepo := repositories.NewUserRepository(db)
	user, _ := userRepo.GetByUserID(callbackQuery.Message.Chat.ID)

	if flag, _ := CompareNumberStrings(user.Amount, deductionAmount); flag < 0 {
		msg := tgbotapi.NewMessage(callbackQuery.Message.Chat.ID,
			"💬"+"<b>"+"用户姓名: "+"</b>"+user.Username+"\n"+
				"👤"+"<b>"+"用户电报ID: "+"</b>"+user.Associates+"\n"+
				"💵"+"<b>"+"USDT余额不足 "+"</b>"+"\n"+
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

	msg := tgbotapi.NewMessage(callbackQuery.Message.Chat.ID, "💬"+"<b>"+"请输入能量接收地址: "+"</b>"+"\n")
	msg.ParseMode = "HTML"
	bot.Send(msg)

	expiration := 1 * time.Minute // 短时间缓存空值

	//设置用户状态
	cache.Set(strconv.FormatInt(callbackQuery.Message.Chat.ID, 10), callbackQuery.Data, expiration)
	//扣款
}

func ExtractBundleService(message *tgbotapi.Message, bot *tgbotapi.BotAPI, db *gorm.DB, status string) bool {
	if !IsValidAddress(message.Text) {
		msg := tgbotapi.NewMessage(message.Chat.ID, "💬"+"<b>"+"地址有误，请重新输入能量接收地址: "+"</b>"+"\n")
		msg.ParseMode = "HTML"
		bot.Send(msg)
		return true
	}

	userRepo := repositories.NewUserRepository(db)
	user, _ := userRepo.GetByUserID(message.Chat.ID)

	fee := status[7:len(status)]
	fmt.Println("status : ", status)
	fmt.Println("fee : ", fee)
	fmt.Println("amount :", user.Amount)

	if CompareStringsWithFloat(fee, user.Amount, 1) {
		//余额不足，需充值
		msg := tgbotapi.NewMessage(message.Chat.ID,
			"💬"+"<b>"+"余额不足: "+"</b>"+"\n"+
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
	} else {
		bundlesRepo := repositories.NewUserOperationBundlesRepository(db)

		bundleRecord, _ := bundlesRepo.Find(context.Background(), fee)
		//10笔（12U）
		bundleNum := bundleRecord.Name
		count, _ := ExtractNumberBeforeBi(bundleNum)

		fmt.Printf("笔数count : %d", count)
		//扣款
		//调用trxfee接口

		//trxfeeHandler := handler.NewTrxfeeHandler()

		//trxfeeHandler.RequestTimesOrder(context.Background(),"","",message.Text,)
		rest, _ := SubtractStringNumbers(user.Amount, fee, 1)
		user.Amount = rest
		userRepo.Update2(context.Background(), &user)
		fmt.Println("rest :", rest)

		msg := tgbotapi.NewMessage(message.Chat.ID,
			"<b>"+"✅笔数套餐订阅成功"+"</b>"+"\n"+
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
	return false
}
