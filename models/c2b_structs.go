package models

import "encoding/json"

func UnmarshalC2BRequest(data []byte) (C2BRequest, error) {
	var r C2BRequest
	err := json.Unmarshal(data, &r)
	return r, err
}

func (r *C2BRequest) Marshal() ([]byte, error) {
	return json.Marshal(r)
}

type C2BRequest struct {
	MerchantCode     string `json:"MerchantCode"`
	NetworkCode      string `json:"NetworkCode"`
	PhoneNumber      string `json:"PhoneNumber"`
	TransactionDesc  string `json:"TransactionDesc"`
	AccountReference string `json:"AccountReference"`
	Currency         string `json:"Currency"`
	Amount           int64  `json:"Amount"`
	CallBackURL      string `json:"CallBackURL"`
}

func UnmarshalC2BResponse(data []byte) (C2BResponse, error) {
	var r C2BResponse
	err := json.Unmarshal(data, &r)
	return r, err
}

func (r *C2BResponse) Marshal() ([]byte, error) {
	return json.Marshal(r)
}

type C2BResponse struct {
	Status              bool   `json:"status"`
	Detail              string `json:"detail"`
	PaymentGateway      string `json:"PaymentGateway"`
	MerchantRequestID   string `json:"MerchantRequestID"`
	CheckoutRequestID   string `json:"CheckoutRequestID"`
	ResponseCode        string `json:"ResponseCode"`
	ResponseDescription string `json:"ResponseDescription"`
	CustomerMessage     string `json:"CustomerMessage"`
}
