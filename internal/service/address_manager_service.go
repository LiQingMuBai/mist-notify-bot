package service

import (
	"context"
	"fmt"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"gorm.io/gorm"
	"strconv"
	"time"
	"ushield_bot/internal/cache"
	"ushield_bot/internal/domain"
	"ushield_bot/internal/infrastructure/repositories"
	. "ushield_bot/internal/infrastructure/tools"
)

func ExtractAddressManager(message *tgbotapi.Message, db *gorm.DB, bot *tgbotapi.BotAPI) {
	if IsValidAddress(message.Text) || IsValidEthereumAddress(message.Text) {
		userRepo := repositories.NewUserAddressMonitorRepo(db)
		var record domain.UserAddressMonitor
		record.ChatID = message.Chat.ID
		record.Address = message.Text
		record.Status = 1
		if IsValidAddress(message.Text) {
			record.Network = "tron"
		}
		if IsValidAddress(message.Text) {
			record.Network = "ethereum"
		}
		errsg := userRepo.Create(context.Background(), &record)
		if errsg != nil {
		}

		msg := tgbotapi.NewMessage(message.Chat.ID, "💬"+"<b>"+"地址添加成功 "+"</b>"+"\n")
		msg.ParseMode = "HTML"
		bot.Send(msg)

	} else {
		msg := tgbotapi.NewMessage(message.Chat.ID, "💬"+"<b>"+"地址有误，请重新输入需添加的地址: "+"</b>"+"\n")
		msg.ParseMode = "HTML"
		bot.Send(msg)
	}
}

func ADDRESS_LIST_TRACE(cache cache.Cache, bot *tgbotapi.BotAPI, callbackQuery *tgbotapi.CallbackQuery, db *gorm.DB) {
	userAddressEventRepo := repositories.NewUserAddressMonitorEventRepo(db)
	addresses, _ := userAddressEventRepo.Query(context.Background(), callbackQuery.Message.Chat.ID)
	// 初始化结果字符串
	var result string

	// 遍历数组并拼接字符串
	for i, item := range addresses {
		if i > 0 {
			result += " ✅\n\n" // 添加分隔符
		}

		restDays := fmt.Sprintf("%d", 30-item.Days)

		result += item.Address + "（剩余" + restDays + "）"
	}
	result += " ✅\n\n" // 添加分隔符
	//查看余额
	userRepo := repositories.NewUserRepository(db)
	user, _ := userRepo.GetByUserID(callbackQuery.Message.Chat.ID)
	if IsEmpty(user.Amount) {
		user.Amount = "0.00"
	}

	if IsEmpty(user.TronAmount) {
		user.TronAmount = "0.00"
	}

	msg := tgbotapi.NewMessage(callbackQuery.Message.Chat.ID, "有服务进行中\n\n📊 当前正在监控的地址：\n\n"+
		result+
		"💼 当前余额："+"\n- "+user.TronAmount+" TRX \n - "+user.Amount+" USDT \n"+
		"📌请保持余额充足，到期将自动续费\n"+
		"如需中止服务，可随时")
	msg.ParseMode = "HTML"

	inlineKeyboard := tgbotapi.NewInlineKeyboardMarkup(
		tgbotapi.NewInlineKeyboardRow(
			//tgbotapi.NewInlineKeyboardButtonData("解绑地址", "free_monitor_address"),
			tgbotapi.NewInlineKeyboardButtonData("停止监控", "stop_freeze_risk"),
		),
		tgbotapi.NewInlineKeyboardRow(
			tgbotapi.NewInlineKeyboardButtonData("第二紧急通知", "user_backup_notify"),
			//tgbotapi.NewInlineKeyboardButtonData("第二紧急通知", ""),
		),
	)
	msg.ReplyMarkup = inlineKeyboard

	bot.Send(msg)

	expiration := 1 * time.Minute // 短时间缓存空值
	//设置用户状态
	cache.Set(strconv.FormatInt(callbackQuery.Message.Chat.ID, 10), "address_list_trace", expiration)
}

func ADDRESS_MANAGER(cache cache.Cache, bot *tgbotapi.BotAPI, callbackQuery *tgbotapi.CallbackQuery, db *gorm.DB) {
	userAddressRepo := repositories.NewUserAddressMonitorRepo(db)

	addresses, _ := userAddressRepo.Query(context.Background(), callbackQuery.Message.Chat.ID)

	result := ""
	for _, item := range addresses {
		result += item.Address + "\n"
	}
	msg := tgbotapi.NewMessage(callbackQuery.Message.Chat.ID, "👇以下监控地址信息列表"+"\n"+result)
	//地址绑定

	msg.ParseMode = "HTML"

	inlineKeyboard := tgbotapi.NewInlineKeyboardMarkup(
		tgbotapi.NewInlineKeyboardRow(
			tgbotapi.NewInlineKeyboardButtonData("➕添加钱包", "address_manager_add"),
			//tgbotapi.NewInlineKeyboardButtonData("设置钱包", "address_manager"),
			tgbotapi.NewInlineKeyboardButtonData("➖删除钱包", "address_manager_remove"),
		),
		tgbotapi.NewInlineKeyboardRow(
			tgbotapi.NewInlineKeyboardButtonData("⬅️返回个人中心", "back_home"),
		),
	)
	msg.ReplyMarkup = inlineKeyboard

	bot.Send(msg)

	expiration := 1 * time.Minute // 短时间缓存空值

	//设置用户状态
	cache.Set(strconv.FormatInt(callbackQuery.Message.Chat.ID, 10), "address_manager", expiration)
}
