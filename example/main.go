package main

import (
	"fmt"
	"testing"

	"github.com/danchengash/sasapay-go-plugin"
	"github.com/danchengash/sasapay-go-plugin/models"
)

var clientSecret = "XXXXX"
var clientId = "XX"
var sp = sasapay.NewSasaPay(clientId, clientSecret, "XX", int(sasapay.Production), true)

func main() {
	sp = sasapay.NewSasaPay(clientId, clientSecret, "XX", int(sasapay.Production), true)

	TestCheckTransactioStatus(&testing.T{})
	// res, _ := sp.CardPayment(models.CardPaymentRequest{

	// 	MerchantCode:         sp.MerchantCode,
	// 	Amount:               "100",
	// 	Reference:            "677",
	// 	Description:          "dtests",
	// 	Currency:             "KES",
	// 	CallbackURL:          "https://c6d5-41-90-115-26.ngrok-free.app",
	// 	SuccessURL:           "",
	// 	FailureURL:           "",
	// 	SasaPayWalletEnabled: true,
	// 	MpesaEnabled:         false,
	// 	CardEnabled:          true,
	// 	AirtelEnabled:        false,
	// })
	// fmt.Println(res.CheckoutURL)
	// res, err := sp.GetTransactions(1, 2, "2023-2-23", "2023-10-28")
	// if err != nil {
	// 	fmt.Println(err.Error())
	// } else {
	// 	for _, v := range res.Data.Transactions {
	// 		fmt.Printf("%f - %f \n", v.TransactionAmount, v.TransactionCharges)
	// 	}
	// }
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
	response, err := sp.CheckTransactionStatus(models.CheckTransactionStatusRequest{CheckoutRequestID: "XXXXX", MerchantTransactionReference: "XXXX", TransactionCode: "SPEJZM2W9YRN75M"})
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
