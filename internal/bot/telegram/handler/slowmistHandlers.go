package handler

import (
	"encoding/json"
	"fmt"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"homework_bot/internal/bot"
	"homework_bot/internal/domain"
	"io"
	"log"
	"net/http"
	"strconv"
	"strings"
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
		b.GetSwitcher().Next(message.Chat.ID)
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
			_addressInfo := getAddressInfo(_symbol, _message)
			_text = getText(_addressInfo)
		}
		if strings.HasPrefix(_message, "T") && len(_message) == 34 {
			_symbol := "USDT-TRC20"
			_addressInfo := getAddressInfo(_symbol, _message)
			_text = getText(_addressInfo)
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

	b.GetSwitcher().Next(message.Chat.ID)
	_ = b.SendMessage(msg, bot.DefaultChannel)
	return nil
}

func getAddressInfo(_symbol string, _address string) SlowMistAddressInfo {
	url := "https://dashboard.misttrack.io/api/v1/address_risk_analysis?coin=" + _symbol + "&address=" + _address
	req, _ := http.NewRequest("GET", url, nil)

	req.Header.Add("accept", "application/json, text/plain, */*")

	req.Header.Add("cookie", "_ga=GA1.1.23337514.1742894564; _bl_uid=O8m7m8ksonwa0Ifjgw0erRqd9147; csrftoken=TxYjGKm5npSBDDIRUseK2kl9orBBbvggNhcxDu0jaWDfjYiIpMqH1SFvM3aiB8QT; sessionid=ob1gj0t1bf3hxzebem4v2775hwv7row4; detect_token=eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJyYW5kb21fc3RyIjoiNTI5MjIzIn0.QNBx0R_ow4ypzT8FbSmjfa1XQVM6Ak7UI8bcKU9wxNM; _ga_SGF4VCWFZY=GS1.1.1743222650.6.0.1743222650.0.0.0; _ga_40VGDGQFCB=GS1.1.1743222654.9.1.1743222703.0.0.0; _ga_5X5Z4KZ7PC=GS1.1.1743222654.9.1.1743222703.0.0.0")
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

	text := _text0 + _text1 + _text2 + _text3 + _text4 + _text5
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
