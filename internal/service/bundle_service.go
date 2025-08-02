package service

import (
	"context"
	"fmt"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"gorm.io/gorm"
	"strconv"
	"strings"
	"time"
	"ushield_bot/internal/cache"
	"ushield_bot/internal/infrastructure/repositories"
	. "ushield_bot/internal/infrastructure/tools"
)

func BUNDLE_CHECK(cache cache.Cache, bot *tgbotapi.BotAPI, callbackQuery *tgbotapi.CallbackQuery, db *gorm.DB) {
	//deductionAmount := callbackQuery.Data[7:len(callbackQuery.Data)]
	userOperationBundlesRepo := repositories.NewUserOperationBundlesRepository(db)
	bundleID := strings.ReplaceAll(callbackQuery.Data, "bundle_", "")
	bundlePackage, err := userOperationBundlesRepo.Query(context.Background(), bundleID)

	if err != nil {

	}

	deductionAmount := bundlePackage.Amount

	//fmt.Printf("deductionAmount: %v\n", deductionAmount)
	userRepo := repositories.NewUserRepository(db)
	user, _ := userRepo.GetByUserID(callbackQuery.Message.Chat.ID)
	if IsEmpty(user.Amount) {
		user.Amount = "0.00"
	}

	if IsEmpty(user.TronAmount) {
		user.TronAmount = "0.00"
	}

	fmt.Printf("user usdt balance : %s\n", user.Amount)
	fmt.Printf("user  trx balance : %s\n", user.TronAmount)
	fmt.Printf("deductionAmount : %s\n", deductionAmount)
	fmt.Printf("Token : %s\n", bundlePackage.Token)

	lessBalance := false
	if bundlePackage.Token == "USDT" {
		//扣usdt
		if flag, _ := CompareNumberStrings(user.Amount, deductionAmount); flag < 0 {
			lessBalance = true
		}
		fmt.Printf("bundle %v is USDT\n", bundlePackage)
	} else if bundlePackage.Token == "TRX" {
		//扣trx
		if flag, _ := CompareNumberStrings(user.TronAmount, deductionAmount); flag < 0 {
			lessBalance = true
		}

		fmt.Printf("bundle %v is trx\n", bundlePackage)
	}

	if lessBalance {
		msg := tgbotapi.NewMessage(callbackQuery.Message.Chat.ID,
			"💬"+"<b>"+"用户姓名: "+"</b>"+user.Username+"\n"+
				"👤"+"<b>"+"用户电报ID: "+"</b>"+user.Associates+"\n"+
				"💵"+"<b>"+"余额不足 "+"</b>"+"\n"+
				"💴"+"<b>"+"当前TRX余额:  "+"</b>"+user.TronAmount+" TRX"+"\n"+
				"💴"+"<b>"+"当前USDT余额:  "+"</b>"+user.Amount+" USDT")

		msg.ParseMode = "HTML"

		inlineKeyboard := tgbotapi.NewInlineKeyboardMarkup(
			tgbotapi.NewInlineKeyboardRow(
				tgbotapi.NewInlineKeyboardButtonData("💵充值", "deposit_amount"),
			),
		)

		msg.ReplyMarkup = inlineKeyboard
		bot.Send(msg)

		return
	}

	msg := tgbotapi.NewMessage(callbackQuery.Message.Chat.ID, "🧾"+"<b>"+"请选择接收能量的地址或重新输入 "+"</b>"+"\n")
	userOperationPackageAddressesRepo := repositories.NewUserOperationPackageAddressesRepo(db)

	addresses, _ := userOperationPackageAddressesRepo.Query(context.Background(), callbackQuery.Message.Chat.ID)

	//msg := tgbotapi.NewMessage(_chatID, "👇请选择要设置的地址："+"\n")
	//地址绑定

	msg.ParseMode = "HTML"

	var allButtons []tgbotapi.InlineKeyboardButton
	var extraButtons []tgbotapi.InlineKeyboardButton
	var keyboard [][]tgbotapi.InlineKeyboardButton
	for _, item := range addresses {
		allButtons = append(allButtons, tgbotapi.NewInlineKeyboardButtonData(item.Address, "apply_bundle_package_"+bundleID+"_"+item.Address))
	}

	extraButtons = append(extraButtons, tgbotapi.NewInlineKeyboardButtonData("🔙返回首页", "back_bundle_package"))

	for i := 0; i < len(allButtons); i += 1 {
		end := i + 1
		if end > len(allButtons) {
			end = len(allButtons)
		}
		row := allButtons[i:end]
		keyboard = append(keyboard, tgbotapi.NewInlineKeyboardRow(row...))
	}

	for i := 0; i < len(extraButtons); i += 1 {
		end := i + 1
		if end > len(extraButtons) {
			end = len(extraButtons)
		}
		row := extraButtons[i:end]
		keyboard = append(keyboard, tgbotapi.NewInlineKeyboardRow(row...))
	}

	// 3. 创建键盘标记
	inlineKeyboard := tgbotapi.NewInlineKeyboardMarkup(keyboard...)

	msg.ReplyMarkup = inlineKeyboard

	msg.ParseMode = "HTML"
	bot.Send(msg)

	expiration := 1 * time.Minute // 短时间缓存空值
	//设置用户状态
	cache.Set(strconv.FormatInt(callbackQuery.Message.Chat.ID, 10), "apply_bundle_package_"+bundleID, expiration)
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
