package trxfee

import (
	"log"
	"testing"
)

func TestAccount(t *testing.T) {
	trxfeeClient := NewTrxfeeClient("https://trxfee.io/", "CC4F20ACDB45AFA10A22D6BDA2AE9F3F", "99144B2AC7ED7F73ECFF59144D46E321F1DC83B373DE2FF6A367423F4CF61FB5")

	accountResp, err := trxfeeClient.Account()

	if err != nil {
		t.Fatalf("failed to call Account(): %v", err)
	}
	log.Println(accountResp.Code)

	log.Println(accountResp.Msg)

	log.Println(accountResp.Data.RechargeAddr)
	log.Println(accountResp.Data.Balance)
	log.Println(accountResp.Data.UsdtBalance)

}
func TestOrder(t *testing.T) {
	//trxfeeClient := NewTrxfeeClient("https://trxfee.io/", "CC4F20ACDB45AFA10A22D6BDA2AE9F3F", "99144B2AC7ED7F73ECFF59144D46E321F1DC83B373DE2FF6A367423F4CF61FB5")
	//
	//accountResp, err := trxfeeClient.Order()
	//
	//if err != nil {
	//	t.Fatalf("failed to call Account(): %v", err)
	//}
	//log.Println(accountResp.Code)
	//
	//log.Println(accountResp.Msg)
	//
	//log.Println(accountResp.Data.RechargeAddr)
	//log.Println(accountResp.Data.Balance)
	//log.Println(accountResp.Data.UsdtBalance)

}
