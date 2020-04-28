/*
* project mini-e-wallet
* created by oktoprima
* email : octoprima93@gmail.com
* github : https://github.com/oktopriima
* created at 01.45
**/

package users

import (
	"github.com/oktopriima/mini-e-wallet/application/request"
)

func (u *userController) FindController(request request.FindByIDRequest) (interface{}, error) {
	user, err := u.userRepo.FindOne(map[string]interface{}{
		"id": request.ID,
	})
	if err != nil {
		return nil, err
	}
	return user, nil

}
