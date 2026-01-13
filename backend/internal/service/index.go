package service

import "github.com/yockii/inkruins/internal/service/impl"

var (
	UserService IUserService
)

func init() {
	UserService = impl.NewUserService()
}
