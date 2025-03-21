package main

import (
	"fmt"
	"testing"

	"github.com/danchengash/sasapay-go-plugin"
	"github.com/danchengash/sasapay-go-plugin/models"
)

var clientSecret = "XXXXXX"
var clientId = "XXXX"
var sp = sasapay.NewSasaPay(clientId, clientSecret, "XXX", int(sasapay.Production), true)

func main() {
	// sp = sasapay.NewSasaPay(clientId, clientSecret, "XXX", int(sasapay.Production), true)
	TestTransactionStatement(&testing.T{})

	TestCheckTransactioStatus(&testing.T{})
	
}

func TestGetBanklist(t *testing.T) {
	sp.GetChannelCodes()

}
func TestTransactionStatement(t *testing.T) {
	reponse, err := sp.GetTransactions(1, 2, "2024-2-23", "2024-10-28", "")
	if err != nil {
		t.Error(err)
	}
	for _, v := range reponse.Data.Transactions {
		fmt.Printf("%v - %f \n", v.ResultDescription, v.TransactionCharges)
	}
}

func TestC2B(t *testing.T) {
	response, err := sp.Customer2Business(models.C2BRequest{
		MerchantCode:     sp.MerchantCode,
		Currency:         "KES",
		NetworkCode:      "63902",
		PhoneNumber:      "2547XXXX",
		TransactionDesc:  "desc",
		AccountReference: "ref",
		Amount:           2,
		CallBackURL:      "https://posthere.io/67df-4d9c-9386",
	})
	if err != nil {
		t.Error(err)
	}

	_, err = sp.C2BProcess(response.CheckoutRequestID, "4345")

	if err != nil {
		t.Error(err)
	}
	t.Log(response.Detail)

}
func TestB2c(t *testing.T) {
	respone, err := sp.Business2Customer(models.B2CRequest{

		MerchantCode:                 sp.MerchantCode,
		MerchantTransactionReference: "dsd",
		Amount:                       1,
		Currency:                     "KES",
		ReceiverNumber:               "254712345677",
		Channel:                      "0",
		Reason:                       "test reason",
		CallBackURL:                  "https://posthere.io/67df-4d9c-9386",
	})
	if err != nil {
		t.Error(err)
	}
	t.Log(respone.Detail)
}

func TestB2B(t *testing.T) {
	response, err := sp.Business2Business(models.B2BRequest{
		MerchantCode:                 sp.MerchantCode,
		NetworkCode:                  "0",
		AccountReference:             "ref00",
		ReceiverAccountType:          "TILL",
		MerchantTransactionReference: "uoiwp10",
		Currency:                     "KES",
		Amount:                       50,
		ReceiverMerchantCode:         "555550",
		CallBackURL:                  "https://webhook.site/e8afa636-daed-49fa-bc1b-aeb9e607955f",
		Reason:                       "test",
	})
	if err != nil {
		fmt.Printf("error ==> %s", err.Error())
	}
	fmt.Println(response.Detail)

}

func TestCheckTransactioStatus(t *testing.T) {
	// pa := models.CheckTransactionStatusRequest{CheckoutRequestID: nil, CallBackUrl: "https://webhook.site/dbeae6a3-9525-42ee-be21-3a9c3c29725e"}
	pa := models.CheckTransactionStatusRequest{CheckoutRequestID: nil}

	re := "R5MeZj_S_1218"
	pa.MerchantTransactionReference = &re
	response, err := sp.CheckTransactionStatus(pa)
	if err != nil {
		t.Error(err)
	}
	t.Log(response.Data.ResultDescription)
}

func TestVerifyTransaction(t *testing.T) {
	response, err := sp.VerifyTransaction("SPEJZM2W9YRN75M")
	if err != nil {
		t.Error(err)
	} else {
		t.Log(response.Data)

	}
}

func TestMerchantBalance(t *testing.T) {
	response, err := sp.CheckMerchantBalance()
	if err != nil {
		t.Error(err)
	}
	t.Log(response.Data)
}

func TestBusiness2Benefiary(t *testing.T) {
	response, err := sp.Business2Benefiary(models.Business2BeneficiaryReq{
		TransactionReference:     "2323",
		SenderMerchantCode:       sp.MerchantCode,
		ReceiverMerchantCode:     "94000",
		BeneficiaryAccountNumber: "254712345672",
		Amount:                   1,
		TransactionFee:           1,
		Reason:                   "test",
		CallBackURL:              "https://posthere.io/67df-4d9c-9386",
	})
	if err != nil {
		t.Error(err)
	}
	t.Log(response.Message)
}
