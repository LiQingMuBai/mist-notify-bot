package handler

//import (
//	"bytes"
//	"encoding/json"
//	"fmt"
//	"io"
//	"log"
//	"net/http"
//	"strings"
//	"testing"
//)
//
//func TestGraphAddress(t *testing.T) {
//
//	_coin := "ETH"
//	_address := "0xf510e53ef8da4e45ffa59eb554511a7410e5efd3"
//
//	lableAddresList := getNotSafeAddress(_coin, _address)
//
//	var index uint64 = 0
//	for _, data := range lableAddresList.GraphDic.NodeList {
//
//		index++
//
//		if strings.Contains(data.Label, "huione") {
//			log.Println(data.Label)
//			log.Println(data.Title)
//			log.Println(data.Title[0:5] + "..." + data.Title[29:34])
//		}
//
//	}
//}
//
//// 2.地址风险查询操作
////
//// 🔍 地址风险查询结果：
////
//// 📍 地址：TJxxxxxxxxxxxxxxxxxxxx
////
//// ⚠️ 风险等级：中风险（55 分）
////
//// 标签：涉嫌与未知地址频繁交易
////
//// 📊 地址概览
////
//// 余额：8,881.99 USDT
////
//// 累计收入：178,806.51 USDT
////
//// 首次活跃时间：2025-02-26
////
//// 最后活跃时间：2025-03-04
////
//// 交易次数：34 笔
////
//// 主要交易对手分析：
////
//// 1. 未知来源 - $341,211（97.84%）
//// 2. OKX - $6,000（1.72%）
//// 3. Huobi - $765（0.22%）
//// 4. Bybit - $755（0.22%）
////
//// 📄 详细分析报告 ➜ 50 TRX
////
//// 👉 /report TJxxxxxxxxxxxx
////
//// 每日免费查询剩余：0 次
////
//// 超额查询 ➜ 10 TRX / 次
////
//// 🛡️ U盾在手，链上无忧！
//func TestSlowMist_AddressInfo(t *testing.T) {
//
//	_coin := "ETH"
//	_address := "0xf510e53ef8da4e45ffa59eb554511a7410e5efd3"
//	addressProfile := getAddressProfile(_coin, _address)
//
//	log.Println("余额：", addressProfile.BalanceUsd)
//	log.Println("累計收入：", addressProfile.TotalReceivedUsd)
//	log.Println("累计支出：", addressProfile.TotalSpentUsd)
//	log.Println("首次活躍時間：", addressProfile.FirstTxTime)
//	log.Println("最後活躍時間：", addressProfile.LastTxTime)
//	log.Println("交易次數：", addressProfile.TxCount+"筆")
//
//}
//
//func TestSlowMist_GraphAddressInfo(t *testing.T) {
//	url := "https://dashboard.misttrack.io/api/v1/address_graph_analysis?coin=ETH&address=0xf510e53ef8da4e45ffa59eb554511a7410e5efd3&time_filter="
//	req, _ := http.NewRequest("GET", url, nil)
//
//	req.Header.Add("accept", "application/json, text/plain, */*")
//
//	//req.Header.Add("cookie", "_ga=GA1.1.23337514.1742894564; _bl_uid=O8m7m8ksonwa0Ifjgw0erRqd9147; _ga_SGF4VCWFZY=GS1.1.1743393981.8.0.1743393981.0.0.0; detect_token=eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJyYW5kb21fc3RyIjoiMzI0Njk1In0.t5lYLE_oSwyNIJUSWAwxL7YrzXN5Di38sh4Vh9gjyJE; csrftoken=AOzVpYUl0Wdyk2gtoIzUQ5uOUEOxRBSMsqlINKjOh30dCmHX2ajNk8EcwFxrWy6g; sessionid=rn1a71d9nkn3coczdn08ahc00u5mw46i; _ga_40VGDGQFCB=GS1.1.1743393983.12.1.1743394123.0.0.0; _ga_5X5Z4KZ7PC=GS1.1.1743393983.12.1.1743394123.0.0.0")
//	req.Header.Add("cookie", "_ga=GA1.1.952339838.1743478159; _bl_uid=5qmId8h8xUwxLhvvIqLy878nX7vz; csrftoken=ZsUzP3PB1b6hFsu7R9hhRsKO5qOSvsvSRMDrqXqq2gRbLywwsr4toHEUZNzTdYk7; sessionid=23qxazzhkz6it7ow8gtz1p3ua2bqx6x3; detect_token=eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJyYW5kb21fc3RyIjoiMzQzNDY5In0.ZYla82HwE6OqaEgJblSdjD08FvRXlWm0YbeermrRhE4; _ga_40VGDGQFCB=GS1.1.1743572931.3.1.1743573087.0.0.0; _ga_5X5Z4KZ7PC=GS1.1.1743572931.3.1.1743573087.0.0.0")
//	req.Header.Add("language", "EN")
//
//	req.Header.Add("referer", "https://dashboard.misttrack.io/address/ETH/0xf510e53ef8da4e45ffa59eb554511a7410e5efd3")
//	req.Header.Add("user-agent", "Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/134.0.0.0 Safari/537.36")
//
//	res, _ := http.DefaultClient.Do(req)
//	defer res.Body.Close()
//	body, _ := io.ReadAll(res.Body)
//
//	log.Println(string(body))
//
//	var addressProfile AddressProfile
//	if err := json.Unmarshal(body, &addressProfile); err != nil { // Parse []byte to go struct pointer
//		fmt.Println("Can not unmarshal JSON")
//	}
//
//	log.Println("余额：", addressProfile.BalanceUsd)
//	log.Println("累计收入：", addressProfile.TotalReceivedUsd)
//	log.Println("累计支出：", addressProfile.TotalSpentUsd)
//	log.Println("首次活跃时间：", addressProfile.FirstTxTime)
//	log.Println("最后活跃时间：", addressProfile.LastTxTime)
//	log.Println("交易次数：", addressProfile.TxCount+"笔")
//	log.Println("交易次数：", addressProfile.TxCount+"笔")
//
//}
//
////GET https://dashboard.misttrack.io/api/v1/address_risk_analysis?coin=USDT-ERC20&address=0xF510e53EF8DA4e45FFA59EB554511a7410E5eFD3
////:authority: dashboard.misttrack.io
////:path:/api/v1/address_risk_analysis?coin=USDT-ERC20&address=0xF510e53EF8DA4e45FFA59EB554511a7410E5eFD3
////:scheme:https
////accept:application/json, text/plain, */*
////accept-encoding:gzip, deflate, br, zstd
////accept-language:en-US,en;q=0.9,zh-CN;q=0.8,zh;q=0.7
////cookie:_ga=GA1.1.23337514.1742894564; _bl_uid=O8m7m8ksonwa0Ifjgw0erRqd9147; csrftoken=TxYjGKm5npSBDDIRUseK2kl9orBBbvggNhcxDu0jaWDfjYiIpMqH1SFvM3aiB8QT; sessionid=ob1gj0t1bf3hxzebem4v2775hwv7row4; detect_token=eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJyYW5kb21fc3RyIjoiNTI5MjIzIn0.QNBx0R_ow4ypzT8FbSmjfa1XQVM6Ak7UI8bcKU9wxNM; _ga_SGF4VCWFZY=GS1.1.1743222650.6.0.1743222650.0.0.0; _ga_40VGDGQFCB=GS1.1.1743222654.9.1.1743222703.0.0.0; _ga_5X5Z4KZ7PC=GS1.1.1743222654.9.1.1743222703.0.0.0
////language:EN
////priority:u=1, i
////referer:https://dashboard.misttrack.io/address/USDT-ERC20/0xF510e53EF8DA4e45FFA59EB554511a7410E5eFD3
////sec-ch-ua:"Chromium";v="134", "Not:A-Brand";v="24", "Google Chrome";v="134"
////sec-ch-ua-mobile:?0
////sec-ch-ua-platform:"Windows"
////sec-fetch-dest:empty
////sec-fetch-mode:cors
////sec-fetch-site:same-origin
////user-agent:Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/134.0.0.0 Safari/537.36
//
//func TestSlowMist_ERC20_Vist(t *testing.T) {
//	log.Println(">>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>TestSlowMistVist<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<")
//	_symbol := "USDT-ERC20"
//	_address := "0xF510e53EF8DA4e45FFA59EB554511a7410E5eFD3"
//	addressInfo := getAddressInfo(_symbol, _address)
//
//	//log.Println(addressInfo.RiskDic.Score)
//	//log.Println(events)
//	text := getText(addressInfo)
//
//	log.Println(text)
//}
//
//func TestSlowMist_TRC20_Vist(t *testing.T) {
//	log.Println(">>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>TestSlowMistVist<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<")
//	_symbol := "USDT-TRC20"
//	_address := "TKKkmmC1evWhPYmxt1HjZot6eEDhkvydBh"
//	addressInfo := getAddressInfo(_symbol, _address)
//	//log.Println("🔍风险评分:" + strconv.Itoa(addressInfo.RiskDic.Score))
//
//	text := getText(addressInfo)
//
//	log.Println(text)
//}
//
//func TestNotifyUser(t *testing.T) {
//	message := map[string]string{
//		"chat_id": "7347235462", // 或直接用 chat_id 如 "123456789"
//		"text":    "🎉 新年快乐！祝你新的一年万事如意，心想事成！ 🎊",
//	}
//
//	botToken := "7916934957:AAEy5cOEhSXdAQk5vQyMTVEs8BMRvonm4Ho"
//
//	// 转换为 JSON
//	jsonData, err := json.Marshal(message)
//	if err != nil {
//		fmt.Println("JSON 编码失败:", err)
//		return
//	}
//
//	// 发送 POST 请求到 Telegram Bot API
//	url := fmt.Sprintf("https://api.telegram.org/bot%s/sendMessage", botToken)
//	resp, err := http.Post(url, "application/json", bytes.NewBuffer(jsonData))
//	if err != nil {
//		fmt.Println("发送消息失败:", err)
//		return
//	}
//	defer resp.Body.Close()
//
//	// 打印响应结果
//	fmt.Println("消息发送状态:", resp.Status)
//}
