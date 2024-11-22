package rest

import (
	"go-payment/modules/transaction/payload"
	util "go-payment/utility"

	"github.com/gin-gonic/gin"
)

func (e endpoint) Transaction(c *gin.Context) {

	// Bind
	payload := payload.ReqCreateTransaction{}
	if err := c.Bind(&payload); err != nil {
		util.ResponseError(c, 200, err, nil, "Error")
		return
	}

	// Call usecase
	res, err := e.useCasetransaction.Transaction(c, payload)
	if err != nil {
		util.ResponseError(c, 200, err, nil, "Failed")
		return
	}

	util.ResponseOK(c, 200, res)

}
