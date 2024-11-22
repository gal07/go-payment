package models

type Transaction struct {
	Xendit_id          string `json:"xendit_id"`
	Xendit_external_id string `json:"xendit_external_id"`
	User_id            string `json:"user_id"`
	Currency           string `json:"currency"`
	Merchant_name      string `json:"merchant_name"`
	Amount             string `json:"amount"`
	Status             string `json:"status"`
	Invoice_url        string `json:"invoice_url"`
	Description        string `json:"description"`
	Expiry_date        string `json:"expiry_date"`
	Created_at         string `json:"created_at"`
	Updated_at         string `json:"updated_at"`
}
