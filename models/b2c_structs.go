package models

import "encoding/json"

func UnmarshalB2CRequest(data []byte) (B2CRequest, error) {
	var r B2CRequest
	err := json.Unmarshal(data, &r)
	return r, err
}

func (r *B2CRequest) Marshal() ([]byte, error) {
	return json.Marshal(r)
}

type B2CRequest struct {
	MerchantCode                 string `json:"MerchantCode"`
	MerchantTransactionReference string `json:"MerchantTransactionReference"`
	Amount                       int64  `json:"Amount"`
	Currency                     string `json:"Currency"`
	ReceiverNumber               string `json:"ReceiverNumber"`
	Channel                      string `json:"Channel"`
	Reason                       string `json:"Reason"`
	CallBackURL                  string `json:"CallBackURL"`
}

func UnmarshalB2CResponse(data []byte) (B2CResponse, error) {
	var r B2CResponse
	err := json.Unmarshal(data, &r)
	return r, err
}

func (r *B2CResponse) Marshal() ([]byte, error) {
	return json.Marshal(r)
}

type B2CResponse struct {
	Status                   bool   `json:"status"`
	Detail                   string `json:"detail"`
	B2CRequestID             string `json:"B2CRequestID"`
	ConversationID           string `json:"ConversationID"`
	OriginatorConversationID string `json:"OriginatorConversationID"`
	ResponseCode             string `json:"ResponseCode"`
	ResponseDescription      string `json:"ResponseDescription"`
}
