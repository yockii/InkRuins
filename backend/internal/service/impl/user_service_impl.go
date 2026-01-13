package impl

import (
	"errors"
	"time"

	gonanoid "github.com/matoous/go-nanoid/v2"
	"github.com/yockii/inkruins/internal/cache"
	"github.com/yockii/inkruins/internal/constant"
	"github.com/yockii/inkruins/internal/dao"
	"github.com/yockii/inkruins/internal/database"
	"github.com/yockii/inkruins/internal/model"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

type UserServiceImpl struct{}

func NewUserService() *UserServiceImpl {
	return &UserServiceImpl{}
}

func (s *UserServiceImpl) CreateUser(user *model.User) error {
	if user == nil {
		return errors.New("user is required")
	}
	if user.Username == "" {
		return errors.New("username is required")
	}
	if user.Password == "" {
		return errors.New("password is required")
	}
	pwd, _ := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.DefaultCost)
	user.Password = string(pwd)
	user.Status = "active"

	return database.DB.Create(user).Error
}

func (s *UserServiceImpl) GetUserByID(id uint64) (*model.User, error) {
	var user model.User
	if err := database.DB.Where(dao.BaseModel.ID.Eq(id)).First(&user).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, nil
		}
		return nil, err
	}
	return &user, nil
}

func (s *UserServiceImpl) GetUserByUsername(username string) (*model.User, error) {
	var user model.User
	if err := database.DB.Where(dao.User.Username.Eq(username)).First(&user).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, nil
		}
		return nil, err
	}
	return &user, nil
}

func (s *UserServiceImpl) GetUserByEmail(email string) (*model.User, error) {
	var user model.User
	if err := database.DB.Where(dao.User.Email.Eq(email)).First(&user).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, nil
		}
		return nil, err
	}
	return &user, nil
}

func (s *UserServiceImpl) Login(username, password string) (string, *model.User, error) {
	user, err := s.GetUserByUsername(username)
	if err != nil {
		return "", nil, err
	}
	if user == nil {
		return "", nil, errors.New("user not found")
	}
	if err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password)); err != nil {
		return "", nil, errors.New("invalid password")
	}

	token, err := gonanoid.New()
	if err != nil {
		return "", nil, err
	}
	cache.Set(constant.CacheKeyUserToken+token, user.ID, time.Hour*24)

	return token, user, nil
}

func (s *UserServiceImpl) Logout(token string) error {
	return cache.Del(constant.CacheKeyUserToken + token)
}

func (s *UserServiceImpl) UpdateUser(user *model.User) error {
	if user == nil {
		return errors.New("user is required")
	}
	if user.Password != "" {
		hashedPassword, err := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.DefaultCost)
		if err != nil {
			return err
		}
		user.Password = string(hashedPassword)
	}
	return database.DB.Omit(dao.User.Password.Column().Name).Save(user).Error
}

func (s *UserServiceImpl) DeleteUser(id uint64) error {
	if id == 0 {
		return errors.New("id is required")
	}
	return database.DB.Delete(&model.User{}, id).Error
}

func (s *UserServiceImpl) GetUserList(condition *model.User, page, pageSize int) ([]*model.User, int64, error) {
	var users []*model.User
	var total int64
	tx := database.DB.Model(&model.User{}).Order(dao.BaseModel.CreatedAt.Desc())
	if condition != nil {
		if condition.Username != "" {
			tx = tx.Where(dao.User.Username.Like("%" + condition.Username + "%"))
		}
		if condition.Email != "" {
			tx = tx.Where(dao.User.Email.Like("%" + condition.Email + "%"))
		}
		if condition.Status != "" {
			tx = tx.Where(dao.User.Status.Eq(condition.Status))
		}
	}
	if err := tx.Count(&total).Error; err != nil {
		return nil, 0, err
	}
	if total == 0 {
		return users, 0, nil
	}
	if page < 1 {
		page = 1
	}
	if pageSize < 1 {
		pageSize = 10
	}
	if err := tx.Offset((page - 1) * pageSize).Limit(pageSize).Find(&users).Error; err != nil {
		return nil, 0, err
	}
	return users, total, nil
}

func (s *UserServiceImpl) GetUserCount(condition *model.User) (int64, error) {
	var total int64

	tx := database.DB.Model(&model.User{})
	if condition != nil {
		if condition.Username != "" {
			tx = tx.Where(dao.User.Username.Like("%" + condition.Username + "%"))
		}
		if condition.Email != "" {
			tx = tx.Where(dao.User.Email.Like("%" + condition.Email + "%"))
		}
	}
	if err := tx.Count(&total).Error; err != nil {
		return 0, err
	}
	return total, nil
}
