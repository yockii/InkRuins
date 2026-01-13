package domain

import "github.com/yockii/inkruins/internal/model"

type RegisterUserReq struct {
	Username string `json:"username"`
	Password string `json:"password"`
	Email    string `json:"email"`
}

type LoginReq struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

type LoginResp struct {
	Token string      `json:"token"`
	User  *model.User `json:"user"`
}
