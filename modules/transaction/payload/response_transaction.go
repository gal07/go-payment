package payload

type ResponseTransaction struct {
	Id                string           `json:"id"`
	Externalid        string           `json:"external_id"`
	Userid            string           `json:"user_id"`
	Status            string           `json:"status"`
	MerchantName      string           `json:"merchant_name"`
	MerchantProfile   string           `json:"merchant_profile_picture_url"`
	Amount            string           `json:"amount"`
	Description       string           `json:"description"`
	ExpiryDate        string           `json:"expiry_date"`
	InvoiceUrl        string           `json:"invoice_url"`
	AvailableBank     []BankThings     `json:"available_banks"`
	AvailableRetail   []RetailThings   `json:"available_retail_outlets"`
	AvailableQr       []QrThings       `json:"available_qr_codes"`
	AvailableEwallet  []EwalletThings  `json:"available_ewallets"`
	AvailableDebit    []DebitThings    `json:"available_direct_debits"`
	AvailablePaylater []PaylaterThings `json:"available_paylaters"`
	ShouldExclude     string           `json:"should_exclude_credit_card"`
	ShouldSendEmail   string           `json:"should_send_email"`
	Created           string           `json:"created"`
	Updated           string           `json:"updated"`
	Customer          Customer         `json:"customer"`
	CustomerNotif     CustomerNotif    `json:"customer_notification_preference"`
	Items             Items            `json:"items"`
	Currency          string           `json:"currency"`
}

type BankThings struct {
	BankCode          string `json:"bank_code"`
	CollectionType    string `json:"collection_type"`
	TransferAmount    int    `json:"transfer_amount"`
	BankBranch        string `json:"bank_branch"`
	AccountHolderName string `json:"account_holder_name"`
	IdentityAmount    int    `json:"identity_amount"`
}

type RetailThings struct {
	Name string `json:"retail_outlet_name"`
}

type EwalletThings struct {
	Name string `json:"ewallet_type"`
}

type QrThings struct {
	Name string `json:"qr_code_type"`
}

type DebitThings struct {
	Name string `json:"direct_debit_type"`
}

type PaylaterThings struct {
	Name string `json:"paylater_type"`
}
