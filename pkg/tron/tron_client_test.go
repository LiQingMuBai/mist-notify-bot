package tron

import (
	"encoding/hex"
	"fmt"
	"github.com/fbsobreira/gotron-sdk/pkg/keys"
	"github.com/magiconair/properties/assert"
	"testing"
)

func TestTronClient_GenerateAddress(t *testing.T) {
	// Hardcoded index of 0 for brandnew account.
	private, _ := keys.FromMnemonicSeedAndPassphrase("abandon ability able about above absent absorb abstract absurd abuse access accident", "", 0)
	pk_bytes := private.Serialize()

	fmt.Println("Privatekey: ", hex.EncodeToString(pk_bytes))

	address0, err := GetTronAddressFromPrivateKey(hex.EncodeToString(pk_bytes))
	if err != nil {
	}
	fmt.Println("address0: ", address0)
	// Hardcoded index of 0 for brandnew account.
	private2, _ := keys.FromMnemonicSeedAndPassphrase("abandon ability able about above absent absorb abstract absurd abuse access accident", "", 1)
	pk_bytes2 := private2.Serialize()

	fmt.Println("Privatekey: ", hex.EncodeToString(pk_bytes2))
	address1, err := GetTronAddressFromPrivateKey(hex.EncodeToString(pk_bytes2))
	if err != nil {
	}
	fmt.Println("address1: ", address1)
	//publicKey := private.PubKey()
	// 5. 生成波场地址
	//tronAddress := address.PubkeyToAddress(publicKey)
	//base58Address := tronAddress.String()
	//
	//// 6. 转换为hex地址格式(41开头)
	//hexAddress := common.BytesToHexString(tronAddress.Bytes())
	//fmt.Printf("TRON Address: %s\n", tronAddress)
	//fmt.Printf("Private Key: %s\n", privateKeyHex)
}
func TestTronHexToBase58(t *testing.T) {

	usdt, _ := TronHexToBase58("a614f803b6fd780986a42c78ec9c7f77e6ded13c")
	//contract, _ := tron.TronHexToBase58(log.Address)
	assert.Equal(t, usdt, "TR7NHqjeKQxGTCi8q8ZY4pL8otSzgjLj6t", "success")
}

func TestTronBase58ToHex(t *testing.T) {

	tronHex := "000000000000000000000000302ac98a7e1db6d55cb77188c4528ac0cd42379b"[24:64]
	fmt.Println(tronHex)
	base58Addr, _ := HexToTronBase58("41" + tronHex)
	fmt.Println(base58Addr)

}

func TestTronClient_GetAccountResources(t *testing.T) {
	//TRON_FULL_NODE := "https://api.trongrid.io"
	TRON_FULL_NODE := "https://young-clean-orb.tron-mainnet.quiknode.pro/9283c9ddb51102d9d22cf1ac5a6e6fc898eeaf77"
	tronClient := NewTronClient(TRON_FULL_NODE)

	_address, _ := Base58ToTronHex("TPiTNfSJdk2Ao12eCDESahv3ZN4npdGBWQ")
	resources, _ := tronClient.GetAccountResources(_address)
	//resources2, _ := tronClient.GetAccountResources("TPiTNfSJdk2Ao12eCDESahv3ZN4npdGBWQ")

	fmt.Println(resources)
	//fmt.Println(resources2)

}

//
//func TestTronClient_BatchGetAddressBalances(t *testing.T) {
//	TRON_FULL_NODE := "https://api.trongrid.io"
//	//TRON_FULL_NODE := "https://young-clean-orb.tron-mainnet.quiknode.pro/9283c9ddb51102d9d22cf1ac5a6e6fc898eeaf77/"
//	tronClient := NewTronClient(TRON_FULL_NODE)
//
//	addresses := []string{"TXtNKWqibAqaFE3HRTKBovUyEiaD8S3Kb9", "TVaHLSfHvdqxCAaiYY5eY5tm1ouHL9zDU1"}
//	addressesValues, err := tronClient.BatchGetAddressBalances(addresses)
//	if err != nil {
//	}
//	for _address, _amount := range addressesValues {
//
//		log.Println(_address)
//		log.Println(_amount.String())
//
//	}
//}
//func TestTronClient_GetNativeBalance(t *testing.T) {
//	TRON_FULL_NODE := "https://api.trongrid.io"
//	//TRON_FULL_NODE := "https://young-clean-orb.tron-mainnet.quiknode.pro/9283c9ddb51102d9d22cf1ac5a6e6fc898eeaf77/"
//	tronClient := NewTronClient(TRON_FULL_NODE)
//
//	//okex hot wallet
//	_address, _ := Base58ToTronHex("TPiTNfSJdk2Ao12eCDESahv3ZN4npdGBWQ")
//	fmt.Println(_address)
//	resources, _ := tronClient.GetNativeBalance(context.Background(), _address)
//
//	fmt.Println(resources.Int64())
//
//	balance := utils.DivideWithPrecision(resources, 4)
//	fmt.Println(balance)
//
//	trxBalance, err := tronClient.GetNativeBalance(context.Background(), "4196c784e3985d35a7133b2a33670871b09e8f86ea")
//	if err != nil {
//		fmt.Errorf("query master's address failed : %w", err)
//	}
//
//	fmt.Println(trxBalance)
//	//fmt.Println(resources.Int64() / 1_000_000)
//}
//
//func TestGetTronAddressFromPrivateKey(t *testing.T) {
//	names := "TG4nheHr9ZsBms6tn3seAsKwGXyG9XWJsV"
//
//	// Split into a slice
//	addresses := strings.Split(names, ",")
//
//	for _, address := range addresses {
//
//		log.Println(address)
//	}
//}
//
//func TestTronClient_GetIncomingTransactions(t *testing.T) {
//	TRON_FULL_NODE := "https://api.trongrid.io"
//	//TRON_FULL_NODE := "https://young-clean-orb.tron-mainnet.quiknode.pro/9283c9ddb51102d9d22cf1ac5a6e6fc898eeaf77/"
//	tronClient := NewTronClient(TRON_FULL_NODE)
//
//	_time := utils.GetTimeDaysAgo(68)
//	txs, err := tronClient.GetIncomingTransactions("TD8TfXPM4gCEVzdRGndD8qem4Kp925DgjH", 200, _time)
//
//	if err != nil {
//	}
//
//	for _, tx := range txs {
//
//		fmt.Println(tx.FromAddress)
//		fmt.Println(tx.ToAddress)
//		fmt.Println(tx.Amount)
//	}
//}
//
//func TestTronClient_GetOutgoingTransactions(t *testing.T) {
//	TRON_FULL_NODE := "https://api.trongrid.io"
//	//TRON_FULL_NODE := "https://young-clean-orb.tron-mainnet.quiknode.pro/9283c9ddb51102d9d22cf1ac5a6e6fc898eeaf77/"
//	tronClient := NewTronClient(TRON_FULL_NODE)
//	_time := utils.GetTimeDaysAgo(68)
//	txs, err := tronClient.GetOutgoingTransactions("TD8TfXPM4gCEVzdRGndD8qem4Kp925DgjH", 200, _time)
//
//	if err != nil {
//	}
//	log.Println(len(txs))
//	for _, tx := range txs {
//
//		fmt.Println(tx.FromAddress)
//		fmt.Println(tx.ToAddress)
//		fmt.Println(tx.Amount)
//		fmt.Println(tx.Timestamp)
//	}
//}
//
//func TestTronClient_GetUSDTTransferCount(t *testing.T) {
//	TRON_FULL_NODE := "https://api.trongrid.io"
//	//TRON_FULL_NODE := "https://young-clean-orb.tron-mainnet.quiknode.pro/9283c9ddb51102d9d22cf1ac5a6e6fc898eeaf77/"
//	tronClient := NewTronClient(TRON_FULL_NODE)
//
//	_count, err := tronClient.GetUSDTTransferCount("TG4nheHr9ZsBms6tn3seAsKwGXyG9XWJsV", "TR7NHqjeKQxGTCi8q8ZY4pL8otSzgjLj6t")
//
//	if err != nil {
//
//	}
//
//	fmt.Println(_count)
//
//	_total, err := getUSDTTransferCount("TDoXUNZ6PajKuiUkcYg3EDSV9bnqGqsbcf")
//
//	if err != nil {
//	}
//
//	fmt.Println(_total)
//}
//
//func TestTronClient_GetLatestBlock(t *testing.T) {
//	TRON_FULL_NODE := "https://api.trongrid.io"
//
//	tronClient := NewTronClient(TRON_FULL_NODE)
//
//	nowBlock := tronClient.GetLatestBlock()
//
//	fmt.Println(nowBlock)
//}
//
//// Base TronGrid API endpoint
//const apiTemplate = "https://api.trongrid.io/v1/accounts/%s/transactions/trc20"
//
//// USDT TRC20 contract address on TRON
//const usdtContract = "TR7NHqjeKQxGTCi8q8ZY4pL8otSzgjLj6t"
//
//// getUSDTTransferCount returns the number of USDT transfers *to* the given address.
//func getUSDTTransferCount(address string) (int, error) {
//	apiURL := fmt.Sprintf(apiTemplate, address)
//	params := url.Values{}
//	params.Set("only_from", "true") // Only incoming transfers
//	params.Set("limit", "1")        // We only care about the total
//	params.Set("contract_address", usdtContract)
//
//	fullURL := apiURL + "?" + params.Encode()
//
//	log.Println(fullURL)
//	resp, err := http.Get(fullURL)
//	if err != nil {
//		return 0, fmt.Errorf("HTTP request failed: %w", err)
//	}
//	defer resp.Body.Close()
//
//	var response Trc20TxResponse
//	if err := json.NewDecoder(resp.Body).Decode(&response); err != nil {
//		return 0, fmt.Errorf("failed to decode JSON: %w", err)
//	}
//
//	return response.Meta.Total, nil
//}
//
//// Structure for the API response
//type Trc20TxResponse struct {
//	Data []struct {
//		TokenInfo struct {
//			Address string `json:"address"`
//			Symbol  string `json:"symbol"`
//		} `json:"token_info"`
//	} `json:"data"`
//	Meta struct {
//		Total int `json:"total"`
//	} `json:"meta"`
//}
