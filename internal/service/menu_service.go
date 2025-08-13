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

func MenuNavigateAddressFreeze(cache cache.Cache, bot *tgbotapi.BotAPI, chatID int64, db *gorm.DB) {

	userRepo := repositories.NewSysDictionariesRepo(db)

	server_trx_price, _ := userRepo.GetDictionaryDetail("server_trx_price")

	server_usdt_price, _ := userRepo.GetDictionaryDetail("server_usdt_price")

	msg := tgbotapi.NewMessage(chatID, "欢迎使用U盾 USDT冻结预警服务\n"+
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
	cache.Set(strconv.FormatInt(chatID, 10), "usdt_risk_monitor", expiration)
}

func MenuNavigateAddressDetection(cache cache.Cache, bot *tgbotapi.BotAPI, chatID int64, db *gorm.DB) {
	userRepo := repositories.NewUserRepository(db)
	user, _ := userRepo.GetByUserID(chatID)

	if IsEmpty(user.Amount) {
		user.Amount = "0.00"
	}

	if IsEmpty(user.TronAmount) {
		user.TronAmount = "0.00"
	}

	dictRepo := repositories.NewSysDictionariesRepo(db)

	address_detection_cost, _ := dictRepo.GetDictionaryDetail("address_detection_cost")
	address_detection_cost_usdt, _ := dictRepo.GetDictionaryDetail("address_detection_cost_usdt")

	msg := tgbotapi.NewMessage(chatID, " 欢迎使用 U盾地址风险检测\n支持 TRON 或 ETH 网络任意地址查询\n系统将基于链上行为、风险标签、关联实体进行评分与分析\n📊 风险等级说明：\n"+
		"🟢低风险(0–30):无异常交易，未关联已知风险实体\n"+
		"🟡中风险(31–70):存在少量高风险交互，对手方不明\n"+
		"🟠高风险(71–90):频繁异常转账，或与恶意地址有关\n"+
		"🔴极高风险(91–100):涉及诈骗、制裁、黑客、洗钱等高风险行为\n\n"+
		"📌 每位用户每天可免费检测 1 次\n"+
		"📌 超出后每次扣除 "+address_detection_cost+"TRX 或 "+address_detection_cost_usdt+"USDT（系统将优先扣除 TRX）\n"+
		"💰当前余额：\n"+
		"- TRX："+user.TronAmount+"\n"+"- USDT："+user.Amount+"\n"+
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
	cache.Set(strconv.FormatInt(chatID, 10), "usdt_risk_query", expiration)
}

func MenuNavigateEnergyExchange(db *gorm.DB, message *tgbotapi.Message, bot *tgbotapi.BotAPI) {
	// 当点击"按钮 1"时显示内联键盘
	inlineKeyboard := tgbotapi.NewInlineKeyboardMarkup(
		tgbotapi.NewInlineKeyboardRow(
			tgbotapi.NewInlineKeyboardButtonData("🖊️笔数套餐", "back_bundle_package"),
		),
	)
	_agent := os.Getenv("Agent")
	sysUserRepo := repositories.NewSysUsersRepository(db)
	receiveAddress, _, _ := sysUserRepo.Find(context.Background(), _agent)

	//dictRepo := repositories.NewSysDictionariesRepo(db)
	//receiveAddress, _ := dictRepo.GetReceiveAddress(_agent)

	dictDetailRepo := repositories.NewSysDictionariesRepo(db)

	energy_cost, _ := dictDetailRepo.GetDictionaryDetail("energy_cost")

	energy_cost_2x, _ := StringMultiply(energy_cost, 2)
	energy_cost_10x, _ := StringMultiply(energy_cost, 10)
	//old_str := "【⚡️能量闪租】\n🔸转账  " + energy_cost + " Trx=  1 笔能量\n🔸转账  " + energy_cost_2x + " Trx=  2 笔能量\n\n单笔 " + energy_cost + " Trx，以此类推，最大10 笔\n" +
	//"1.向无U地址转账，需要双倍能量。\n2.请在1小时内转账，否则过期回收。\n\n🔸闪租能量收款地址:\n"

	//old_str = "【⚡️能量闪租】\n\n 转账 3 TRX，系统自动按原路返还一笔能量，\n 如需向无U地址转账 ，请转账 6 TRX（返还两笔能量）\n\n"

	old_str := "欢迎使用U盾能量闪兑\n🔸转账  " + energy_cost + " Trx=  1 笔能量\n🔸转账  " + energy_cost_2x + " Trx=  2 笔能量\n🔸闪兑收款地址: "
	msg := tgbotapi.NewMessage(message.Chat.ID, old_str+
		"<code>"+receiveAddress+"</code>"+"\n"+
		"➖➖➖➖"+"点击复制"+"➖➖➖➖\n重要提示："+"\n"+
		"1.单笔 "+energy_cost+"Trx，以此类推，一次最大 10笔（"+energy_cost_10x+"TRX，超出不予入账）\n"+
		"2.向无U地址转账，需要购买两笔能量\n"+
		"3.向闪兑地址转账成功后能量将即时按充值地址原路完成闪兑\n"+
		"4.禁止使用交易所钱包提币使用",
	)
	msg.ReplyMarkup = inlineKeyboard
	msg.ParseMode = "HTML"
	//msg.DisableWebPagePreview = true
	bot.Send(msg)
}
func MenuNavigateBundlePackage(db *gorm.DB, _chatID int64, bot *tgbotapi.BotAPI, token string) {
	bundlesRepo := repositories.NewUserOperationBundlesRepository(db)

	trxlist, err := bundlesRepo.ListByToken(context.Background(), token)

	if err != nil {

	}

	var allButtons []tgbotapi.InlineKeyboardButton
	var extraButtons []tgbotapi.InlineKeyboardButton
	var onlyButtons []tgbotapi.InlineKeyboardButton
	var keyboard [][]tgbotapi.InlineKeyboardButton
	for _, trx := range trxlist {

		allButtons = append(allButtons, tgbotapi.NewInlineKeyboardButtonData("👝"+trx.Name, CombineInt64AndString("bundle_", trx.Id)))
	}

	if token == "TRX" {
		onlyButtons = append(onlyButtons,
			tgbotapi.NewInlineKeyboardButtonData("🛠️切换到USDT支付", "click_switch_usdt"),
		)
	}
	if token == "USDT" {
		onlyButtons = append(onlyButtons,
			tgbotapi.NewInlineKeyboardButtonData("🛠️切换到TRX支付", "click_switch_trx"),
		)
	}

	extraButtons = append(extraButtons,
		tgbotapi.NewInlineKeyboardButtonData("🧾地址列表", "click_bundle_package_address_stats"),
		tgbotapi.NewInlineKeyboardButtonData("➕添加地址", "click_bundle_package_address_management"),
		tgbotapi.NewInlineKeyboardButtonData("📜笔数套餐扣款记录", "click_bundle_package_cost_records"),
	)

	for i := 0; i < len(allButtons); i += 2 {
		end := i + 2
		if end > len(allButtons) {
			end = len(allButtons)
		}
		row := allButtons[i:end]
		keyboard = append(keyboard, tgbotapi.NewInlineKeyboardRow(row...))
	}
	for i := 0; i < len(onlyButtons); i += 1 {
		end := i + 1
		if end > len(onlyButtons) {
			end = len(onlyButtons)
		}
		row := onlyButtons[i:end]
		keyboard = append(keyboard, tgbotapi.NewInlineKeyboardRow(row...))
	}

	for i := 0; i < len(extraButtons); i += 2 {
		end := i + 2
		if end > len(extraButtons) {
			end = len(extraButtons)
		}
		row := extraButtons[i:end]
		keyboard = append(keyboard, tgbotapi.NewInlineKeyboardRow(row...))
	}

	// 3. 创建键盘标记
	inlineKeyboard := tgbotapi.NewInlineKeyboardMarkup(keyboard...)

	userRepo := repositories.NewUserRepository(db)
	user, _ := userRepo.GetByUserID(_chatID)
	if IsEmpty(user.Amount) {
		user.Amount = "0.00"
	}

	if IsEmpty(user.TronAmount) {
		user.TronAmount = "0.00"
	}

	msg := tgbotapi.NewMessage(_chatID,
		"欢迎使用 U盾能量笔数套餐\n"+
			"一次购买 · 多地址使用 · 一键发能 · 快捷高效！\n"+
			"⚙️ 功能介绍\n    📍 地址列表\n •最多可同时管理 4 个接收地址\n "+
			"•可随时设置、修改默认地址\n"+
			"⚡️ 发能管理\n "+
			"•自动发能开启后系统会自动检测地址能量余量，不足时自动补充（每次消耗 1 笔），默认关闭，可在“地址列表”中开启/关闭。\n "+
			"•一键发能：可向地址列表中任意地址或自定义地址快速发放 1 笔能量\n"+
			"⏳ 能量有效期 1 小时，过期将自动回收并扣除笔数。\n"+
			"🆔 用户ID: "+user.Associates+"\n"+
			"👤 用户名: @"+user.Username+"\n"+
			"💰 当前余额: "+"\n"+"- TRX：   "+user.TronAmount+"-  USDT："+user.Amount)
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
			tgbotapi.NewInlineKeyboardButtonData("🔗第二通知人", "click_backup_account"),
			tgbotapi.NewInlineKeyboardButtonData("📄充值账单", "click_my_recepit"),
			//tgbotapi.NewInlineKeyboardButtonData("🛠️我的服务", "click_my_service"),
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
		//id, _ := strconv.ParseInt(user.BackupChatID, 10, 64)
		//backup_user, _ := userRepo.GetByUserID(id)
		str = "🔗 第二通知人：  " + "@" + user.BackupChatID
	} else {
		str = "第二通知人：（无）"
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
