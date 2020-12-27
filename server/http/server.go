package server

import (
	"sky_blue_demo/api/test"

	"github.com/zdao-pro/sky_blue/pkg/net/http/gin"
)

//Server web
var Server *gin.Engine

//NewServer return web
func NewServer() *gin.Engine {
	Server = gin.Default()
	test.RegisterDemoBMServer(Server)
	return Server
}
