/*
* project mini-e-wallet
* created by oktoprima
* email : octoprima93@gmail.com
* github : https://github.com/oktopriima
* created at 01.50
**/

package users

import (
	"net/http"

	"github.com/gin-gonic/gin"

	"github.com/oktopriima/mini-e-wallet/application/request"
	"github.com/oktopriima/mini-e-wallet/domain/response"
)

func (u *userHandler) FindHandler(ctx *gin.Context) {
	var req request.FindByIDRequest

	if err := ctx.ShouldBindUri(&req); err != nil {
		response.NewErrorResponse(ctx, http.StatusBadRequest, err)
		return
	}

	resp, err := u.user.FindController(req)
	if err != nil {
		response.NewErrorResponse(ctx, http.StatusNotAcceptable, err)
		return
	}

	response.NewSuccessResponse(ctx, resp)
}
