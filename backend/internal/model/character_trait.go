package model

// 角色可变属性定义，每个项目可以定义很多变化的属性
type CharacterTrait struct {
	BaseModel
	ProjectID    uint64  `gorm:"not null;index" json:"project_id,string"`
	Name         string  `gorm:"size:200" json:"name"`         // 名称,如信任度、愤怒值、希望指数
	Description  string  `gorm:"size:2000" json:"description"` // 描述
	MinValue     float64 `gorm:"default:0" json:"min_value"`
	MaxValue     float64 `gorm:"default:100" json:"max_value"`
	DefaultValue float64 `gorm:"default:50" json:"default_value"`
}

func init() {
	Models = append(Models, &CharacterTrait{})
}
