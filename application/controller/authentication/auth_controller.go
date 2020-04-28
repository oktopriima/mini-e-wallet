/*
* project mini-e-wallet
* created by oktoprima
* email : octoprima93@gmail.com
* github : https://github.com/oktopriima
* created at 02.54
**/

package authentication

import (
	"github.com/jinzhu/gorm"

	"github.com/oktopriima/mini-e-wallet/application/request"
	"github.com/oktopriima/mini-e-wallet/domain/config"
	"github.com/oktopriima/mini-e-wallet/domain/core/repo"
)

type AuthController interface {
	LoginController(request request.LoginRequest) (interface{}, error)
}

type authController struct {
	db       *gorm.DB
	conf     config.Config
	userRepo repo.UserRepository
}

func NewAuthController(db *gorm.DB,
	conf config.Config,
	userRepo repo.UserRepository) AuthController {
	return &authController{
		db:       db,
		conf:     conf,
		userRepo: userRepo,
	}
}
