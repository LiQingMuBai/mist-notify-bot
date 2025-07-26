package handler

import (
	"bytes"
	"context"
	"crypto/hmac"
	"crypto/sha256"
	"encoding/hex"
	"fmt"
	jsoniter "github.com/json-iterator/go"
	"io/ioutil"
	"net/http"
	"strconv"
	"time"
)

const (
	//APIKey    = "你的APIKEY"
	//APISecret = "你的APISECRET"
	URL = "https://trxfee.io/v1/timesOrder"
)

type Data struct {
	RentTimes         int    `json:"rentTimes"`
	RecvAddr          string `json:"recvAddr"`
	FreePause         int    `json:"freePause"`
	ResourceReplenish string `json:"resourceReplenish"`
}
type TrxfeeHandler struct{}

func NewTrxfeeHandler() *TrxfeeHandler {
	return &TrxfeeHandler{}
}

func (c *TrxfeeHandler) RequestTimesOrder(
	ctx context.Context,
	APIKey string, //trxfee
	APISecret string, //trxfee
	_toAddress string,
	_times int,
) {
	timestamp := strconv.FormatInt(time.Now().Unix(), 10)

	data := Data{
		RentTimes:         _times,
		RecvAddr:          _toAddress,
		FreePause:         2,
		ResourceReplenish: "1",
	}

	jordered_data := map[string]interface{}{
		"rentTimes":         data.RentTimes,
		"recvAddr":          data.RecvAddr,
		"freePause":         data.FreePause,
		"resourceReplenish": data.ResourceReplenish,
	}

	json := jsoniter.ConfigCompatibleWithStandardLibrary
	b, err := json.Marshal(jordered_data)
	if err != nil {
		panic(err)
	}
	json_data := string(b)

	message := timestamp + "&" + json_data
	signature := createHmac(message, APISecret)

	client := &http.Client{}
	req, err := http.NewRequest("POST", URL, bytes.NewBuffer([]byte(json_data)))
	if err != nil {
		panic(err)
	}

	req.Header.Set("API-KEY", APIKey)
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
