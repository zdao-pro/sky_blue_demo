package test

import "github.com/zdao-pro/sky_blue/pkg/net/http/gin"

func test(c *gin.Context) {
	c.JSON(200, gin.H{
		"message": "pong",
	})
}
