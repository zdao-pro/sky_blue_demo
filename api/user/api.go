package test

import (
	userSercice "sky_blue_demo/service/user_service"

	"github.com/zdao-pro/sky_blue/pkg/net/http/gin"
)

//RegisterInfo ..
type RegisterInfo struct {
	Name string `form:"name" validate:"required"`
	Age  int    `form:"age" validate:"required min=1"` //required:表示必传 min表示最小值
}

//注册Controller
func registerController(c *gin.Context) {
	var registerParm RegisterInfo
	err := c.ShouldBindQuery(&registerParm)
	if nil != err {
		c.AbortWithError(406, err)
	}
	//service
	userServiceObj := userSercice.NewUserService()
	e := userServiceObj.RegisterHandle(registerParm.Name, registerParm.Age)
	if nil != e {
		c.AbortWithError(406, err)
	}
}
