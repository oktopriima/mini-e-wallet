/*
* project mini-e-wallet
* created by oktoprima
* email : octoprima93@gmail.com
* github : https://github.com/oktopriima
* created at 13.06
**/

package repo

import (
	"github.com/jinzhu/gorm"

	"github.com/oktopriima/mini-e-wallet/domain/core/model"
)

type UserBalanceHistoryRepository interface {
	Create(history *model.UserBalanceHistory, tx *gorm.DB) (*model.UserBalanceHistory, error)
}
