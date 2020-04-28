/*
 * Name : Okto Prima Jaya
 * Github : https://github.com/oktopriima
 * Email : octoprima93@gmail.com
 * Created At : 18/11/2019, 20:51
 * Copyright (c) 2019
 */

package middleware

import (
	"net/http"

	"github.com/gin-gonic/gin"

	"github.com/oktopriima/mini-e-wallet/domain/config"
)

func CORSMiddleware() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		conf := config.NewConfig()

		ctx.Writer.Header().Set("Access-Control-Allow-Origin", conf.GetString("cors.allowed_origins"))
		ctx.Writer.Header().Set("Access-Control-Allow-Headers", conf.GetString("cors.allowed_headers"))
		ctx.Writer.Header().Set("Access-Control-Allow-Methods", conf.GetString("cors.allowed_methods"))

		if ctx.Request.Method == "OPTIONS" {
			ctx.AbortWithStatus(http.StatusOK)
		}

		ctx.Next()
	}
}
