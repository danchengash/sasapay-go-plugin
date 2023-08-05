package models

import "encoding/json"

func UnmarshalVerifyTransactionResponse(data []byte) (VerifyTransactionResponse, error) {
	var r VerifyTransactionResponse
	err := json.Unmarshal(data, &r)
	return r, err
}

func (r *VerifyTransactionResponse) Marshal() ([]byte, error) {
	return json.Marshal(r)
}

type VerifyTransactionResponse struct {
	StatusCode int64      `json:"statusCode"`
	Detail     string     `json:"detail"`
	Data       DataVerify `json:"data"`
}

type DataVerify struct {
	TransactionType   string `json:"TransactionType"`
	TransID           string `json:"TransID"`
	TransTime         string `json:"TransTime"`
	TransAmount       string `json:"TransAmount"`
	MerchantCode      string `json:"MerchantCode"`
	BillRefNumber     string `json:"BillRefNumber"`
	InvoiceNumber     string `json:"InvoiceNumber"`
	OrgAccountBalance string `json:"OrgAccountBalance"`
	CustomerMobile    string `json:"CustomerMobile"`
	FirstName         string `json:"FirstName"`
	MiddleName        string `json:"MiddleName"`
	LastName          string `json:"LastName"`
}
