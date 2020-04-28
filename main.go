/*
 * Name : Okto Prima Jaya
 * Github : https://github.com/oktopriima
 * Email : octoprima93@gmail.com
 * Created At : 18/11/2019, 21:14
 * Copyright (c) 2019
 */

package main

import (
	"os"

	_ "github.com/go-sql-driver/mysql"

	"github.com/oktopriima/mini-e-wallet/domain/config"
	"github.com/oktopriima/mini-e-wallet/domain/middleware"
	"github.com/oktopriima/mini-e-wallet/routes"
)

func init() {
	// set time zone
	_ = os.Setenv("TZ", "Asia/Jakarta")
}

func main() {
	var err error

	c := BuildContainer()

	cfg := config.NewConfig()
	if err, _ = middleware.NewMiddlewareConfig(cfg); err != nil {
		panic(err)
	}

	if err = c.Invoke(routes.InvokeRoute); err != nil {
		panic(err)
	}

	if err = c.Provide(NewRoute); err != nil {
		panic(err)
	}

	if err = c.Invoke(func(s *ServerRoute) {
		s.Run()
	}); err != nil {
		panic(err)
	}
}
