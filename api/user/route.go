package test

import "github.com/zdao-pro/sky_blue/pkg/net/http/gin"

//RegisterDemoBMServer ..
func RegisterDemoBMServer(e *gin.Engine) {
	e.GET("/register", registerController)
}
