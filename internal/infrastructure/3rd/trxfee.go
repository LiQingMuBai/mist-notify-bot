package trxfee

import (
	"bytes"
	"crypto/hmac"
	"crypto/sha256"
	"encoding/hex"
	"encoding/json"
	"fmt"
	jsoniter "github.com/json-iterator/go"

	"io/ioutil"
	"net/http"
	"strconv"
	"time"
)

type TrxfeeClient struct {
	APIKey    string
	APISecret string
	URL       string
}

func NewTrxfeeClient(_url, apiKey, apiSecret string) *TrxfeeClient {
	return &TrxfeeClient{
		URL:       _url,
		APIKey:    apiKey,
		APISecret: apiSecret,
	}
}

type Data struct {
	EnergyAmount   int    `json:"energy_amount"`
	Period         string `json:"period"`
	ReceiveAddress string `json:"receive_address"`
	CallbackURL    string `json:"callback_url"`
	OutTradeNo     string `json:"out_trade_no"`
}

type AccountDataResp struct {
	Msg  string `json:"msg"`
	Code int    `json:"code"`
	Data struct {
		Balance      float64 `json:"balance"`
		UsdtBalance  float64 `json:"usdtBalance"`
		RechargeAddr string  `json:"rechargeAddr"`
	} `json:"data"`
}

func (c *TrxfeeClient) Account() (resp *AccountDataResp, err error) {
	url := c.URL + "/v1/account"

	req, _ := http.NewRequest("GET", url, nil)

	req.Header.Add("API-KEY", c.APIKey)

	res, _ := http.DefaultClient.Do(req)

	defer res.Body.Close()
	body, _ := ioutil.ReadAll(res.Body)

	fmt.Println(res)
	fmt.Println(string(body))

	var accountResp AccountDataResp

	if err := json.Unmarshal(body, &accountResp); err != nil {
		return nil, fmt.Errorf("failed to unmarshal account response: %w", err)
	}
	return &accountResp, nil

}
func (c *TrxfeeClient) Order(_outTradeNo, _receiveAddress string, _energyAmount int) {

	time.Sleep(1 * time.Second)
	timestamp := strconv.FormatInt(time.Now().Unix(), 10)

	data := Data{
		EnergyAmount:   _energyAmount,
		Period:         "1H",
		ReceiveAddress: _receiveAddress,
		CallbackURL:    "http://{mydomain}/callback",
		OutTradeNo:     _outTradeNo,
	}

	ordered_data := map[string]interface{}{
		"energy_amount":   data.EnergyAmount,
		"period":          data.Period,
		"receive_address": data.ReceiveAddress,
		"callback_url":    data.CallbackURL,
		"out_trade_no":    data.OutTradeNo,
	}

	json := jsoniter.ConfigCompatibleWithStandardLibrary
	b, err := json.Marshal(ordered_data)
	if err != nil {
		panic(err)
	}
	json_data := string(b)

	message := timestamp + "&" + json_data
	signature := createHmac(message, c.APISecret)

	client := &http.Client{}
	req, err := http.NewRequest("POST", c.URL+"/v1/api", bytes.NewBuffer([]byte(json_data)))
	if err != nil {
		panic(err)
	}

	req.Header.Set("API-KEY", c.APIKey)
	req.Header.Set("TIMESTAMP", timestamp)
	req.Header.Set("SIGNATURE", signature)
	req.Header.Set("Content-Type", "application/json")

	resp, err := client.Do(req)
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()

	respBody, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		panic(err)
	}
	fmt.Println(string(respBody))

}

func createHmac(message string, secret string) string {
	h := hmac.New(sha256.New, []byte(secret))
	h.Write([]byte(message))
	return hex.EncodeToString(h.Sum(nil))
}
