package main

// JSONReturn is a representation of a datajson
type JSONReturn struct {
	StatusCode        string `json:"status_code"`
	StatusMessage     string `json:"status_message"`
	TransactionID     string `json:"transaction_id"`
	OrderID           string `json:"order_id"`
	GrossAmount       string `json:"gross_amount"`
	PaymentType       string `json:"payment_type"`
	TransactionTime   string `json:"transaction_time"`
	TransactionStatus string `json:"transaction_status"`
	FraudStatus       string `json:"fraud_status"`
	BillKey           string `json:"bill_key"`
	BillerCode        string `json:"biller_code"`
	PdfURL            string `json:"pdf_url"`
	FinishRedirectURL string `json:"finish_redirect_url"`
}

// JSONBody is a representation of a datajson
type JSONBody struct {
	TransactionDetails struct {
		OrderID     string `json:"order_id"`
		GrossAmount int    `json:"gross_amount"`
	} `json:"transaction_details"`
	ItemDetails []struct {
		ID           string `json:"id"`
		Price        int    `json:"price"`
		Quantity     int    `json:"quantity"`
		Name         string `json:"name"`
		Brand        string `json:"brand"`
		Category     string `json:"category"`
		MerchantName string `json:"merchant_name"`
	} `json:"item_details"`
	CustomerDetails struct {
		FirstName      string `json:"first_name"`
		LastName       string `json:"last_name"`
		Email          string `json:"email"`
		Phone          string `json:"phone"`
		BillingAddress struct {
			FirstName string `json:"first_name"`
			LastName  string `json:"last_name"`
			//Email       string `json:"email"`
			Phone       string `json:"phone"`
			Address     string `json:"address"`
			City        string `json:"city"`
			PostalCode  string `json:"postal_code"`
			CountryCode string `json:"country_code"`
		} `json:"billing_address"`
		ShippingAddress struct {
			FirstName string `json:"first_name"`
			LastName  string `json:"last_name"`
			//Email       string `json:"email"`
			Phone       string `json:"phone"`
			Address     string `json:"address"`
			City        string `json:"city"`
			PostalCode  string `json:"postal_code"`
			CountryCode string `json:"country_code"`
		} `json:"shipping_address"`
	} `json:"customer_details"`
}

// JSONNotif is a representation of a datajson
type JSONNotif struct {
	VaNumbers []struct {
		VaNumber string `json:"va_number"`
		Bank     string `json:"bank"`
	} `json:"va_numbers"`
	TransactionTime   string        `json:"transaction_time"`
	TransactionStatus string        `json:"transaction_status"`
	TransactionID     string        `json:"transaction_id"`
	StatusMessage     string        `json:"status_message"`
	StatusCode        string        `json:"status_code"`
	SignatureKey      string        `json:"signature_key"`
	PaymentType       string        `json:"payment_type"`
	PaymentAmounts    []interface{} `json:"payment_amounts"`
	OrderID           string        `json:"order_id"`
	GrossAmount       string        `json:"gross_amount"`
	FraudStatus       string        `json:"fraud_status"`
}

// Test is a representation of a datajson
type Test struct {
	MaskedCard        string `json:"masked_card,omitempty"`
	ApprovalCode      string `json:"approval_code,omitempty"`
	Bank              string `json:"bank,omitempty"`
	Eci               string `json:"eci,omitempty"`
	TransactionTime   string `json:"transaction_time,omitempty"`
	GrossAmount       string `json:"gross_amount,omitempty"`
	OrderID           string `json:"order_id,omitempty"`
	PaymentType       string `json:"payment_type,omitempty"`
	SignatureKey      string `json:"signature_key,omitempty"`
	StatusCode        string `json:"status_code,omitempty"`
	TransactionID     string `json:"transaction_id,omitempty"`
	TransactionStatus string `json:"transaction_status,omitempty"`
	FraudStatus       string `json:"fraud_status,omitempty"`
	StatusMessage     string `json:"status_message,omitempty"`
}
