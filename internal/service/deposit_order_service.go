package service

import (
	"context"
	"fmt"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"gorm.io/gorm"
	"log"
	"os"
	"strconv"
	"strings"
	"time"
	"ushield_bot/internal/cache"
	"ushield_bot/internal/domain"
	"ushield_bot/internal/infrastructure/repositories"
	. "ushield_bot/internal/infrastructure/tools"
)

func DepositPrevUSDTOrder(cache cache.Cache, bot *tgbotapi.BotAPI, callbackQuery *tgbotapi.CallbackQuery, db *gorm.DB) {
	transferAmount := callbackQuery.Data[13:len(callbackQuery.Data)]

	fmt.Printf("transferAmount: %s\n", transferAmount)

	usdtPlaceholderRepo := repositories.NewUserUsdtPlaceholdersRepository(db)
	placeholder, esg := usdtPlaceholderRepo.Query(context.Background())

	//err := trxPlaceholderRepo.Update(context.Background(), placeholder.Id, 1)
	if esg != nil {
		fmt.Printf("Failed to update user: " + esg.Error())
		msg := tgbotapi.NewMessage(callbackQuery.Message.Chat.ID,
			"由于波场(TRON)网络出现不稳定情况，可能导致交易延迟或失败。"+
				"为保障用户资产安全，我们决定暂时关闭波场(TRON)网络的充值通道，待网络稳定后重新开放。"+
				"\n✅ 其他功能：预警、检测、笔数套餐等业务均正常运作，不受影响。\n"+
				"建议：\n🔹 如需充值，请等待10分钟后再尝试。\n\n"+
				"我们正在密切关注波场(TRON)网络情况，由此带来的不便，敬请谅解！")

		inlineKeyboard := tgbotapi.NewInlineKeyboardMarkup(
			tgbotapi.NewInlineKeyboardRow(
				tgbotapi.NewInlineKeyboardButtonData("🕣取消订单", "cancel_order"),
				tgbotapi.NewInlineKeyboardButtonData("🔙返回个人中心", "back_home"),
			))
		msg.ReplyMarkup = inlineKeyboard
		msg.ParseMode = "HTML"
		//msg.DisableWebPagePreview = true
		bot.Send(msg)
		return

	}
	if placeholder.Id == 0 {
		msg := tgbotapi.NewMessage(callbackQuery.Message.Chat.ID,
			"由于波场(TRON)网络出现不稳定情况，可能导致交易延迟或失败。"+
				"为保障用户资产安全，我们决定暂时关闭波场(TRON)网络的充值通道，待网络稳定后重新开放。"+
				"\n✅ 其他功能：预警、检测、笔数套餐等业务均正常运作，不受影响。\n"+
				"建议：\n🔹 如需充值，请等待10分钟后再尝试。\n\n"+
				"我们正在密切关注波场(TRON)网络情况，由此带来的不便，敬请谅解！")

		inlineKeyboard := tgbotapi.NewInlineKeyboardMarkup(
			tgbotapi.NewInlineKeyboardRow(
				tgbotapi.NewInlineKeyboardButtonData("🕣取消订单", "cancel_order"),
				tgbotapi.NewInlineKeyboardButtonData("🔙返回个人中心", "back_home"),
			))
		msg.ReplyMarkup = inlineKeyboard
		msg.ParseMode = "HTML"
		//msg.DisableWebPagePreview = true
		bot.Send(msg)

		return
	}

	err := usdtPlaceholderRepo.Update(context.Background(), placeholder.Id, 1)
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
	usdtDeposit.Amount = transferAmount
	usdtDeposit.CreatedAt = time.Now()

	errsg := usdtDepositRepo.Create(context.Background(), &usdtDeposit)
	if errsg != nil {
		log.Printf("Error creating usdtDeposit: %v", errsg)
	}

	msg := tgbotapi.NewMessage(callbackQuery.Message.Chat.ID,

		"支付金额："+"<code>"+realTransferAmount+"</code>"+" usdt （点击复制）"+"\n"+
			"收款地址："+"<code>"+usdtDeposit.Address+"</code>"+"（点击复制）"+"\n"+
			"订单号：#TOPUP-"+usdtDeposit.OrderNO+"\n"+
			"有效期：10 分钟"+"\n"+
			"充值时间："+Format4Chinesese(usdtDeposit.CreatedAt)+"\n"+
			"⚠️ 系统会自动为订单金额添加识别尾数，请务必输入完整金额，否则无法入账！"+"\n")
	//"⚠️注意："+"\n"+
	//"▫️注意小数点 "+realTransferAmount+" usdt 转错金额不能到账"+"\n"+
	//"▫️请在10分钟完成付款，转错金额不能到账。"+"\n"+
	//"转账10分钟后没到账及时联系"+"\n")

	inlineKeyboard := tgbotapi.NewInlineKeyboardMarkup(
		tgbotapi.NewInlineKeyboardRow(
			tgbotapi.NewInlineKeyboardButtonData("🕣取消订单", "cancel_order"),
			tgbotapi.NewInlineKeyboardButtonData("🔙返回个人中心", "back_home"),
		))
	msg.ReplyMarkup = inlineKeyboard
	msg.ParseMode = "HTML"
	//msg.DisableWebPagePreview = true
	bot.Send(msg)

	expiration := 1 * time.Minute // 短时间缓存空值

	//设置用户状态
	cache.Set(strconv.FormatInt(callbackQuery.Message.Chat.ID, 10)+"_order_no", "USDT_"+usdtDeposit.OrderNO, expiration)
}

func DepositCancelOrder(cache cache.Cache, bot *tgbotapi.BotAPI, callbackQuery *tgbotapi.CallbackQuery, db *gorm.DB) {
	//设置用户状态
	orderNO, _ := cache.Get(strconv.FormatInt(callbackQuery.Message.Chat.ID, 10) + "_order_no")
	msg_order := tgbotapi.NewMessage(callbackQuery.Message.Chat.ID,
		"订单号：#TOPUP-"+orderNO+" 订单已取消")
	msg_order.ParseMode = "HTML"
	//msg.DisableWebPagePreview = true
	bot.Send(msg_order)

	if strings.Contains(orderNO, "TRX_") {

		_orderNO := strings.ReplaceAll(orderNO, "TRX_", "")
		userTRXDepositsRepo := repositories.NewUserTRXDepositsRepository(db)
		record, _ := userTRXDepositsRepo.Query(context.Background(), _orderNO)

		//update
		fmt.Printf("record: %v\n", record)

		userTRXPlaceholdersRepo := repositories.NewUserTRXPlaceholdersRepository(db)
		userTRXPlaceholdersRepo.UpdateByPlaceholder(context.Background(), record.Placeholder, 0)
		fmt.Printf("placeholder重置 %s\n", record.Placeholder)
	}

	if strings.Contains(orderNO, "USDT_") {
		_orderNO := strings.ReplaceAll(orderNO, "USDT_", "")
		userUSDTDepositsRepo := repositories.NewUserUSDTDepositsRepository(db)
		record, _ := userUSDTDepositsRepo.Find(context.Background(), _orderNO)
		//update
		fmt.Printf("record: %v\n", record)
		userUSDTPlaceholdersRepo := repositories.NewUserUsdtPlaceholdersRepository(db)
		userUSDTPlaceholdersRepo.UpdateByPlaceholder(context.Background(), record.Placeholder, 0)
		fmt.Printf("placeholder重置 %s\n", record.Placeholder)
	}

	inlineKeyboard := tgbotapi.NewInlineKeyboardMarkup(
		//tgbotapi.NewInlineKeyboardRow(
		//	tgbotapi.NewInlineKeyboardButtonData("🆔我的账户", "click_my_account"),
		//	tgbotapi.NewInlineKeyboardButtonData("💳充值", "click_my_deposit"),
		//),
		tgbotapi.NewInlineKeyboardRow(
			tgbotapi.NewInlineKeyboardButtonData("💳充值", "deposit_amount"),
			tgbotapi.NewInlineKeyboardButtonData("🔗第二通知人", "click_backup_account"),
			tgbotapi.NewInlineKeyboardButtonData("📄充值账单", "click_my_recepit"),
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
		//id, _ := strconv.ParseInt(user.BackupChatID, 10, 64)
		//backup_user, _ := userRepo.GetByUserID(id)
		str = "🔗 第二通知人：  " + "@" + user.BackupChatID
	} else {
		str = "第二通知人：（无）"
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

func DepositPrevOrder(cache cache.Cache, bot *tgbotapi.BotAPI, callbackQuery *tgbotapi.CallbackQuery, db *gorm.DB) {
	transferAmount := callbackQuery.Data[12:len(callbackQuery.Data)]

	fmt.Printf("transferAmount: %s\n", transferAmount)

	trxPlaceholderRepo := repositories.NewUserTRXPlaceholdersRepository(db)
	placeholder, esg := trxPlaceholderRepo.Query(context.Background())

	//err := trxPlaceholderRepo.Update(context.Background(), placeholder.Id, 1)
	if esg != nil {
		fmt.Printf("Failed to update user: " + esg.Error())
		msg := tgbotapi.NewMessage(callbackQuery.Message.Chat.ID,
			"由于波场(TRON)网络出现不稳定情况，可能导致交易延迟或失败。"+
				"为保障用户资产安全，我们决定暂时关闭波场(TRON)网络的充值通道，待网络稳定后重新开放。"+
				"\n✅ 其他功能：预警、检测、笔数套餐等业务均正常运作，不受影响。\n"+
				"建议：\n🔹 如需充值，请等待10分钟后再尝试。\n\n"+
				"我们正在密切关注波场(TRON)网络情况，由此带来的不便，敬请谅解！")

		inlineKeyboard := tgbotapi.NewInlineKeyboardMarkup(
			tgbotapi.NewInlineKeyboardRow(
				tgbotapi.NewInlineKeyboardButtonData("🕣取消订单", "cancel_order"),
				tgbotapi.NewInlineKeyboardButtonData("🔙返回个人中心", "back_home"),
			))
		msg.ReplyMarkup = inlineKeyboard
		msg.ParseMode = "HTML"
		//msg.DisableWebPagePreview = true
		bot.Send(msg)

		return

	}
	if placeholder.Id == 0 {
		msg := tgbotapi.NewMessage(callbackQuery.Message.Chat.ID,
			"由于波场(TRON)网络出现不稳定情况，可能导致交易延迟或失败。"+
				"为保障用户资产安全，我们决定暂时关闭波场(TRON)网络的充值通道，待网络稳定后重新开放。"+
				"\n✅ 其他功能：预警、检测、笔数套餐等业务均正常运作，不受影响。\n"+
				"建议：\n🔹 如需充值，请等待10分钟后再尝试。\n\n"+
				"我们正在密切关注波场(TRON)网络情况，由此带来的不便，敬请谅解！")

		inlineKeyboard := tgbotapi.NewInlineKeyboardMarkup(
			tgbotapi.NewInlineKeyboardRow(
				tgbotapi.NewInlineKeyboardButtonData("🕣取消订单", "cancel_order"),
				tgbotapi.NewInlineKeyboardButtonData("🔙返回个人中心", "back_home"),
			))
		msg.ReplyMarkup = inlineKeyboard
		msg.ParseMode = "HTML"
		//msg.DisableWebPagePreview = true
		bot.Send(msg)

		return
	}

	err := trxPlaceholderRepo.Update(context.Background(), placeholder.Id, 1)
	if err != nil {
		log.Printf("Error updating trx placeholder: %v", err)
	}
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
	trxDeposit.Amount = transferAmount
	trxDeposit.CreatedAt = time.Now()

	errsg := trxDepositRepo.Create(context.Background(), &trxDeposit)
	if errsg != nil {
		log.Printf("Error creating trxDeposit: %v", errsg)
	}

	//msg := tgbotapi.NewMessage(callbackQuery.Message.Chat.ID,
	//	"订单号：#TOPUP-"+trxDeposit.OrderNO+"\n"+
	//		"转账金额："+"<code>"+realTransferAmount+"</code>"+" TRX （点击即可复制）"+"\n"+
	//		"转账地址："+"<code>"+trxDeposit.Address+"</code>"+"（点击即可复制）"+"\n"+
	//		"充值时间："+Format4Chinesese(trxDeposit.CreatedAt)+"\n"+
	//		"⚠️注意："+"\n"+
	//		"▫️注意小数点 "+realTransferAmount+" TRX 转错金额不能到账"+"\n"+
	//		"▫️请在10分钟完成付款，转错金额不能到账。"+"\n"+
	//		"转账10分钟后没到账及时联系"+"\n")

	msg := tgbotapi.NewMessage(callbackQuery.Message.Chat.ID,

		"支付金额："+"<code>"+realTransferAmount+"</code>"+" usdt （点击复制）"+"\n"+
			"收款地址："+"<code>"+trxDeposit.Address+"</code>"+"（点击复制）"+"\n"+
			"订单号：#TOPUP-"+trxDeposit.OrderNO+"\n"+
			"有效期：10 分钟"+"\n"+
			"充值时间："+Format4Chinesese(trxDeposit.CreatedAt)+"\n"+
			"⚠️ 系统会自动为订单金额添加识别尾数，请务必输入完整金额，否则无法入账！"+"\n")
	//"⚠️注意："+"\n"+
	//"▫️注意小数点 "+realTransferAmount+" usdt 转错金额不能到账"+"\n"+
	//"▫️请在10分钟完成付款，转错金额不能到账。"+"\n"+
	//"转账10分钟后没到账及时联系"+"\n")
	inlineKeyboard := tgbotapi.NewInlineKeyboardMarkup(
		tgbotapi.NewInlineKeyboardRow(
			tgbotapi.NewInlineKeyboardButtonData("🕣取消订单", "cancel_order"),
			tgbotapi.NewInlineKeyboardButtonData("🔙返回个人中心", "back_home"),
		))
	msg.ReplyMarkup = inlineKeyboard
	msg.ParseMode = "HTML"
	//msg.DisableWebPagePreview = true
	bot.Send(msg)
	expiration := 1 * time.Minute // 短时间缓存空值

	//设置用户状态
	cache.Set(strconv.FormatInt(callbackQuery.Message.Chat.ID, 10)+"_order_no", "TRX_"+trxDeposit.OrderNO, expiration)
}
