package telegram

import (
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	bot "homework_bot/internal/bot"
)

func getCommandMenu() tgbotapi.SetMyCommandsConfig {
	menu := tgbotapi.NewSetMyCommands(
		tgbotapi.BotCommand{
			Command:     bot.CommandStart,
			Description: "開始與機器人聊天",
		},
		tgbotapi.BotCommand{
			Command:     bot.CommandScoreEnergy,
			Description: "USDT地址風險查詢",
		},
		tgbotapi.BotCommand{
			Command:     bot.ADDRESS_BEHAVIOR_REPORT,
			Description: "地址行爲分析報告",
		},
		tgbotapi.BotCommand{
			Command:     bot.MONITOR_ADDRESS,
			Description: "跟踪地址",
		},
		tgbotapi.BotCommand{
			Command:     bot.GET_TODAY_FROZEN_ADDRESSES,
			Description: "統計今日凍結地址列表",
		},
		tgbotapi.BotCommand{
			Command:     bot.GET_PENDING_FROZEN_ADDRESSES,
			Description: "統計即將凍結地址列表",
		},
		tgbotapi.BotCommand{
			Command:     bot.GET_HISTORICAL_STATS,
			Description: "歷史統計信息",
		},
		tgbotapi.BotCommand{
			Command:     bot.GET_VIP,
			Description: "昇級vip用戶",
		},
		tgbotapi.BotCommand{
			Command:     bot.CommandExchangeEnergy,
			Description: "能量兑换",
		},
		tgbotapi.BotCommand{
			Command:     bot.CommandGetAccount,
			Description: "账户信息",
		},

		tgbotapi.BotCommand{
			Command:     bot.CommandHelp,
			Description: "客服",
		},
	)
	return menu
}
