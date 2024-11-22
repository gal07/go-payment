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
)

func (s *TransactionUseCase) Transaction(ctx context.Context, req payload.ReqCreateTransaction) (res interface{}, err error) {

	// Fill optional requirement
	req.CustomerNotif.InvoiceCreated = []string{"whatsapp", "email"}
	req.CustomerNotif.InvoicePaid = []string{"whatsapp", "email"}
	req.Description = "Pembelian barang di revstore"

	reqBodyBytes := new(bytes.Buffer)
	json.NewEncoder(reqBodyBytes).Encode(req)
	reqBodyBytes.Bytes() // this is the []byte

	// Execute
	basicauth := base64.StdEncoding.EncodeToString([]byte("xnd_production_BFiEGSPcIPjrmnBBslcAtIyYLIUFaeA3lYnd1IrhSgTznmAmdqY6Eo16jwq4WO0b:"))
	request, err := http.NewRequest("POST", "https://api.xendit.co/v2/invoices", bytes.NewBuffer(reqBodyBytes.Bytes()))
	request.Header.Set("Content-Type", "application/json; charset=UTF-8")
	request.Header.Set("Authorization", "Basic "+basicauth)

	client := &http.Client{}
	response, err := client.Do(request)

	if err != nil {
		panic(err)
	}
	defer response.Body.Close()

	fmt.Println("response Status:", response.Status)
	fmt.Println("response Headers:", response.Header)
	body, _ := io.ReadAll(response.Body)
	fmt.Println("response Body:", string(body))
	res = body

	return res, err

}
