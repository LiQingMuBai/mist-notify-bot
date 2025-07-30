package service

import (
	"context"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"gorm.io/gorm"
	"strings"
	"ushield_bot/internal/infrastructure/repositories"
	. "ushield_bot/internal/infrastructure/tools"
)

func ExtractAddressDetection(db *gorm.DB, callbackQuery *tgbotapi.CallbackQuery) tgbotapi.MessageConfig {
	userRepo := repositories.NewUserRepository(db)
	user, _ := userRepo.GetByUserID(callbackQuery.Message.Chat.ID)
	if IsEmpty(user.Amount) {
		user.Amount = "0.00"
	}

	if IsEmpty(user.TronAmount) {
		user.TronAmount = "0.00"
	}

	usdtDepositRepo := repositories.NewUserUSDTDepositsRepository(db)
	usdtlist, _ := usdtDepositRepo.ListAll(context.Background(), callbackQuery.Message.Chat.ID, 1)

	trxDepositRepo := repositories.NewUserTRXDepositsRepository(db)
	trxlist, _ := trxDepositRepo.ListAll(context.Background(), callbackQuery.Message.Chat.ID, 1)

	var builder strings.Builder
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

	var builder2 strings.Builder
	//- [6.29] +3000 TRX（订单 #TOPUP-92308）
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

	// 去除最后一个空格
	result2 := strings.TrimSpace(builder2.String())
	msg := tgbotapi.NewMessage(callbackQuery.Message.Chat.ID, "🧾扣款记录\n\n "+
		result+"\n"+
		result2+"\n")
	msg.ParseMode = "HTML"
	inlineKeyboard := tgbotapi.NewInlineKeyboardMarkup(
		tgbotapi.NewInlineKeyboardRow(
			tgbotapi.NewInlineKeyboardButtonData("上一页", "click_deposit_trx_records"),
			tgbotapi.NewInlineKeyboardButtonData("下一页", "click_cost_records"),
		),
		tgbotapi.NewInlineKeyboardRow(
			//tgbotapi.NewInlineKeyboardButtonData("解绑地址", "free_monitor_address"),
			tgbotapi.NewInlineKeyboardButtonData("🔙返回个人中心", "back_home"),
		),
	)
	msg.ReplyMarkup = inlineKeyboard
	return msg
}
