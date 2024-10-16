package models

import "encoding/json"

func UnmarshalBankListResp(data []byte) (BankListResp, error) {
	var r BankListResp
	err := json.Unmarshal(data, &r)
	return r, err
}

func (r *BankListResp) Marshal() ([]byte, error) {
	return json.Marshal(r)
}

type BankListResp struct {
	Status       bool       `json:"status"`
	ResponseCode string     `json:"ResponseCode"`
	Detail       string     `json:"detail"`
	Data         []BankList `json:"data"`
}

type BankList struct {
	BankName string `json:"bank_name"`
	BankCode string `json:"bank_code"`
}
