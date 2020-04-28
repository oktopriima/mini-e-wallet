/*
* project mini-e-wallet
* created by oktoprima
* email : octoprima93@gmail.com
* github : https://github.com/oktopriima
* created at 03.13
**/

package authentication

import (
	"net/http"

	"github.com/gin-gonic/gin"

	"github.com/oktopriima/mini-e-wallet/application/request"
	"github.com/oktopriima/mini-e-wallet/domain/response"
)

func (a *authHandler) LoginHandler(ctx *gin.Context) {
	var req request.LoginRequest

	if err := ctx.ShouldBindJSON(&req); err != nil {
		response.NewErrorResponse(ctx, http.StatusBadRequest, err)
		return
	}

	resp, err := a.controller.LoginController(req)
	if err != nil {
		response.NewErrorResponse(ctx, http.StatusNotAcceptable, err)
		return
	}

	response.NewSuccessResponse(ctx, resp)
}
