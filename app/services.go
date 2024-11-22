package main

import (
	transactionRest "go-payment/modules/transaction/delivery/rest"
	transactionEntity "go-payment/modules/transaction/entity"
	transactionUsecase "go-payment/modules/transaction/usecase"

	"github.com/gin-gonic/gin"
)

func service(
	route *gin.Engine,
) {

	// call db instance
	sqlCon, err := MysqlConnect()
	if err != nil {
		panic(err)
	}

	// call entity
	transactionEntity, err := transactionEntity.NewEntity(sqlCon)
	if err != nil {
		panic(err)
	}

	// call usecase
	transactionUsecase, err := transactionUsecase.NewUseCase(transactionEntity)
	if err != nil {
		panic(err)
	}

	// call endpoint
	transactionRest.NewEndPoint(route, transactionUsecase)

}
