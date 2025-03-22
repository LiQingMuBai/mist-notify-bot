package telegram

import (
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	bot "homework_bot/internal/bot"
)

func getCommandMenu() tgbotapi.SetMyCommandsConfig {
	menu := tgbotapi.NewSetMyCommands(
		tgbotapi.BotCommand{
			Command:     bot.CommandStart,
			Description: "开始与机器人聊天",
		},
		tgbotapi.BotCommand{
			Command:     bot.CommandScoreEnergy,
			Description: "地址风险评估",
		},
		tgbotapi.BotCommand{
			Command:     bot.GET_TODAY_FROZEN_ADDRESSES,
			Description: "统计今日冻结地址列表",
		},
		tgbotapi.BotCommand{
			Command:     bot.GET_PENDING_FROZEN_ADDRESSES,
			Description: "统计即将冻结地址列表",
		},
		tgbotapi.BotCommand{
			Command:     bot.GET_HISTORICAL_STATS,
			Description: "历史统计信息",
		},
		tgbotapi.BotCommand{
			Command:     bot.ASSOCIATED_RECOMMENDATION_RELATION,
			Description: "绑定推荐关系",
		},
		//tgbotapi.BotCommand{
		//	Command:     bot.CommandAskGroup,
		//	Description: "Задать группу",
		//},
		//tgbotapi.BotCommand{
		//	Command:     bot.CommandScheduleWeek,
		//	Description: "Расписание на неделю",
		//},
		//tgbotapi.BotCommand{
		//	Command:     bot.CommandScheduleToday,
		//	Description: "Расписание на cегодня",
		//},
		//tgbotapi.BotCommand{
		//	Command:     bot.CommandScheduleTomorrow,
		//	Description: "Расписание на завтра",
		//},
		//tgbotapi.BotCommand{
		//	Command:     bot.CommandScheduleDate,
		//	Description: "Расписание на день",
		//},
		//tgbotapi.BotCommand{
		//	Command:     bot.CommandAdd,
		//	Description: "Добавить новую запись",
		//},
		//tgbotapi.BotCommand{
		//	Command:     bot.CommandUpdate,
		//	Description: "Обновить запись",
		//},
		//tgbotapi.BotCommand{
		//	Command:     bot.CommandDelete,
		//	Description: "Удалить запись",
		//},
		//tgbotapi.BotCommand{
		//	Command:     bot.CommandGetAll,
		//	Description: "Всё дз",
		//},
		//tgbotapi.BotCommand{
		//	Command:     bot.CommandGetOnId,
		//	Description: "Получить дз по id",
		//},
		//tgbotapi.BotCommand{
		//	Command:     bot.CommandGetOnDate,
		//	Description: "Дз на дату",
		//},
		//tgbotapi.BotCommand{
		//	Command:     bot.CommandGetOnToday,
		//	Description: "Дз на сегодня",
		//},
		//tgbotapi.BotCommand{
		//	Command:     bot.CommandGetOnTomorrow,
		//	Description: "Дз на завтра",
		//},
		//tgbotapi.BotCommand{
		//	Command:     bot.CommandGetOnWeek,
		//	Description: "Дз на неделю",
		//},
		//tgbotapi.BotCommand{
		//	Command:     bot.CommandHelp,
		//	Description: "Инструкция",
		//},
		//tgbotapi.BotCommand{
		//	Command:     bot.CommandScheduleNextWeek,
		//	Description: "Расписание на след. неделю",
		//},
		//tgbotapi.BotCommand{
		//	Command:     bot.CommandExchangeEnergy,
		//	Description: "波场Gas兑换",
		//},

		tgbotapi.BotCommand{
			Command:     bot.CommandHelp,
			Description: "客服",
		},
	)
	return menu
}
