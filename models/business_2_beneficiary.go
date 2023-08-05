package models

import "encoding/json"

func UnmarshalBusiness2BeneficiaryReq(data []byte) (Business2BeneficiaryReq, error) {
	var r Business2BeneficiaryReq
	err := json.Unmarshal(data, &r)
	return r, err
}

func (r *Business2BeneficiaryReq) Marshal() ([]byte, error) {
	return json.Marshal(r)
}

type Business2BeneficiaryReq struct {
	TransactionReference     string `json:"TransactionReference"`
	SenderMerchantCode       string `json:"SenderMerchantCode"`
	ReceiverMerchantCode     string `json:"ReceiverMerchantCode"`
	BeneficiaryAccountNumber string `json:"BeneficiaryAccountNumber"`
	Amount                   int64  `json:"Amount"`
	TransactionFee           int64  `json:"TransactionFee"`
	Reason                   string `json:"Reason"`
	CallBackURL              string `json:"CallBackUrl"`
}


func UnmarshalBusiness2BeneficiaryResp(data []byte) (Business2BeneficiaryResp, error) {
	var r Business2BeneficiaryResp
	err := json.Unmarshal(data, &r)
	return r, err
}

func (r *Business2BeneficiaryResp) Marshal() ([]byte, error) {
	return json.Marshal(r)
}

type Business2BeneficiaryResp struct {
	Status            bool   `json:"status"`
	Message           string `json:"message"`
	MerchantRequestID string `json:"MerchantRequestID"`
	CheckoutRequestID string `json:"CheckoutRequestID"`
}
