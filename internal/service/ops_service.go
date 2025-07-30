package service

import (
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"strconv"
	"time"
	"ushield_bot/internal/cache"
)

func ExtraQA(cache cache.Cache, bot *tgbotapi.BotAPI, callbackQuery *tgbotapi.CallbackQuery) {
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
}
