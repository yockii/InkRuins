package model

// ItemTrait 道具属性定义（EAV模式）
// 设计理念：每个项目可以自定义道具属性维度
//
// 属性示例：
// - 攻击力（MinValue: 0, MaxValue: 10000）
// - 防御力（MinValue: 0, MaxValue: 10000）
// - 重量（MinValue: 0, MaxValue: 1000）
// - 耐久度（MinValue: 0, MaxValue: 100）
// - 魔法值（MinValue: 0, MaxValue: 10000）
// - 特殊效果（文本描述）
// - 价值（MinValue: 0, MaxValue: 10000000）
type ItemTrait struct {
	BaseModel
	ProjectID    uint64   `gorm:"not null;index" json:"project_id,string"`
	Name         string   `gorm:"size:200;not null" json:"name"`   // 属性名称
	Description  string   `gorm:"size:2000" json:"description"`    // 属性描述
	MinValue     *float64 `gorm:"default:0" json:"min_value"`      // 最小值
	MaxValue     *float64 `gorm:"default:100" json:"max_value"`    // 最大值
	DefaultValue *float64 `gorm:"default:50" json:"default_value"` // 默认值
	Remark       string   `gorm:"size:2000" json:"remark"`         // 备注
}

func init() {
	Models = append(Models, &ItemTrait{})
}
