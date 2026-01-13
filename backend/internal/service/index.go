package service

import "github.com/yockii/inkruins/internal/service/impl"

var (
	UserService    IUserService
	ProjectService IProjectService
)

func init() {
	UserService = impl.NewUserService()
	ProjectService = impl.NewProjectService()
}
