package service

import (
	"context"
	"fmt"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"gorm.io/gorm"
	"log"
	"strconv"
	"strings"
	"time"
	"ushield_bot/internal/cache"
	"ushield_bot/internal/domain"
	. "ushield_bot/internal/global"
	"ushield_bot/internal/infrastructure/repositories"
	. "ushield_bot/internal/infrastructure/tools"
	"ushield_bot/internal/request"
)

func CLICK_BUNDLE_PACKAGE_ADDRESS_MANAGER_REMOVE(cache cache.Cache, bot *tgbotapi.BotAPI, message *tgbotapi.Message, db *gorm.DB) bool {
	if !IsValidAddress(message.Text) {
		msg := tgbotapi.NewMessage(message.Chat.ID, "💬"+"<b>"+"地址有误，请重新输入地址: "+"</b>"+"\n")
		msg.ParseMode = "HTML"
		bot.Send(msg)
		return true
	}

	userOperationPackageAddressesRepo := repositories.NewUserOperationPackageAddressesRepo(db)

	var record domain.UserOperationPackageAddresses
	record.Status = 0
	record.Address = message.Text
	record.ChatID = message.Chat.ID

	errsg := userOperationPackageAddressesRepo.Remove(context.Background(), message.Chat.ID, message.Text)
	if errsg != nil {
		log.Printf("errsg: %s", errsg)
		return true
	}
	msg := tgbotapi.NewMessage(message.Chat.ID, "✅"+"<b>"+"地址删除成功 "+"</b>"+"\n")
	msg.ParseMode = "HTML"
	bot.Send(msg)
	CLICK_BUNDLE_PACKAGE_ADDRESS_MANAGEMENT(cache, bot, message.Chat.ID, db)
	return false
}
func CLICK_BUNDLE_PACKAGE_ADDRESS_MANAGER_ADD(cache cache.Cache, bot *tgbotapi.BotAPI, message *tgbotapi.Message, db *gorm.DB) bool {
	if !IsValidAddress(message.Text) {
		msg := tgbotapi.NewMessage(message.Chat.ID, "💬"+"<b>"+"地址有误，请重新输入地址: "+"</b>"+"\n")
		msg.ParseMode = "HTML"
		bot.Send(msg)
		return true
	}

	userOperationPackageAddressesRepo := repositories.NewUserOperationPackageAddressesRepo(db)

	var record domain.UserOperationPackageAddresses
	record.Status = 0
	record.Address = message.Text
	record.ChatID = message.Chat.ID

	errsg := userOperationPackageAddressesRepo.Create(context.Background(), &record)
	if errsg != nil {
		log.Printf("errsg: %s", errsg)
		return true
	}
	msg := tgbotapi.NewMessage(message.Chat.ID, "✅"+"<b>"+"地址添加成功 "+"</b>"+"\n")
	msg.ParseMode = "HTML"
	bot.Send(msg)
	CLICK_BUNDLE_PACKAGE_ADDRESS_MANAGEMENT(cache, bot, message.Chat.ID, db)
	return false
}

func APPLY_BUNDLE_PACKAGE(cache cache.Cache, bot *tgbotapi.BotAPI, message *tgbotapi.Message, db *gorm.DB, status string) bool {
	if !IsValidAddress(message.Text) {
		msg := tgbotapi.NewMessage(message.Chat.ID, "💬"+"<b>"+"地址有误，请重新输入地址: "+"</b>"+"\n")
		msg.ParseMode = "HTML"
		bot.Send(msg)
		return true
	}

	bundleID := strings.ReplaceAll(status, "apply_bundle_package_", "")
	userOperationBundlesRepo := repositories.NewUserOperationBundlesRepository(db)
	bundlePackage, err := userOperationBundlesRepo.Query(context.Background(), bundleID)

	if err != nil {
		fmt.Println(err)
	}
	userRepo := repositories.NewUserRepository(db)
	user, _ := userRepo.GetByUserID(message.Chat.ID)

	//扣錢
	if bundlePackage.Token == "TRX" {
		balance, _ := SubtractStringNumbers(user.TronAmount, bundlePackage.Amount, 1)
		fmt.Printf("TRX balance %s", balance)
		user.TronAmount = balance
	} else if bundlePackage.Token == "USDT" {
		balance, _ := SubtractStringNumbers(user.Amount, bundlePackage.Amount, 1)
		fmt.Printf("USDT balance %s", balance)

		user.Amount = balance
	}

	err = userRepo.Update2(context.Background(), &user)
	if err != nil {

	}

	//加入訂閲記錄
	userPackageSubscriptionsRepo := repositories.NewUserPackageSubscriptionsRepository(db)
	var record domain.UserPackageSubscriptions
	record.ChatID = message.Chat.ID
	record.Address = message.Text
	bundle, _ := strconv.ParseInt(bundleID, 10, 64)

	record.BundleID = bundle
	record.Status = 2
	record.Amount = bundlePackage.Amount
	record.Times = ExtractLeadingInt64(bundlePackage.Name)
	record.BundleName = bundlePackage.Name

	err = userPackageSubscriptionsRepo.Create(context.Background(), &record)
	if err != nil {
		return true
	}
	msg := tgbotapi.NewMessage(message.Chat.ID, "✅"+"🧾笔数套餐订单购买成功\n\n"+
		"套餐："+bundlePackage.Name+"\n\n"+
		"支付金额："+bundlePackage.Amount+" "+bundlePackage.Token+"\n\n"+
		"地址："+message.Text+"\n\n"+
		"订单号："+fmt.Sprintf("%d", record.Id)+""+"\n\n")
	msg.ParseMode = "HTML"
	// 当点击"按钮 1"时显示内联键盘
	inlineKeyboard := tgbotapi.NewInlineKeyboardMarkup(
		tgbotapi.NewInlineKeyboardRow(
			tgbotapi.NewInlineKeyboardButtonData("🧾地址列表", "click_bundle_package_address_stats"),
			tgbotapi.NewInlineKeyboardButtonData("🔙️返回首页", "back_bundle_package"),
		),
	)
	msg.ReplyMarkup = inlineKeyboard

	bot.Send(msg)

	expiration := 1 * time.Minute // 短时间缓存空值

	//设置用户状态
	cache.Set(strconv.FormatInt(message.Chat.ID, 10), "null_apply_bundle_package_address", expiration)
	return false
}
func CLICK_BUNDLE_PACKAGE_ADDRESS_STATS(db *gorm.DB, chatID int64) tgbotapi.MessageConfig {

	//fmt.Println("ExtractBundlePackage")
	userAddressDetectionRepo := repositories.NewUserPackageSubscriptionsRepository(db)
	var info request.UserAddressDetectionSearch

	info.Page = 1
	info.PageSize = 5
	orderlist, total, err := userAddressDetectionRepo.GetUserPackageSubscriptionsInfoList(context.Background(), info, chatID)
	if err != nil {

		fmt.Println("能量笔数套餐空", err)
	}
	var builder strings.Builder
	if total > 0 {
		//- [6.29] +3000 TRX（订单 #TOPUP-92308）
		for _, order := range orderlist {
			builder.WriteString("地址：")
			builder.WriteString("<code>" + order.Address + "</code>")
			builder.WriteString("\n")
			builder.WriteString("状态：")
			//0默认初始化状态  1 自动派送 2 手动 3 结束
			if order.Status == 3 {
				builder.WriteString("<b>" + "已结束" + "</b>")
			} else if order.Status == 2 {
				builder.WriteString("<b>" + "已停止" + "</b>")
			} else if order.Status == 1 {
				builder.WriteString("<b>" + "已开启" + "</b>")
			} else if order.Status == 0 {
				builder.WriteString("<b>" + "初始化" + "</b>")
			}

			builder.WriteString("\n")

			builder.WriteString("剩余：")
			builder.WriteString(strconv.FormatInt(order.Times, 10))
			builder.WriteString("笔")

			usedTimes := ExtractLeadingInt64(order.BundleName) - order.Times
			builder.WriteString("          已用：")
			builder.WriteString(strconv.FormatInt(usedTimes, 10))
			builder.WriteString("笔")

			//builder.WriteString(" （能量笔数套餐）")

			builder.WriteString("\n\n") // 添加分隔符
			if order.Times > 0 {
				if order.Status == 2 {
					builder.WriteString("开启自动发能： /startAutoDispatch")
					builder.WriteString(strconv.FormatInt(order.Id, 10))
				}
				if order.Status == 1 {
					builder.WriteString("关闭自动发能： /stopAutoDispatch")
					builder.WriteString(strconv.FormatInt(order.Id, 10))
				}
				builder.WriteString("\n") // 添加分隔符
				builder.WriteString("手工发能：/dispatchNow")
				builder.WriteString(strconv.FormatInt(order.Id, 10))
				builder.WriteString("\n") // 添加分隔符
				builder.WriteString("发能其他用户：/dispatchOthers")
				builder.WriteString(strconv.FormatInt(order.Id, 10))
				builder.WriteString("\n") // 添加分隔符
			}
			builder.WriteString("\n")
			builder.WriteString("➖➖➖➖➖➖➖➖➖➖➖➖➖➖➖") // 添加分隔符
			builder.WriteString("\n")              // 添加分隔符
		}
	} else {
		builder.WriteString("\n\n") // 添加分隔符
	}

	// 去除最后一个空格
	result := strings.TrimSpace(builder.String())

	msg := tgbotapi.NewMessage(chatID, "🧾<b>转账笔数 地址列表：</b>\n\n "+
		result+"\n")
	msg.ParseMode = "HTML"
	inlineKeyboard := tgbotapi.NewInlineKeyboardMarkup(
		tgbotapi.NewInlineKeyboardRow(
			tgbotapi.NewInlineKeyboardButtonData("上一页", "next_bundle_package_address_stats"),
			tgbotapi.NewInlineKeyboardButtonData("下一页", "prev_bundle_package_address_stats"),
		),
		tgbotapi.NewInlineKeyboardRow(
			//tgbotapi.NewInlineKeyboardButtonData("解绑地址", "free_monitor_address"),
			tgbotapi.NewInlineKeyboardButtonData("🔙️返回首页", "back_bundle_package"),
		),
	)
	msg.ReplyMarkup = inlineKeyboard
	return msg
}
func NEXT_BUNDLE_PACKAGE_ADDRESS_STATS(callbackQuery *tgbotapi.CallbackQuery, db *gorm.DB, bot *tgbotapi.BotAPI) bool {
	state := DepositStates[callbackQuery.Message.Chat.ID]
	if state == nil {
		var state2 DepositState
		state2.CurrentPage = 1
		state = &state2
	}
	//if state != nil && state.CurrentPage > 1 {
	state.CurrentPage = state.CurrentPage + 1
	userAddressDetectionRepo := repositories.NewUserPackageSubscriptionsRepository(db)
	var info request.UserAddressDetectionSearch
	info.PageInfo.Page = state.CurrentPage
	info.PageInfo.PageSize = 10
	orderlist, total, _ := userAddressDetectionRepo.GetUserPackageSubscriptionsInfoList(context.Background(), info, callbackQuery.Message.Chat.ID)

	fmt.Printf("currentpage : %d", state.CurrentPage)
	fmt.Printf("total: %v\n", total)
	totalPages := (total + 5 - 1) / 5

	fmt.Printf("totalPages : %d", totalPages)
	if int64(state.CurrentPage) > totalPages {
		state.CurrentPage = totalPages
		return true
	}
	var builder strings.Builder
	if total > 0 {
		//- [6.29] +3000 TRX（订单 #TOPUP-92308）
		for _, order := range orderlist {
			builder.WriteString("地址：")
			builder.WriteString("<code>" + order.Address + "</code>")
			builder.WriteString("\n")
			builder.WriteString("状态：")
			//0默认初始化状态  1 自动派送 2 手动 3 结束
			if order.Status == 3 {
				builder.WriteString("<b>" + "已结束" + "</b>")
			} else if order.Status == 2 {
				builder.WriteString("<b>" + "已停止" + "</b>")
			} else if order.Status == 1 {
				builder.WriteString("<b>" + "已开启" + "</b>")
			} else if order.Status == 0 {
				builder.WriteString("<b>" + "初始化" + "</b>")
			}

			builder.WriteString("\n")

			builder.WriteString("剩余：")
			builder.WriteString(strconv.FormatInt(order.Times, 10))
			builder.WriteString("笔")

			usedTimes := ExtractLeadingInt64(order.BundleName) - order.Times
			builder.WriteString("          已用：")
			builder.WriteString(strconv.FormatInt(usedTimes, 10))
			builder.WriteString("笔")

			//builder.WriteString(" （能量笔数套餐）")

			builder.WriteString("\n\n") // 添加分隔符
			if order.Times > 0 {
				if order.Status == 2 {
					builder.WriteString("开启自动发能： /startAutoDispatch")
					builder.WriteString(strconv.FormatInt(order.Id, 10))
				}
				if order.Status == 1 {
					builder.WriteString("关闭自动发能： /stopAutoDispatch")
					builder.WriteString(strconv.FormatInt(order.Id, 10))
				}
				builder.WriteString("\n") // 添加分隔符
				builder.WriteString("手工发能：/dispatchNow")
				builder.WriteString(strconv.FormatInt(order.Id, 10))

				builder.WriteString("\n") // 添加分隔符
				builder.WriteString("发能其他用户：/dispatchOthers")
				builder.WriteString(strconv.FormatInt(order.Id, 10))
				builder.WriteString("\n") // 添加分隔符
			}
			builder.WriteString("\n")
			builder.WriteString("➖➖➖➖➖➖➖➖➖➖➖➖➖➖➖") // 添加分隔符
			builder.WriteString("\n")              // 添加分隔符
		}
	} else {
		builder.WriteString("\n\n") // 添加分隔符
	}

	// 去除最后一个空格
	result := strings.TrimSpace(builder.String())
	msg := tgbotapi.NewMessage(callbackQuery.Message.Chat.ID, "🧾<b>转账笔数 地址列表：</b>\n\n "+
		result+"\n")
	msg.ParseMode = "HTML"
	inlineKeyboard := tgbotapi.NewInlineKeyboardMarkup(
		tgbotapi.NewInlineKeyboardRow(
			tgbotapi.NewInlineKeyboardButtonData("上一页", "next_bundle_package_address_stats"),
			tgbotapi.NewInlineKeyboardButtonData("下一页", "prev_bundle_package_address_stats"),
		),
		tgbotapi.NewInlineKeyboardRow(
			//tgbotapi.NewInlineKeyboardButtonData("解绑地址", "free_monitor_address"),
			tgbotapi.NewInlineKeyboardButtonData("🔙️返回首页", "back_bundle_package"),
		),
	)
	msg.ReplyMarkup = inlineKeyboard
	bot.Send(msg)
	//}
	fmt.Printf("state: %v\n", state)

	DepositStates[callbackQuery.Message.Chat.ID] = state
	return false
}

func PREV_BUNDLE_PACKAGE_ADDRESS_STATS(callbackQuery *tgbotapi.CallbackQuery, db *gorm.DB, bot *tgbotapi.BotAPI) (*DepositState, bool) {
	state := DepositStates[callbackQuery.Message.Chat.ID]

	if state != nil && state.CurrentPage == 1 {
		return nil, true
	}
	if state == nil {
		var state DepositState
		state.CurrentPage = 1
		DepositStates[callbackQuery.Message.Chat.ID] = &state
		userAddressDetectionRepo := repositories.NewUserPackageSubscriptionsRepository(db)
		var info request.UserAddressDetectionSearch
		info.PageInfo.Page = 1
		info.PageInfo.PageSize = 10
		orderlist, total, _ := userAddressDetectionRepo.GetUserPackageSubscriptionsInfoList(context.Background(), info, callbackQuery.Message.Chat.ID)
		var builder strings.Builder
		if total > 0 {
			//- [6.29] +3000 TRX（订单 #TOPUP-92308）
			for _, order := range orderlist {
				builder.WriteString("地址：")
				builder.WriteString("<code>" + order.Address + "</code>")
				builder.WriteString("\n")
				builder.WriteString("状态：")
				//0默认初始化状态  1 自动派送 2 手动 3 结束
				if order.Status == 3 {
					builder.WriteString("<b>" + "已结束" + "</b>")
				} else if order.Status == 2 {
					builder.WriteString("<b>" + "已停止" + "</b>")
				} else if order.Status == 1 {
					builder.WriteString("<b>" + "已开启" + "</b>")
				} else if order.Status == 0 {
					builder.WriteString("<b>" + "初始化" + "</b>")
				}

				builder.WriteString("\n")

				builder.WriteString("剩余：")
				builder.WriteString(strconv.FormatInt(order.Times, 10))
				builder.WriteString("笔")

				usedTimes := ExtractLeadingInt64(order.BundleName) - order.Times
				builder.WriteString("          已用：")
				builder.WriteString(strconv.FormatInt(usedTimes, 10))
				builder.WriteString("笔")

				//builder.WriteString(" （能量笔数套餐）")

				builder.WriteString("\n\n") // 添加分隔符
				if order.Times > 0 {
					if order.Status == 2 {
						builder.WriteString("开启自动发能： /startAutoDispatch")
						builder.WriteString(strconv.FormatInt(order.Id, 10))
					}
					if order.Status == 1 {
						builder.WriteString("关闭自动发能： /stopAutoDispatch")
						builder.WriteString(strconv.FormatInt(order.Id, 10))
					}
					builder.WriteString("\n") // 添加分隔符
					builder.WriteString("手工发能：/dispatchNow")
					builder.WriteString(strconv.FormatInt(order.Id, 10))
					builder.WriteString("\n") // 添加分隔符
					builder.WriteString("发能其他用户：/dispatchOthers")
					builder.WriteString(strconv.FormatInt(order.Id, 10))
					builder.WriteString("\n") // 添加分隔符
				}
				builder.WriteString("\n")
				builder.WriteString("➖➖➖➖➖➖➖➖➖➖➖➖➖➖➖") // 添加分隔符
				builder.WriteString("\n")              // 添加分隔符
			}
		} else {
			builder.WriteString("\n\n") // 添加分隔符
		}

		// 去除最后一个空格
		result := strings.TrimSpace(builder.String())
		msg := tgbotapi.NewMessage(callbackQuery.Message.Chat.ID, "🧾<b>转账笔数 地址列表：</b>\n\n "+
			result+"\n")
		msg.ParseMode = "HTML"
		inlineKeyboard := tgbotapi.NewInlineKeyboardMarkup(
			tgbotapi.NewInlineKeyboardRow(
				tgbotapi.NewInlineKeyboardButtonData("上一页", "next_bundle_package_address_stats"),
				tgbotapi.NewInlineKeyboardButtonData("下一页", "prev_bundle_package_address_stats"),
			),
			tgbotapi.NewInlineKeyboardRow(
				//tgbotapi.NewInlineKeyboardButtonData("解绑地址", "free_monitor_address"),
				tgbotapi.NewInlineKeyboardButtonData("🔙️返回首页", "back_bundle_package"),
			),
		)
		msg.ReplyMarkup = inlineKeyboard
		bot.Send(msg)
	} else {
		state.CurrentPage = state.CurrentPage - 1
		userAddressDetectionRepo := repositories.NewUserPackageSubscriptionsRepository(db)
		var info request.UserAddressDetectionSearch
		info.PageInfo.Page = state.CurrentPage
		info.PageInfo.PageSize = 10
		orderlist, total, _ := userAddressDetectionRepo.GetUserPackageSubscriptionsInfoList(context.Background(), info, callbackQuery.Message.Chat.ID)
		var builder strings.Builder
		if total > 0 {
			//- [6.29] +3000 TRX（订单 #TOPUP-92308）
			for _, order := range orderlist {
				builder.WriteString("地址：")
				builder.WriteString("<code>" + order.Address + "</code>")
				builder.WriteString("\n")
				builder.WriteString("状态：")
				//0默认初始化状态  1 自动派送 2 手动 3 结束
				if order.Status == 3 {
					builder.WriteString("<b>" + "已结束" + "</b>")
				} else if order.Status == 2 {
					builder.WriteString("<b>" + "已停止" + "</b>")
				} else if order.Status == 1 {
					builder.WriteString("<b>" + "已开启" + "</b>")
				} else if order.Status == 0 {
					builder.WriteString("<b>" + "初始化" + "</b>")
				}

				builder.WriteString("\n")

				builder.WriteString("剩余：")
				builder.WriteString(strconv.FormatInt(order.Times, 10))
				builder.WriteString("笔")

				usedTimes := ExtractLeadingInt64(order.BundleName) - order.Times
				builder.WriteString("          已用：")
				builder.WriteString(strconv.FormatInt(usedTimes, 10))
				builder.WriteString("笔")

				//builder.WriteString(" （能量笔数套餐）")

				builder.WriteString("\n\n") // 添加分隔符
				if order.Times > 0 {
					if order.Status == 2 {
						builder.WriteString("开启自动发能： /startAutoDispatch")
						builder.WriteString(strconv.FormatInt(order.Id, 10))
					}
					if order.Status == 1 {
						builder.WriteString("关闭自动发能： /stopAutoDispatch")
						builder.WriteString(strconv.FormatInt(order.Id, 10))
					}
					builder.WriteString("\n") // 添加分隔符
					builder.WriteString("手工发能：/dispatchNow")
					builder.WriteString(strconv.FormatInt(order.Id, 10))
					builder.WriteString("\n") // 添加分隔符
					builder.WriteString("发能其他用户：/dispatchOthers")
					builder.WriteString(strconv.FormatInt(order.Id, 10))
					builder.WriteString("\n") // 添加分隔符
				}
				builder.WriteString("\n")
				builder.WriteString("➖➖➖➖➖➖➖➖➖➖➖➖➖➖➖") // 添加分隔符
				builder.WriteString("\n")              // 添加分隔符
			}
		} else {
			builder.WriteString("\n\n") // 添加分隔符
		}

		// 去除最后一个空格
		result := strings.TrimSpace(builder.String())
		msg := tgbotapi.NewMessage(callbackQuery.Message.Chat.ID, "🧾<b>转账笔数 地址列表：</b>\n\n "+
			result+"\n")
		msg.ParseMode = "HTML"
		inlineKeyboard := tgbotapi.NewInlineKeyboardMarkup(
			tgbotapi.NewInlineKeyboardRow(
				tgbotapi.NewInlineKeyboardButtonData("上一页", "next_bundle_package_address_stats"),
				tgbotapi.NewInlineKeyboardButtonData("下一页", "prev_bundle_package_address_stats"),
			),
			tgbotapi.NewInlineKeyboardRow(
				//tgbotapi.NewInlineKeyboardButtonData("解绑地址", "free_monitor_address"),
				tgbotapi.NewInlineKeyboardButtonData("🔙️返回首页", "back_bundle_package"),
			),
		)
		msg.ReplyMarkup = inlineKeyboard
		bot.Send(msg)
	}
	return state, false
}
