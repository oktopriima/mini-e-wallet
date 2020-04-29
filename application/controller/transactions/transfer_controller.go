/*
* project mini-e-wallet
* created by oktoprima
* email : octoprima93@gmail.com
* github : https://github.com/oktopriima
* created at 14.16
**/

package transactions

import (
	"fmt"

	"github.com/leekchan/accounting"

	"github.com/oktopriima/mini-e-wallet/application/request"
	"github.com/oktopriima/mini-e-wallet/domain/core/model"
)

func (t *transactionController) TransferController(request request.TransferRequest) (interface{}, error) {
	// debit transaction from source user
	ac := accounting.Accounting{
		Symbol:    "Rp",
		Precision: 2,
	}

	// get authenticated user
	authenticatedUser, err := t.userRepo.FindOne(map[string]interface{}{
		"id": request.AuthenticatedUser,
	})
	if err != nil {
		return nil, err
	}

	// get destination user
	destinationUser, err := t.userRepo.FindOne(map[string]interface{}{
		"id": request.UserDestination,
	})
	if err != nil {
		return nil, err
	}

	// get authenticated user balance
	currentBalance := t.getCurrentBalance(authenticatedUser.ID)
	newBalance := currentBalance - request.Amount

	if currentBalance < request.Amount {
		return nil, fmt.Errorf("your balance is not enough")
	}

	// create debit transaction for authenticated user
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

	// create history balance for debit
	h := new(model.UserBalanceHistory)
	h.UserBalanceID = ub.ID
	h.BalanceBefore = currentBalance
	h.BalanceAfter = newBalance
	h.Activity = fmt.Sprintf("transfer balance of %s to %s by %s",
		ac.FormatMoney(request.Amount),
		destinationUser.Username,
		authenticatedUser.Username)
	h.Type = "debit"
	h.Location = request.Location
	h.IP = request.IP
	h.UserAgent = request.UserAgent
	h.Author = authenticatedUser.Username

	if _, err = t.userBalanceHistoryRepo.Create(h, tx); err != nil {
		return nil, err
	}

	tx.Commit()

	if err = t.destinationTransferBalance(destinationUser, authenticatedUser, request); err != nil {
		return nil, err
	}

	return struct {
		Message string `json:"message"`
	}{
		Message: "success transfer balance",
	}, nil
}

func (t *transactionController) destinationTransferBalance(
	user *model.User,
	authenticatedUser *model.User,
	request request.TransferRequest) error {
	var err error

	ac := accounting.Accounting{
		Symbol:    "Rp",
		Precision: 2,
	}

	currentBalance := t.getCurrentBalance(user.ID)
	newBalance := currentBalance + request.Amount

	ub := new(model.UserBalance)
	ub.UserID = user.ID
	ub.Balance = newBalance
	ub.BalanceAchieve = request.Amount

	tx := t.db.Begin()
	defer tx.Rollback()

	ub, err = t.userBalanceRepo.Create(ub, tx)
	if err != nil {
		return err
	}

	h := new(model.UserBalanceHistory)
	h.UserBalanceID = ub.ID
	h.BalanceBefore = currentBalance
	h.BalanceAfter = newBalance
	h.Activity = fmt.Sprintf("transfer balance of %s from %s",
		ac.FormatMoney(request.Amount),
		authenticatedUser.Username)
	h.Type = "credit"
	h.Location = request.Location
	h.IP = request.IP
	h.UserAgent = request.UserAgent
	h.Author = user.Username

	if _, err = t.userBalanceHistoryRepo.Create(h, tx); err != nil {
		return err
	}

	tx.Commit()

	return nil
}
