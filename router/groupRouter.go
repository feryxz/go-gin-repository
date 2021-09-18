package router

import (
	routingPingV1 "github.com/faizalnurrozi/go-gin-repository/router/v1"
	routingPingV2 "github.com/faizalnurrozi/go-gin-repository/router/v2"
	"github.com/gin-gonic/gin"
)

type GroupRouting struct {
	Engine *gin.Engine
}

func NewGroupRouting(routing *gin.Engine) *GroupRouting {
	return &GroupRouting{
		Engine: routing,
	}
}

func (group GroupRouting) RegisterRoute() {

	v1 := group.Engine.Group("/v1")
	{
		ping1 := routingPingV1.NewRoutersPing1(v1)
		ping1.RegisterRoute()

		ping2 := routingPingV1.NewRoutersPing2(v1)
		ping2.RegisterRoute()
	}

	v2 := group.Engine.Group("/v2")
	{
		ping1 := routingPingV2.NewRoutersPing1(v2)
		ping1.RegisterRoute()

		ping2 := routingPingV2.NewRoutersPing2(v2)
		ping2.RegisterRoute()
	}
}
