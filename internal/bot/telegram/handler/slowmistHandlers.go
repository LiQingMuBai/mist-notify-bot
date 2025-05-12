package handler

import (
	"encoding/json"
	"fmt"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"io"
	"log"
	"net/http"
	"strconv"
	"strings"
	"ushield_bot/internal/bot"
	"ushield_bot/internal/domain"
)

type MisttrackHandler struct{}

func NewMisttrackHandler() *MisttrackHandler {
	return &MisttrackHandler{}
}

func (h *MisttrackHandler) Handle(b bot.IBot, message *tgbotapi.Message) error {

	userName := message.From.UserName
	user, err := b.GetServices().IUserService.GetByUsername(userName)
	msg := domain.MessageToSend{
		ChatId: message.Chat.ID,
		Text:   "系統錯誤，請重新輸入地址",
	}
	if err != nil {
		_ = b.SendMessage(msg, bot.DefaultChannel)
		return nil
	}
	if user.Times == 1 {
		msg = domain.MessageToSend{
			ChatId: message.Chat.ID,
			Text: "🔍普通用戶每日贈送 1 次地址風險查詢\n" +
				"📞聯繫客服 @Ushield001\n",
		}

	} else {
		_message := message.Text
		_text := "系統錯誤，請重新輸入地址，📞聯繫客服 @Ushield001\n"
		if strings.HasPrefix(_message, "0x") && len(_message) == 42 {
			_symbol := "USDT-ERC20"
			_addressInfo := getAddressInfo(_symbol, _message, b.GetCookie())
			_text = getText(_addressInfo)

			//_coin := "ETH"
			addressProfile := getAddressProfile(_symbol, _message, b.GetCookie())
			_text7 := "余額：" + addressProfile.BalanceUsd + "\n"
			//log.Println("余额：", addressProfile.BalanceUsd)
			//log.Println("累计收入：", addressProfile.TotalReceivedUsd)
			//log.Println("累计支出：", addressProfile.TotalSpentUsd)
			//log.Println("首次活跃时间：", addressProfile.FirstTxTime)
			//log.Println("最后活跃时间：", addressProfile.LastTxTime)
			//log.Println("交易次数：", addressProfile.TxCount+"笔")
			_text8 := "累計收入：" + addressProfile.TotalReceivedUsd + "\n"
			_text9 := "累计支出：" + addressProfile.TotalSpentUsd + "\n"
			_text10 := "首次活躍時間：" + addressProfile.FirstTxTime + "\n"
			_text11 := "最後活躍時間：" + addressProfile.LastTxTime + "\n"
			_text12 := "交易次數：" + addressProfile.TxCount + "筆" + "\n"

			_text13 := "📄 详细分析报告 ➜ 50 TRX" + "\n"

			_text99 := "主要交易对手分析：" + "\n"

			_text14 := "每日免费查询剩余：0 次" + "\n"

			_text15 := "超额查询 ➜ 10 TRX / 次" + "\n"
			_text16 := "🛡️ U盾在手，链上无忧！" + "\n"

			_text = _text + _text7 + _text8 + _text9 + _text10 + _text11 + _text12 + _text13 + _text99 + _text14 + _text15 + _text16

		}
		if strings.HasPrefix(_message, "T") && len(_message) == 34 {
			_symbol := "USDT-TRC20"
			_addressInfo := getAddressInfo(_symbol, _message, b.GetCookie())
			_text = getText(_addressInfo)

			addressProfile := getAddressProfile(_symbol, _message, b.GetCookie())
			_text7 := "余額：" + addressProfile.BalanceUsd + "\n"
			//log.Println("余额：", addressProfile.BalanceUsd)
			//log.Println("累计收入：", addressProfile.TotalReceivedUsd)
			//log.Println("累计支出：", addressProfile.TotalSpentUsd)
			//log.Println("首次活跃时间：", addressProfile.FirstTxTime)
			//log.Println("最后活跃时间：", addressProfile.LastTxTime)
			//log.Println("交易次数：", addressProfile.TxCount+"笔")
			_text8 := "累計收入：" + addressProfile.TotalReceivedUsd + "\n"
			_text9 := "累计支出：" + addressProfile.TotalSpentUsd + "\n"
			_text10 := "首次活躍時間：" + addressProfile.FirstTxTime + "\n"
			_text11 := "最後活躍時間：" + addressProfile.LastTxTime + "\n"
			_text12 := "交易次數：" + addressProfile.TxCount + "筆" + "\n"

			_text13 := "📄 详细分析报告 ➜ 50 TRX" + "\n"

			_text99 := "危险交易对手分析：" + "\n"

			lableAddresList := getNotSafeAddress(_symbol, _message, b.GetCookie())

			_text100 := ""
			if len(lableAddresList.GraphDic.NodeList) > 0 {
				for _, data := range lableAddresList.GraphDic.NodeList {
					if strings.Contains(data.Label, "huione") {
						_text100 = _text100 + data.Title[0:5] + "..." + data.Title[29:34] + "\n"
					}
				}
			}

			_text14 := "每日免费查询剩余：0 次" + "\n"

			_text15 := "超额查询 ➜ 10 TRX / 次" + "\n"
			_text16 := "🛡️ U盾在手，链上无忧！" + "\n"

			_text = _text + _text7 + _text8 + _text9 + _text10 + _text11 + _text12 + _text13 + _text99 + _text100 + _text14 + _text15 + _text16

		}
		msg = domain.MessageToSend{
			ChatId: message.Chat.ID,
			Text:   _text,
		}

		err := b.GetServices().IUserService.UpdateTimes(1, userName)

		if err != nil {

			msg = domain.MessageToSend{
				ChatId: message.Chat.ID,
				Text:   "系統錯誤，請重新輸入地址，📞聯繫客服 @Ushield001\n",
			}
		}
	}
	_ = b.SendMessage(msg, bot.DefaultChannel)
	return nil
}

func getNotSafeAddress(_coin string, _address, _cookie string) LableAddresList {
	url := "https://dashboard.misttrack.io/api/v1/address_graph_analysis?coin=" + _coin + "&address=" + _address + "&time_filter="
	req, _ := http.NewRequest("GET", url, nil)

	req.Header.Add("accept", "application/json, text/plain, */*")

	//req.Header.Add("cookie", "_ga=GA1.1.23337514.1742894564; _bl_uid=O8m7m8ksonwa0Ifjgw0erRqd9147; _ga_SGF4VCWFZY=GS1.1.1743393981.8.0.1743393981.0.0.0; detect_token=eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJyYW5kb21fc3RyIjoiMzI0Njk1In0.t5lYLE_oSwyNIJUSWAwxL7YrzXN5Di38sh4Vh9gjyJE; csrftoken=AOzVpYUl0Wdyk2gtoIzUQ5uOUEOxRBSMsqlINKjOh30dCmHX2ajNk8EcwFxrWy6g; sessionid=rn1a71d9nkn3coczdn08ahc00u5mw46i; _ga_40VGDGQFCB=GS1.1.1743393983.12.1.1743394123.0.0.0; _ga_5X5Z4KZ7PC=GS1.1.1743393983.12.1.1743394123.0.0.0")
	req.Header.Add("cookie", _cookie)
	req.Header.Add("language", "EN")

	//req.Header.Add("referer", "https://dashboard.misttrack.io/address/ETH/0xf510e53ef8da4e45ffa59eb554511a7410e5efd3")
	req.Header.Add("user-agent", "Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/134.0.0.0 Safari/537.36")

	res, _ := http.DefaultClient.Do(req)
	defer res.Body.Close()
	body, _ := io.ReadAll(res.Body)

	log.Println(string(body))

	var lableAddresList LableAddresList
	if err := json.Unmarshal(body, &lableAddresList); err != nil { // Parse []byte to go struct pointer
		fmt.Println("Can not unmarshal JSON")
	}
	return lableAddresList
}

type LableAddresList struct {
	Success  bool   `json:"success"`
	Msg      string `json:"msg"`
	GraphDic struct {
		NodeList []struct {
			ID        string `json:"id"`
			Label     string `json:"label"`
			Title     string `json:"title"`
			Layer     int    `json:"layer"`
			Addr      string `json:"addr"`
			Track     string `json:"track"`
			Pid       int    `json:"pid"`
			Color     string `json:"color,omitempty"`
			Shape     string `json:"shape,omitempty"`
			Expanded  bool   `json:"expanded"`
			Malicious int    `json:"malicious,omitempty"`
			Dex       int    `json:"dex"`
		} `json:"node_list"`
		EdgeList []struct {
			From       string   `json:"from"`
			To         string   `json:"to"`
			Label      string   `json:"label"`
			Val        float64  `json:"val"`
			TxHashList []string `json:"tx_hash_list"`
			TxTime     string   `json:"tx_time"`
			Color      struct {
				Color     string `json:"color"`
				Highlight string `json:"highlight"`
			} `json:"color"`
		} `json:"edge_list"`
		TxCount                 int    `json:"tx_count"`
		FirstTxDatetime         string `json:"first_tx_datetime"`
		LatestTxDatetime        string `json:"latest_tx_datetime"`
		AddressFirstTxDatetime  string `json:"address_first_tx_datetime"`
		AddressLatestTxDatetime string `json:"address_latest_tx_datetime"`
	} `json:"graph_dic"`
	AddressFirstTxDatetime  string `json:"address_first_tx_datetime"`
	AddressLatestTxDatetime string `json:"address_latest_tx_datetime"`
}

type AddressProfile struct {
	Success          bool   `json:"success"`
	Msg              string `json:"msg"`
	Balance          string `json:"balance"`
	TxCount          string `json:"tx_count"`
	FirstTxTime      string `json:"first_tx_time"`
	LastTxTime       string `json:"last_tx_time"`
	TotalReceived    string `json:"total_received"`
	TotalSpent       string `json:"total_spent"`
	ReceivedCount    string `json:"received_count"`
	SpentCount       string `json:"spent_count"`
	TotalReceivedUsd string `json:"total_received_usd"`
	TotalSpentUsd    string `json:"total_spent_usd"`
	BalanceUsd       string `json:"balance_usd"`
}

func getAddressInfo(_symbol string, _address, _cookie string) SlowMistAddressInfo {
	url := "https://dashboard.misttrack.io/api/v1/address_risk_analysis?coin=" + _symbol + "&address=" + _address
	req, _ := http.NewRequest("GET", url, nil)

	req.Header.Add("accept", "application/json, text/plain, */*")
	req.Header.Add("cookie", _cookie)
	req.Header.Add("language", "EN")

	req.Header.Add("referer", "https://dashboard.misttrack.io/address/"+_symbol+"/"+_address)
	req.Header.Add("user-agent", "Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/134.0.0.0 Safari/537.36")

	res, _ := http.DefaultClient.Do(req)
	defer res.Body.Close()
	body, _ := io.ReadAll(res.Body)

	log.Println(string(body))
	var addressInfo SlowMistAddressInfo
	if err := json.Unmarshal(body, &addressInfo); err != nil { // Parse []byte to go struct pointer
		fmt.Println("Can not unmarshal JSON")
	}
	return addressInfo
}

func getText(addressInfo SlowMistAddressInfo) string {
	_item0 := addressInfo.RiskDic.TriangleLevel[0]
	_item1 := addressInfo.RiskDic.TriangleLevel[1]
	_item2 := addressInfo.RiskDic.TriangleLevel[2]

	_text0 := "🔍風險評分:" + strconv.Itoa(addressInfo.RiskDic.Score) + "\n"
	_text1 := ""
	_text2 := ""
	_text3 := ""
	_text4 := ""
	if _item0 > 1 {
		//log.Println("⚠️有與疑似惡意地址交互")
		_text1 = "⚠️有與疑似惡意地址交互\n"
	}
	if _item1 > 1 {
		//log.Println("⚠️️有與惡意地址交互")
		_text2 = "⚠️️有與惡意地址交互\n"
	}
	if _item2 > 1 {
		//log.Println("⚠️️有與高風險標籤地址交互")
		_text3 = "⚠️️有與高風險標籤地址交互\n"
	}

	_banned_item := addressInfo.RiskDic.HackingEvent

	if _banned_item != "" {
		//log.Println("⚠️️受制裁實體")
		_text4 = "⚠️️受制裁實體\n"
	}
	//msg = domain.MessageToSend{
	//	ChatId: message.Chat.ID,
	//	Text: "🔍風險評分:87\n" +
	//		"⚠️有與疑似惡意地址交互\n" +
	//		"⚠️️有與惡意地址交互\n" +
	//		"⚠️️有與高風險標籤地址交互\n" +
	//		"⚠️️受制裁實體\n" +
	//		"📢📢📢更詳細報告請聯繫客服@ushield001\n",
	//}
	//log.Println(events)
	_text5 := "📢📢📢更詳細報告請聯繫客服 @Ushield001\n"
	_text6 := "📊 地址概览\n"

	text := _text0 + _text1 + _text2 + _text3 + _text4 + _text5 + _text6
	return text
}

type SlowMistAddressInfo struct {
	Success bool   `json:"success"`
	Msg     string `json:"msg"`
	RiskDic struct {
		Score         int    `json:"score"`
		RiskList      []any  `json:"risk_list"`
		TriangleLevel []int  `json:"triangle_level"`
		HackingEvent  string `json:"hacking_event"`
		RiskDetail    []any  `json:"risk_detail"`
		ChkPhishDn    int    `json:"chk_phish_dn"`
		Upgrade       int    `json:"upgrade"`
	} `json:"risk_dic"`
}

func getAddressProfile(_coin string, _address, _cookie string) AddressProfile {
	url := "https://dashboard.misttrack.io/api/v1/address_overview?coin=" + _coin + "&address=" + _address
	req, _ := http.NewRequest("GET", url, nil)

	req.Header.Add("accept", "application/json, text/plain, */*")

	//req.Header.Add("cookie", "_ga=GA1.1.23337514.1742894564; _bl_uid=O8m7m8ksonwa0Ifjgw0erRqd9147; _ga_SGF4VCWFZY=GS1.1.1743393981.8.0.1743393981.0.0.0; detect_token=eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJyYW5kb21fc3RyIjoiMzI0Njk1In0.t5lYLE_oSwyNIJUSWAwxL7YrzXN5Di38sh4Vh9gjyJE; csrftoken=AOzVpYUl0Wdyk2gtoIzUQ5uOUEOxRBSMsqlINKjOh30dCmHX2ajNk8EcwFxrWy6g; sessionid=rn1a71d9nkn3coczdn08ahc00u5mw46i; _ga_40VGDGQFCB=GS1.1.1743393983.12.1.1743394123.0.0.0; _ga_5X5Z4KZ7PC=GS1.1.1743393983.12.1.1743394123.0.0.0")
	req.Header.Add("cookie", _cookie)
	req.Header.Add("language", "EN")

	req.Header.Add("referer", "https://dashboard.misttrack.io/address/"+_coin+"/"+_address)
	req.Header.Add("user-agent", "Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/134.0.0.0 Safari/537.36")

	res, _ := http.DefaultClient.Do(req)
	defer res.Body.Close()
	body, _ := io.ReadAll(res.Body)

	log.Println(string(body))

	var addressProfile AddressProfile
	if err := json.Unmarshal(body, &addressProfile); err != nil { // Parse []byte to go struct pointer
		fmt.Println("Can not unmarshal JSON")
	}
	return addressProfile
}
