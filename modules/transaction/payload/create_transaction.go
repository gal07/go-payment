package payload

type ReqCreateTransaction struct {
	External_id string `json:"external_id"`
	Amount      string `json:"amount"`
	Customer
}

type Customer struct {
	Given_name    string `json:"given_names"`
	Surname       string `json:"surname"`
	Email         string `json:"email"`
	Mobile_number string `json:"mobile_number"`
}
