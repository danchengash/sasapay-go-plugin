package models

import "encoding/json"

func UnmarshalMerchantBalanceResp(data []byte) (MerchantBalanceResp, error) {
	var r MerchantBalanceResp
	err := json.Unmarshal(data, &r)
	return r, err
}

func (r *MerchantBalanceResp) Marshal() ([]byte, error) {
	return json.Marshal(r)
}

type MerchantBalanceResp struct {
	StatusCode string       `json:"statusCode"`
	Message    string       `json:"message"`
	Data       DataMerchant `json:"data"`
}

type DataMerchant struct {
	CurrencyCode      string    `json:"CurrencyCode"`
	OrgAccountBalance float64   `json:"OrgAccountBalance"`
	Accounts          []Account `json:"Accounts"`
}

type Account struct {
	AccountLabel   string  `json:"account_label"`
	AccountBalance float64 `json:"account_balance"`
}
