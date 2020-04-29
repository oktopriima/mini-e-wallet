/*
* project mini-e-wallet
* created by oktoprima
* email : octoprima93@gmail.com
* github : https://github.com/oktopriima
* created at 02.43
**/

package services

import (
	"encoding/json"

	"github.com/jinzhu/gorm"

	"github.com/oktopriima/mini-e-wallet/domain/core/model"
	"github.com/oktopriima/mini-e-wallet/domain/core/repo"
)

type userBalanceServices struct {
	db *gorm.DB
}

func (u *userBalanceServices) FindLastBalance(criteria map[string]interface{}) (*model.UserBalance, error) {
	balance := new(model.UserBalance)

	if err := u.db.Where(criteria).Order("id DESC").Limit(1).Find(&balance).Error; err != nil {
		return nil, err
	}
	return balance, nil
}

func (u *userBalanceServices) Create(balance *model.UserBalance, tx *gorm.DB) (*model.UserBalance, error) {
	db := tx.Create(&balance)
	m := new(model.UserBalance)

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

func NewUserBalanceServices(db *gorm.DB) repo.UserBalanceRepository {
	return &userBalanceServices{db: db}
}
