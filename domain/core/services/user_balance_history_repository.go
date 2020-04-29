/*
* project mini-e-wallet
* created by oktoprima
* email : octoprima93@gmail.com
* github : https://github.com/oktopriima
* created at 13.07
**/

package services

import (
	"encoding/json"

	"github.com/jinzhu/gorm"

	"github.com/oktopriima/mini-e-wallet/domain/core/model"
	"github.com/oktopriima/mini-e-wallet/domain/core/repo"
)

type userBalanceHistoryServices struct {
	db *gorm.DB
}

func (u *userBalanceHistoryServices) Create(history *model.UserBalanceHistory, tx *gorm.DB) (*model.UserBalanceHistory, error) {
	db := tx.Create(&history)
	m := new(model.UserBalanceHistory)

	if err := db.Error; err != nil {
		return nil, err
	}

	byteData, err := json.Marshal(db.Value)
	if err != nil {
		return nil, err
	}

	if err = json.Unmarshal(byteData, &m); err != nil {
		return nil, err
	}

	return m, nil
}

func NewUserBalanceHistoryServices(db *gorm.DB) repo.UserBalanceHistoryRepository {
	return &userBalanceHistoryServices{db: db}
}
