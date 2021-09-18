package main

import (
	"github.com/faizalnurrozi/go-gin-repository/router"
	"github.com/gin-gonic/gin"
)

type GroupRouting struct {
	engine *gin.Engine
}

func (group GroupRouting) RegisterRoute() {
	pingRoute1 := router.RoutersPing1{
		Engine: group.engine,
	}
	pingRoute1.RegisterRoute()

	pingRoute2 := router.RoutersPing2{
		Engine: group.engine,
	}
	pingRoute2.RegisterRoute()
}
