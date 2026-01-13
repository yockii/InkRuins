package model

type AIProvider struct {
	BaseModel
	Name    string `gorm:"not null" json:"name"`
	APIKey  string `gorm:"not null" json:"api_key"`
	BaseUrl string `gorm:"" json:"base_url"`
}

func init() {
	Models = append(Models, &AIProvider{})
}
