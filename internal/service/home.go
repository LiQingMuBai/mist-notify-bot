package service

import (
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"gorm.io/gorm"
	"strconv"
	"ushield_bot/internal/infrastructure/repositories"
	. "ushield_bot/internal/infrastructure/tools"
)

func BackHOME(db *gorm.DB, callbackQuery *tgbotapi.CallbackQuery, bot *tgbotapi.BotAPI) {
	inlineKeyboard := tgbotapi.NewInlineKeyboardMarkup(
		//tgbotapi.NewInlineKeyboardRow(
		//	tgbotapi.NewInlineKeyboardButtonData("🆔我的账户", "click_my_account"),
		//	tgbotapi.NewInlineKeyboardButtonData("💳充值", "click_my_deposit"),
		//),
		tgbotapi.NewInlineKeyboardRow(
			tgbotapi.NewInlineKeyboardButtonData("💳充值", "deposit_amount"),
			tgbotapi.NewInlineKeyboardButtonData("📄账单", "click_my_recepit"),
			//	tgbotapi.NewInlineKeyboardButtonData("🛠️我的服务", "click_my_service"),
		),
		tgbotapi.NewInlineKeyboardRow(
			//tgbotapi.NewInlineKeyboardButtonData("🔗绑定备用帐号", "click_backup_account"),
			tgbotapi.NewInlineKeyboardButtonData("👥商务合作", "click_business_cooperation"),
			tgbotapi.NewInlineKeyboardButtonData("🛎️客服", "click_callcenter"),
			tgbotapi.NewInlineKeyboardButtonData("❓常见问题FAQ", "click_QA"),
		),
		//tgbotapi.NewInlineKeyboardRow(
		//	tgbotapi.NewInlineKeyboardButtonData("👥商务合作", "click_business_cooperation"),
		//),
	)
	userRepo := repositories.NewUserRepository(db)
	user, _ := userRepo.GetByUserID(callbackQuery.Message.Chat.ID)

	if IsEmpty(user.Amount) {
		user.Amount = "0.00"
	}

	if IsEmpty(user.TronAmount) {
		user.TronAmount = "0.00"
	}

	str := ""
	if len(user.BackupChatID) > 0 {
		id, _ := strconv.ParseInt(user.BackupChatID, 10, 64)
		backup_user, _ := userRepo.GetByUserID(id)
		str = "🔗 已绑定备用账号  " + "@" + backup_user.Username + "（权限：观察者模式）"
	} else {
		str = "未绑定备用帐号"
	}

	msg := tgbotapi.NewMessage(callbackQuery.Message.Chat.ID, "📇 我的账户\n\n🆔 用户ID："+user.Associates+"\n\n👤 用户名：@"+user.Username+"\n\n"+
		str+"\n\n💰 "+
		"当前余额：\n\n"+
		"- TRX："+user.TronAmount+"\n"+
		"- USDT："+user.Amount)
	//msg := tgbotapi.NewMessage(callbackQuery.Message.Chat.ID, "📇 我的账户\n\n🆔 用户ID：123456789\n\n👤 用户名：@YourUsername\n\n🔗 已绑定备用账号/未绑定备用帐号\n\n@BackupUser01（权限：观察者模式）\n\n💰 当前余额：\n\n- TRX：73.50\n- USDT：2.00")
	msg.ReplyMarkup = inlineKeyboard
	msg.ParseMode = "HTML"
	bot.Send(msg)
}
