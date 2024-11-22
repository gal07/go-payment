package rest

import (
	"go-payment/modules/transaction"

	"github.com/gin-gonic/gin"
)

type endpoint struct {
	useCasetransaction transaction.ITransactionUseCase
}

func NewEndPoint(
	engine *gin.Engine,
	useCasetransaction transaction.ITransactionUseCase,
) error {

	// declare endpoint
	var edp = endpoint{
		useCasetransaction: useCasetransaction,
	}

	// Basic Auth
	const rootEndpoint = "/api/v1/transaction"
	r := engine.Group(rootEndpoint)
	r.POST("/", edp.Transaction)
	r.Use()

	return nil

}
