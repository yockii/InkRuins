package data

import (
	"github.com/yockii/inkruins/internal/dao"
	"github.com/yockii/inkruins/internal/database"
	"github.com/yockii/inkruins/internal/model"
	"golang.org/x/crypto/bcrypt"
)

func initializeUser() {
	users := []*model.User{
		{
			Username: "admin",
			Password: "123456",
		},
	}
	for _, user := range users {
		b, _ := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.DefaultCost)
		user.Password = string(b)
		database.DB.Where(dao.User.Username.Eq(user.Username)).FirstOrCreate(user)
	}
}
