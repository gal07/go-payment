package entity

import (
	"database/sql"
	"go-payment/modules/transaction"
	"go-payment/modules/transaction/entity/database"

	_ "github.com/go-sql-driver/mysql"
)

type TransactionEntity struct {
	TransactionRepo transaction.ITransactionRepo
}

func NewEntity(Db *sql.DB) (TransactionEntity, error) {
	return TransactionEntity{
		TransactionRepo: database.NewTransactionRepo(Db)}, nil
}
