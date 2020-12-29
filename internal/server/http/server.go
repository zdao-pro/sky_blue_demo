package server

import (
	"net/http"
	test "sky_blue_demo/api/user"

	"github.com/zdao-pro/sky_blue/pkg/net/http/gin"
)

//Server web
var Server *gin.Engine

//NewServer return web
func NewServer() *gin.Engine {
	Server = gin.Default()
	Server.GET("/ping", func(c *gin.Context) {
		c.AbortWithStatus(http.StatusOK)
	})
	//注册路由到gin
	test.RegisterDemoBMServer(Server)
	return Server
}
