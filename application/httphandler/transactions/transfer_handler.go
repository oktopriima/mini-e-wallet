/*
* project mini-e-wallet
* created by oktoprima
* email : octoprima93@gmail.com
* github : https://github.com/oktopriima
* created at 14.33
**/

package transactions

import (
	"net/http"

	"github.com/gin-gonic/gin"

	"github.com/oktopriima/mini-e-wallet/application/request"
	"github.com/oktopriima/mini-e-wallet/domain/helper"
	"github.com/oktopriima/mini-e-wallet/domain/response"
)

func (t *transactionHandler) TransferHandler(ctx *gin.Context) {
	var req request.TransferRequest

	if err := ctx.ShouldBindJSON(&req); err != nil {
		response.NewErrorResponse(ctx, http.StatusBadRequest, err)
		return
	}

	authenticatedUser, err := helper.GetAuthenticatedUser(ctx.Request)
	if err != nil {
		response.NewErrorResponse(ctx, http.StatusForbidden, err)
		return
	}
	req.AuthenticatedUser = int(authenticatedUser)
	req.IP = helper.GetIP(ctx.Request)
	req.UserAgent = ctx.Request.Header.Get("User-Agent")

	resp, err := t.controller.TransferController(req)
	if err != nil {
		response.NewErrorResponse(ctx, http.StatusUnprocessableEntity, err)
		return
	}

	response.NewSuccessResponse(ctx, resp)
}
