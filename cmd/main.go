package main

import (
	"context"
	"fmt"
	"github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"github.com/joho/godotenv"
	"github.com/sirupsen/logrus"
	"github.com/spf13/viper"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"log"
	"os"
	"strconv"
	"strings"
	"time"
	"ushield_bot/internal/service"

	"ushield_bot/internal/cache"
	"ushield_bot/internal/domain"
	"ushield_bot/internal/infrastructure/repositories"
	. "ushield_bot/internal/infrastructure/tools"
)

func initConfig() error {
	viper.AddConfigPath("configs")
	viper.SetConfigName("config")
	return viper.ReadInConfig()
}
func main() {

	logrus.SetFormatter(&logrus.JSONFormatter{})

	if err := initConfig(); err != nil {
		logrus.Fatalf("init configs err: %s", err.Error())
	}

	if err := godotenv.Load(); err != nil {
		logrus.Fatalf("load .env file err: %s", err.Error())
	}

	// Database connection string
	host := viper.GetString("db.host")
	port := viper.GetString("db.port")
	username := viper.GetString("db.username")
	password := viper.GetString("db.password")
	dbname := viper.GetString("db.dbname")

	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=Local", username, password, host, port, dbname)
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})

	if err != nil {
		panic("Failed to connect to the database: " + err.Error())
	}
	TG_BOT_API := os.Getenv("TG_BOT_API")
	bot, err := tgbotapi.NewBotAPI(TG_BOT_API)
	if err != nil {
		log.Panic(err)
	}

	bot.Debug = true
	log.Printf("Authorized on account %s", bot.Self.UserName)

	_cookie1 := os.Getenv("COOKIE1")
	_cookie2 := os.Getenv("COOKIE2")
	_cookie3 := os.Getenv("COOKIE3")

	// 1. 创建字符串数组
	cookies := []string{_cookie1, _cookie2, _cookie3}

	fmt.Printf("cookies: %s\n", cookies)

	_cookie := RandomCookiesString(cookies)

	cache := cache.NewMemoryCache()
	// 设置命令
	_, err = bot.Request(tgbotapi.NewSetMyCommands(
		tgbotapi.BotCommand{Command: "start", Description: "启动"},
		tgbotapi.BotCommand{Command: "hide", Description: "隐藏键盘"},
	))
	if err != nil {
		log.Printf("Error setting commands: %v", err)
	}

	u := tgbotapi.NewUpdate(0)
	u.Timeout = 60
	updates := bot.GetUpdatesChan(u)

	for update := range updates {
		if update.Message != nil {
			if update.Message.IsCommand() {
				switch {
				case strings.HasPrefix(update.Message.Command(), "startAutoDispatch"):
					subscribeBundleID := strings.ReplaceAll(update.Message.Command(), "startAutoDispatch", "")
					log.Println("subscribeBundleID :" + subscribeBundleID)
					log.Println(subscribeBundleID + "startAutoDispatch command")
					userPackageSubscriptionsRepo := repositories.NewUserPackageSubscriptionsRepository(db)
					subscribeBundlePackageID, _ := strconv.ParseInt(subscribeBundleID, 10, 64)

					userPackageSubscriptionsRepo.UpdateStatus(context.Background(), subscribeBundlePackageID, 1)
					msg := service.CLICK_BUNDLE_PACKAGE_ADDRESS_STATS(db, update.Message.Chat.ID)
					bot.Send(msg)
				case strings.HasPrefix(update.Message.Command(), "dispatchNow"):
					subscribeBundleID := strings.ReplaceAll(update.Message.Command(), "dispatchNow", "")
					log.Println("subscribeBundleID :" + subscribeBundleID)
					log.Println(subscribeBundleID + "dispatchNow command")

					//手工发能

					//trxfee
					userPackageSubscriptionsRepo := repositories.NewUserPackageSubscriptionsRepository(db)
					record, _ := userPackageSubscriptionsRepo.Query(context.Background(), subscribeBundleID)

					restTimes := record.Times - 1
					userPackageSubscriptionsRepo.UpdateTimes(context.Background(), record.Id, restTimes)

					//
					msg2 := service.CLICK_BUNDLE_PACKAGE_ADDRESS_STATS(db, update.Message.Chat.ID)
					bot.Send(msg2)

					msg := tgbotapi.NewMessage(update.Message.Chat.ID, "📢【✅ U盾成功发送一笔能量】\n\n"+
						"接收地址："+record.Address+"\n\n"+
						"剩余笔数："+strconv.FormatInt(restTimes, 10)+"\n\n")
					msg.ParseMode = "HTML"
					bot.Send(msg)

				case strings.HasPrefix(update.Message.Command(), "stopAutoDispatch"):
					subscribeBundleID := strings.ReplaceAll(update.Message.Command(), "stopAutoDispatch", "")
					log.Println("subscribeBundleID :" + subscribeBundleID)
					log.Println(subscribeBundleID + "stopAutoDispatch command")
					userPackageSubscriptionsRepo := repositories.NewUserPackageSubscriptionsRepository(db)

					subscribeBundlePackageID, _ := strconv.ParseInt(subscribeBundleID, 10, 64)

					userPackageSubscriptionsRepo.UpdateStatus(context.Background(), subscribeBundlePackageID, 2)
					msg := service.CLICK_BUNDLE_PACKAGE_ADDRESS_STATS(db, update.Message.Chat.ID)
					bot.Send(msg)

				case strings.HasPrefix(update.Message.Command(), "dispatchOthers"):
					subscribeBundleID := strings.ReplaceAll(update.Message.Command(), "dispatchOthers", "")
					log.Println("subscribeBundleID :" + subscribeBundleID)
					log.Println(subscribeBundleID + "dispatchOthers command")
					//userPackageSubscriptionsRepo := repositories.NewUserPackageSubscriptionsRepository(db)

					//subscribeBundlePackageID, _ := strconv.ParseInt(subscribeBundleID, 10, 64)
					//userPackageSubscriptionsRepo.UpdateStatus(context.Background(), subscribeBundlePackageID, 2)
					//msg := service.CLICK_BUNDLE_PACKAGE_ADDRESS_STATS(db, update.Message.Chat.ID)
					//bot.Send(msg)
					//

					service.DispatchOthers(subscribeBundleID, cache, bot, update.Message.Chat.ID, db)

				case update.Message.Command() == "start":
					log.Printf("1")

					//存用户
					userRepo := repositories.NewUserRepository(db)
					_, err := userRepo.GetByUserID(update.Message.Chat.ID)
					if err != nil {
						//增加用户
						var user domain.User
						user.Associates = strconv.FormatInt(update.Message.Chat.ID, 10)
						user.Username = update.Message.Chat.UserName
						err := userRepo.Create2(context.Background(), &user)
						if err != nil {
							return
						}
					}

					handleStartCommand(cache, bot, update.Message)
				case update.Message.Command() == "hide":
					log.Printf("2")
					handleHideCommand(cache, bot, update.Message)
				}
			} else {

				log.Printf("3")
				log.Printf("来自于自发的信息[%s] %s", update.Message.From.UserName, update.Message.Text)
				handleRegularMessage(cache, bot, update.Message, db, _cookie)
			}
		} else if update.CallbackQuery != nil {
			log.Printf("4")
			handleCallbackQuery(cache, bot, update.CallbackQuery, db)
		}
	}
}

// 处理 /start 命令 - 显示永久键盘
func handleStartCommand(cache cache.Cache, bot *tgbotapi.BotAPI, message *tgbotapi.Message) {
	// 创建永久性回复键盘
	keyboard := tgbotapi.NewReplyKeyboard(
		tgbotapi.NewKeyboardButtonRow(
			tgbotapi.NewKeyboardButton("⚡能量闪兑"),
			tgbotapi.NewKeyboardButton("🖊️笔数套餐"),
		),
		tgbotapi.NewKeyboardButtonRow(
			tgbotapi.NewKeyboardButton("🔍地址检测"),
			tgbotapi.NewKeyboardButton("🚨USDT冻结预警"),
		),
		tgbotapi.NewKeyboardButtonRow(
			tgbotapi.NewKeyboardButton("👤个人中心"),
		),
	)

	// 关键设置：确保键盘一直存在
	keyboard.OneTimeKeyboard = false
	keyboard.ResizeKeyboard = true
	keyboard.Selective = false

	msg := tgbotapi.NewMessage(message.Chat.ID, "🛡️U盾，做您链上资产的护盾！\n我们不仅关注低价能量，更专注于交易安全！\n让每一笔转账都更安心，让每一次链上交互都值得信任！\n🤖 三大实用功能，助您安全、高效地管理链上资产\n🔋 波场能量闪兑, 节省超过80%!\n🕵️ 地址风险检测, 让每一笔转账都更安心!\n🚨 USDT冻结预警,秒级响应，让您的U永不冻结！\n🎉新用户福利：每日一次免费地址风险查询")
	msg.ReplyMarkup = keyboard
	msg.ParseMode = "HTML"
	bot.Send(msg)
}

// 处理 /hide 命令 - 隐藏键盘
func handleHideCommand(cache cache.Cache, bot *tgbotapi.BotAPI, message *tgbotapi.Message) {
	hideKeyboard := tgbotapi.NewRemoveKeyboard(true)
	msg := tgbotapi.NewMessage(message.Chat.ID, "键盘已隐藏，发送 /start 重新显示")
	msg.ReplyMarkup = hideKeyboard
	bot.Send(msg)
}

// 处理普通消息（键盘按钮点击）
func handleRegularMessage(cache cache.Cache, bot *tgbotapi.BotAPI, message *tgbotapi.Message, db *gorm.DB, _cookie string) {
	switch message.Text {
	case "🔍地址检测":
		service.MenuNavigateAddressDetection(cache, bot, message.Chat.ID, db)
	case "🚨USDT冻结预警":
		service.MenuNavigateAddressFreeze(cache, bot, message.Chat.ID, db)
	case "🖊️笔数套餐":
		service.MenuNavigateBundlePackage(db, message.Chat.ID, bot, "TRX")
	case "⚡能量闪兑":
		service.MenuNavigateEnergyExchange(db, message, bot)
	case "👤个人中心":
		service.MenuNavigateHome(db, message, bot)
	default:
		status, _ := cache.Get(strconv.FormatInt(message.Chat.ID, 10))

		log.Printf("用户状态staus %s", status)
		switch {
		case strings.HasPrefix(status, "user_backup_notify"):

			if service.ExtractBackup(message, bot, db) {
				return
			}
		case strings.HasPrefix(status, "start_freeze_risk"):

			if !IsValidAddress(message.Text) && !IsValidEthereumAddress(message.Text) {
				msg := tgbotapi.NewMessage(message.Chat.ID, "💬"+"<b>"+"地址有误，请重新输入地址: "+"</b>"+"\n")
				msg.ParseMode = "HTML"
				bot.Send(msg)
				return
			}

			dictRepo := repositories.NewSysDictionariesRepo(db)

			server_trx_price, _ := dictRepo.GetDictionaryDetail("server_trx_price")

			server_usdt_price, _ := dictRepo.GetDictionaryDetail("server_usdt_price")
			msg := tgbotapi.NewMessage(message.Chat.ID, "为以下地址开启冻结预警： "+"\n"+"地址："+message.Text+"\n\n"+
				"🎯 服务开启后U盾将 24 小时不间断保护您的资产安全。\n"+
				"⏰ 系统将在冻结前启动预警机制，持续 10 分钟每分钟推送提醒，通知您及时转移资产。\n"+
				"📌 服务费用："+server_trx_price+" TRX / 30 天 或 "+server_usdt_price+" USDT / 30 天\n是否确认开启该服务？")
			msg.ParseMode = "HTML"
			inlineKeyboard := tgbotapi.NewInlineKeyboardMarkup(
				tgbotapi.NewInlineKeyboardRow(
					tgbotapi.NewInlineKeyboardButtonData("✅ 确认开启", "confirm_freeze_risk_"+message.Text),
					tgbotapi.NewInlineKeyboardButtonData("❌ 取消操作", "back_risk_home"),
				),
				tgbotapi.NewInlineKeyboardRow(
					tgbotapi.NewInlineKeyboardButtonData("🔙️返回首页", "back_risk_home"),
				),
			)
			msg.ReplyMarkup = inlineKeyboard
			bot.Send(msg)
			expiration := 1 * time.Minute // 短时间缓存空值
			//设置用户状态
			cache.Set(strconv.FormatInt(message.Chat.ID, 10), "start_freeze_risk_status", expiration)

		case strings.HasPrefix(status, "address_list_trace"):

		case strings.HasPrefix(status, "address_manager_remove"):
			if IsValidAddress(message.Text) || IsValidEthereumAddress(message.Text) {
				userRepo := repositories.NewUserAddressMonitorRepo(db)
				err := userRepo.Remove(context.Background(), message.Chat.ID, message.Text)
				if err != nil {
				}
				msg := tgbotapi.NewMessage(message.Chat.ID, "✅ "+"<b>"+"地址删除成功 "+"</b>"+"\n")
				msg.ParseMode = "HTML"
				bot.Send(msg)

				service.ADDRESS_MANAGER(cache, bot, message.Chat.ID, db)

			} else {
				msg := tgbotapi.NewMessage(message.Chat.ID, "💬"+"<b>"+"地址有误，请重新输入需删除的地址: "+"</b>"+"\n")
				msg.ParseMode = "HTML"
				bot.Send(msg)
			}

		case strings.HasPrefix(status, "address_manager_add"):
			service.ExtractAddressManager(message, db, bot)

			service.ADDRESS_MANAGER(cache, bot, message.Chat.ID, db)

		case strings.HasPrefix(status, "bundle_"):
			fmt.Printf(">>>>>>>>>>>>>>>>>>>>bundle: %s", status)

			if service.ExtractBundleService(message, bot, db, status) {
				return
			}

		case strings.HasPrefix(status, "usdt_risk_monitor"):
			//fmt.Printf("bundle: %s", status)

			if !IsValidAddress(message.Text) {
				msg := tgbotapi.NewMessage(message.Chat.ID, "💬"+"<b>"+"地址有误，请重新输入地址: "+"</b>"+"\n")
				msg.ParseMode = "HTML"
				bot.Send(msg)
			}

			msg := tgbotapi.NewMessage(message.Chat.ID, "")

			//msg.ReplyMarkup = inlineKeyboard
			msg.ParseMode = "HTML"

			bot.Send(msg)

		case strings.HasPrefix(status, "click_bundle_package_address_manager_remove"):
			if service.CLICK_BUNDLE_PACKAGE_ADDRESS_MANAGER_REMOVE(cache, bot, message, db) {
				return
			}

		case strings.HasPrefix(status, "click_bundle_package_address_manager_add"):
			if service.CLICK_BUNDLE_PACKAGE_ADDRESS_MANAGER_ADD(cache, bot, message, db) {
				return
			}

		case strings.HasPrefix(status, "apply_bundle_package_"):
			if service.APPLY_BUNDLE_PACKAGE(cache, bot, message, db, status) {
				return
			}

		case strings.HasPrefix(status, "click_backup_account"):

			if !strings.Contains(message.Text, "@") {
				msg := tgbotapi.NewMessage(message.Chat.ID, "❌ 用户名格式有误，请重新输入")
				msg.ParseMode = "HTML"
				bot.Send(msg)
				return
			}
			userName := strings.ReplaceAll(message.Text, "@", "")

			userRepo := repositories.NewUserRepository(db)
			user, err := userRepo.GetByUsername(userName)

			if err != nil {
				msg := tgbotapi.NewMessage(message.Chat.ID, "❌ 用户名格式有误，请重新输入")
				msg.ParseMode = "HTML"
				bot.Send(msg)
				return
			}

			if user.Id == 0 {
				msg := tgbotapi.NewMessage(message.Chat.ID, "❌ 用户名格式有误，请重新输入")
				msg.ParseMode = "HTML"
				bot.Send(msg)
				return
			}

			user.BackupChatID = userName

			err2 := userRepo.UpdateBackupChat(context.Background(), userName, message.Chat.ID)
			if err2 == nil {
				msg := tgbotapi.NewMessage(message.Chat.ID, "✅ 成功绑定第二紧急联系人: "+message.Text)
				msg.ParseMode = "HTML"
				bot.Send(msg)
				//return true
			}

			service.BackHOME(db, message.Chat.ID, bot)

		case strings.HasPrefix(status, "usdt_risk_query"):
			//fmt.Printf("bundle: %s", status)
			service.ExtractSlowMistRiskQuery(message, db, _cookie, bot)
		}
	}
}

// 处理内联键盘回调
func handleCallbackQuery(cache cache.Cache, bot *tgbotapi.BotAPI, callbackQuery *tgbotapi.CallbackQuery, db *gorm.DB) {
	// 先应答回调
	callback := tgbotapi.NewCallback(callbackQuery.ID, "已选择: "+callbackQuery.Data)
	if _, err := bot.Request(callback); err != nil {
		log.Printf("Error answering callback: %v", err)
	}

	// 根据回调数据执行不同操作
	var responseText string
	switch {

	case callbackQuery.Data == "back_address_detection_home":

		service.MenuNavigateAddressDetection(cache, bot, callbackQuery.Message.Chat.ID, db)

	case strings.HasPrefix(callbackQuery.Data, "dispatch_others_"):
		bundleAddress := strings.ReplaceAll(callbackQuery.Data, "dispatch_others_", "")

		bundleID := strings.Split(bundleAddress, "_")[0]
		address := strings.Split(bundleAddress, "_")[1]

		fmt.Printf("bundleID %s\n", bundleID)
		fmt.Printf("address %s\n", address)

		//手工发能

		//trxfee
		userPackageSubscriptionsRepo := repositories.NewUserPackageSubscriptionsRepository(db)
		record, _ := userPackageSubscriptionsRepo.Query(context.Background(), bundleID)

		restTimes := record.Times - 1
		userPackageSubscriptionsRepo.UpdateTimes(context.Background(), record.Id, restTimes)

		//
		msg2 := service.CLICK_BUNDLE_PACKAGE_ADDRESS_STATS(db, callbackQuery.Message.Chat.ID)
		bot.Send(msg2)

		msg := tgbotapi.NewMessage(callbackQuery.Message.Chat.ID, "📢【✅ U盾成功发送一笔能量】\n\n"+
			"接收地址："+record.Address+"\n\n"+
			"剩余笔数："+strconv.FormatInt(restTimes, 10)+"\n\n")
		msg.ParseMode = "HTML"
		bot.Send(msg)

	case strings.HasPrefix(callbackQuery.Data, "confirm_freeze_risk_"):
		address := strings.ReplaceAll(callbackQuery.Data, "confirm_freeze_risk_", "")

		fmt.Printf("address : %s\n", address)
		sysDictionariesRepo := repositories.NewSysDictionariesRepo(db)
		server_trx_price, _ := sysDictionariesRepo.GetDictionaryDetail("server_trx_price")
		server_usdt_price, _ := sysDictionariesRepo.GetDictionaryDetail("server_usdt_price")
		userRepo := repositories.NewUserRepository(db)
		user, _ := userRepo.GetByUserID(callbackQuery.Message.Chat.ID)
		if !CompareStringsWithFloat(user.TronAmount, server_trx_price, 1) && !CompareStringsWithFloat(user.Amount, server_usdt_price, 1) {
			msg := tgbotapi.NewMessage(callbackQuery.Message.Chat.ID, "⚠️ 当前余额不足，无法开启冻结预警服务\n\n")
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
		fmt.Println("余额充足")
		var COST_FROM_TRX bool
		var COST_FROM_USDT bool
		if CompareStringsWithFloat(user.TronAmount, server_trx_price, 1) || CompareStringsWithFloat(user.Amount, server_usdt_price, 1) {

			if CompareStringsWithFloat(user.TronAmount, server_trx_price, float64(1)) {
				rest, _ := SubtractStringNumbers(user.TronAmount, server_trx_price, float64(1))

				user.TronAmount = rest
				userRepo.Update2(context.Background(), &user)
				fmt.Printf("rest: %s", rest)
				COST_FROM_TRX = true
				//扣usdt
			} else if CompareStringsWithFloat(user.Amount, server_usdt_price, float64(1)) {
				rest, _ := SubtractStringNumbers(user.Amount, server_usdt_price, float64(1))
				fmt.Printf("rest: %s", rest)
				user.Amount = rest
				userRepo.Update2(context.Background(), &user)
				COST_FROM_USDT = true
			}

			//添加记录
			userAddressEventRepo := repositories.NewUserAddressMonitorEventRepo(db)

			var event domain.UserAddressMonitorEvent
			event.ChatID = callbackQuery.Message.Chat.ID
			event.Status = 1
			event.Address = address

			if len(address) == 42 {
				event.Network = "Ethereum"
			}
			if len(address) == 34 {
				event.Network = "Tron"
			}

			event.Days = 1
			if COST_FROM_TRX {
				event.Amount = server_trx_price + " TRX"
			}
			if COST_FROM_USDT {
				event.Amount = server_usdt_price + " USDT"
			}
			userAddressEventRepo.Create(context.Background(), &event)

			//后台跟踪起来
			//user, _ := userRepo.GetByUserID(callbackQuery.Message.Chat.ID)
			msg := tgbotapi.NewMessage(callbackQuery.Message.Chat.ID,
				"✅"+"地址开启冻结预警监测成功：\n"+
					"地址："+address+"\n"+
					"网络："+event.Network)
			msg.ParseMode = "HTML"
			inlineKeyboard := tgbotapi.NewInlineKeyboardMarkup(
				tgbotapi.NewInlineKeyboardRow(
					tgbotapi.NewInlineKeyboardButtonData("预警监控列表", "address_list_trace"),
					tgbotapi.NewInlineKeyboardButtonData("🔙️返回首页", "back_risk_home"),
				),
			)
			msg.ReplyMarkup = inlineKeyboard
			bot.Send(msg)

		}

	case strings.HasPrefix(callbackQuery.Data, "set_bundle_package_default_"):
		target := strings.ReplaceAll(callbackQuery.Data, "set_bundle_package_default_", "")
		userOperationPackageAddressesRepo := repositories.NewUserOperationPackageAddressesRepo(db)

		errsg := userOperationPackageAddressesRepo.Update(context.Background(), callbackQuery.Message.Chat.ID, target)
		if errsg != nil {
			log.Printf("errsg: %s", errsg)
			return
		}
		msg := tgbotapi.NewMessage(callbackQuery.Message.Chat.ID, "✅"+"<b>"+"设置默认地址成功 "+"</b>"+"\n")
		msg.ParseMode = "HTML"
		bot.Send(msg)
		service.CLICK_BUNDLE_PACKAGE_ADDRESS_MANAGEMENT(cache, bot, callbackQuery.Message.Chat.ID, db)

	case strings.HasPrefix(callbackQuery.Data, "remove_bundle_package_"):
		target := strings.ReplaceAll(callbackQuery.Data, "remove_bundle_package_", "")
		userOperationPackageAddressesRepo := repositories.NewUserOperationPackageAddressesRepo(db)

		var record domain.UserOperationPackageAddresses
		record.Status = 0
		record.Address = target
		record.ChatID = callbackQuery.Message.Chat.ID

		errsg := userOperationPackageAddressesRepo.Remove(context.Background(), callbackQuery.Message.Chat.ID, target)
		if errsg != nil {
			log.Printf("errsg: %s", errsg)
			return
		}
		msg := tgbotapi.NewMessage(callbackQuery.Message.Chat.ID, "✅"+"<b>"+"地址删除成功 "+"</b>"+"\n")
		msg.ParseMode = "HTML"
		bot.Send(msg)
		service.CLICK_BUNDLE_PACKAGE_ADDRESS_MANAGEMENT(cache, bot, callbackQuery.Message.Chat.ID, db)

	case strings.HasPrefix(callbackQuery.Data, "close_freeze_risk_"):
		target := strings.ReplaceAll(callbackQuery.Data, "close_freeze_risk_", "")
		//⚠️ 确认停止监控以下地址？
		//
		//地址：TX8kY...5a9rP
		//
		//当前剩余天数：12 天
		//
		//停止监控后将立即终止监控，服务时间不予退还
		//
		//🔒 为避免误操作，请再次确认：
		//
		//✅ 确认解绑
		//
		//❌ 取消操作
		//
		//✅ 地址监控已停止

		log.Println("target:", target)
		userAddressEventRepo := repositories.NewUserAddressMonitorEventRepo(db)
		event, _ := userAddressEventRepo.Find(context.Background(), target)

		restDays := fmt.Sprintf("%d", 30-event.Days)

		msg := tgbotapi.NewMessage(callbackQuery.Message.Chat.ID, "️ ⚠️ 确认停止监控以下地址？"+"\n"+
			"地址："+event.Address+"\n"+
			"当前剩余天数："+restDays+" 天\n"+
			"停止监控后将立即终止监控，服务时间不予退还\n"+"🔒 为避免误操作，请再次确认：")
		msg.ParseMode = "HTML"
		// 当点击"按钮 1"时显示内联键盘
		inlineKeyboard := tgbotapi.NewInlineKeyboardMarkup(
			tgbotapi.NewInlineKeyboardRow(
				tgbotapi.NewInlineKeyboardButtonData("✅ 确认停止", "close_risk_"+target),
				tgbotapi.NewInlineKeyboardButtonData("❌ 取消操作", "back_risk_home"),
			),
			tgbotapi.NewInlineKeyboardRow(
				tgbotapi.NewInlineKeyboardButtonData("🔙️返回首页", "back_risk_home"),
			),
		)
		msg.ReplyMarkup = inlineKeyboard

		bot.Send(msg)

	case strings.HasPrefix(callbackQuery.Data, "close_risk_"):
		target := strings.ReplaceAll(callbackQuery.Data, "close_risk_", "")
		log.Println("target:", target)
		userAddressEventRepo := repositories.NewUserAddressMonitorEventRepo(db)
		err := userAddressEventRepo.Close(context.Background(), target)
		if err == nil {
			msg := tgbotapi.NewMessage(callbackQuery.Message.Chat.ID, "️ ✅ 地址监控已停止")
			msg.ParseMode = "HTML"
			// 当点击"按钮 1"时显示内联键盘
			inlineKeyboard := tgbotapi.NewInlineKeyboardMarkup(
				tgbotapi.NewInlineKeyboardRow(
					tgbotapi.NewInlineKeyboardButtonData("预警监控列表", "address_list_trace"),
					tgbotapi.NewInlineKeyboardButtonData("🔙️返回首页", "back_risk_home"),
				),
			)
			msg.ReplyMarkup = inlineKeyboard

			bot.Send(msg)
		}
	case strings.HasPrefix(callbackQuery.Data, "apply_bundle_package_"):

		target := strings.ReplaceAll(callbackQuery.Data, "apply_bundle_package_", "")
		service.APPLY_BUNDLE_PACKAGE_ADDRESS(target, cache, bot, callbackQuery.Message, db)

	case strings.HasPrefix(callbackQuery.Data, "config_bundle_package_address_"):

		target := strings.ReplaceAll(callbackQuery.Data, "config_bundle_package_address_", "")
		service.CONFIG_BUNDLE_PACKAGE_ADDRESS(target, cache, bot, callbackQuery.Message, db)
	case callbackQuery.Data == "click_backup_account":

		msg := tgbotapi.NewMessage(callbackQuery.Message.Chat.ID, "👥欢迎使用第二通知人服务"+"\n"+
			"为确保实时接收预警信息，您可绑定一个第二通知人TG帐号。"+"\n"+
			"绑定前请确保第二通知人已与本机器人互动，绑定后该账号将同步接收预警信息，第二通知人替换请重复绑定步骤，系统将自动替换。请输入的第二通知人TG帐号@用户名 👇")
		msg.ParseMode = "HTML"

		inlineKeyboard := tgbotapi.NewInlineKeyboardMarkup(
			tgbotapi.NewInlineKeyboardRow(
				tgbotapi.NewInlineKeyboardButtonData("返回个人中心", "back_home"),
				//tgbotapi.NewInlineKeyboardButtonData("第二紧急通知", ""),
			),
		)
		msg.ReplyMarkup = inlineKeyboard

		bot.Send(msg)

		expiration := 1 * time.Minute // 短时间缓存空值

		//设置用户状态
		cache.Set(strconv.FormatInt(callbackQuery.Message.Chat.ID, 10), "click_backup_account", expiration)

	case callbackQuery.Data == "back_risk_home":
		service.MenuNavigateAddressFreeze(cache, bot, callbackQuery.Message.Chat.ID, db)
	case callbackQuery.Data == "click_switch_trx":
		service.MenuNavigateBundlePackage(db, callbackQuery.Message.Chat.ID, bot, "TRX")
	case callbackQuery.Data == "click_switch_usdt":
		service.MenuNavigateBundlePackage(db, callbackQuery.Message.Chat.ID, bot, "USDT")
	case callbackQuery.Data == "back_bundle_package":
		service.MenuNavigateBundlePackage(db, callbackQuery.Message.Chat.ID, bot, "TRX")
	case callbackQuery.Data == "click_bundle_package_address_manager_config":
		service.CLICK_BUNDLE_PACKAGE_ADDRESS_MANAGER_CONFIG(cache, bot, callbackQuery.Message.Chat.ID, db)
	case callbackQuery.Data == "click_bundle_package_address_manager_remove":
		msg := tgbotapi.NewMessage(callbackQuery.Message.Chat.ID, "💬"+"<b>"+"请输入需要删除的地址: "+"</b>"+"\n")
		msg.ParseMode = "HTML"
		bot.Send(msg)

		expiration := 1 * time.Minute // 短时间缓存空值

		//设置用户状态
		cache.Set(strconv.FormatInt(callbackQuery.Message.Chat.ID, 10), callbackQuery.Data, expiration)

	case callbackQuery.Data == "click_bundle_package_address_manager_add":

		userOperationPackageAddressesRepo := repositories.NewUserOperationPackageAddressesRepo(db)

		list, _ := userOperationPackageAddressesRepo.Query(context.Background(), callbackQuery.Message.Chat.ID)
		if len(list) >= 4 {
			msg := tgbotapi.NewMessage(callbackQuery.Message.Chat.ID, "<b>"+"❌ 添加新地址失败，地址已达上限，请先删除一个旧地址 。"+"</b>"+"\n")
			msg.ParseMode = "HTML"
			bot.Send(msg)
			return
		}
		msg := tgbotapi.NewMessage(callbackQuery.Message.Chat.ID, "<b>"+"为方便用户管理地址，系统默认最多添加4个地址，请输入新地址👇: "+"</b>"+"\n")
		msg.ParseMode = "HTML"
		bot.Send(msg)

		expiration := 1 * time.Minute // 短时间缓存空值

		//设置用户状态
		cache.Set(strconv.FormatInt(callbackQuery.Message.Chat.ID, 10), callbackQuery.Data, expiration)
		//笔数套餐地址列表
	case callbackQuery.Data == "click_bundle_package_address_stats":
		msg := service.CLICK_BUNDLE_PACKAGE_ADDRESS_STATS(db, callbackQuery.Message.Chat.ID)
		bot.Send(msg)

	case callbackQuery.Data == "next_bundle_package_address_stats":
		if service.NEXT_BUNDLE_PACKAGE_ADDRESS_STATS(callbackQuery, db, bot) {
			return
		}
	case callbackQuery.Data == "prev_bundle_package_address_stats":
		state, done := service.PREV_BUNDLE_PACKAGE_ADDRESS_STATS(callbackQuery, db, bot)
		if done {
			return
		}
		fmt.Printf("state: %v\n", state)

	case callbackQuery.Data == "click_bundle_package_address_management":
		service.CLICK_BUNDLE_PACKAGE_ADDRESS_MANAGEMENT(cache, bot, callbackQuery.Message.Chat.ID, db)
	case callbackQuery.Data == "address_list_trace":
		service.ADDRESS_LIST_TRACE(cache, bot, callbackQuery, db)
	case callbackQuery.Data == "back_home":
		service.BackHOME(db, callbackQuery.Message.Chat.ID, bot)
	case callbackQuery.Data == "click_business_cooperation":
		service.ClickBusinessCooperation(callbackQuery, bot)
	case callbackQuery.Data == "click_callcenter":
		service.ClickCallCenter(callbackQuery, bot)
	case callbackQuery.Data == "click_my_recepit":
		service.CLICK_MY_RECEPIT(db, callbackQuery, bot)
	case callbackQuery.Data == "address_freeze_risk_records":
		msg := service.ExtractAddressRiskQuery(db, callbackQuery)
		bot.Send(msg)
	case callbackQuery.Data == "user_detection_cost_records":
		msg := service.ExtractAddressDetection(db, callbackQuery)
		bot.Send(msg)
	case callbackQuery.Data == "click_bundle_package_cost_records":
		msg := service.ExtractBundlePackage(db, callbackQuery)
		bot.Send(msg)
	case callbackQuery.Data == "click_bundle_package_management":
		msg := service.ExtractBundlePackage(db, callbackQuery)
		bot.Send(msg)
	case callbackQuery.Data == "click_deposit_usdt_records":
		service.CLICK_DEPOSIT_USDT_RECORDS(db, callbackQuery, bot)
	case callbackQuery.Data == "click_deposit_trx_records":
		service.CLICK_DEPOSIT_TRX_RECORDS(db, callbackQuery, bot)
	case callbackQuery.Data == "next_address_detection_page":
		if service.EXTRACT_NEXT_ADDRESS_DETECTION_PAGE(callbackQuery, db, bot) {
			return
		}
	case callbackQuery.Data == "prev_address_detection_page":
		state, done := service.EXTRACT_PREV_ADDRESS_DETECTION_PAGE(callbackQuery, db, bot)
		if done {
			return
		}
		fmt.Printf("state: %v\n", state)
	case callbackQuery.Data == "prev_deposit_usdt_page":
		state, done := service.EXTRACT_PREV_DEPOSIT_USDT_PAGE(callbackQuery, db, bot)
		if done {
			return
		}
		fmt.Printf("state: %v\n", state)
	case callbackQuery.Data == "prev_deposit_trx_page":
		state, done := service.EXTRACT_PREV_DEPOSIT_TRX_PAGE(callbackQuery, db, bot)
		if done {
			return
		}
		fmt.Printf("state: %v\n", state)
	case callbackQuery.Data == "prev_address_risk_page":
		state, done := service.EXTRACT_PREV_ADDRESS_RISK_PAGE(callbackQuery, db, bot)
		if done {
			return
		}
		fmt.Printf("state: %v\n", state)

	case callbackQuery.Data == "next_address_risk_page":
		if service.ExtraNextAddressRiskPage(callbackQuery, db, bot) {
			return
		}
	case callbackQuery.Data == "next_deposit_usdt_page":
		if service.ExtraNextDepositUSDTPage(callbackQuery, db, bot) {
			return
		}
	case callbackQuery.Data == "next_deposit_trx_page":
		if service.ExtracNextDepositTrxPage(callbackQuery, db, bot) {
			return
		}

	case callbackQuery.Data == "prev_bundle_package_page":
		state, done := service.EXTRACT_PREV_BUNDLE_PACKAGE_PAGE(callbackQuery, db, bot)
		if done {
			return
		}
		fmt.Printf("state: %v\n", state)

	case callbackQuery.Data == "next_bundle_package_page":
		if service.EXTRACT_NEXT_BUNDLE_PACKAGE_PAGE(callbackQuery, db, bot) {
			return
		}

	case callbackQuery.Data == "click_QA":
		service.ExtraQA(cache, bot, callbackQuery)

	case callbackQuery.Data == "user_backup_notify":
		msg := tgbotapi.NewMessage(callbackQuery.Message.Chat.ID, "💬"+"<b>"+"请输入需添加的第二紧急通知用户电报ID: "+"</b>"+"\n")
		msg.ParseMode = "HTML"
		bot.Send(msg)

		expiration := 1 * time.Minute // 短时间缓存空值

		//设置用户状态
		cache.Set(strconv.FormatInt(callbackQuery.Message.Chat.ID, 10), callbackQuery.Data, expiration)
	case callbackQuery.Data == "start_freeze_risk_1":
		//查看余额
		service.START_FREEZE_RISK_1(cache, db, callbackQuery, bot)

	case callbackQuery.Data == "click_my_service":
		msg := tgbotapi.NewMessage(callbackQuery.Message.Chat.ID, "🛡 当前服务状态：\n\n🔋 能量闪兑\n\n- 剩余笔数：12\n- 自动补能：关闭 /开启\n\n➡️ /闪兑\n\n➡️ /笔数套餐\n\n➡️ /手动发能（1笔）\n\n➡️ /开启/关闭自动发能\n\n📍 地址风险检测\n\n- 今日免费次数：已用完\n\n➡️ /地址风险检测\n\n🚨 USDT冻结预警\n\n- 地址1：TX8kY...5a9rP（剩余12天）✅\n- 地址2：TEw9Q...iS6Ht（剩余28天）✅")
		msg.ParseMode = "HTML"

		inlineKeyboard := tgbotapi.NewInlineKeyboardMarkup(
			tgbotapi.NewInlineKeyboardRow(
				tgbotapi.NewInlineKeyboardButtonData("预警监控列表", "address_list_trace"),
				//	tgbotapi.NewInlineKeyboardButtonData("地址管理", "address_manager"),
			),
		)
		msg.ReplyMarkup = inlineKeyboard

		bot.Send(msg)

		expiration := 1 * time.Minute // 短时间缓存空值

		//设置用户状态
		cache.Set(strconv.FormatInt(callbackQuery.Message.Chat.ID, 10), "usdt_risk_monitor", expiration)

	case callbackQuery.Data == "stop_freeze_risk_1":

		//删除event表里面
		userAddressEventRepo := repositories.NewUserAddressMonitorEventRepo(db)

		userAddressEventRepo.RemoveAll(context.Background(), callbackQuery.Message.Chat.ID)

		msg := tgbotapi.NewMessage(callbackQuery.Message.Chat.ID, "已经暂停所有监控")
		msg.ParseMode = "HTML"

		bot.Send(msg)

		expiration := 1 * time.Minute // 短时间缓存空值

		//设置用户状态
		cache.Set(strconv.FormatInt(callbackQuery.Message.Chat.ID, 10), "reset", expiration)

	case callbackQuery.Data == "start_freeze_risk_0":

		sysDictionariesRepo := repositories.NewSysDictionariesRepo(db)

		server_trx_price, _ := sysDictionariesRepo.GetDictionaryDetail("server_trx_price")

		server_usdt_price, _ := sysDictionariesRepo.GetDictionaryDetail("server_usdt_price")

		msg := tgbotapi.NewMessage(callbackQuery.Message.Chat.ID, "欢迎使用U盾 USDT冻结预警服务\n"+
			"🛡️ U盾，做您链上资产的护盾！\n"+
			"地址一旦被链上风控冻，资产将难以追回，损失巨大！\n"+
			"每天都有数百个 USDT 钱包地址被冻结锁定，风险就在身边！\n"+
			"✅ 适用于经常收付款 / 被制裁地址感染/与诈骗地址交互\n"+
			"✅ 支持TRON/ETH网络的USDT 钱包地址\n"+
			"📌 服务价格（每地址）：\n • "+server_trx_price+" TRX / 30天\n • "+
			" 或 "+server_usdt_price+" USDT / 30天\n"+
			"🎯 服务开启后U盾将24 小时不间断保护您的资产安全。\n"+
			"⏰ 系统将在冻结前启动预警机制，持续 10 分钟每分钟推送提醒，通知您及时转移资产。\n"+
			"📩 所有预警信息将通过 Telegram 实时推送")
		msg.ParseMode = "HTML"

		inlineKeyboard := tgbotapi.NewInlineKeyboardMarkup(
			tgbotapi.NewInlineKeyboardRow(
				tgbotapi.NewInlineKeyboardButtonData("开启冻结预警", "start_freeze_risk"),
				//tgbotapi.NewInlineKeyboardButtonData("地址管理", "address_manager"),
			),
			tgbotapi.NewInlineKeyboardRow(
				tgbotapi.NewInlineKeyboardButtonData("预警监控列表", "address_list_trace"),
				tgbotapi.NewInlineKeyboardButtonData("冻结预警扣款记录", "address_freeze_risk_records"),
			),
			//tgbotapi.NewInlineKeyboardRow(
			//	tgbotapi.NewInlineKeyboardButtonData("冻结预警扣款记录", "address_freeze_risk_records"),
			//	//tgbotapi.NewInlineKeyboardButtonData("第二紧急通知", ""),
			//),
		)
		msg.ReplyMarkup = inlineKeyboard

		bot.Send(msg)

		expiration := 1 * time.Minute // 短时间缓存空值

		//设置用户状态
		cache.Set(strconv.FormatInt(callbackQuery.Message.Chat.ID, 10), "usdt_risk_monitor", expiration)
	case callbackQuery.Data == "stop_freeze_risk":

		userAddressEventRepo := repositories.NewUserAddressMonitorEventRepo(db)
		addresses, _ := userAddressEventRepo.Query(context.Background(), callbackQuery.Message.Chat.ID)

		//msg.ParseMode = "HTML"

		var allButtons []tgbotapi.InlineKeyboardButton
		var extraButtons []tgbotapi.InlineKeyboardButton
		var keyboard [][]tgbotapi.InlineKeyboardButton
		for _, item := range addresses {
			allButtons = append(allButtons, tgbotapi.NewInlineKeyboardButtonData(item.Address, "close_freeze_risk_"+fmt.Sprintf("%d", item.Id)))
		}

		extraButtons = append(extraButtons, tgbotapi.NewInlineKeyboardButtonData("🔙返回首页", "back_risk_home"))

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

		//msg.ReplyMarkup = inlineKeyboard
		//
		//bot.Send(msg)
		//
		//expiration := 1 * time.Minute // 短时间缓存空值
		//
		////设置用户状态
		//cache.Set(strconv.FormatInt(_chatID, 10), "start_freeze_risk", expiration)
		//
		//msg := tgbotapi.NewMessage(callbackQuery.Message.Chat.ID, "📡 是否确认停止该服务？")
		//msg.ParseMode = "HTML"
		//
		//inlineKeyboard := tgbotapi.NewInlineKeyboardMarkup(
		//	tgbotapi.NewInlineKeyboardRow(
		//		tgbotapi.NewInlineKeyboardButtonData("✅ 确认停止", "stop_freeze_risk_1"),
		//		tgbotapi.NewInlineKeyboardButtonData("❌ 取消操作", "start_freeze_risk_0"),
		//	),
		//tgbotapi.NewInlineKeyboardRow(
		//	tgbotapi.NewInlineKeyboardButtonData("地址", ""),
		//),
		//)
		msg := tgbotapi.NewMessage(callbackQuery.Message.Chat.ID, "预警地址列表如下："+"\n\n")
		//地址绑定

		msg.ParseMode = "HTML"

		msg.ReplyMarkup = inlineKeyboard

		bot.Send(msg)

		//expiration := 1 * time.Minute // 短时间缓存空值

		//设置用户状态
		//cache.Set(strconv.FormatInt(callbackQuery.Message.Chat.ID, 10), "stop_freeze_risk", expiration)

	case callbackQuery.Data == "start_freeze_risk":

		sysDictionariesRepo := repositories.NewSysDictionariesRepo(db)
		server_trx_price, _ := sysDictionariesRepo.GetDictionaryDetail("server_trx_price")
		server_usdt_price, _ := sysDictionariesRepo.GetDictionaryDetail("server_usdt_price")
		userRepo := repositories.NewUserRepository(db)
		user, _ := userRepo.GetByUserID(callbackQuery.Message.Chat.ID)
		if !CompareStringsWithFloat(user.TronAmount, server_trx_price, 1) && !CompareStringsWithFloat(user.Amount, server_usdt_price, 1) {
			msg := tgbotapi.NewMessage(callbackQuery.Message.Chat.ID, "⚠️ 当前余额不足，无法开启冻结预警服务\n\n")
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

		//sysDictionariesRepo := repositories.NewSysDictionariesRepo(db)
		//
		//server_trx_price, _ := sysDictionariesRepo.GetDictionaryDetail("server_trx_price")
		//
		//server_usdt_price, _ := sysDictionariesRepo.GetDictionaryDetail("server_usdt_price")
		//
		//msg := tgbotapi.NewMessage(callbackQuery.Message.Chat.ID, "🎯 服务开启后U盾将 24 小时不间断保护您的资产安全。\n"+
		//	"⏰ 系统将在冻结前启动预警机制，持续 10 分钟每分钟推送提醒，通知您及时转移资产。\n"+
		//	"📌 服务价格（每地址）：\n • "+server_trx_price+" TRX / 30天\n • "+
		//	" 或 "+server_usdt_price+" USDT / 30天\n"+
		//	"是否确认开启该服务？")
		//msg.ParseMode = "HTML"
		//
		//inlineKeyboard := tgbotapi.NewInlineKeyboardMarkup(
		//	tgbotapi.NewInlineKeyboardRow(
		//		tgbotapi.NewInlineKeyboardButtonData("✅ 确认开启", "start_freeze_risk_1"),
		//		tgbotapi.NewInlineKeyboardButtonData("❌ 取消操作", "back_risk_home"),
		//	),
		//	tgbotapi.NewInlineKeyboardRow(
		//		tgbotapi.NewInlineKeyboardButtonData("🔙️返回首页", "back_risk_home"),
		//	),
		//)
		//msg.ReplyMarkup = inlineKeyboard
		//
		//bot.Send(msg)

		msg := tgbotapi.NewMessage(callbackQuery.Message.Chat.ID, "请输入要预警的地址 👇")
		msg.ParseMode = "HTML"
		bot.Send(msg)
		expiration := 1 * time.Minute // 短时间缓存空值

		//设置用户状态
		cache.Set(strconv.FormatInt(callbackQuery.Message.Chat.ID, 10), "start_freeze_risk", expiration)

	case callbackQuery.Data == "address_manager_return":

		sysDictionariesRepo := repositories.NewSysDictionariesRepo(db)

		server_trx_price, _ := sysDictionariesRepo.GetDictionaryDetail("server_trx_price")

		server_usdt_price, _ := sysDictionariesRepo.GetDictionaryDetail("server_usdt_price")

		msg := tgbotapi.NewMessage(callbackQuery.Message.Chat.ID, "欢迎使用U盾 USDT冻结预警服务\n"+
			"🛡️ U盾，做您链上资产的护盾！\n"+
			"地址一旦被链上风控冻，资产将难以追回，损失巨大！\n"+
			"每天都有数百个 USDT 钱包地址被冻结锁定，风险就在身边！\n"+
			"✅ 适用于经常收付款 / 被制裁地址感染/与诈骗地址交互\n"+
			"✅ 支持TRON/ETH网络的USDT 钱包地址\n"+
			"📌 服务价格（每地址）：\n • "+server_trx_price+" TRX / 30天\n • "+
			" 或 "+server_usdt_price+" USDT / 30天\n"+
			"🎯 服务开启后U盾将24 小时不间断保护您的资产安全。\n"+
			"⏰ 系统将在冻结前启动预警机制，持续 10 分钟每分钟推送提醒，通知您及时转移资产。\n"+
			"📩 所有预警信息将通过 Telegram 实时推送")
		msg.ParseMode = "HTML"

		inlineKeyboard := tgbotapi.NewInlineKeyboardMarkup(
			tgbotapi.NewInlineKeyboardRow(
				tgbotapi.NewInlineKeyboardButtonData("开启冻结预警", "start_freeze_risk"),
				//	tgbotapi.NewInlineKeyboardButtonData("地址管理", "address_manager"),
			),
			tgbotapi.NewInlineKeyboardRow(
				tgbotapi.NewInlineKeyboardButtonData("预警监控列表", "address_list_trace"),
				tgbotapi.NewInlineKeyboardButtonData("冻结预警扣款记录", "address_freeze_risk_records"),
			),
			//tgbotapi.NewInlineKeyboardRow(
			//	tgbotapi.NewInlineKeyboardButtonData("冻结预警扣款记录", "address_freeze_risk_records"),
			//	//tgbotapi.NewInlineKeyboardButtonData("第二紧急通知", ""),
			//),
		)
		msg.ReplyMarkup = inlineKeyboard

		bot.Send(msg)

		expiration := 1 * time.Minute // 短时间缓存空值

		//设置用户状态
		cache.Set(strconv.FormatInt(callbackQuery.Message.Chat.ID, 10), "usdt_risk_monitor", expiration)

	case callbackQuery.Data == "address_manager_add":
		msg := tgbotapi.NewMessage(callbackQuery.Message.Chat.ID, "💬"+"<b>"+"请输入需添加的地址: "+"</b>"+"\n")
		msg.ParseMode = "HTML"
		bot.Send(msg)

		expiration := 1 * time.Minute // 短时间缓存空值

		//设置用户状态
		cache.Set(strconv.FormatInt(callbackQuery.Message.Chat.ID, 10), callbackQuery.Data, expiration)
	case callbackQuery.Data == "address_manager_remove":
		msg := tgbotapi.NewMessage(callbackQuery.Message.Chat.ID, "💬"+"<b>"+"请输入需删除的地址: "+"</b>"+"\n")
		msg.ParseMode = "HTML"
		bot.Send(msg)

		expiration := 1 * time.Minute // 短时间缓存空值

		//设置用户状态
		cache.Set(strconv.FormatInt(callbackQuery.Message.Chat.ID, 10), callbackQuery.Data, expiration)
	case callbackQuery.Data == "address_manager":
		service.ADDRESS_MANAGER(cache, bot, callbackQuery.Message.Chat.ID, db)

	case callbackQuery.Data == "deposit_amount":

		service.DEPOSIT_AMOUNT(db, callbackQuery, bot)

	case strings.HasPrefix(callbackQuery.Data, "bundle_"):
		service.BUNDLE_CHECK(cache, bot, callbackQuery, db)
		//调用trxfee接口进行笔数扣款
	case strings.HasPrefix(callbackQuery.Data, "deposit_usdt"):
		service.DepositPrevUSDTOrder(cache, bot, callbackQuery, db)
		//responseText = "你选择了选项 A"
	case strings.HasPrefix(callbackQuery.Data, "deposit_trx"):
		service.DepositPrevOrder(cache, bot, callbackQuery, db)
	case callbackQuery.Data == "cancel_order":
		service.DepositCancelOrder(cache, bot, callbackQuery, db)
	case callbackQuery.Data == "forward_deposit_usdt":
		usdtSubscriptionsRepo := repositories.NewUserUsdtSubscriptionsRepository(db)

		usdtlist, err := usdtSubscriptionsRepo.ListAll(context.Background())

		if err != nil {

		}
		var allButtons []tgbotapi.InlineKeyboardButton
		var extraButtons []tgbotapi.InlineKeyboardButton
		var keyboard [][]tgbotapi.InlineKeyboardButton
		for _, usdtRecord := range usdtlist {
			allButtons = append(allButtons, tgbotapi.NewInlineKeyboardButtonData("💰"+usdtRecord.Name, "deposit_usdt_"+usdtRecord.Amount))
		}

		extraButtons = append(extraButtons, tgbotapi.NewInlineKeyboardButtonData("🔘切换到TRX充值", "deposit_amount"), tgbotapi.NewInlineKeyboardButtonData("🔙返回个人中心", "back_home"))

		for i := 0; i < len(allButtons); i += 2 {
			end := i + 2
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

		userRepo := repositories.NewUserRepository(db)

		user, _ := userRepo.GetByUserID(callbackQuery.Message.Chat.ID)
		if IsEmpty(user.Amount) {
			user.Amount = "0.00"
		}

		if IsEmpty(user.TronAmount) {
			user.TronAmount = "0.00"
		}

		msg := tgbotapi.NewMessage(callbackQuery.Message.Chat.ID,
			"🆔 用户ID: "+user.Associates+"\n"+
				"👤 用户名: @"+user.Username+"\n"+
				"💰 当前余额: "+"\n"+
				"- TRX：   "+user.TronAmount+"\n"+
				"-  USDT："+user.Amount)

		msg.ReplyMarkup = inlineKeyboard
		msg.ParseMode = "HTML"

		bot.Send(msg)

	default:
		responseText = "未知选项"
	}

	// 发送新消息作为响应
	bot.Send(tgbotapi.NewMessage(callbackQuery.Message.Chat.ID, responseText))
}
