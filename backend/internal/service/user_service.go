package service

import "github.com/yockii/inkruins/internal/model"

type IUserService interface {
	CreateUser(user *model.User) error
	GetUserByID(id uint64) (*model.User, error)
	GetUserByUsername(username string) (*model.User, error)
	GetUserByEmail(email string) (*model.User, error)
	Login(username, password string) (string, *model.User, error)
	Logout(token string) error
	UpdateUser(user *model.User) error
	DeleteUser(id uint64) error
	GetUserList(condition *model.User, page, pageSize int) ([]*model.User, int64, error)
	GetUserCount(condition *model.User) (int64, error)
}
