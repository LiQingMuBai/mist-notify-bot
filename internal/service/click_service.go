package service

import (
	"context"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"gorm.io/gorm"
	"strings"
	"ushield_bot/internal/infrastructure/repositories"
	. "ushield_bot/internal/infrastructure/tools"
	"ushield_bot/internal/request"
)

func CLICK_DEPOSIT_USDT_RECORDS(db *gorm.DB, callbackQuery *tgbotapi.CallbackQuery, bot *tgbotapi.BotAPI) {
	userRepo := repositories.NewUserRepository(db)
	user, _ := userRepo.GetByUserID(callbackQuery.Message.Chat.ID)
	if IsEmpty(user.Amount) {
		user.Amount = "0.00"
	}

	if IsEmpty(user.TronAmount) {
		user.TronAmount = "0.00"
	}

	usdtDepositRepo := repositories.NewUserUSDTDepositsRepository(db)

	//trxDepositRepo := repositories.NewUserTRXDepositsRepository(db)
	var info request.UserUsdtDepositsSearch
	info.PageInfo.Page = 1
	info.PageInfo.PageSize = 5
	//trxlist, _, _ := trxDepositRepo.GetUserTrxDepositsInfoList(context.Background(), info, callbackQuery.Message.Chat.ID)
	usdtlist, _, _ := usdtDepositRepo.GetUserUsdtDepositsInfoList(context.Background(), info, callbackQuery.Message.Chat.ID)

	var builder strings.Builder
	builder.WriteString("\n") // 添加分隔符

	// 去除最后一个空格
	result := strings.TrimSpace(builder.String())

	for _, word := range usdtlist {
		builder.WriteString("[")
		builder.WriteString(word.CreatedDate)
		builder.WriteString("]")
		builder.WriteString("+")
		builder.WriteString(word.Amount)
		builder.WriteString(" USDT ")
		builder.WriteString(" （订单 #TOPUP- ")
		builder.WriteString(word.OrderNO)
		builder.WriteString("）")

		builder.WriteString("\n") // 添加分隔符
	}
	//
	//// 去除最后一个空格
	result = strings.TrimSpace(builder.String())
	msg := tgbotapi.NewMessage(callbackQuery.Message.Chat.ID, "🧾充值记录\n\n "+
		result+"\n")
	msg.ParseMode = "HTML"
	inlineKeyboard := tgbotapi.NewInlineKeyboardMarkup(
		tgbotapi.NewInlineKeyboardRow(
			tgbotapi.NewInlineKeyboardButtonData("上一页", "prev_deposit_usdt_page"),
			tgbotapi.NewInlineKeyboardButtonData("下一页", "next_deposit_usdt_page"),
		),
		tgbotapi.NewInlineKeyboardRow(
			//tgbotapi.NewInlineKeyboardButtonData("解绑地址", "free_monitor_address"),
			tgbotapi.NewInlineKeyboardButtonData("🔙返回个人中心", "back_home"),
		),
	)
	msg.ReplyMarkup = inlineKeyboard
	bot.Send(msg)
}

func ClickBusinessCooperation(callbackQuery *tgbotapi.CallbackQuery, bot *tgbotapi.BotAPI) {
	msg := tgbotapi.NewMessage(callbackQuery.Message.Chat.ID, "👥加入商务合作VIP群：https://t.me/+OCevU0Q12V8wZGY1\n")
	msg.ParseMode = "HTML"
	inlineKeyboard := tgbotapi.NewInlineKeyboardMarkup(
		tgbotapi.NewInlineKeyboardRow(
			//tgbotapi.NewInlineKeyboardButtonData("解绑地址", "free_monitor_address"),
			tgbotapi.NewInlineKeyboardButtonData("返回个人中心", "back_home"),
		),
	)
	msg.ReplyMarkup = inlineKeyboard
	bot.Send(msg)
}
func ClickCallCenter(callbackQuery *tgbotapi.CallbackQuery, bot *tgbotapi.BotAPI) {
	msg := tgbotapi.NewMessage(callbackQuery.Message.Chat.ID, "📞联系客服：@Ushield001\n")
	msg.ParseMode = "HTML"
	inlineKeyboard := tgbotapi.NewInlineKeyboardMarkup(
		tgbotapi.NewInlineKeyboardRow(
			//tgbotapi.NewInlineKeyboardButtonData("解绑地址", "free_monitor_address"),
			tgbotapi.NewInlineKeyboardButtonData("返回个人中心", "back_home"),
		),
	)
	msg.ReplyMarkup = inlineKeyboard
	bot.Send(msg)
}

func CLICK_DEPOSIT_TRX_RECORDS(db *gorm.DB, callbackQuery *tgbotapi.CallbackQuery, bot *tgbotapi.BotAPI) {
	userRepo := repositories.NewUserRepository(db)
	user, _ := userRepo.GetByUserID(callbackQuery.Message.Chat.ID)
	if IsEmpty(user.Amount) {
		user.Amount = "0.00"
	}

	if IsEmpty(user.TronAmount) {
		user.TronAmount = "0.00"
	}

	//usdtDepositRepo := repositories.NewUserUSDTDepositsRepository(db)
	//usdtlist, _ := usdtDepositRepo.ListAll(context.Background(), callbackQuery.Message.Chat.ID, 1)

	trxDepositRepo := repositories.NewUserTRXDepositsRepository(db)
	var info request.UserTrxDepositsSearch
	info.PageInfo.Page = 1
	info.PageInfo.PageSize = 5
	trxlist, _, _ := trxDepositRepo.GetUserTrxDepositsInfoList(context.Background(), info, callbackQuery.Message.Chat.ID)

	var builder strings.Builder
	builder.WriteString("\n") // 添加分隔符
	//- [6.29] +3000 TRX（订单 #TOPUP-92308）
	for _, word := range trxlist {
		builder.WriteString("[")
		builder.WriteString(word.CreatedDate)
		builder.WriteString("]")
		builder.WriteString("+")
		builder.WriteString(word.Amount)
		builder.WriteString(" TRX ")
		builder.WriteString(" （订单 #TOPUP- ")
		builder.WriteString(word.OrderNO)
		builder.WriteString("）")

		builder.WriteString("\n") // 添加分隔符
	}

	// 去除最后一个空格
	result := strings.TrimSpace(builder.String())

	msg := tgbotapi.NewMessage(callbackQuery.Message.Chat.ID, "🧾充值记录\n\n "+
		result+"\n")
	msg.ParseMode = "HTML"
	inlineKeyboard := tgbotapi.NewInlineKeyboardMarkup(
		tgbotapi.NewInlineKeyboardRow(
			tgbotapi.NewInlineKeyboardButtonData("上一页", "prev_deposit_trx_page"),
			tgbotapi.NewInlineKeyboardButtonData("下一页", "next_deposit_trx_page"),
		),
		tgbotapi.NewInlineKeyboardRow(
			//tgbotapi.NewInlineKeyboardButtonData("解绑地址", "free_monitor_address"),
			tgbotapi.NewInlineKeyboardButtonData("🔙返回个人中心", "back_home"),
		),
	)
	msg.ReplyMarkup = inlineKeyboard
	bot.Send(msg)
}
func CLICK_MY_RECEPIT(db *gorm.DB, callbackQuery *tgbotapi.CallbackQuery, bot *tgbotapi.BotAPI) {
	userRepo := repositories.NewUserRepository(db)
	user, _ := userRepo.GetByUserID(callbackQuery.Message.Chat.ID)
	if IsEmpty(user.Amount) {
		user.Amount = "0.00"
	}

	if IsEmpty(user.TronAmount) {
		user.TronAmount = "0.00"
	}

	msg := tgbotapi.NewMessage(callbackQuery.Message.Chat.ID, "🧾 我的账单记录\n\n📌 "+
		"当前余额：\n\n- TRX："+user.TronAmount+"\n- USDT："+user.Amount+"\n")

	msg.ParseMode = "HTML"
	inlineKeyboard := tgbotapi.NewInlineKeyboardMarkup(
		tgbotapi.NewInlineKeyboardRow(
			tgbotapi.NewInlineKeyboardButtonData("⬇️TRX充值记录", "click_deposit_trx_records"),
			tgbotapi.NewInlineKeyboardButtonData("⬇️USDT充值记录", "click_deposit_usdt_records"),
		),
		tgbotapi.NewInlineKeyboardRow(
			//tgbotapi.NewInlineKeyboardButtonData("解绑地址", "free_monitor_address"),
			tgbotapi.NewInlineKeyboardButtonData("🔙返回个人中心", "back_home"),
		),
	)
	msg.ReplyMarkup = inlineKeyboard
	bot.Send(msg)
}
