/*
* project mini-e-wallet
* created by oktoprima
* email : octoprima93@gmail.com
* github : https://github.com/oktopriima
* created at 21.45
**/

package users

import (
	"github.com/jinzhu/gorm"

	"github.com/oktopriima/mini-e-wallet/application/request"
	"github.com/oktopriima/mini-e-wallet/domain/core/repo"
)

type UserController interface {
	CreateController(request request.UserRequest) (interface{}, error)
	FindPagedController(request request.FindPagedRequest) (interface{}, error)
	FindController(request request.FindByIDRequest) (interface{}, error)
	UpdateController(request request.UserRequest) (interface{}, error)
}
type userController struct {
	db       *gorm.DB
	userRepo repo.UserRepository
}

func NewUserController(db *gorm.DB,
	userRepo repo.UserRepository) UserController {
	return &userController{db, userRepo}
}
