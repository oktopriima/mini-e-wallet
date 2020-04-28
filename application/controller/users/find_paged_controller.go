/*
* project mini-e-wallet
* created by oktoprima
* email : octoprima93@gmail.com
* github : https://github.com/oktopriima
* created at 00.49
**/

package users

import (
	"github.com/oktopriima/mini-e-wallet/application/request"
	"github.com/oktopriima/mini-e-wallet/domain/response"
)

func (u *userController) FindPagedController(request request.FindPagedRequest) (interface{}, error) {
	users, err := u.userRepo.FindPaged(nil, request.Page, request.Size)
	if err != nil {
		return nil, err
	}

	resp := new(response.ResponsePaged)
	resp.Size = request.Size
	resp.Page = request.Page
	resp.Total = u.userRepo.Count(nil)
	resp.Data = users

	return resp, nil
}
