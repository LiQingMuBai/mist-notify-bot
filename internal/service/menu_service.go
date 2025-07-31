package service

import (
	"context"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"gorm.io/gorm"
	"os"
	"strconv"
	"time"
	"ushield_bot/internal/cache"
	"ushield_bot/internal/infrastructure/repositories"
	. "ushield_bot/internal/infrastructure/tools"
)

func MenuNavigateAddressFreeze(cache cache.Cache, bot *tgbotapi.BotAPI, message *tgbotapi.Message) {
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
		tgbotapi.NewInlineKeyboardRow(
			tgbotapi.NewInlineKeyboardButtonData("冻结预警扣款记录", "address_freeze_risk_records"),
			//tgbotapi.NewInlineKeyboardButtonData("第二紧急通知", ""),
		),
	)
	msg.ReplyMarkup = inlineKeyboard

	bot.Send(msg)

	expiration := 1 * time.Minute // 短时间缓存空值

	//设置用户状态
	cache.Set(strconv.FormatInt(message.Chat.ID, 10), "usdt_risk_monitor", expiration)
}

func MenuNavigateAddressDetection(cache cache.Cache, bot *tgbotapi.BotAPI, message *tgbotapi.Message, db *gorm.DB) {
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
			tgbotapi.NewInlineKeyboardButtonData("💴地址检测扣款记录", "user_detection_cost_records"),
		),
	)
	msg.ReplyMarkup = inlineKeyboard

	bot.Send(msg)

	expiration := 1 * time.Minute // 短时间缓存空值

	//设置用户状态
	cache.Set(strconv.FormatInt(message.Chat.ID, 10), "usdt_risk_query", expiration)
}

func MenuNavigateEnergyExchange(db *gorm.DB, message *tgbotapi.Message, bot *tgbotapi.BotAPI) {
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
}
func MenuNavigateBundlePackage(db *gorm.DB, message *tgbotapi.Message, bot *tgbotapi.BotAPI) {
	bundlesRepo := repositories.NewUserOperationBundlesRepository(db)

	trxlist, err := bundlesRepo.ListAll(context.Background())

	if err != nil {

	}

	var allButtons []tgbotapi.InlineKeyboardButton
	var extraButtons []tgbotapi.InlineKeyboardButton
	var keyboard [][]tgbotapi.InlineKeyboardButton
	for _, trx := range trxlist {
		allButtons = append(allButtons, tgbotapi.NewInlineKeyboardButtonData("👝"+trx.Name, "bundle_"+trx.Amount))
	}

	extraButtons = append(extraButtons, tgbotapi.NewInlineKeyboardButtonData("笔数套餐扣款记录", "click_bundle_package_cost_records"))

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
}

func MenuNavigateHome(db *gorm.DB, message *tgbotapi.Message, bot *tgbotapi.BotAPI) {
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
}
