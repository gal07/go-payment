package usecase

import (
	"bytes"
	"context"
	"encoding/base64"
	"encoding/json"
	"fmt"
	"go-payment/modules/transaction/payload"
	"io"
	"net/http"
	"os"
)

func (s *TransactionUseCase) Transaction(ctx context.Context, req payload.ReqCreateTransaction) (res interface{}, err error) {

	// Fill optional requirement
	req.CustomerNotif.InvoiceCreated = []string{"whatsapp", "email"}
	req.CustomerNotif.InvoicePaid = []string{"whatsapp", "email"}
	req.Description = "Pembelian barang di revstore"
	req.InvoiceDuration = 600

	reqBodyBytes := new(bytes.Buffer)
	json.NewEncoder(reqBodyBytes).Encode(req)
	reqBodyBytes.Bytes() // this is the []byte

	// Execute
	basicauth := base64.StdEncoding.EncodeToString([]byte(os.Getenv("GP_XENDIT_API_SANDBOX")))
	request, err := http.NewRequest("POST", "https://api.xendit.co/v2/invoices", bytes.NewBuffer(reqBodyBytes.Bytes()))
	request.Header.Set("Content-Type", "application/json; charset=UTF-8")
	request.Header.Set("Authorization", "Basic "+basicauth)

	client := &http.Client{}
	response, err := client.Do(request)

	if err != nil {
		panic(err)
	}
	defer response.Body.Close()

	var result payload.ResponseTransaction
	body, _ := io.ReadAll(response.Body)
	if err := json.Unmarshal(body, &result); err != nil { // Parse []byte to go struct pointer
		fmt.Println("Can not unmarshal JSON")
	}
	res = body

	return result, err

}
