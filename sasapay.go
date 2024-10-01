package sasapay

import (
	"encoding/base64"
	"encoding/json"
	"errors"
	"fmt"
	"time"

	"github.com/danchengash/sasapay-go-plugin/helpers"
	"github.com/danchengash/sasapay-go-plugin/models"
)

type SasaPay struct {
	Environment  int
	ClientId     string
	ClientSecret string
	MerchantCode string
	cacheToken   models.AccessTokenResponse
	Showlogs     bool
}

func NewSasaPay(clientId string, clientSecret string, merchantcode string, environment int, showLogs bool) *SasaPay {
	var accessToken = models.AccessTokenResponse{}

	return &SasaPay{
		ClientId:     clientId,
		ClientSecret: clientSecret,
		MerchantCode: merchantcode,
		Environment:  environment,
		cacheToken:   accessToken,
		Showlogs:     showLogs,
	}

}

// setAccessToken returns a time bound access token to call allowed APIs.
// This token should be used in all other subsequent responses to the APIs
func (s *SasaPay) setAccessToken() (string, error) {
	if time.Until(s.cacheToken.ExpiresAt.UTC()).Seconds() > 0 {
		fmt.Println("cache token")
		return s.cacheToken.AccessToken, nil
	}
	url := s.baseURL() + SetAccessTokenUrl
	b := []byte(s.ClientId + ":" + s.ClientSecret)
	encoded := base64.StdEncoding.EncodeToString(b)
	headers := make(map[string]string)
	headers["Authorization"] = "Basic " + encoded
	res, err := helpers.NewReq(url, nil, &headers, false)
	if err != nil {
		fmt.Println(err)
		return "", &models.RequestError{StatusCode: res.StatusCode(), Message: string(res.Body()), Url: res.LocalAddr().String()}
	}
	if res.StatusCode() >= 200 && res.StatusCode() <= 300 {
		s.cacheToken, err = models.UnmarshalAccessTokenResponse(res.Body())
		s.cacheToken.ExpiresAt = time.Now().Add(time.Duration(s.cacheToken.ExpiresIn) * time.Second)
		if err != nil {
			return "", err
		}
	}
	return s.cacheToken.AccessToken, nil
}

func (s *SasaPay) baseURL() string {
	if s.Environment == int(Production) {
		return "https://api.sasapay.app/api/v1/"
	}
	return "https://api.sasapay.me/api/v1/"
}

func (s *SasaPay) RegisterCallBackUrl(confirmationUrl string) (*models.RegisterConfirmationURLResponse, error) {
	token, err := s.setAccessToken()
	if err != nil {
		return nil, err
	}

	headers := make(map[string]string)
	params := make(map[string]interface{})
	params["MerchantCode"] = s.MerchantCode
	params["ConfirmationUrl"] = confirmationUrl
	reqEntityBytes, _ := json.Marshal(params)
	headers["Authorization"] = "Bearer " + token
	url := s.baseURL() + registerCallBack

	res, err := helpers.NewReq(url, &reqEntityBytes, &headers, s.Showlogs)
	if err != nil {
		return nil, &models.RequestError{StatusCode: res.StatusCode(), Message: string(res.Body()), Url: res.LocalAddr().String()}
	}
	resp, err := models.UnmarshalRegisterConfirmationURLResponse(res.Body())
	if err != nil {
		return nil, err
	}

	return &resp, nil
}

func (s *SasaPay) Customer2Business(body models.C2BRequest) (*models.C2BResponse, error) {
	token, err := s.setAccessToken()
	if err != nil {
		return nil, err
	}
	headers := make(map[string]string)
	headers["Authorization"] = "Bearer " + token
	url := s.baseURL() + c2burl
	reqbody, err := body.Marshal()
	if err != nil {
		return nil, err
	}

	resp, err := helpers.NewReq(url, &reqbody, &headers, s.Showlogs)
	if err != nil {
		return nil, err
	}
	if resp.StatusCode() != 200 {
		errRespose, err := models.UnmarshalAPIResponse(resp.Body())
		if err != nil {
			return nil, errors.New(string(resp.Body()))
		}
		return nil, errors.New(errRespose.Detail)
	}
	response, err := models.UnmarshalC2BResponse(resp.Body())
	if err != nil {
		return nil, err
	}

	return &response, nil

}
func (s *SasaPay) C2BProcess(checkoutRequestID string, otpCode string) (*models.APIResponse, error) {
	params := make(map[string]interface{})
	params["CheckoutRequestID"] = checkoutRequestID
	params["MerchantCode"] = s.MerchantCode
	params["VerificationCode"] = otpCode

	token, err := s.setAccessToken()
	if err != nil {
		return nil, err
	}
	headers := make(map[string]string)
	headers["Authorization"] = "Bearer " + token
	url := s.baseURL() + c2bProcess

	paramsBytes, _ := json.Marshal(params)
	resp, err := helpers.NewReq(url, &paramsBytes, &headers, s.Showlogs)
	if err != nil {
		return nil, err
	}
	if resp.StatusCode() != 200 {
		errRespose, err := models.UnmarshalAPIResponse(resp.Body())
		if err != nil {
			return nil, errors.New(string(resp.Body()))
		}
		return nil, errors.New(errRespose.Detail)
	}
	response, err := models.UnmarshalAPIResponse(resp.Body())
	if err != nil {
		return nil, err
	}
	return &response, nil

}

/// look at the README File to get a list of all the channel codes

func (s *SasaPay) Business2Customer(params models.B2CRequest) (*models.B2CResponse, error) {
	token, err := s.setAccessToken()
	if err != nil {
		return nil, err
	}
	headers := make(map[string]string)
	headers["Authorization"] = "Bearer " + token
	url := s.baseURL() + b2cUrl

	body, _ := params.Marshal()

	resp, err := helpers.NewReq(url, &body, &headers, s.Showlogs)
	if err != nil {
		return nil, err
	}

	if resp.StatusCode() != 200 {
		errRespose, err := models.UnmarshalAPIResponse(resp.Body())
		if err != nil {
			return nil, errors.New(string(resp.Body()))
		}
		return nil, errors.New(errRespose.Detail)
	}
	response, err := models.UnmarshalB2CResponse(resp.Body())
	if err != nil {
		return nil, err
	}

	return &response, nil
}

func (s *SasaPay) Business2Business(params models.B2BRequest) (*models.B2BResponse, error) {
	token, err := s.setAccessToken()
	if err != nil {
		return nil, err
	}
	headers := make(map[string]string)
	headers["Authorization"] = "Bearer " + token
	url := s.baseURL() + b2bUrl

	body, _ := params.Marshal()
	resp, err := helpers.NewReq(url, &body, &headers, s.Showlogs)
	if err != nil {
		return nil, err
	}
	if resp.StatusCode() != 200 {
		errRespose, err := models.UnmarshalAPIResponse(resp.Body())
		if err != nil {
			return nil, errors.New(string(resp.Body()))
		}
		return nil, errors.New(errRespose.Detail)
	}
	response, err := models.UnmarshalB2BResponse(resp.Body())
	if err != nil {
		return nil, err
	}

	return &response, nil
}

func (s *SasaPay) CheckTransactionStatus(checkSta models.CheckTransactionStatusRequest) (*models.TransactionStatusResponse, error) {
	token, err := s.setAccessToken()
	if err != nil {
		return nil, err
	}
	headers := make(map[string]string)
	params := make(map[string]interface{})

	headers["Authorization"] = "Bearer " + token
	url := s.baseURL() + check_transaction_status
	if checkSta.CallBackUrl != "" {
		url = s.baseURL() + check_transaction_status_query
		params["CallbackUrl"] = checkSta.CallBackUrl
	}
	checkSta.MerchantCode = s.MerchantCode
	if checkSta.CheckoutRequestID != nil {
		params["CheckoutRequestId"] = checkSta.CheckoutRequestID
	}
	params["MerchantCode"] = s.MerchantCode
	params["MerchantTransactionReference"] = checkSta.MerchantTransactionReference
	params["TransactionCode"] = checkSta.TransactionCode
	paramsBytes, _ := json.Marshal(params)
	resp, err := helpers.NewReq(url, &paramsBytes, &headers, s.Showlogs)
	if err != nil {
		return nil, err
	}
	if resp.StatusCode() != 200 {
		// if resp.StatusCode() == 400 {
		// 	return nil, errors.New(string(resp.Body()))
		// }
		errRespose, err := models.UnmarshalAPIResponse(resp.Body())
		if err != nil {
			return nil, fmt.Errorf("%s-%s", err.Error(), resp.Body())
		}
		if errRespose.Detail == "" {
			return nil, errors.New(errRespose.Message)
		} else {
			return nil, errors.New(errRespose.Detail)
		}
	}
	response, err := models.UnmarshalTransactionStatusResponse(resp.Body())
	if err != nil {
		return nil, err
	}
	return &response, nil
}

func (s *SasaPay) VerifyTransaction(transactionCode string) (*models.VerifyTransactionResponse, error) {
	token, err := s.setAccessToken()
	if err != nil {
		return nil, err
	}
	headers := make(map[string]string)
	headers["Authorization"] = "Bearer " + token
	url := s.baseURL() + verify_transaction
	params := make(map[string]interface{})
	params["MerchantCode"] = s.MerchantCode
	params["TransactionCode"] = transactionCode
	paramsBytes, _ := json.Marshal(params)
	resp, err := helpers.NewReq(url, &paramsBytes, &headers, s.Showlogs)
	if err != nil {
		return nil, err
	}
	if resp.StatusCode() != 200 {
		errRespose, err := models.UnmarshalAPIRespSecond(resp.Body())
		if err != nil {
			return nil, errors.New(string(resp.Body()))
		}
		return nil, errors.New(errRespose.Message)
	}
	response, err := models.UnmarshalVerifyTransactionResponse(resp.Body())
	if err != nil {
		return nil, err
	}
	return &response, nil
}
func (s *SasaPay) CheckMerchantBalance() (*models.MerchantBalanceResp, error) {
	token, err := s.setAccessToken()
	if err != nil {
		return nil, err
	}
	headers := make(map[string]string)
	headers["Authorization"] = "Bearer " + token
	url := s.baseURL() + check_merchant_balance + s.MerchantCode
	resp, err := helpers.NewReq(url, nil, &headers, s.Showlogs)
	if err != nil {
		return nil, err
	}
	if resp.StatusCode() != 200 {
		errRespose, err := models.UnmarshalAPIResponse(resp.Body())
		if err != nil {
			return nil, errors.New(string(resp.Body()))
		}
		return nil, errors.New(errRespose.Detail)
	}
	response, err := models.UnmarshalMerchantBalanceResp(resp.Body())
	if err != nil {
		return nil, err
	}
	return &response, nil
}
func (s *SasaPay) Business2Benefiary(params models.Business2BeneficiaryReq) (*models.Business2BeneficiaryResp, error) {
	token, err := s.setAccessToken()
	if err != nil {
		return nil, err
	}
	headers := make(map[string]string)
	headers["Authorization"] = "Bearer " + token
	url := s.baseURL() + business2Beneficiary

	body, _ := params.Marshal()
	resp, err := helpers.NewReq(url, &body, &headers, s.Showlogs)
	if err != nil {
		return nil, err
	}
	if resp.StatusCode() != 200 {
		errRespose, err := models.UnmarshalAPIRespSecond(resp.Body())
		if err != nil {
			return nil, errors.New(string(resp.Body()))
		}
		return nil, errors.New(errRespose.Message)
	}
	response, err := models.UnmarshalBusiness2BeneficiaryResp(resp.Body())
	if err != nil {
		return nil, err
	}
	return &response, nil
}
func (s *SasaPay) AccountValidate(acct string, channel string) (*models.AccountValidationRes, error) {
	token, err := s.setAccessToken()
	if err != nil {
		return nil, err
	}
	headers := make(map[string]string)
	headers["Authorization"] = "Bearer " + token
	url := s.baseURL() + accountValidation

	params := make(map[string]interface{})
	params["merchant_code"] = s.MerchantCode
	params["channel_code"] = channel
	params["account_number"] = acct

	paramsBytes, _ := json.Marshal(params)

	resp, err := helpers.NewReq(url, &paramsBytes, &headers, s.Showlogs)
	if err != nil {
		return nil, err
	}
	if resp.StatusCode() != 200 {
		errRespose, err := models.UnmarshalAPIResponse(resp.Body())
		if err != nil {
			return nil, errors.New(string(resp.Body()))
		}
		return nil, errors.New(errRespose.Detail)
	}
	response, err := models.UnmarshalAccountValidationRes(resp.Body())
	if err != nil {
		return nil, err
	}
	return &response, nil
}
func (s *SasaPay) GetTransactions(page int, pageSize int, startEndate string, endDate string, transactionCode string) (*models.TransactionsResponse, error) {

	token, err := s.setAccessToken()
	if err != nil {
		return nil, err
	}
	headers := make(map[string]string)
	headers["Authorization"] = "Bearer " + token
	url := s.baseURL() + transactionsUrl + s.MerchantCode + fmt.Sprintf("&page=%d&page_size=%d&start_date=%s&end_date=%s", page, pageSize, startEndate, endDate)
	if transactionCode != "" {
		url = url + fmt.Sprintf("&transaction_code=%s", transactionCode)
	}
	resp, err := helpers.NewReq(url, nil, &headers, s.Showlogs)
	if err != nil {
		return nil, err
	}
	if resp.StatusCode() != 200 {
		errRespose, err := models.UnmarshalAPIRespSecond(resp.Body())
		if err != nil {
			return nil, errors.New(string(resp.Body()))
		}
		return nil, errors.New(errRespose.Message)
	}
	response, err := models.UnmarshalTransactionsResponse(resp.Body())
	if err != nil {
		return nil, err
	}
	return &response, nil
}

func (s *SasaPay) CardPayment(param models.CardPaymentRequest) (*models.CardPaymentResponse, error) {
	token, err := s.setAccessToken()
	if err != nil {
		return nil, err
	}
	headers := make(map[string]string)
	headers["Authorization"] = "Bearer " + token
	url := s.baseURL() + cardPaymentUrl
	paramsBytes, _ := param.Marshal()
	resp, err := helpers.NewReq(url, &paramsBytes, &headers, s.Showlogs)
	if err != nil {
		return nil, err
	}
	if resp.StatusCode() != 200 {
		errRespose, err := models.UnmarshalAPIResponse(resp.Body())
		if err != nil {
			return nil, errors.New(string(resp.Body()))
		}
		return nil, errors.New(errRespose.Detail)
	}
	response, err := models.UnmarshalCardPaymentResponse(resp.Body())
	if err != nil {
		return nil, err
	}
	return &response, nil
}
