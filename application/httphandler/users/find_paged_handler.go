/*
* project mini-e-wallet
* created by oktoprima
* email : octoprima93@gmail.com
* github : https://github.com/oktopriima
* created at 01.05
**/

package users

import (
	"net/http"

	"github.com/gin-gonic/gin"

	"github.com/oktopriima/mini-e-wallet/application/request"
	"github.com/oktopriima/mini-e-wallet/domain/response"
)

func (u *userHandler) FindPagedHandler(ctx *gin.Context) {
	var req request.FindPagedRequest

	// binding query parameter into request structure
	if err := ctx.ShouldBindQuery(&req); err != nil {
		response.NewErrorResponse(ctx, http.StatusBadRequest, err)
		return
	}

	resp, err := u.user.FindPagedController(req)

	// check possibility error from query
	if err != nil {
		response.NewErrorResponse(ctx, http.StatusNotAcceptable, err)
		return
	}

	// return success
	response.NewSuccessResponse(ctx, resp)
}
