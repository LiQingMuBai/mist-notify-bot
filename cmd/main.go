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
	"ushield_bot/internal/handler"
	"ushield_bot/internal/request"

	"ushield_bot/internal/cache"
	"ushield_bot/internal/domain"
	"ushield_bot/internal/infrastructure/repositories"
	. "ushield_bot/internal/infrastructure/tools"
)

// BotState 存储每个聊天中的分页状态
type DepositState struct {
	CurrentPage int
	TotalPages  int
}
type CostState struct {
	CurrentPage int
	TotalPages  int
}

var (
	depositStates = make(map[int64]*DepositState) // 按ChatID存储状态
	costStates    = make(map[int64]*CostState)    // 按ChatID存储状态
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

	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s", username, password, host, port, dbname)
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

	_cookie := os.Getenv("COOKIE")

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
				switch update.Message.Command() {
				case "start":
					log.Printf("1")

					//存用户
					userRepo := repositories.NewUserRepository(db)

					_, err := userRepo.GetByUserID(update.Message.Chat.ID)
					if err != nil {
						//增加用户
						var user domain.User
						user.Associates = strconv.FormatInt(update.Message.Chat.ID, 10)
						user.Username = update.Message.Chat.UserName
						//user.CreatedAt = time.Now()
						//user.UpdatedAt = time.Now()
						err := userRepo.Create2(context.Background(), &user)
						if err != nil {
							return
						}
					}

					handleStartCommand(cache, bot, update.Message)
				case "hide":
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

	msg := tgbotapi.NewMessage(message.Chat.ID, "U盾，做您链上资产的护盾！\n\n我们不仅关注低价能量，更专注于交易安全！\n\n让每一笔转账都更安心，让每一次链上交互都值得信任！\n\n🤖 "+
		"三大实用功能，助您安全、高效地管理链上资产\n\n🔋 波场能量闪兑\n\n🕵️ 地址风险检测\n\n🚨 USDT冻结预警\n\n您好："+message.Chat.UserName+" 欢迎使用U盾机器人\nU盾，做您链上资产的护盾！\n\n🔋 波场能量闪兑, 节省超过70%!\n🕵️ 地址风险检测, 让每一笔转账都更安心!\n"+
		"🚨 USDT冻结预警,秒级响应，让您的U永不冻结！\n新用户福利：\n每日一次地址风险查询\n常用指令：\n个人中心\n能量闪兑\n地址风险检测\nUSDT冻结预警\n客服：@Ushield001")
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

		userRepo := repositories.NewUserRepository(db)
		user, _ := userRepo.GetByUserID(message.Chat.ID)

		if IsEmpty(user.Amount) {
			user.Amount = "0.00"
		}

		if IsEmpty(user.TronAmount) {
			user.TronAmount = "0.00"
		}

		msg := tgbotapi.NewMessage(message.Chat.ID, "🔍 欢迎使用 U盾地址风险检测\n\n支持 TRON 或 ETH 网络任意地址查询\n\n系统将基于链上行为、风险标签、关联实体进行评分与分析\n\n📊 风险等级说明：\n🟢 低风险（0–30）：无异常交易，未关联已知风险实体\n\n🟡 中风险（31–70）：存在少量高风险交互，对手方不明\n\n🟠 高风险（71–90）：频繁异常转账，或与恶意地址有关\n\n🔴 极高风险（91–100）：涉及诈骗、制裁、黑客、洗钱等高风险行为\n\n📌 每位用户每天可免费检测 1 次\n\n💰 超出后每次扣除 4 TRX 或 1 USDT（系统将优先扣除 TRX）\n\n💼 当前余额：\n\n"+
			"- TRX："+user.TronAmount+"\n"+
			"- USDT："+user.Amount+"\n"+
			//"\n🔋 快速充值：\n➡️ 充值TRX\n➡️ 充值USDT\n\n请输入要检测的地址 👇")
			"请输入要检测的地址 👇")
		msg.ParseMode = "HTML"
		// 当点击"按钮 1"时显示内联键盘
		inlineKeyboard := tgbotapi.NewInlineKeyboardMarkup(
			tgbotapi.NewInlineKeyboardRow(
				tgbotapi.NewInlineKeyboardButtonData("💵充值", "deposit_amount"),
			),
		)
		msg.ReplyMarkup = inlineKeyboard

		bot.Send(msg)

		expiration := 1 * time.Minute // 短时间缓存空值

		//设置用户状态
		cache.Set(strconv.FormatInt(message.Chat.ID, 10), "usdt_risk_query", expiration)

	case "🚨USDT冻结预警":
		msg := tgbotapi.NewMessage(message.Chat.ID, "🛡️ U盾，做您链上资产的护盾！实时守护您的资产安全！\n\n地址一旦被链上风控冻，资产将难以追回，损失巨大！\n\n每天都有数百个 USDT 钱包地址被冻结锁定，风险就在身边！\n\nU盾将为您的地址提供 24 小时不间断监控\n\n⏰ 系统将在冻结前持续 10 分钟启动预警机制，每分钟推送提醒，通知您及时转移资产\n\n✅ 适用于经常收付款 / 高频交易 / 风险暴露地址\n\n✅ 支持在TRON网络下的USDT 钱包地址\n\n📌 服务价格（每地址）：\n\n- 2800 TRX / 30天\n- 或 800 USDT / 30天\n\n🎯 服务开启后系统将 24 小时不间断监控\n\n📩 所有预警信息将通过 Telegram 实时推送\n\n点击下方按钮开始 👇")
		msg.ParseMode = "HTML"

		inlineKeyboard := tgbotapi.NewInlineKeyboardMarkup(
			tgbotapi.NewInlineKeyboardRow(
				tgbotapi.NewInlineKeyboardButtonData("开启冻结预警", "start_freeze_risk"),
				tgbotapi.NewInlineKeyboardButtonData("地址管理", "address_manager"),
			),
			tgbotapi.NewInlineKeyboardRow(
				tgbotapi.NewInlineKeyboardButtonData("地址监控列表", "address_list_trace"),
			),
			//tgbotapi.NewInlineKeyboardRow(
			//	tgbotapi.NewInlineKeyboardButtonData("第二紧急通知", "address_manager_return"),
			//	//tgbotapi.NewInlineKeyboardButtonData("第二紧急通知", ""),
			//),
		)
		msg.ReplyMarkup = inlineKeyboard

		bot.Send(msg)

		expiration := 1 * time.Minute // 短时间缓存空值

		//设置用户状态
		cache.Set(strconv.FormatInt(message.Chat.ID, 10), "usdt_risk_monitor", expiration)

	case "🖊️笔数套餐":

		bundlesRepo := repositories.NewUserOperationBundlesRepository(db)

		trxlist, err := bundlesRepo.ListAll(context.Background())

		if err != nil {

		}

		var allButtons []tgbotapi.InlineKeyboardButton
		//var extraButtons []tgbotapi.InlineKeyboardButton
		var keyboard [][]tgbotapi.InlineKeyboardButton
		for _, trx := range trxlist {
			allButtons = append(allButtons, tgbotapi.NewInlineKeyboardButtonData("👝"+trx.Name, "bundle_"+trx.Amount))
		}

		//extraButtons = append(extraButtons, tgbotapi.NewInlineKeyboardButtonData("⚖️切换到USDT充值", "forward_deposit_usdt"), tgbotapi.NewInlineKeyboardButtonData("🔙返回上一级", "back_deposit_trx"))

		for i := 0; i < len(allButtons); i += 2 {
			end := i + 2
			if end > len(allButtons) {
				end = len(allButtons)
			}
			row := allButtons[i:end]
			keyboard = append(keyboard, tgbotapi.NewInlineKeyboardRow(row...))
		}

		// 3. 创建键盘标记
		inlineKeyboard := tgbotapi.NewInlineKeyboardMarkup(keyboard...)

		userRepo := repositories.NewUserRepository(db)
		user, _ := userRepo.GetByUserID(message.Chat.ID)
		if IsEmpty(user.Amount) {
			user.Amount = "0.00"
		}

		if IsEmpty(user.TronAmount) {
			user.TronAmount = "0.00"
		}

		msg := tgbotapi.NewMessage(message.Chat.ID,
			"💬"+"<b>"+"用户姓名: "+"</b>"+user.Username+"\n"+
				"👤"+"<b>"+"用户电报ID: "+"</b>"+user.Associates+"\n"+
				"💵"+"<b>"+"TRX余额:  "+"</b>"+user.TronAmount+" TRX"+"\n"+
				"💴"+"<b>"+"USDT余额:  "+"</b>"+user.Amount+" USDT"+"\n"+
				"【✏️笔数套餐】：\n"+
				"🔶赠送350带宽到地址，从此不在消耗0.35TRX\n"+
				"🔶按笔数计费的能量租用方式。\n"+
				"🔶每笔发送131K能量，对方地址无U也是扣一笔\n\n"+
				"🔶不限时，24小时内有一笔以上转账，不额外扣费！\n"+
				"1.24小时内未转账，会扣除一笔计数。\n"+
				"2.长时间不转账，可以在地址列表关闭笔数套餐\n\n🔥【真】【假】笔数套餐科普：\n"+
				"✅无论65K或者131K（对方地址是否有U），只扣一笔！\n"+
				"✅【🌈带宽笔笔送】\n"+
				//"🔸目前为促销ING,每笔赠送350带宽，从此不再消耗0.35 TRX，每笔节省0.35 TRX费用！\n"+
				"👆满足以上条件，才可称之为：【✏️笔数套餐】\n"+
				"➖➖➖➖➖➖➖➖➖\n"+
				"以下按钮可以选择不同的笔数套餐方案：")
		msg.ReplyMarkup = inlineKeyboard
		msg.ParseMode = "HTML"

		bot.Send(msg)

	case "⚡能量闪兑":
		// 当点击"按钮 1"时显示内联键盘
		inlineKeyboard := tgbotapi.NewInlineKeyboardMarkup(
			tgbotapi.NewInlineKeyboardRow(
				tgbotapi.NewInlineKeyboardButtonData("💵充值", "deposit_amount"),
			),
		)
		_agent := os.Getenv("Agent")
		sysUserRepo := repositories.NewSysUsersRepository(db)
		receiveAddress, _, _ := sysUserRepo.Find(context.Background(), _agent)

		//dictRepo := repositories.NewSysDictionariesRepo(db)
		//receiveAddress, _ := dictRepo.GetReceiveAddress(_agent)

		old_str := "【⚡️能量闪租】\n🔸转账  3 Trx=  1 笔能量\n🔸转账  6 Trx=  2 笔能量\n\n单笔 3 Trx，以此类推，最大 5 笔\n" +
			"1.向无U地址转账，需要双倍能量。\n2.请在1小时内转账，否则过期回收。\n\n🔸闪租能量收款地址:\n"

		old_str = "【⚡️能量闪租】\n\n 转账 3 TRX，系统自动按原路返还一笔能量，\n 如需向无U地址转账 ，请转账 6 TRX（返还两笔能量）\n\n"
		msg := tgbotapi.NewMessage(message.Chat.ID, old_str+
			//"```\n"+
			//"TQSrBJjbzgUThwE3N1ZJWoQ2mYgB581xij"+
			//"```\n\n"+
			"<code>"+receiveAddress+"</code>"+"\n"+
			"➖➖➖➖➖➖➖➖➖\n以下按钮可以选择其他能量租用模式：\n温馨提醒：\n闪租地址保存地址本要打上醒目标识，以免转账转错！")
		msg.ReplyMarkup = inlineKeyboard
		msg.ParseMode = "HTML"
		//msg.DisableWebPagePreview = true
		bot.Send(msg)

	case "钱包":
		userRepo := repositories.NewUserRepository(db)
		user, _ := userRepo.GetByUserID(message.Chat.ID)

		if IsEmpty(user.Amount) {
			user.Amount = "0.00"
		}

		if IsEmpty(user.TronAmount) {
			user.TronAmount = "0.00"
		}

		msg := tgbotapi.NewMessage(message.Chat.ID,
			"💬"+"<b>"+"用户姓名: "+"</b>"+user.Username+"\n"+
				"👤"+"<b>"+"用户电报ID: "+"</b>"+user.Associates+"\n"+
				"💵"+"<b>"+"TRX余额:  "+"</b>"+user.TronAmount+" TRX"+"\n"+
				"💴"+"<b>"+"USDT余额:  "+"</b>"+user.Amount+" USDT")
		msg.ParseMode = "HTML"
		bot.Send(msg)
	case "充值":

		// 当点击"按钮 1"时显示内联键盘
		inlineKeyboard := tgbotapi.NewInlineKeyboardMarkup(
			tgbotapi.NewInlineKeyboardRow(
				tgbotapi.NewInlineKeyboardButtonData("🕣充值", "deposit_amount"),
			),
		)

		userRepo := repositories.NewUserRepository(db)

		user, _ := userRepo.GetByUserID(message.Chat.ID)
		if IsEmpty(user.Amount) {
			user.Amount = "0.00"
		}

		if IsEmpty(user.TronAmount) {
			user.TronAmount = "0.00"
		}

		msg := tgbotapi.NewMessage(message.Chat.ID,
			"💬"+"<b>"+"用户姓名: "+"</b>"+user.Username+"\n"+
				"👤"+"<b>"+"用户电报ID: "+"</b>"+user.Associates+"\n"+
				"💵"+"<b>"+"TRX余额:  "+"</b>"+user.TronAmount+" TRX"+"\n"+
				"💴"+"<b>"+"USDT余额:  "+"</b>"+user.Amount+" USDT")

		msg.ReplyMarkup = inlineKeyboard
		msg.ParseMode = "HTML"

		bot.Send(msg)
	case "👤个人中心":

		inlineKeyboard := tgbotapi.NewInlineKeyboardMarkup(
			//tgbotapi.NewInlineKeyboardRow(
			//	tgbotapi.NewInlineKeyboardButtonData("🆔我的账户", "click_my_account"),
			//
			//),
			tgbotapi.NewInlineKeyboardRow(
				tgbotapi.NewInlineKeyboardButtonData("💳充值", "deposit_amount"),
				tgbotapi.NewInlineKeyboardButtonData("📄账单", "click_my_recepit"),
				tgbotapi.NewInlineKeyboardButtonData("🛠️我的服务", "click_my_service"),
			),
			tgbotapi.NewInlineKeyboardRow(
				//tgbotapi.NewInlineKeyboardButtonData("🔗绑定备用帐号", "click_backup_account"),
				tgbotapi.NewInlineKeyboardButtonData("👥商务合作", "click_business_cooperation"),
				tgbotapi.NewInlineKeyboardButtonData("🛎️客服", "click_callcenter"),
				tgbotapi.NewInlineKeyboardButtonData("❓常见问题FAQ", "click_QA"),
			),
			//tgbotapi.NewInlineKeyboardRow(),
		)

		userRepo := repositories.NewUserRepository(db)
		user, _ := userRepo.GetByUserID(message.Chat.ID)

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

		msg := tgbotapi.NewMessage(message.Chat.ID, "📇 我的账户\n\n🆔 用户ID："+user.Associates+"\n\n👤 用户名：@"+user.Username+"\n\n"+
			str+"\n\n💰 "+
			"当前余额：\n\n"+
			"- TRX："+user.TronAmount+"\n"+
			"- USDT："+user.Amount)
		msg.ReplyMarkup = inlineKeyboard
		msg.ParseMode = "HTML"
		bot.Send(msg)

	case "客服":
		msg := tgbotapi.NewMessage(message.Chat.ID, "📞联系客服：@Ushield001\n")
		msg.ParseMode = "HTML"

		bot.Send(msg)

	case "账单":

		userRepo := repositories.NewUserRepository(db)
		user, _ := userRepo.GetByUserID(message.Chat.ID)
		if IsEmpty(user.Amount) {
			user.Amount = "0.00"
		}

		if IsEmpty(user.TronAmount) {
			user.TronAmount = "0.00"
		}

		usdtDepositRepo := repositories.NewUserUSDTDepositsRepository(db)
		usdtlist, _ := usdtDepositRepo.ListAll(context.Background(), message.Chat.ID, 1)

		trxDepositRepo := repositories.NewUserTRXDepositsRepository(db)
		trxlist, _ := trxDepositRepo.ListAll(context.Background(), message.Chat.ID, 1)

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

		msg := tgbotapi.NewMessage(message.Chat.ID, "🧾 我的账单记录\n\n📌 "+
			"当前余额：\n\n- TRX："+user.TronAmount+"\n- USDT："+user.Amount+"\n\n📥 "+
			"充值记录：\n "+
			result+"\n"+
			result2+"\n")
		msg.ParseMode = "HTML"

		bot.Send(msg)

	case "帮助":
		bot.Send(tgbotapi.NewMessage(message.Chat.ID, "帮助信息：\n- 点击'按钮 1'显示内联菜单\n- 使用 /start 重新显示键盘\n- 使用 /hide 隐藏键盘"))
	default:
		status, _ := cache.Get(strconv.FormatInt(message.Chat.ID, 10))

		log.Printf("用户状态staus %s", status)
		switch {
		case strings.HasPrefix(status, "user_backup_notify"):

			chat_ID, err := strconv.ParseInt(message.Text, 10, 64)
			if err != nil {

				msg := tgbotapi.NewMessage(message.Chat.ID, "请输入正确的对方👤用户电报ID？")
				msg.ParseMode = "HTML"
				bot.Send(msg)
				return
			}

			//用户电报ID
			userRepo := repositories.NewUserRepository(db)
			backupUser, esg := userRepo.GetByUserID(chat_ID)
			if esg != nil {
				msg := tgbotapi.NewMessage(message.Chat.ID, "👤用户电报ID未在机器人发现，请让对方用户电报登录机器人")
				msg.ParseMode = "HTML"
				bot.Send(msg)
				return
			}
			user, _ := userRepo.GetByUserID(message.Chat.ID)
			user.BackupChatID = backupUser.Associates
			err2 := userRepo.Update2(context.Background(), &user)
			if err2 == nil {
				msg := tgbotapi.NewMessage(message.Chat.ID, "✅ 成功绑定第二紧急联系人: "+backupUser.Associates)
				msg.ParseMode = "HTML"
				bot.Send(msg)
				return
			}
		case strings.HasPrefix(status, "start_freeze_risk"):
			//msg := tgbotapi.NewMessage(message.Chat.ID, "📡 系统将自动启动 24 小时预警服务\n如检测到潜在冻结风险，系统将在冻结前持续 10 分钟预警\n每分钟推送提醒，通知您及时转移资产，避免冻结损失\n📌 服务费用：2800 TRX / 30 天 或 800 USDT / 30 天\n是否确认启用该服务？")
			//msg.ParseMode = "HTML"
			//
			//inlineKeyboard := tgbotapi.NewInlineKeyboardMarkup(
			//	tgbotapi.NewInlineKeyboardRow(
			//		tgbotapi.NewInlineKeyboardButtonData("✅ 确认开通", "start_freeze_risk"),
			//		tgbotapi.NewInlineKeyboardButtonData("❌ 取消操作", "address_manager"),
			//	),
			//	//tgbotapi.NewInlineKeyboardRow(
			//	//	tgbotapi.NewInlineKeyboardButtonData("地址", ""),
			//	//),
			//)
			//msg.ReplyMarkup = inlineKeyboard
			//
			//bot.Send(msg)
			//
			//expiration := 1 * time.Minute // 短时间缓存空值
			//
			////设置用户状态
			//cache.Set(strconv.FormatInt(message.Chat.ID, 10), "usdt_risk_monitor", expiration)
		case strings.HasPrefix(status, "address_list_trace"):
		//msg := tgbotapi.NewMessage(message.Chat.ID, "📡 系统将自动启动 24 小时预警服务\n如检测到潜在冻结风险，系统将在冻结前持续 10 分钟预警\n每分钟推送提醒，通知您及时转移资产，避免冻结损失\n📌 服务费用：2800 TRX / 30 天 或 800 USDT / 30 天\n是否确认启用该服务？")
		//msg.ParseMode = "HTML"
		//
		//inlineKeyboard := tgbotapi.NewInlineKeyboardMarkup(
		//	tgbotapi.NewInlineKeyboardRow(
		//		tgbotapi.NewInlineKeyboardButtonData("✅ 确认开通", "start_freeze_risk"),
		//		tgbotapi.NewInlineKeyboardButtonData("❌ 取消操作", "address_manager"),
		//	),
		//	//tgbotapi.NewInlineKeyboardRow(
		//	//	tgbotapi.NewInlineKeyboardButtonData("地址", ""),
		//	//),
		//)
		//msg.ReplyMarkup = inlineKeyboard
		//
		//bot.Send(msg)
		//
		//expiration := 1 * time.Minute // 短时间缓存空值
		//
		////设置用户状态
		//cache.Set(strconv.FormatInt(message.Chat.ID, 10), "usdt_risk_monitor", expiration)
		case strings.HasPrefix(status, "address_manager_remove"):
			if IsValidAddress(message.Text) || IsValidEthereumAddress(message.Text) {
				userRepo := repositories.NewUserAddressMonitorRepo(db)
				err := userRepo.Remove(context.Background(), message.Chat.ID, message.Text)
				if err != nil {
				}
				msg := tgbotapi.NewMessage(message.Chat.ID, "💬"+"<b>"+"地址删除成功 "+"</b>"+"\n")
				msg.ParseMode = "HTML"
				bot.Send(msg)
			} else {
				msg := tgbotapi.NewMessage(message.Chat.ID, "💬"+"<b>"+"地址有误，请重新输入需删除的地址: "+"</b>"+"\n")
				msg.ParseMode = "HTML"
				bot.Send(msg)
			}

		case strings.HasPrefix(status, "address_manager_add"):
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

		case strings.HasPrefix(status, "bundle_"):
			fmt.Printf(">>>>>>>>>>>>>>>>>>>>bundle: %s", status)

			if !IsValidAddress(message.Text) {
				msg := tgbotapi.NewMessage(message.Chat.ID, "💬"+"<b>"+"地址有误，请重新输入能量接收地址: "+"</b>"+"\n")
				msg.ParseMode = "HTML"
				bot.Send(msg)
				return
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
			//userRepo := repositories.NewUserRepository(db)
			//user, _ := userRepo.GetByUserID(message.Chat.ID)

			//if IsEmpty(user.Amount) {
			//	user.Amount = "0.00"
			//}
			//
			//if IsEmpty(user.TronAmount) {
			//	user.TronAmount = "0.00"
			//}

			//msg := tgbotapi.NewMessage(message.Chat.ID,
			//	"💬"+"<b>"+"用户姓名: "+"</b>"+user.Username+"\n"+
			//		"👤"+"<b>"+"用户电报ID: "+"</b>"+user.Associates+"\n"+
			//		"💵"+"<b>"+"TRX余额:  "+"</b>"+user.TronAmount+" TRX"+"\n"+
			//		"💴"+"<b>"+"USDT余额:  "+"</b>"+user.Amount+" USDT")
			//msg.ParseMode = "HTML"
			//bot.Send(msg)

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

		case strings.HasPrefix(status, "usdt_risk_query"):
			//fmt.Printf("bundle: %s", status)

			if IsValidAddress(message.Text) || IsValidEthereumAddress(message.Text) {
				userRepo := repositories.NewUserRepository(db)
				user, _ := userRepo.GetByUserID(message.Chat.ID)
				//if strings.Contains(message.Chat.UserName, "Ushield") {
				//	user.Times = 10000
				//}

				if user.Times == 1 {

					//需要扣钱 4trx或者1u
					if CompareStringsWithFloat(user.Amount, "1", 1) || CompareStringsWithFloat(user.TronAmount, "4", 1) {

						if CompareStringsWithFloat(user.Amount, "1", 1) {
							amount, _ := SubtractStringNumbers(user.Amount, "1", 1)
							user.Amount = amount
							userRepo.Update2(context.Background(), &user)
						}

						if CompareStringsWithFloat(user.TronAmount, "4", 1) {
							tronAmount, _ := SubtractStringNumbers(user.TronAmount, "4", 1)
							user.TronAmount = tronAmount
							userRepo.Update2(context.Background(), &user)
						}
						_text := ""
						if strings.HasPrefix(message.Text, "0x") && len(message.Text) == 42 {
							_symbol := "USDT-ERC20"
							_addressInfo := handler.GetAddressInfo(_symbol, message.Text, _cookie)
							_text = handler.GetText(_addressInfo)

							addressProfile := handler.GetAddressProfile(_symbol, message.Text, _cookie)
							_text7 := "余額：" + addressProfile.BalanceUsd + "\n"
							_text8 := "累計收入：" + addressProfile.TotalReceivedUsd + "\n"
							_text9 := "累计支出：" + addressProfile.TotalSpentUsd + "\n"
							_text10 := "首次活躍時間：" + addressProfile.FirstTxTime + "\n"
							_text11 := "最後活躍時間：" + addressProfile.LastTxTime + "\n"
							_text12 := "交易次數：" + addressProfile.TxCount + "筆" + "\n"
							_text99 := "主要交易对手分析：" + "\n"
							_text5 := "📢更多查询請聯繫客服 @Ushield001\n"
							_text16 := "🛡️ U盾在手，链上无忧！" + "\n"

							_text = _text + _text7 + _text8 + _text9 + _text10 + _text11 + _text12 + _text99 + _text5 + _text16

						}
						if strings.HasPrefix(message.Text, "T") && len(message.Text) == 34 {
							_symbol := "USDT-TRC20"
							_addressInfo := handler.GetAddressInfo(_symbol, message.Text, _cookie)
							_text = handler.GetText(_addressInfo)

							addressProfile := handler.GetAddressProfile(_symbol, message.Text, _cookie)
							_text7 := "余額：" + addressProfile.BalanceUsd + "\n"
							_text8 := "累計收入：" + addressProfile.TotalReceivedUsd + "\n"
							_text9 := "累计支出：" + addressProfile.TotalSpentUsd + "\n"
							_text10 := "首次活躍時間：" + addressProfile.FirstTxTime + "\n"
							_text11 := "最後活躍時間：" + addressProfile.LastTxTime + "\n"
							_text12 := "交易次數：" + addressProfile.TxCount + "筆" + "\n"
							_text99 := "危险交易对手分析：" + "\n"
							lableAddresList := handler.GetNotSafeAddress(_symbol, message.Text, _cookie)

							_text100 := ""
							if len(lableAddresList.GraphDic.NodeList) > 0 {
								for _, data := range lableAddresList.GraphDic.NodeList {
									if strings.Contains(data.Label, "huione") {
										_text100 = _text100 + data.Title[0:5] + "..." + data.Title[29:34] + "\n"
									}
								}
							}
							_text5 := "📢更多查询請聯繫客服 @Ushield001\n"
							_text16 := "🛡️ U盾在手，链上无忧！" + "\n"

							_text = _text + _text7 + _text8 + _text9 + _text10 + _text11 + _text12 + _text99 + _text100 + _text5 + _text16

						}
						msg := tgbotapi.NewMessage(message.Chat.ID, _text)
						//msg.ReplyMarkup = inlineKeyboard
						msg.ParseMode = "HTML"
						bot.Send(msg)
						userRepo.UpdateTimesByChatID(1, message.Chat.ID)
					} else {
						//msg := tgbotapi.NewMessage(message.Chat.ID,
						//	"🔍普通用戶每日贈送 1 次地址風險查詢\n"+
						//		"📞聯繫客服 @Ushield001\n")
						//msg.ReplyMarkup = inlineKeyboard

						msg := tgbotapi.NewMessage(message.Chat.ID,
							"💬"+"<b>"+"🔍普通用戶每日贈送 1 次地址風險查詢 "+"</b>"+user.Username+"\n"+
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
						//bot.Send(msg)

						msg.ParseMode = "HTML"
						bot.Send(msg)
					}
				} else {
					_text := ""
					if strings.HasPrefix(message.Text, "0x") && len(message.Text) == 42 {
						_symbol := "USDT-ERC20"
						_addressInfo := handler.GetAddressInfo(_symbol, message.Text, _cookie)
						_text = handler.GetText(_addressInfo)

						addressProfile := handler.GetAddressProfile(_symbol, message.Text, _cookie)
						_text7 := "余額：" + addressProfile.BalanceUsd + "\n"
						_text8 := "累計收入：" + addressProfile.TotalReceivedUsd + "\n"
						_text9 := "累计支出：" + addressProfile.TotalSpentUsd + "\n"
						_text10 := "首次活躍時間：" + addressProfile.FirstTxTime + "\n"
						_text11 := "最後活躍時間：" + addressProfile.LastTxTime + "\n"
						_text12 := "交易次數：" + addressProfile.TxCount + "筆" + "\n"
						_text99 := "主要交易对手分析：" + "\n"
						_text5 := "📢更多查询請聯繫客服 @Ushield001\n"
						_text16 := "🛡️ U盾在手，链上无忧！" + "\n"

						_text = _text + _text7 + _text8 + _text9 + _text10 + _text11 + _text12 + _text99 + _text5 + _text16

					}
					if strings.HasPrefix(message.Text, "T") && len(message.Text) == 34 {
						_symbol := "USDT-TRC20"
						_addressInfo := handler.GetAddressInfo(_symbol, message.Text, _cookie)
						_text = handler.GetText(_addressInfo)

						addressProfile := handler.GetAddressProfile(_symbol, message.Text, _cookie)
						_text7 := "余額：" + addressProfile.BalanceUsd + "\n"
						_text8 := "累計收入：" + addressProfile.TotalReceivedUsd + "\n"
						_text9 := "累计支出：" + addressProfile.TotalSpentUsd + "\n"
						_text10 := "首次活躍時間：" + addressProfile.FirstTxTime + "\n"
						_text11 := "最後活躍時間：" + addressProfile.LastTxTime + "\n"
						_text12 := "交易次數：" + addressProfile.TxCount + "筆" + "\n"
						_text99 := "危险交易对手分析：" + "\n"
						lableAddresList := handler.GetNotSafeAddress(_symbol, message.Text, _cookie)

						_text100 := ""
						if len(lableAddresList.GraphDic.NodeList) > 0 {
							for _, data := range lableAddresList.GraphDic.NodeList {
								if strings.Contains(data.Label, "huione") {
									_text100 = _text100 + data.Title[0:5] + "..." + data.Title[29:34] + "\n"
								}
							}
						}
						_text5 := "📢更多查询請聯繫客服 @Ushield001\n"
						_text16 := "🛡️ U盾在手，链上无忧！" + "\n"

						_text = _text + _text7 + _text8 + _text9 + _text10 + _text11 + _text12 + _text99 + _text100 + _text5 + _text16

					}
					msg := tgbotapi.NewMessage(message.Chat.ID, _text)
					//msg.ReplyMarkup = inlineKeyboard
					msg.ParseMode = "HTML"
					bot.Send(msg)
					userRepo.UpdateTimesByChatID(1, message.Chat.ID)
				}

			} else {
				msg := tgbotapi.NewMessage(message.Chat.ID, "💬"+"<b>"+"地址有误，请重新输入地址: "+"</b>"+"\n")
				msg.ParseMode = "HTML"
				bot.Send(msg)
			}

		}

		//bot.Send(tgbotapi.NewMessage(message.Chat.ID, "收到消息: "+message.Text))
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
	case callbackQuery.Data == "address_list_trace":
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
	case callbackQuery.Data == "back_home":

		inlineKeyboard := tgbotapi.NewInlineKeyboardMarkup(
			//tgbotapi.NewInlineKeyboardRow(
			//	tgbotapi.NewInlineKeyboardButtonData("🆔我的账户", "click_my_account"),
			//	tgbotapi.NewInlineKeyboardButtonData("💳充值", "click_my_deposit"),
			//),
			tgbotapi.NewInlineKeyboardRow(
				tgbotapi.NewInlineKeyboardButtonData("💳充值", "deposit_amount"),
				tgbotapi.NewInlineKeyboardButtonData("📄账单", "click_my_recepit"),
				tgbotapi.NewInlineKeyboardButtonData("🛠️我的服务", "click_my_service"),
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

	case callbackQuery.Data == "click_business_cooperation":
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

	case callbackQuery.Data == "click_callcenter":
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

	case callbackQuery.Data == "click_my_recepit":
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
				tgbotapi.NewInlineKeyboardButtonData("⬇️充值记录", "click_deposit_records"),
				tgbotapi.NewInlineKeyboardButtonData("⬆️服务扣款记录", "click_cost_records"),
			),
			tgbotapi.NewInlineKeyboardRow(
				//tgbotapi.NewInlineKeyboardButtonData("解绑地址", "free_monitor_address"),
				tgbotapi.NewInlineKeyboardButtonData("🔙返回个人中心", "back_home"),
			),
		)
		msg.ReplyMarkup = inlineKeyboard
		bot.Send(msg)
	case callbackQuery.Data == "click_cost_records":
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
				tgbotapi.NewInlineKeyboardButtonData("上一页", "click_deposit_records"),
				tgbotapi.NewInlineKeyboardButtonData("下一页", "click_cost_records"),
			),
			tgbotapi.NewInlineKeyboardRow(
				//tgbotapi.NewInlineKeyboardButtonData("解绑地址", "free_monitor_address"),
				tgbotapi.NewInlineKeyboardButtonData("🔙返回个人中心", "back_home"),
			),
		)
		msg.ReplyMarkup = inlineKeyboard
		bot.Send(msg)
	case callbackQuery.Data == "click_deposit_records":
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
		msg := tgbotapi.NewMessage(callbackQuery.Message.Chat.ID, "🧾充值记录\n\n "+
			result+"\n"+
			result2+"\n")
		msg.ParseMode = "HTML"
		inlineKeyboard := tgbotapi.NewInlineKeyboardMarkup(
			tgbotapi.NewInlineKeyboardRow(
				tgbotapi.NewInlineKeyboardButtonData("上一页", "prev_deposit_page"),
				tgbotapi.NewInlineKeyboardButtonData("下一页", "next_deposit_page"),
			),
			tgbotapi.NewInlineKeyboardRow(
				//tgbotapi.NewInlineKeyboardButtonData("解绑地址", "free_monitor_address"),
				tgbotapi.NewInlineKeyboardButtonData("🔙返回个人中心", "back_home"),
			),
		)
		msg.ReplyMarkup = inlineKeyboard
		bot.Send(msg)
	case callbackQuery.Data == "prev_deposit_page":
		state := depositStates[callbackQuery.Message.Chat.ID]
		if state == nil {
			var state DepositState
			state.CurrentPage = 1
			depositStates[callbackQuery.Message.Chat.ID] = &state
			trxDepositRepo := repositories.NewUserTRXDepositsRepository(db)
			var info request.UserTrxDepositsSearch
			info.PageInfo.Page = 1
			info.PageInfo.PageSize = 20
			trxDepositRepo.GetUserTrxDepositsInfoList(context.Background(), info, callbackQuery.Message.Chat.ID)
		} else {
			state.CurrentPage = state.CurrentPage + 1
			trxDepositRepo := repositories.NewUserTRXDepositsRepository(db)
			var info request.UserTrxDepositsSearch
			info.PageInfo.Page = 1
			info.PageInfo.PageSize = 20
			trxDepositRepo.GetUserTrxDepositsInfoList(context.Background(), info, callbackQuery.Message.Chat.ID)
		}
		fmt.Printf("state: %v\n", state)

	case callbackQuery.Data == "next_deposit_page":
		state := depositStates[callbackQuery.Message.Chat.ID]
		if state != nil && state.CurrentPage > 1 {
			state.CurrentPage = state.CurrentPage - 1
			trxDepositRepo := repositories.NewUserTRXDepositsRepository(db)
			var info request.UserTrxDepositsSearch
			info.PageInfo.Page = 1
			info.PageInfo.PageSize = 20
			trxDepositRepo.GetUserTrxDepositsInfoList(context.Background(), info, callbackQuery.Message.Chat.ID)
		}
		fmt.Printf("state: %v\n", state)

	case callbackQuery.Data == "click_QA":
		msg := tgbotapi.NewMessage(callbackQuery.Message.Chat.ID, "🧠 常见问题帮助\n\n💰充值\n\n1️⃣充值金额输错未到账怎么办？\n\n➡️联系客服，客服将为您提供一笔小额确认金额订单（例如：1.003 TRX），用于验证您对原转账钱包的控制权。请提供原错误订单的转账截图和新的身份验证订单转账截图一并发给客服，待确认身份后客服将给予入账。\n\n🔋 能量闪兑\n\n1️⃣ 为什么我转了 3 TRX 没有收到能量？\n\n➡️ 请确认是否向正确地址转账，系统只识别官方闪兑地址，请核对官方闪兑地址TENERGYxxxxxxxxxxxxx。\n\n2️⃣ 笔数套餐如何查看剩余？\n\n➡️ 点击 个人中心/我的服务 查看剩余笔数与补能设置。\n\n3️⃣ 为什么 购买能量后USDT 转账时还是扣除了钱包的TRX作为手续费\n\n➡️ 可能因向无U地址转账导致当前钱包能量不足，请根据钱包转账最后的手续费提示，如需扣除TRX建议再次购买一笔能量以抵扣手续费。\n\n📍 地址检测\n\n1️⃣ 每天免费次数是多少？\n\n➡️ 每位用户每天可免费检测 1 次，之后需付费。\n\n2️⃣ 检测时余额不足怎么办？\n\n➡️ 系统将提示充值链接并生成支付订单。\n\n3️⃣ 地址风险评分是如何判断的？\n\n➡️ 基于链上行为、交互对象与风险标签等维度综合评分。\n\n🚨 冻结预警\n\n1️⃣ 如何判断地址是否被冻结？\n\n➡️ 预警服务采用多个服务综合判断确保地址在冻结前 持续10 分钟发送连续警报提醒用户转移资产。\n\n2️⃣ 服务能否转移到其他地址？\n\n➡️ 当前按地址计费，不支持转移或换绑。\n\n3️⃣ 到期是否自动续费？\n\n➡️ 系统将尝试自动扣费，余额不足会提前通知用户。\n\n4️⃣一个账号能绑定多个地址同时进行监控吗？\n\n➡️是的，单个账号可绑定多个地址进行服务监控\n\n每个地址单独计费。\n\n👥帐号问题\n\n1️⃣ 观察者模式与全局模式的区别\n\n➡️观察者模式只可接收冻结预警无法执行服务操作（如发能、查询、解绑）， 全局模式等同主账号权限，可进行所有操作（如检测、续费、管理服务），您可随时通过 /解绑地址 或 /更改权限进行调整。\n\n2️⃣  主账号被盗，丢失不可用应急说明\n\n➡️若备用账号为「全局模式」，可使用备用帐号正常继续使用所有服务  。\n\n➡️ 若为「观察者模式」，仅能查看推送，无法操作服务。\n\n如需更改备用帐号权限请准备主账号最近一个月有充值记录的钱包并联系客服确认身份")
		inlineKeyboard := tgbotapi.NewInlineKeyboardMarkup(

			tgbotapi.NewInlineKeyboardRow(
				//tgbotapi.NewInlineKeyboardButtonData("解绑地址", "free_monitor_address"),
				tgbotapi.NewInlineKeyboardButtonData("🔙返回个人中心", "back_home"),
			),
		)
		msg.ReplyMarkup = inlineKeyboard
		msg.ParseMode = "HTML"
		bot.Send(msg)

		expiration := 1 * time.Minute // 短时间缓存空值

		//设置用户状态
		cache.Set(strconv.FormatInt(callbackQuery.Message.Chat.ID, 10), callbackQuery.Data, expiration)

	case callbackQuery.Data == "user_backup_notify":
		msg := tgbotapi.NewMessage(callbackQuery.Message.Chat.ID, "💬"+"<b>"+"请输入需添加的第二紧急通知用户电报ID: "+"</b>"+"\n")
		msg.ParseMode = "HTML"
		bot.Send(msg)

		expiration := 1 * time.Minute // 短时间缓存空值

		//设置用户状态
		cache.Set(strconv.FormatInt(callbackQuery.Message.Chat.ID, 10), callbackQuery.Data, expiration)
	case callbackQuery.Data == "start_freeze_risk_1":
		//查看余额
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
		if CompareStringsWithFloat(user.TronAmount, "2800", float64(nums)) || CompareStringsWithFloat(user.Amount, "800", float64(nums)) {
			//扣减

			//扣trx
			if CompareStringsWithFloat(user.TronAmount, "2800", float64(nums)) {
				rest, _ := SubtractStringNumbers(user.TronAmount, "2800", float64(nums))

				user.TronAmount = rest
				userRepo.Update2(context.Background(), &user)
				fmt.Printf("rest: %s", rest)

			} else if CompareStringsWithFloat(user.Amount, "800", float64(nums)) {
				rest, _ := SubtractStringNumbers(user.Amount, "800", float64(nums))
				fmt.Printf("rest: %s", rest)
				user.Amount = rest
				userRepo.Update2(context.Background(), &user)
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

	case callbackQuery.Data == "click_my_service":
		msg := tgbotapi.NewMessage(callbackQuery.Message.Chat.ID, "🛡 当前服务状态：\n\n🔋 能量闪兑\n\n- 剩余笔数：12\n- 自动补能：关闭 /开启\n\n➡️ /闪兑\n\n➡️ /笔数套餐\n\n➡️ /手动发能（1笔）\n\n➡️ /开启/关闭自动发能\n\n📍 地址风险检测\n\n- 今日免费次数：已用完\n\n➡️ /地址风险检测\n\n🚨 USDT冻结预警\n\n- 地址1：TX8kY...5a9rP（剩余12天）✅\n- 地址2：TEw9Q...iS6Ht（剩余28天）✅")
		msg.ParseMode = "HTML"

		inlineKeyboard := tgbotapi.NewInlineKeyboardMarkup(
			tgbotapi.NewInlineKeyboardRow(
				tgbotapi.NewInlineKeyboardButtonData("地址监控列表", "address_list_trace"),
				tgbotapi.NewInlineKeyboardButtonData("地址管理", "address_manager"),
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
		msg := tgbotapi.NewMessage(callbackQuery.Message.Chat.ID, "🛡️ U盾，做您链上资产的护盾！实时守护您的资产安全！\n\n地址一旦被链上风控冻，资产将难以追回，损失巨大！\n\n每天都有数百个 USDT 钱包地址被冻结锁定，风险就在身边！\n\nU盾将为您的地址提供 24 小时不间断监控\n\n⏰ 系统将在冻结前持续 10 分钟启动预警机制，每分钟推送提醒，通知您及时转移资产\n\n✅ 适用于经常收付款 / 高频交易 / 风险暴露地址\n\n✅ 支持在TRON网络下的USDT 钱包地址\n\n📌 服务价格（每地址）：\n\n- 2800 TRX / 30天\n- 或 800 USDT / 30天\n\n🎯 服务开启后系统将 24 小时不间断监控\n\n📩 所有预警信息将通过 Telegram 实时推送\n\n点击下方按钮开始 👇")
		msg.ParseMode = "HTML"

		inlineKeyboard := tgbotapi.NewInlineKeyboardMarkup(
			tgbotapi.NewInlineKeyboardRow(
				tgbotapi.NewInlineKeyboardButtonData("开启冻结预警", "start_freeze_risk"),
				tgbotapi.NewInlineKeyboardButtonData("地址管理", "address_manager"),
			),
			tgbotapi.NewInlineKeyboardRow(
				tgbotapi.NewInlineKeyboardButtonData("地址监控列表", "address_list_trace"),
			),
		)
		msg.ReplyMarkup = inlineKeyboard

		bot.Send(msg)

		expiration := 1 * time.Minute // 短时间缓存空值

		//设置用户状态
		cache.Set(strconv.FormatInt(callbackQuery.Message.Chat.ID, 10), "usdt_risk_monitor", expiration)
	case callbackQuery.Data == "stop_freeze_risk":
		msg := tgbotapi.NewMessage(callbackQuery.Message.Chat.ID, "📡 是否确认停止该服务？")
		msg.ParseMode = "HTML"

		inlineKeyboard := tgbotapi.NewInlineKeyboardMarkup(
			tgbotapi.NewInlineKeyboardRow(
				tgbotapi.NewInlineKeyboardButtonData("✅ 确认停止", "stop_freeze_risk_1"),
				tgbotapi.NewInlineKeyboardButtonData("❌ 取消操作", "start_freeze_risk_0"),
			),
			//tgbotapi.NewInlineKeyboardRow(
			//	tgbotapi.NewInlineKeyboardButtonData("地址", ""),
			//),
		)
		msg.ReplyMarkup = inlineKeyboard

		bot.Send(msg)

		expiration := 1 * time.Minute // 短时间缓存空值

		//设置用户状态
		cache.Set(strconv.FormatInt(callbackQuery.Message.Chat.ID, 10), "start_freeze_risk", expiration)

	case callbackQuery.Data == "start_freeze_risk":
		msg := tgbotapi.NewMessage(callbackQuery.Message.Chat.ID, "📡 系统将自动启动 24 小时预警服务\n如检测到潜在冻结风险，系统将在冻结前持续 10 分钟预警\n每分钟推送提醒，通知您及时转移资产，避免冻结损失\n📌 服务费用：2800 TRX / 30 天 或 800 USDT / 30 天\n是否确认启用该服务？")
		msg.ParseMode = "HTML"

		inlineKeyboard := tgbotapi.NewInlineKeyboardMarkup(
			tgbotapi.NewInlineKeyboardRow(
				tgbotapi.NewInlineKeyboardButtonData("✅ 确认开通", "start_freeze_risk_1"),
				tgbotapi.NewInlineKeyboardButtonData("❌ 取消操作", "start_freeze_risk_0"),
			),
			//tgbotapi.NewInlineKeyboardRow(
			//	tgbotapi.NewInlineKeyboardButtonData("地址", ""),
			//),
		)
		msg.ReplyMarkup = inlineKeyboard

		bot.Send(msg)

		expiration := 1 * time.Minute // 短时间缓存空值

		//设置用户状态
		cache.Set(strconv.FormatInt(callbackQuery.Message.Chat.ID, 10), "start_freeze_risk", expiration)

	case callbackQuery.Data == "address_manager_return":
		msg := tgbotapi.NewMessage(callbackQuery.Message.Chat.ID, "🛡️ U盾，做您链上资产的护盾！实时守护您的资产安全！\n\n地址一旦被链上风控冻，资产将难以追回，损失巨大！\n\n每天都有数百个 USDT 钱包地址被冻结锁定，风险就在身边！\n\nU盾将为您的地址提供 24 小时不间断监控\n\n⏰ 系统将在冻结前持续 10 分钟启动预警机制，每分钟推送提醒，通知您及时转移资产\n\n✅ 适用于经常收付款 / 高频交易 / 风险暴露地址\n\n✅ 支持在TRON网络下的USDT 钱包地址\n\n📌 服务价格（每地址）：\n\n- 2800 TRX / 30天\n- 或 800 USDT / 30天\n\n🎯 服务开启后系统将 24 小时不间断监控\n\n📩 所有预警信息将通过 Telegram 实时推送\n\n点击下方按钮开始 👇")
		msg.ParseMode = "HTML"

		inlineKeyboard := tgbotapi.NewInlineKeyboardMarkup(
			tgbotapi.NewInlineKeyboardRow(
				tgbotapi.NewInlineKeyboardButtonData("开启冻结预警", "start_freeze_risk"),
				tgbotapi.NewInlineKeyboardButtonData("地址管理", "address_manager"),
			),
			tgbotapi.NewInlineKeyboardRow(
				tgbotapi.NewInlineKeyboardButtonData("地址监控列表", "address_list_trace"),
			),
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
				tgbotapi.NewInlineKeyboardButtonData("⬅️返回", "address_manager_return"),
			),
		)
		msg.ReplyMarkup = inlineKeyboard

		bot.Send(msg)

		expiration := 1 * time.Minute // 短时间缓存空值

		//设置用户状态
		cache.Set(strconv.FormatInt(callbackQuery.Message.Chat.ID, 10), "address_manager", expiration)

	case callbackQuery.Data == "deposit_amount":

		trxSubscriptionsRepo := repositories.NewUserTRXSubscriptionsRepository(db)

		trxlist, _ := trxSubscriptionsRepo.ListAll(context.Background())

		//if err != nil {
		//
		//}
		var allButtons []tgbotapi.InlineKeyboardButton
		var extraButtons []tgbotapi.InlineKeyboardButton
		var keyboard [][]tgbotapi.InlineKeyboardButton
		for _, trx := range trxlist {
			allButtons = append(allButtons, tgbotapi.NewInlineKeyboardButtonData("🏦"+trx.Name, "deposit_trx_"+trx.Amount))
		}

		extraButtons = append(extraButtons, tgbotapi.NewInlineKeyboardButtonData("⚖️切换到USDT充值", "forward_deposit_usdt"), tgbotapi.NewInlineKeyboardButtonData("🔙返回个人中心", "back_home"))

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
			"💬"+"<b>"+"用户姓名: "+"</b>"+user.Username+"\n"+
				"👤"+"<b>"+"用户电报ID: "+"</b>"+user.Associates+"\n"+
				"💵"+"<b>"+"TRX余额:  "+"</b>"+user.TronAmount+" TRX"+"\n"+
				"💴"+"<b>"+"USDT余额:  "+"</b>"+user.Amount+" USDT")
		msg.ReplyMarkup = inlineKeyboard
		msg.ParseMode = "HTML"

		bot.Send(msg)

	case strings.HasPrefix(callbackQuery.Data, "bundle_"):

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
		//调用trxfee接口进行笔数扣款

	case strings.HasPrefix(callbackQuery.Data, "deposit_usdt"):

		transferAmount := callbackQuery.Data[13:len(callbackQuery.Data)]

		fmt.Printf("transferAmount: %s\n", transferAmount)

		usdtPlaceholderRepo := repositories.NewUserUsdtPlaceholdersRepository(db)
		placeholder, _ := usdtPlaceholderRepo.Find(context.Background())

		err := usdtPlaceholderRepo.Update(context.Background(), placeholder.Id, 0)
		if err != nil {
			log.Printf("Error updating usdt placeholder: %v", err)
		}
		realTransferAmount := AddStringsAsFloats(placeholder.Placeholder, transferAmount)

		fmt.Printf("realTransferAmount: %s\n", realTransferAmount)

		//生成订单
		usdtDepositRepo := repositories.NewUserUSDTDepositsRepository(db)

		orderNO := Generate6DigitOrderNo()
		var usdtDeposit domain.UserUSDTDeposits
		usdtDeposit.OrderNO = orderNO
		usdtDeposit.UserID = callbackQuery.Message.Chat.ID
		usdtDeposit.Status = 0
		usdtDeposit.Placeholder = placeholder.Placeholder

		//dictRepo := repositories.NewSysDictionariesRepo(db)
		_agent := os.Getenv("Agent")
		//depositAddress, _ := dictRepo.GetDepositAddress(_agent)
		//_agent := os.Getenv("Agent")
		sysUserRepo := repositories.NewSysUsersRepository(db)
		_, depositAddress, _ := sysUserRepo.Find(context.Background(), _agent)
		usdtDeposit.Address = depositAddress
		usdtDeposit.Amount = realTransferAmount
		usdtDeposit.CreatedAt = time.Now()

		errsg := usdtDepositRepo.Create(context.Background(), &usdtDeposit)
		if errsg != nil {
			log.Printf("Error creating usdtDeposit: %v", errsg)
		}

		msg := tgbotapi.NewMessage(callbackQuery.Message.Chat.ID,
			"<b>"+"订单号：#TOPUP-"+"</b>"+usdtDeposit.OrderNO+"\n"+
				"<b>"+"转账金额："+"</b>"+"<code>"+usdtDeposit.Amount+"</code>"+" usdt （点击即可复制）"+"\n"+
				"<b>"+"转账地址："+"</b>"+"<code>"+usdtDeposit.Address+"</code>"+"（点击即可复制）"+"\n"+
				"<b>"+"充值时间："+"</b>"+Format4Chinesese(usdtDeposit.CreatedAt)+"\n"+
				"<b>"+"⚠️注意："+"</b>"+"\n"+
				"▫️注意小数点 "+usdtDeposit.Amount+" usdt 转错金额不能到账"+"\n"+
				"<b>"+"▫️请在10分钟完成付款，转错金额不能到账。"+"</b>"+"\n"+
				"转账10分钟后没到账及时联系"+"\n")

		inlineKeyboard := tgbotapi.NewInlineKeyboardMarkup(
			tgbotapi.NewInlineKeyboardRow(
				tgbotapi.NewInlineKeyboardButtonData("🕣取消订单", "cancel_order"),
				tgbotapi.NewInlineKeyboardButtonData("⬅️返回个人中心", "back_home"),
			))
		msg.ReplyMarkup = inlineKeyboard
		msg.ParseMode = "HTML"
		//msg.DisableWebPagePreview = true
		bot.Send(msg)

		expiration := 1 * time.Minute // 短时间缓存空值

		//设置用户状态
		cache.Set(strconv.FormatInt(callbackQuery.Message.Chat.ID, 10)+"_order_no", usdtDeposit.OrderNO, expiration)
		//responseText = "你选择了选项 A"
	case strings.HasPrefix(callbackQuery.Data, "deposit_trx"):

		transferAmount := callbackQuery.Data[12:len(callbackQuery.Data)]

		fmt.Printf("transferAmount: %s\n", transferAmount)

		trxPlaceholderRepo := repositories.NewUserTRXPlaceholdersRepository(db)
		placeholder, _ := trxPlaceholderRepo.Find(context.Background())

		//err := trxPlaceholderRepo.Update(context.Background(), placeholder.Id, 1)
		//if err != nil {
		//	log.Printf("Error updating trx placeholder: %v", err)
		//}
		realTransferAmount := AddStringsAsFloats(placeholder.Placeholder, transferAmount)

		fmt.Printf("realTransferAmount: %s\n", realTransferAmount)

		//生成订单
		trxDepositRepo := repositories.NewUserTRXDepositsRepository(db)

		orderNO := Generate6DigitOrderNo()
		var trxDeposit domain.UserTRXDeposits
		trxDeposit.OrderNO = orderNO
		trxDeposit.UserID = callbackQuery.Message.Chat.ID
		trxDeposit.Status = 0
		trxDeposit.Placeholder = placeholder.Placeholder

		//dictRepo := repositories.NewSysDictionariesRepo(db)
		_agent := os.Getenv("Agent")
		//depositAddress, _ := dictRepo.GetDepositAddress(_agent)
		sysUserRepo := repositories.NewSysUsersRepository(db)
		_, depositAddress, _ := sysUserRepo.Find(context.Background(), _agent)
		trxDeposit.Address = depositAddress
		trxDeposit.Amount = realTransferAmount
		trxDeposit.CreatedAt = time.Now()

		errsg := trxDepositRepo.Create(context.Background(), &trxDeposit)
		if errsg != nil {
			log.Printf("Error creating trxDeposit: %v", errsg)
		}

		msg := tgbotapi.NewMessage(callbackQuery.Message.Chat.ID,
			"<b>"+"订单号：#TOPUP-"+"</b>"+trxDeposit.OrderNO+"\n"+
				"<b>"+"转账金额："+"</b>"+"<code>"+trxDeposit.Amount+"</code>"+" TRX （点击即可复制）"+"\n"+
				"<b>"+"转账地址："+"</b>"+"<code>"+trxDeposit.Address+"</code>"+"（点击即可复制）"+"\n"+
				"<b>"+"充值时间："+"</b>"+Format4Chinesese(trxDeposit.CreatedAt)+"\n"+
				"<b>"+"⚠️注意："+"</b>"+"\n"+
				"▫️注意小数点 "+trxDeposit.Amount+" TRX 转错金额不能到账"+"\n"+
				"<b>"+"▫️请在10分钟完成付款，转错金额不能到账。"+"</b>"+"\n"+
				"转账10分钟后没到账及时联系"+"\n")

		inlineKeyboard := tgbotapi.NewInlineKeyboardMarkup(
			tgbotapi.NewInlineKeyboardRow(
				tgbotapi.NewInlineKeyboardButtonData("🕣取消订单", "cancel_order"),
				tgbotapi.NewInlineKeyboardButtonData("⬅️返回个人中心", "back_home"),
			))
		msg.ReplyMarkup = inlineKeyboard
		msg.ParseMode = "HTML"
		//msg.DisableWebPagePreview = true
		bot.Send(msg)
		expiration := 1 * time.Minute // 短时间缓存空值

		//设置用户状态
		cache.Set(strconv.FormatInt(callbackQuery.Message.Chat.ID, 10)+"_order_no", trxDeposit.OrderNO, expiration)
	case callbackQuery.Data == "cancel_order":

		//设置用户状态
		orderNO, _ := cache.Get(strconv.FormatInt(callbackQuery.Message.Chat.ID, 10) + "_order_no")
		msg_order := tgbotapi.NewMessage(callbackQuery.Message.Chat.ID,
			"订单号：#TOPUP-"+orderNO+" 订单已取消")
		msg_order.ParseMode = "HTML"
		//msg.DisableWebPagePreview = true
		bot.Send(msg_order)

		inlineKeyboard := tgbotapi.NewInlineKeyboardMarkup(
			//tgbotapi.NewInlineKeyboardRow(
			//	tgbotapi.NewInlineKeyboardButtonData("🆔我的账户", "click_my_account"),
			//	tgbotapi.NewInlineKeyboardButtonData("💳充值", "click_my_deposit"),
			//),
			tgbotapi.NewInlineKeyboardRow(
				tgbotapi.NewInlineKeyboardButtonData("💳充值", "deposit_amount"),
				tgbotapi.NewInlineKeyboardButtonData("📄账单", "click_my_recepit"),
				tgbotapi.NewInlineKeyboardButtonData("🛠️我的服务", "click_my_service"),
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

	case callbackQuery.Data == "forward_deposit_usdt":
		usdtSubscriptionsRepo := repositories.NewUserUsdtSubscriptionsRepository(db)

		usdtlist, err := usdtSubscriptionsRepo.ListAll(context.Background())

		if err != nil {

		}
		var allButtons []tgbotapi.InlineKeyboardButton
		var extraButtons []tgbotapi.InlineKeyboardButton
		var keyboard [][]tgbotapi.InlineKeyboardButton
		for _, usdtRecord := range usdtlist {
			allButtons = append(allButtons, tgbotapi.NewInlineKeyboardButtonData("🏦"+usdtRecord.Name, "deposit_usdt_"+usdtRecord.Amount))
		}

		extraButtons = append(extraButtons, tgbotapi.NewInlineKeyboardButtonData("⚖️切换到TRX充值", "forward_deposit_usdt"), tgbotapi.NewInlineKeyboardButtonData("🔙返回个人中心", "back_home"))

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
			"💬"+"<b>"+"用户姓名: "+"</b>"+user.Username+"\n"+
				"👤"+"<b>"+"用户电报ID: "+"</b>"+user.Associates+"\n"+
				"💵"+"<b>"+"TRX余额:  "+"</b>"+user.TronAmount+" TRX"+"\n"+
				"💴"+"<b>"+"USDT余额:  "+"</b>"+user.Amount+" USDT")

		msg.ReplyMarkup = inlineKeyboard
		msg.ParseMode = "HTML"

		bot.Send(msg)

	default:
		responseText = "未知选项"
	}

	// 发送新消息作为响应
	bot.Send(tgbotapi.NewMessage(callbackQuery.Message.Chat.ID, responseText))

	// 可以编辑原始内联键盘消息（可选）
	//editMsg := tgbotapi.NewEditMessageText(
	//	callbackQuery.Message.Chat.ID,
	//	callbackQuery.Message.MessageID,
	//	"你已选择: "+callbackQuery.Data,
	//)
	//bot.Send(editMsg)
}
