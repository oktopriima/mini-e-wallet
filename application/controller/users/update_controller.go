/*
* project mini-e-wallet
* created by oktoprima
* email : octoprima93@gmail.com
* github : https://github.com/oktopriima
* created at 01.59
**/

package users

import (
	"fmt"

	"github.com/jinzhu/copier"

	"github.com/oktopriima/mini-e-wallet/application/request"
)

func (u *userController) UpdateController(request request.UserRequest) (interface{}, error) {
	user, err := u.userRepo.FindOne(map[string]interface{}{
		"id": request.ID,
	})
	if err != nil {
		return nil, fmt.Errorf("users %s", err.Error())
	}

	if err := copier.Copy(&user, request); err != nil {
		return nil, err
	}

	tx := u.db.Begin()
	defer tx.Rollback()

	if err := u.userRepo.Update(user, tx); err != nil {
		return nil, err
	}

	tx.Commit()
	return user, nil
}
