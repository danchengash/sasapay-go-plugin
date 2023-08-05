package models

import (
	"encoding/json"
	"fmt"
	"time"
)

func UnmarshalAccessTokenResponse(data []byte) (AccessTokenResponse, error) {
	var r AccessTokenResponse
	err := json.Unmarshal(data, &r)
	return r, err
}

func (r *AccessTokenResponse) Marshal() ([]byte, error) {
	return json.Marshal(r)
}

type AccessTokenResponse struct {
	Status      bool   `json:"status"`
	Detail      string `json:"detail"`
	AccessToken string `json:"access_token"`
	ExpiresIn   int    `json:"expires_in"`
	ExpiresAt   time.Time
	TokenType   string `json:"token_type"`
	Scope       string `json:"scope"`
}

func UnmarshalInvalidAccessTokenResponse(data []byte) (InvalidAccessTokenResponse, error) {
	var r InvalidAccessTokenResponse
	err := json.Unmarshal(data, &r)
	return r, err
}

func (r *InvalidAccessTokenResponse) Marshal() ([]byte, error) {
	return json.Marshal(r)
}

type InvalidAccessTokenResponse struct {
	Detail     string `json:"detail"`
	Status     bool   `json:"status"`
	StatusCode int64  `json:"status_code"`
}

type RequestError struct {
	StatusCode int

	Message string
	Url     string
}

func (r *RequestError) Error() string {
	return fmt.Sprintf("url is: %s \n status code is: %d \n  and body is : %s", r.Url, r.StatusCode, r.Message)

}

func UnmarshalRegisterConfirmationURLResponse(data []byte) (RegisterConfirmationURLResponse, error) {
	var r RegisterConfirmationURLResponse
	err := json.Unmarshal(data, &r)
	return r, err
}

func (r *RegisterConfirmationURLResponse) Marshal() ([]byte, error) {
	return json.Marshal(r)
}

type RegisterConfirmationURLResponse struct {
	Status bool   `json:"status"`
	Detail string `json:"detail"`
}
