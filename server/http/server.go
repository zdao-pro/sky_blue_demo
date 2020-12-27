package server

import (
	"github.com/zdao-pro/sky_blue/pkg/net/http/gin"
)

//Server web
var Server *gin.Engine

//NewServer return web
func NewServer() *gin.Engine {
	Server = gin.Default()
	return Server
}
