package transaction

import (
	"context"
	"go-payment/modules/transaction/models"
	"go-payment/modules/transaction/payload"
)

type ITransactionUseCase interface {
	Transaction(ctx context.Context, req payload.ReqCreateTransaction) (res interface{}, err error)
}

type ITransactionRepo interface {
	Create(ctx context.Context, req models.Transaction) (res models.Transaction, err error)
	Get(ctx context.Context, req models.Transaction) (res models.Transaction, err error)
}
