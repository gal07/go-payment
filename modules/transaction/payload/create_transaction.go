package payload

type ReqCreateTransaction struct {
	Externalid      string        `json:"external_id"`
	Amount          string        `json:"amount"`
	Description     string        `json:"description"`
	Customer        Customer      `json:"customer"`
	CustomerNotif   CustomerNotif `json:"customer_notification_preference"`
	Items           []Items       `json:"items"`
	InvoiceDuration int           `json:"invoice_duration"`
}

type Customer struct {
	Givenname    string `json:"given_names"`
	Surname      string `json:"surname"`
	Email        string `json:"email"`
	Mobilenumber string `json:"mobile_number"`
}

type CustomerNotif struct {
	InvoiceCreated []string `json:"invoice_created"`
	InvoicePaid    []string `json:"invoice_paid"`
}

type Items struct {
	Name     string `json:"name"`
	Quantity int    `json:"quantity"`
	Price    int    `json:"price"`
}
