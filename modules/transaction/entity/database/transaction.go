package database

import (
	"context"
	"database/sql"
	"go-payment/modules/transaction"
	"go-payment/modules/transaction/models"

	_ "github.com/go-sql-driver/mysql"
)

type accessTransaction struct {
	Db *sql.DB
}

func NewTransactionRepo(Db *sql.DB) transaction.ITransactionRepo {
	return accessTransaction{Db: Db}
}

func (s accessTransaction) Create(ctx context.Context, req models.Transaction) (res models.Transaction, err error) {

	// Execute
	// var datetime = time.Now()
	// dt := datetime.Format(time.RFC3339)
	_, err = s.Db.Exec("insert into mt_transaction (xendit_id,xendit_external_id,user_id,currency,status,merchant_name,amount,description,expiry_date,invoice_url,created_at,updated_at) values (?,?)",
		&req.Xendit_id, &req.Xendit_external_id, &req.User_id, &req.Currency, &req.Status, &req.Merchant_name, &req.Amount, &req.Description, &req.Expiry_date, &req.Invoice_url, &req.Created_at, &req.Updated_at)
	if err != nil {
		return res, err
	}
	return res, err

}

func (s accessTransaction) Get(ctx context.Context, req models.Transaction) (res models.Transaction, err error) {
	return res, err
}
