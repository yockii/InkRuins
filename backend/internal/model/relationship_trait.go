package model

// 关系维度定义
type RelationshipTrait struct {
	BaseModel
	ProjectID    uint64   `gorm:"not null;index" json:"project_id,string"`
	Name         string   `gorm:"not null;index" json:"name"`
	Description  string   `gorm:"not null;index" json:"description"`
	MinValue     *float64 `gorm:"default:0" json:"min_value"`
	MaxValue     *float64 `gorm:"default:100" json:"max_value"`
	DefaultValue *float64 `gorm:"default:50" json:"default_value"`
	Remark       string   `gorm:"size:2000" json:"remark"` // 备注
}

func init() {
	Models = append(Models, &RelationshipTrait{})
}
