package model

type User struct {
	BaseDeleteModel

	Username string `gorm:"uniqueIndex;size:100;not null" json:"username"`
	Email    string `gorm:"index;size:255" json:"email"`
	Password string `gorm:"size:255;not null" json:"password"` // bcrypt哈希

	// OAuth信息 以后再实现，单独建表，一个用户可能绑定多个OAuth账号
	// OAuthProvider string `gorm:"index;size:50" json:"oauth_provider"` // linuxdo, github, etc.
	// OAuthID       string `gorm:"index;size:255" json:"oauth_id"`

	// 用户信息
	DisplayName string `gorm:"size:100" json:"display_name"`
	AvatarURL   string `gorm:"size:500" json:"avatar_url"`

	// 状态
	Status string `gorm:"index;default:'active'" json:"status"` // active, suspended

}

func init() {
	Models = append(Models, &User{})
}
