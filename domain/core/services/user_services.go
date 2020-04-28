/*
 * Name : Okto Prima Jaya
 * Github : https://github.com/oktopriima
 * Email : octoprima93@gmail.com
 * Created At : 30/11/2019, 19:53
 * Copyright (c) 2019
 */

package services

import (
	"encoding/json"

	"github.com/jinzhu/gorm"

	"github.com/oktopriima/mini-e-wallet/domain/core/model"
	"github.com/oktopriima/mini-e-wallet/domain/core/repo"
	"github.com/oktopriima/mini-e-wallet/domain/helper"
)

type userServices struct {
	db *gorm.DB
}

func NewUserServices(db *gorm.DB) repo.UserRepository {
	return &userServices{db}
}

func (srv *userServices) Create(user *model.User, tx *gorm.DB) (*model.User, error) {
	db := tx.Create(&user)
	m := new(model.User)

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

func (srv *userServices) FindPaged(criteria map[string]interface{}, page, size int) ([]*model.User, error) {
	var users []*model.User

	limit, offset := helper.GetLimitOffset(page, size)
	if err := srv.db.Where(criteria).Limit(limit).Offset(offset).Find(&users).Error; err != nil {
		return nil, err
	}

	return users, nil
}

func (srv *userServices) Count(criteria map[string]interface{}) int {
	var count int
	srv.db.Model(model.User{}).Where(criteria).Count(&count)
	return count
}

func (srv *userServices) FindOne(criteria map[string]interface{}) (*model.User, error) {
	user := new(model.User)
	if err := srv.db.Where(criteria).Find(&user).Error; err != nil {
		return nil, err
	}
	return user, nil
}

func (srv *userServices) Update(user *model.User, tx *gorm.DB) error {
	err := tx.Save(user).Error
	return err
}
