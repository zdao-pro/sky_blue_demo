package test

import (
	userSercice "sky_blue_demo/service/user_service"

	"github.com/zdao-pro/sky_blue/pkg/ecode"
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
	//解析请求参数到registerParm
	err := c.ShouldBindQuery(&registerParm)
	if nil != err {
		//Exit函数,框架使用ecode包作为错误码传递,第二个参数可以为error or struct,可不穿
		c.Exit(int(ecode.ParamInvaidErr), err)
	}
	//service实例
	userServiceObj := userSercice.NewUserService()
	//执行注册逻辑
	e := userServiceObj.RegisterHandle(registerParm.Name, registerParm.Age)
	if nil != e {
		c.Exit(int(ecode.ParamInvaidErr), e)
	}
	c.Exit(int(ecode.OK))
}
