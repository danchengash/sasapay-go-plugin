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
	TransactionType    string `json:"TransactionType"`
	TransactionDate    string `json:"TransactionDate"`
	CheckoutID         string `json:"CheckoutId"`
	TransactionAmount  int64  `json:"TransactionAmount"`
	Paid               bool   `json:"Paid"`
	AmountPaid         int64  `json:"AmountPaid"`
	PaidDate           string `json:"PaidDate"`
	SourceChannel      string `json:"SourceChannel"`
	DestinationChannel string `json:"DestinationChannel"`
}
