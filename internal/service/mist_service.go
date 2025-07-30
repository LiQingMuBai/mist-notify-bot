package service

import (
	"context"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"gorm.io/gorm"
	"strings"
	"ushield_bot/internal/handler"
	"ushield_bot/internal/infrastructure/repositories"
	. "ushield_bot/internal/infrastructure/tools"
)

func ExtractSlowMistRiskQuery(message *tgbotapi.Message, db *gorm.DB, _cookie string, bot *tgbotapi.BotAPI) {
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
