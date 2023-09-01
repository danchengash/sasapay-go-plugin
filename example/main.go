package main

import (
	"fmt"
	"testing"

	"github.com/danchengash/sasapay-go-plugin"
	"github.com/danchengash/sasapay-go-plugin/models"
)

var clientId = "XXXXXX"
var clientSecret = "XXXXXXXXXXXXXXXXXXXX"
var sp = sasapay.NewSasaPay(clientId, clientSecret, "1234", int(sasapay.Production), false)

func main() {
	response, err := sp.VerifyTransaction("SPEJZM2W9YRN75M")
	if err != nil {
		fmt.Println(err.Error())
	} else {
		fmt.Printf("amount %f", response.Data.TransAmount)
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
		MerchantTransactionReference: "uoiwp",
		Currency:                     "KES",
		Amount:                       1,
		ReceiverMerchantCode:         "94000",
		CallBackURL:                  "https://posthere.io/67df-4d9c-9386",
		Reason:                       "test",
	})
	if err != nil {
		t.Error(err)
	}
	t.Log(response.Detail)
}

func TestCheckTransactioStatus(t *testing.T) {
	response, err := sp.CheckTransactionStatus("6e1e251f-afb0-****-a097-f1ae0e0b2ce6")
	if err != nil {
		t.Error(err)
	}
	t.Log(response.Data)
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
