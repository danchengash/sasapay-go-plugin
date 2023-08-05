package models

import "encoding/json"

func UnmarshalB2BRequest(data []byte) (B2BRequest, error) {
	var r B2BRequest
	err := json.Unmarshal(data, &r)
	return r, err
}

func (r *B2BRequest) Marshal() ([]byte, error) {
	return json.Marshal(r)
}

type B2BRequest struct {
	MerchantCode                 string `json:"MerchantCode"`
	MerchantTransactionReference string `json:"MerchantTransactionReference"`
	Currency                     string `json:"Currency"`
	Amount                       int64  `json:"Amount"`
	ReceiverMerchantCode         string `json:"ReceiverMerchantCode"`
	CallBackURL                  string `json:"CallBackURL"`
	Reason                       string `json:"Reason"`
}

func UnmarshalB2BResponse(data []byte) (B2BResponse, error) {
	var r B2BResponse
	err := json.Unmarshal(data, &r)
	return r, err
}

func (r *B2BResponse) Marshal() ([]byte, error) {
	return json.Marshal(r)
}

type B2BResponse struct {
	Status                   bool   `json:"status"`
	Detail                   string `json:"detail"`
	B2BRequestID             string `json:"B2BRequestID"`
	ConversationID           string `json:"ConversationID"`
	OriginatorConversationID string `json:"OriginatorConversationID"`
	ResponseCode             string `json:"ResponseCode"`
	ResponseDescription      bool   `json:"ResponseDescription"`
}
