/*
 * Name : Okto Prima Jaya
 * Github : https://github.com/oktopriima
 * Email : octoprima93@gmail.com
 * Created At : 18/11/2019, 21:36
 * Copyright (c) 2019
 */

package main

import (
	"log"

	"github.com/gin-gonic/gin"

	"github.com/oktopriima/mini-e-wallet/domain/config"
)

type ServerRoute struct {
	cfg    config.Config
	engine *gin.Engine
}

func NewRoute(cfg config.Config, engine *gin.Engine) *ServerRoute {
	return &ServerRoute{cfg, engine}
}

func (s *ServerRoute) Run() {
	if err := s.engine.Run(s.cfg.GetString("server.address")); err != nil {
		log.Fatal(err)
	}
}
