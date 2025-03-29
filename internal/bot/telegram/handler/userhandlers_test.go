package handler

import (
	"log"
	"testing"
)

//GET https://dashboard.misttrack.io/api/v1/address_risk_analysis?coin=USDT-ERC20&address=0xF510e53EF8DA4e45FFA59EB554511a7410E5eFD3
//:authority: dashboard.misttrack.io
//:path:/api/v1/address_risk_analysis?coin=USDT-ERC20&address=0xF510e53EF8DA4e45FFA59EB554511a7410E5eFD3
//:scheme:https
//accept:application/json, text/plain, */*
//accept-encoding:gzip, deflate, br, zstd
//accept-language:en-US,en;q=0.9,zh-CN;q=0.8,zh;q=0.7
//cookie:_ga=GA1.1.23337514.1742894564; _bl_uid=O8m7m8ksonwa0Ifjgw0erRqd9147; csrftoken=TxYjGKm5npSBDDIRUseK2kl9orBBbvggNhcxDu0jaWDfjYiIpMqH1SFvM3aiB8QT; sessionid=ob1gj0t1bf3hxzebem4v2775hwv7row4; detect_token=eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJyYW5kb21fc3RyIjoiNTI5MjIzIn0.QNBx0R_ow4ypzT8FbSmjfa1XQVM6Ak7UI8bcKU9wxNM; _ga_SGF4VCWFZY=GS1.1.1743222650.6.0.1743222650.0.0.0; _ga_40VGDGQFCB=GS1.1.1743222654.9.1.1743222703.0.0.0; _ga_5X5Z4KZ7PC=GS1.1.1743222654.9.1.1743222703.0.0.0
//language:EN
//priority:u=1, i
//referer:https://dashboard.misttrack.io/address/USDT-ERC20/0xF510e53EF8DA4e45FFA59EB554511a7410E5eFD3
//sec-ch-ua:"Chromium";v="134", "Not:A-Brand";v="24", "Google Chrome";v="134"
//sec-ch-ua-mobile:?0
//sec-ch-ua-platform:"Windows"
//sec-fetch-dest:empty
//sec-fetch-mode:cors
//sec-fetch-site:same-origin
//user-agent:Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/134.0.0.0 Safari/537.36

func TestSlowMist_ERC20_Vist(t *testing.T) {
	log.Println(">>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>TestSlowMistVist<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<")
	_symbol := "USDT-ERC20"
	_address := "0xF510e53EF8DA4e45FFA59EB554511a7410E5eFD3"
	addressInfo := getAddressInfo(_symbol, _address)

	//log.Println(addressInfo.RiskDic.Score)
	//log.Println(events)
	text := getText(addressInfo)

	log.Println(text)
}

func TestSlowMist_TRC20_Vist(t *testing.T) {
	log.Println(">>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>TestSlowMistVist<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<")
	_symbol := "USDT-TRC20"
	_address := "TKKkmmC1evWhPYmxt1HjZot6eEDhkvydBh"
	addressInfo := getAddressInfo(_symbol, _address)
	//log.Println("🔍風險評分:" + strconv.Itoa(addressInfo.RiskDic.Score))

	text := getText(addressInfo)

	log.Println(text)
}
