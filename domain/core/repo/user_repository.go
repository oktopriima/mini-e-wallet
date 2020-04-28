/*
 * Name : Okto Prima Jaya
 * Github : https://github.com/oktopriima
 * Email : octoprima93@gmail.com
 * Created At : 30/11/2019, 19:52
 * Copyright (c) 2019
 */

package repo

import (
	"github.com/jinzhu/gorm"

	"github.com/oktopriima/mini-e-wallet/domain/core/model"
)

type UserRepository interface {
	Create(user *model.User, tx *gorm.DB) (*model.User, error)
	FindPaged(criteria map[string]interface{}, page, size int) ([]*model.User, error)
	Count(criteria map[string]interface{}) int
	FindOne(criteria map[string]interface{}) (*model.User, error)
	Update(user *model.User, tx *gorm.DB) error
}
