/*
 * Name : Okto Prima Jaya
 * Github : https://github.com/oktopriima
 * Email : octoprima93@gmail.com
 * Created At : 18/11/2019, 20:52
 * Copyright (c) 2019
 */

package response

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type ResponseError struct {
	ErrorCode int    `json:"error_code"`
	Message   string `json:"message"`
	Status    int    `json:"status"`
}

type ResponsePaged struct {
	Data  interface{} `json:"data"`
	Page  int         `json:"page"`
	Size  int         `json:"size"`
	Total int         `json:"total"`
}

type ResponseObject struct {
	Data interface{} `json:"data"`
}

func NewErrorResponse(c *gin.Context, code int, err error) {
	response := new(ResponseError)
	response.Status = code
	response.Message = err.Error()
	response.ErrorCode = code
	c.JSON(code, &response)
	return
}

func NewSuccessResponse(c *gin.Context, data interface{}) {
	c.JSON(http.StatusOK, data)
	return
}

func NewCreatedResponse(ctx *gin.Context, data interface{}) {
	ctx.JSON(http.StatusCreated, data)
	return
}
