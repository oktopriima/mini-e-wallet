/*
* project mini-e-wallet
* created by oktoprima
* email : octoprima93@gmail.com
* github : https://github.com/oktopriima
* created at 02.42
**/

package repo

import (
	"github.com/jinzhu/gorm"

	"github.com/oktopriima/mini-e-wallet/domain/core/model"
)

type UserBalanceRepository interface {
	Create(balance *model.UserBalance, tx *gorm.DB) (*model.UserBalance, error)
	FindLastBalance(criteria map[string]interface{}) (*model.UserBalance, error)
}
