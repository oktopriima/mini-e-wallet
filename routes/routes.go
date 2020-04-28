/*
 * Name : Okto Prima Jaya
 * Github : https://github.com/oktopriima
 * Email : octoprima93@gmail.com
 * Created At : 18/11/2019, 21:33
 * Copyright (c) 2019
 */

package routes

import (
	"github.com/gin-gonic/gin"

	"github.com/oktopriima/mini-e-wallet/application/httphandler/users"
	"github.com/oktopriima/mini-e-wallet/domain/config"
	"github.com/oktopriima/mini-e-wallet/domain/middleware"
)

func InvokeRoute(
	engine *gin.Engine,
	user users.UserHandler,
) {
	conf := config.NewConfig()
	route := engine.Group("api/" + conf.GetString("app.version.tag") + conf.GetString("app.version.value"))

	route.Use(gin.Logger())
	route.Use(gin.Recovery())
	route.Use(gin.ErrorLogger())

	route.OPTIONS("/*path", middleware.CORSMiddleware())

	// user route
	{
		userRoute := route.Group("users")
		userRoute.Use(middleware.MyAuth())
		userRoute.POST("", user.CreateHandler)
		userRoute.GET("", user.FindPagedHandler)
		userRoute.GET(":id", user.FindHandler)
		userRoute.PUT(":id", user.UpdateHandler)
	}

}
