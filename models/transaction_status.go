package models

import "encoding/json"

func UnmarshalTransactionStatusResponse(data []byte) (TransactionStatusResponse, error) {
	var r TransactionStatusResponse
	err := json.Unmarshal(data, &r)
	return r, err
}

func (r *TransactionStatusResponse) Marshal() ([]byte, error) {
	return json.Marshal(r)
}

type TransactionStatusResponse struct {
	Status  bool   `json:"status"`
	Message string `json:"message"`
	Data    Data   `json:"data"`
}

type Data struct {
	TransactionType           string  `json:"TransactionType"`
	TransactionDate           string  `json:"TransactionDate"`
	CheckoutID                string  `json:"CheckoutId"`
	MerchantReference         string  `json:"MerchantReference"`
	TransactionAmount         float64 `json:"TransactionAmount"`
	Paid                      bool    `json:"Paid"`
	AmountPaid                float64 `json:"AmountPaid"`
	PaidDate                  string  `json:"PaidDate"`
	SourceChannel             string  `json:"SourceChannel"`
	DestinationChannel        string  `json:"DestinationChannel"`
	TransID                   string  `json:"TransID"`
	TransactionCode           string  `json:"TransactionCode"`
	ThirdPartyTransactionCode string  `json:"ThirdPartyTransactionCode"`
	ResultCode                string  `json:"ResultCode,omitempty" bson:"result_code"`
	ResultDescription         string  `json:"ResultDescription,omitempty" bson:"result_description"`
}

func UnmarshalCheckTransactionStatusRequest(data []byte) (CheckTransactionStatusRequest, error) {
	var r CheckTransactionStatusRequest
	err := json.Unmarshal(data, &r)
	return r, err
}

func (r *CheckTransactionStatusRequest) Marshal() ([]byte, error) {
	return json.Marshal(r)
}

type CheckTransactionStatusRequest struct {
	MerchantCode                 string  `json:"MerchantCode"`
	CheckoutRequestID            *string `json:"CheckoutRequestId"`
	MerchantTransactionReference *string `json:"MerchantTransactionReference"`
	TransactionCode              *string `json:"TransactionCode"`
	CallBackUrl                  string  `json:"CallBackUrl,omitempty" bson:"call_back_url"`
}
