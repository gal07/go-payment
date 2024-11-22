package usecase

import (
	"bytes"
	"context"
	"fmt"
	"go-payment/modules/transaction/payload"
	"io"
	"net/http"
)

func (s *TransactionUseCase) Transaction(ctx context.Context, req payload.ReqCreateTransaction) (res interface{}, err error) {

	var jsonData = []byte(`{
		"name": "morpheus",
		"job": "leader"
	}`)

	// Execute
	request, err := http.NewRequest("POST", "https://api.xendit.co/v2/invoices", bytes.NewBuffer(jsonData))
	request.Header.Set("Content-Type", "application/json; charset=UTF-8")

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

	return res, err

}
