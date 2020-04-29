/*
* project mini-e-wallet
* created by oktoprima
* email : octoprima93@gmail.com
* github : https://github.com/oktopriima
* created at 02.50
**/

package transactions

import (
	"github.com/jinzhu/gorm"

	"github.com/oktopriima/mini-e-wallet/application/request"
	"github.com/oktopriima/mini-e-wallet/domain/core/repo"
)

type TransactionController interface {
	TopUpController(request request.TopUpRequest) (interface{}, error)
	TransferController(request request.TransferRequest) (interface{}, error)
}

type transactionController struct {
	db                     *gorm.DB
	userBalanceRepo        repo.UserBalanceRepository
	userBalanceHistoryRepo repo.UserBalanceHistoryRepository
	userRepo               repo.UserRepository
}

func NewTransactionController(db *gorm.DB,
	userBalanceRepo repo.UserBalanceRepository,
	userBalanceHistoryRepo repo.UserBalanceHistoryRepository,
	userRepo repo.UserRepository) TransactionController {
	return &transactionController{
		db:                     db,
		userBalanceRepo:        userBalanceRepo,
		userBalanceHistoryRepo: userBalanceHistoryRepo,
		userRepo:               userRepo,
	}
}
