package models

// This file was generated from JSON Schema using quicktype, do not modify it directly.
// To parse and unparse this JSON data, add this code to your project and do:
//
//    transactionsResponse, err := UnmarshalTransactionsResponse(bytes)
//    bytes, err = transactionsResponse.Marshal()

import "encoding/json"

func UnmarshalTransactionsResponse(data []byte) (TransactionsResponse, error) {
	var r TransactionsResponse
	err := json.Unmarshal(data, &r)
	return r, err
}

func (r *TransactionsResponse) Marshal() ([]byte, error) {
	return json.Marshal(r)
}

type TransactionsResponse struct {
	Status      bool            `json:"status"`
	Message     string          `json:"message"`
	Detail      string          `json:"detail"`
	Links       Links           `json:"links"`
	Count       int64           `json:"count"`
	CurrentPage int64           `json:"current_page"`
	Pages       int64           `json:"pages"`
	Data        DataTransaction `json:"data"`
}

type DataTransaction struct {
	Transactions []Transaction `json:"transactions"`
}

type Transaction struct {
	ID                     int64          `json:"id"`
	MerchantCode           string         `json:"merchant_code"`
	TransactionAmount      float64        `json:"transaction_amount"`
	TransactionCharges     float64        `json:"transaction_charges"`
	TransactionType        string         `json:"transaction_type"`
	TransactionCode        string         `json:"transaction_code"`
	TransactionDescription string         `json:"transaction_description"`
	TransactionReference   string         `json:"transaction_reference"`
	TransactionDate        string         `json:"transaction_date"`
	PaymentDetails         PaymentDetails `json:"payment_details"`
	ResultCode             string         `json:"result_code"`
	ResultDescription      string         `json:"result_description"`
	ReversalStatus         string         `json:"reversal_status"`
	CreatedDate            string         `json:"created_date"`
}

type PaymentDetails struct {
	PartyBAccountNumber         string `json:"party_B_account_number"`
	PartyBAccountName           string `json:"party_B_account_name"`
	ChannelName                 string `json:"channel_name"`
	ChannelTransactionReference string `json:"channel_transaction_reference"`
}

type Links struct {
	Next     string  `json:"next"`
	Previous *string `json:"previous"`
}
