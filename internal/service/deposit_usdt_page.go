package service

import (
	"context"
	"fmt"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"gorm.io/gorm"
	"strings"
	. "ushield_bot/internal/global"
	"ushield_bot/internal/infrastructure/repositories"
	"ushield_bot/internal/request"
)

func EXTRACT_PREV_DEPOSIT_USDT_PAGE(callbackQuery *tgbotapi.CallbackQuery, db *gorm.DB, bot *tgbotapi.BotAPI) (*DepositState, bool) {
	state := DepositStates[callbackQuery.Message.Chat.ID]

	if state != nil && state.CurrentPage == 1 {
		return nil, true
	}
	if state == nil {
		var state DepositState
		state.CurrentPage = 1
		DepositStates[callbackQuery.Message.Chat.ID] = &state
		usdtDepositRepo := repositories.NewUserUSDTDepositsRepository(db)
		var info request.UserUsdtDepositsSearch
		info.PageInfo.Page = 1
		info.PageInfo.PageSize = 10
		trxlist, _, _ := usdtDepositRepo.GetUserUsdtDepositsInfoList(context.Background(), info, callbackQuery.Message.Chat.ID)
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
	} else {
		state.CurrentPage = state.CurrentPage - 1
		usdtDepositRepo := repositories.NewUserUSDTDepositsRepository(db)
		var info request.UserUsdtDepositsSearch
		info.PageInfo.Page = state.CurrentPage
		info.PageInfo.PageSize = 10
		trxlist, _, _ := usdtDepositRepo.GetUserUsdtDepositsInfoList(context.Background(), info, callbackQuery.Message.Chat.ID)
		var builder strings.Builder
		builder.WriteString("\n") // 添加分隔符
		//- [6.29] +3000 TRX（订单 #TOPUP-92308）
		for _, word := range trxlist {
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
		result := strings.TrimSpace(builder.String())
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
	return state, false
}

func ExtraNextDepositUSDTPage(callbackQuery *tgbotapi.CallbackQuery, db *gorm.DB, bot *tgbotapi.BotAPI) bool {
	state := DepositStates[callbackQuery.Message.Chat.ID]
	if state == nil {
		var state2 DepositState
		state2.CurrentPage = 1
		state = &state2
	}
	//if state != nil && state.CurrentPage > 1 {
	state.CurrentPage = state.CurrentPage + 1
	usdtDepositRepo := repositories.NewUserUSDTDepositsRepository(db)
	var info request.UserUsdtDepositsSearch
	info.PageInfo.Page = state.CurrentPage
	info.PageInfo.PageSize = 10
	trxlist, total, _ := usdtDepositRepo.GetUserUsdtDepositsInfoList(context.Background(), info, callbackQuery.Message.Chat.ID)

	fmt.Printf("currentpage : %d", state.CurrentPage)
	fmt.Printf("total: %v\n", total)
	totalPages := (total + 5 - 1) / 5

	fmt.Printf("totalPages : %d", totalPages)
	if int64(state.CurrentPage) > totalPages {
		state.CurrentPage = totalPages
		return true
	}
	var builder strings.Builder
	builder.WriteString("\n") // 添加分隔符
	//- [6.29] +3000 TRX（订单 #TOPUP-92308）
	for _, word := range trxlist {
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
	result := strings.TrimSpace(builder.String())
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
	//}
	fmt.Printf("state: %v\n", state)

	DepositStates[callbackQuery.Message.Chat.ID] = state
	return false
}
