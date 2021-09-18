package router

import "github.com/gin-gonic/gin"

type RoutersPing1 struct {
	routerGroup *gin.RouterGroup
}

func NewRoutersPing1(rg *gin.RouterGroup) *RoutersPing1 {
	return &RoutersPing1{
		routerGroup: rg,
	}
}

func (ping RoutersPing1) RegisterRoute() {
	ping.routerGroup.GET("/ping", func(context *gin.Context) {
		context.JSON(200, gin.H{
			"message": "Ping first ...",
		})
	})
}