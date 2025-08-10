package handler

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"strconv"
)

func GetNotSafeAddress(_coin string, _address, _cookie string) LableAddresList {
	//_cookie = "_bl_uid=1wmz8eCq1445tmhU8hktzps2hC51; detect_token=eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJyYW5kb21fc3RyIjoiODY3NzM5In0.DaCjSesFMsjGWQkB7iHA1EI5Lp2s3-DTPmxB7nNIPKI; csrftoken=u5xzDP2pcMqbACyYHyVUlJNtmNlr4pIn5i6ullnZNtNunsFbIHvHZk9rteAcyq2l; sessionid=uqs748r6gmq6cjjrqig5461rw8nc3gq9"
	//_coin = "ETH"
	//
	//_address = "0xf510e53ef8da4e45ffa59eb554511a7410e5efd3"
	url := "https://dashboard.misttrack.io/api/v1/address_graph_analysis?coin=" + _coin + "&address=" + _address + "&time_filter="
	req, _ := http.NewRequest("GET", url, nil)
	//https://dashboard.misttrack.io/api/v1/address_graph_analysis?coin=ETH&address=0xf510e53ef8da4e45ffa59eb554511a7410e5efd3&time_filter=
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
			ID    string `json:"id"`
			Label string `json:"label"`
			Title string `json:"title"`
			Layer int    `json:"layer"`
			Addr  string `json:"addr"`
			Track string `json:"track"`
			//Pid       string `json:"pid"`
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

func GetAddressInfo(_symbol string, _address, _cookie string) SlowMistAddressInfo {
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

func GetText(addressInfo SlowMistAddressInfo) string {
	_item0 := addressInfo.RiskDic.TriangleLevel[0]
	_item1 := addressInfo.RiskDic.TriangleLevel[1]
	_item2 := addressInfo.RiskDic.TriangleLevel[2]

	_text0 := "🔍风险评分:" + strconv.Itoa(addressInfo.RiskDic.Score)

	if addressInfo.RiskDic.Score <= 3 {
		_text0 += " 🟢" + "\n"
	}
	if addressInfo.RiskDic.Score > 3 && addressInfo.RiskDic.Score <= 60 {
		_text0 += " 🟡" + "\n"
	}
	if addressInfo.RiskDic.Score > 60 {
		_text0 += " 🔴" + "\n"
	}
	_text1 := ""
	_text2 := ""
	_text3 := ""
	_text4 := ""
	if _item0 > 1 {
		//log.Println("⚠️有與疑似惡意地址交互")
		_text1 = "⚠️有與疑似惡意地址交互\n"
	}
	if _item1 > 1 {
		//log.Println("⚠️️有与恶意地址交互")
		_text2 = "⚠️️有与恶意地址交互\n"
	}
	if _item2 > 1 {
		//log.Println("⚠️️有與高風險標籤地址交互")
		_text3 = "⚠️️有與高風險標籤地址交互\n"
	}

	_banned_item := addressInfo.RiskDic.HackingEvent

	if _banned_item != "" {
		//log.Println("⚠️️受制裁实体")
		_text4 = "⚠️️受制裁实体\n"
	}
	//msg = domain.MessageToSend{
	//	ChatId: message.Chat.ID,
	//	Text: "🔍风险评分:87\n" +
	//		"⚠️有與疑似惡意地址交互\n" +
	//		"⚠️️有与恶意地址交互\n" +
	//		"⚠️️有與高風險標籤地址交互\n" +
	//		"⚠️️受制裁实体\n" +
	//		"📢📢📢更詳細報告請聯繫客服@ushield001\n",
	//}
	//log.Println(events)

	_text6 := "📊 地址概览\n"

	text := _text0 + _text1 + _text2 + _text3 + _text4 + _text6
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

func GetAddressProfile(_coin string, _address, _cookie string) AddressProfile {
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
