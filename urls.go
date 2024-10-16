package sasapay

const (
	SetAccessTokenUrl              = "auth/token/?grant_type=client_credentials"
	registerCallBack               = "payments/register-ipn-url/"
	c2burl                         = "payments/request-payment/"
	c2bProcess                     = "payments/process-payment/"
	b2cUrl                         = "payments/b2c/"
	b2bUrl                         = "payments/b2b/"
	check_transaction_status       = "transactions/status/"
	check_transaction_status_query = "transactions/status-query/"

	verify_transaction     = "transactions/verify/"
	check_merchant_balance = "payments/check-balance/?MerchantCode="
	business2Beneficiary   = "payments/b2c/beneficiary/"
	accountValidation      = "accounts/account-validation/"
	transactionsUrl        = "transactions/?merchant_code="
	cardPaymentUrl         = "payments/card-payments/"
	getBankList            = "payments/channel-codes/"
)
