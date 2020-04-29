/*
* project mini-e-wallet
* created by oktoprima
* email : octoprima93@gmail.com
* github : https://github.com/oktopriima
* created at 09.43
**/

package transactions

import (
	"fmt"

	"github.com/leekchan/accounting"

	"github.com/oktopriima/mini-e-wallet/application/request"
	"github.com/oktopriima/mini-e-wallet/domain/core/model"
)

func (t *transactionController) TopUpController(request request.TopUpRequest) (interface{}, error) {
	var err error
	ac := accounting.Accounting{
		Symbol:    "Rp",
		Precision: 2,
	}

	user, err := t.userRepo.FindOne(map[string]interface{}{
		"id": request.AuthenticatedUser,
	})
	if err != nil {
		return nil, err
	}

	currentBalance := t.getCurrentBalance(request.AuthenticatedUser)
	newBalance := currentBalance + request.Amount

	ub := new(model.UserBalance)
	ub.UserID = request.AuthenticatedUser
	ub.Balance = newBalance
	ub.BalanceAchieve = request.Amount

	tx := t.db.Begin()
	defer tx.Rollback()

	ub, err = t.userBalanceRepo.Create(ub, tx)
	if err != nil {
		return nil, err
	}

	h := new(model.UserBalanceHistory)
	h.UserBalanceID = ub.ID
	h.BalanceBefore = currentBalance
	h.BalanceAfter = newBalance
	h.Activity = fmt.Sprintf("top up balance of %s by %s", ac.FormatMoney(request.Amount), user.Username)
	h.Type = "credit"
	h.Location = request.Location
	h.IP = request.IP
	h.UserAgent = request.UserAgent
	h.Author = user.Username

	if _, err = t.userBalanceHistoryRepo.Create(h, tx); err != nil {
		return nil, err
	}

	tx.Commit()

	return struct {
		Message string `json:"message"`
	}{
		Message: "success top up balance",
	}, nil
}

func (t *transactionController) getCurrentBalance(userID int) float64 {
	balance, err := t.userBalanceRepo.FindLastBalance(map[string]interface{}{
		"user_id": userID,
	})
	if err != nil {
		return 0
	}

	return balance.Balance
}
