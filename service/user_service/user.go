package userservice

import (
	"sky_blue_demo/internal/model"
	"time"
)

//UserService ..
type UserService struct {
}

//NewUserService ..
func NewUserService() (u UserService) {
	u = UserService{}
	return
}

//RegisterHandle ..
func (us UserService) RegisterHandle(name string, age int) (err error) {
	u := model.NewUserInfo()
	err = u.Insert(name, age, int(time.Now().UnixNano()/1000))
	return
}
