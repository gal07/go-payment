package usecase

import (
	"go-payment/modules/transaction"
	"go-payment/modules/transaction/entity"
)

type TransactionUseCase struct {
	TransactionEntity entity.TransactionEntity
}

func NewUseCase(TransactionEntity entity.TransactionEntity) (transaction.ITransactionUseCase, error) {
	return &TransactionUseCase{
		TransactionEntity: TransactionEntity}, nil
}
