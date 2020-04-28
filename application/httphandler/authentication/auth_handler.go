/*
* project mini-e-wallet
* created by oktoprima
* email : octoprima93@gmail.com
* github : https://github.com/oktopriima
* created at 03.11
**/

package authentication

import (
	"github.com/gin-gonic/gin"

	"github.com/oktopriima/mini-e-wallet/application/controller/authentication"
)

type AuthHandler interface {
	LoginHandler(ctx *gin.Context)
}

type authHandler struct {
	controller authentication.AuthController
}

func NewAuthHandler(controller authentication.AuthController) AuthHandler {
	return &authHandler{controller: controller}
}
