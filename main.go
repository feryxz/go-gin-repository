package main

import (
	"fmt"
	"github.com/faizalnurrozi/go-gin-repository/router"
	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()

	pingRoute := router.RoutersPing{
		Engine: r,
	}
	pingRoute.RegisterRoute()

	err := r.Run()
	if err != nil {
		fmt.Println(err)
	}
}
