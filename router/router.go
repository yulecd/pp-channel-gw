package router

import (
	"npp-channel-gw/app/api"

	"github.com/gin-gonic/gin"
	"github.com/yulecd/pp-common/middleware"
)

var route gin.IRouter

func InitRouter() *gin.Engine {
	r := gin.New()

	r.Use(gin.Recovery(), middleware.LogReqResp)

	r.GET("/", api.ApiGroupApp.BaseApi.Greeter)

	route = r.Group("/v1/")

	registerRoute()

	return r
}

func registerRoute() {
	route.GET("/hello", api.ApiGroupApp.BaseApi.Greeter)
}
