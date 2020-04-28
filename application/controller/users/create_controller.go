/*
* project mini-e-wallet
* created by oktoprima
* email : octoprima93@gmail.com
* github : https://github.com/oktopriima
* created at 21.49
**/

package users

import (
	"time"

	"github.com/jinzhu/copier"
	"golang.org/x/crypto/bcrypt"

	"github.com/oktopriima/mini-e-wallet/application/request"
	"github.com/oktopriima/mini-e-wallet/domain/core/model"
)

func (u userController) CreateController(request request.UserRequest) (interface{}, error) {
	user := new(model.User)

	// copy request to user model
	if err := copier.Copy(&user, &request); err != nil {
		return nil, err
	}

	// fill the model
	user.IsActive = true
	user.IsVerified = false
	user.LastLogin = time.Now()

	// give rand char of password
	rand := []byte(request.Password)
	pass, err := bcrypt.GenerateFromPassword(rand, bcrypt.DefaultCost)
	if err != nil {
		return nil, err
	}
	user.Password = string(pass)

	// start transaction
	tx := u.db.Begin()

	// close if any error
	defer tx.Rollback()

	m, err := u.userRepo.Create(user, tx)
	if err != nil {
		return nil, err
	}

	// commit if there's no error anymore
	tx.Commit()

	return m, nil
}
