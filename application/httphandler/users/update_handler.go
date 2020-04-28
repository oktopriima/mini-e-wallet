/*
* project mini-e-wallet
* created by oktoprima
* email : octoprima93@gmail.com
* github : https://github.com/oktopriima
* created at 02.05
**/

package users

import (
	"net/http"

	"github.com/gin-gonic/gin"

	"github.com/oktopriima/mini-e-wallet/application/request"
	"github.com/oktopriima/mini-e-wallet/domain/response"
)

func (u *userHandler) UpdateHandler(ctx *gin.Context) {
	var req request.UserRequest
	var err error

	// read URL parameter
	if err = ctx.ShouldBindUri(&req); err != nil {
		response.NewErrorResponse(ctx, http.StatusBadRequest, err)
		return
	}

	// read body
	if err = ctx.ShouldBindJSON(&req); err != nil {
		response.NewErrorResponse(ctx, http.StatusBadRequest, err)
		return
	}

	resp, err := u.user.UpdateController(req)
	if err != nil {
		response.NewErrorResponse(ctx, http.StatusNotAcceptable, err)
		return
	}

	response.NewSuccessResponse(ctx, resp)

}
