package userservice

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
	return
}
