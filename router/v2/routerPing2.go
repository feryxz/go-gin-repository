package v1

import "github.com/gin-gonic/gin"

type RoutersPing2 struct {
	routerGroup *gin.RouterGroup
}

func NewRoutersPing2(rg *gin.RouterGroup) *RoutersPing2 {
	return &RoutersPing2{
		routerGroup: rg,
	}
}

func (ping RoutersPing2) RegisterRoute() {
	ping.routerGroup.GET("/ping2", func(context *gin.Context) {
		context.JSON(200, gin.H{
			"message": "Ping second (v1) ...",
		})
	})
}