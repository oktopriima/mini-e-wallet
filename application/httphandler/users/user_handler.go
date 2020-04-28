/*
* project mini-e-wallet
* created by oktoprima
* email : octoprima93@gmail.com
* github : https://github.com/oktopriima
* created at 22.03
**/

package users

import (
	"github.com/gin-gonic/gin"

	"github.com/oktopriima/mini-e-wallet/application/controller/users"
)

type UserHandler interface {
	CreateHandler(ctx *gin.Context)
	FindPagedHandler(ctx *gin.Context)
	FindHandler(ctx *gin.Context)
	UpdateHandler(ctx *gin.Context)
}

type userHandler struct {
	user users.UserController
}

func NewUserHandler(user users.UserController) UserHandler {
	return &userHandler{user}
}
