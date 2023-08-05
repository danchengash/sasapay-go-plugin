package models

import "encoding/json"

func UnmarshalAccountValidationRes(data []byte) (AccountValidationRes, error) {
	var r AccountValidationRes
	err := json.Unmarshal(data, &r)
	return r, err
}

func (r *AccountValidationRes) Marshal() ([]byte, error) {
	return json.Marshal(r)
}

type AccountValidationRes struct {
	Status         bool            `json:"status"`
	Detail         string          `json:"detail"`
	AccountDetails *AccountDetails `json:"account_details"`
}

type AccountDetails struct {
	ChannelCode   string `json:"channel_code"`
	ChannelName   string `json:"channel_name"`
	AccountNumber string `json:"account_number"`
	AccountName   string `json:"account_name"`
}
