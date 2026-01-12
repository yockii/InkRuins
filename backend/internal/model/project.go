package model

// Project 项目表
// 设计理念：存储小说项目的基础信息和世界观设定
type Project struct {
	BaseDeleteModel
	UserID      uint64 `gorm:"not null;index" json:"user_id,string"` // 所属用户ID
	Title       string `gorm:"size:200;not null" json:"title"`       // 项目标题
	Description string `gorm:"type:text" json:"description"`         // 项目简介
	Genre       string `gorm:"size:50" json:"genre"`                 // 小说类型（玄幻/都市/科幻等）

	// 世界观设定
	WorldTimePeriod string `gorm:"type:text" json:"world_time_period"` // 时间背景
	WorldLocation   string `gorm:"type:text" json:"world_location"`    // 地理位置
	WorldAtmosphere string `gorm:"type:text" json:"world_atmosphere"`  // 氛围基调
	WorldRules      string `gorm:"type:text" json:"world_rules"`       // 世界规则

	// 项目统计
	TargetWords  int    `gorm:"default:0" json:"target_words"`          // 目标字数
	CurrentWords int    `gorm:"default:0" json:"current_words"`         // 当前字数
	Status       string `gorm:"index;default:'planning'" json:"status"` // 状态：planning/writing/completed/paused

	// 叙事视角
	NarrativePerspective string `gorm:"size:50" json:"narrative_perspective"` // first_person/third_person/omniscient

	// 写作风格（可选）
	WritingStyleID *uint64 `gorm:"index" json:"writing_style_id,string"` // 关联的写作风格ID
}

func init() {
	Models = append(Models, &Project{})
}
