package models

// This file was generated from JSON Schema using quicktype, do not modify it directly.
// To parse and unparse this JSON data, add this code to your project and do:
//
//    cardPaymentRequest, err := UnmarshalCardPaymentRequest(bytes)
//    bytes, err = cardPaymentRequest.Marshal()

import "encoding/json"

func UnmarshalCardPaymentRequest(data []byte) (CardPaymentRequest, error) {
	var r CardPaymentRequest
	err := json.Unmarshal(data, &r)
	return r, err
}

func (r *CardPaymentRequest) Marshal() ([]byte, error) {
	return json.Marshal(r)
}

type CardPaymentRequest struct {
	MerchantCode         string `json:"MerchantCode"`
	Amount               string `json:"Amount"`
	Reference            string `json:"Reference"`
	Description          string `json:"Description"`
	Currency             string `json:"Currency"`
	CallbackURL          string `json:"CallbackUrl"`
	SuccessURL           string `json:"SuccessUrl"`
	FailureURL           string `json:"FailureUrl"`
	SasaPayWalletEnabled bool   `json:"SasaPayWalletEnabled"`
	MpesaEnabled         bool   `json:"MpesaEnabled"`
	CardEnabled          bool   `json:"CardEnabled"`
	AirtelEnabled        bool   `json:"AirtelEnabled"`
}



func UnmarshalCardPaymentResponse(data []byte) (CardPaymentResponse, error) {
	var r CardPaymentResponse
	err := json.Unmarshal(data, &r)
	return r, err
}

func (r *CardPaymentResponse) Marshal() ([]byte, error) {
	return json.Marshal(r)
}

type CardPaymentResponse struct {
	Status              bool   `json:"status"`
	Detail              string `json:"detail"`
	PaymentGateway      string `json:"PaymentGateway"`
	MerchantRequestID   string `json:"MerchantRequestID"`
	CheckoutRequestID   string `json:"CheckoutRequestID"`
	ResponseCode        string `json:"ResponseCode"`
	ResponseDescription string `json:"ResponseDescription"`
	CheckoutURL         string `json:"CheckoutUrl"`
}
