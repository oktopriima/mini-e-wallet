/*
* project mini-e-wallet
* created by oktoprima
* email : octoprima93@gmail.com
* github : https://github.com/oktopriima
* created at 09.45
**/

package transactions

import (
	"github.com/gin-gonic/gin"

	"github.com/oktopriima/mini-e-wallet/application/controller/transactions"
)

type TransactionHandler interface {
	TopupHandler(ctx *gin.Context)
}

type transactionHandler struct {
	controller transactions.TransactionController
}

func NewTransactionHandler(controller transactions.TransactionController) TransactionHandler {
	return &transactionHandler{controller: controller}
}
