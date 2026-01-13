package model

// Location 地点表
// 设计理念：管理小说中的地理位置和场所
//
// 地点类型示例：
// - 行政区域（国家、省、城市、村庄）
// - 自然地理（山脉、河流、森林、沙漠、海洋）
// - 建筑/场所（宫殿、客栈、酒馆、店铺、寺庙）
// - 虚拟地点（秘境、洞穴、传送门、异世界）
type Location struct {
	BaseDeleteModel
	ProjectID uint64 `gorm:"not null;index" json:"project_id,string"`

	// 基本信息
	Name        string `gorm:"size:200;not null" json:"name"` // 地点名称
	Type        string `gorm:"size:50;index" json:"type"`     // 地点类型
	SubType     string `gorm:"size:50" json:"sub_type"`       // 子类型
	Description string `gorm:"type:text" json:"description"`  // 地点描述

	// 层级关系（支持树形结构）
	ParentID *uint64 `gorm:"index" json:"parent_id,string"` // 父地点ID
	Path     string  `gorm:"type:text;index" json:"path"`   // 路径：/uuid1/uuid2/uuid3

	// 地理位置（可选，用于地图）
	X *float64 `gorm:"default:0" json:"x"` // 经度或X坐标
	Y *float64 `gorm:"default:0" json:"y"` // 纬度或Y坐标
	Z *float64 `gorm:"default:0" json:"z"` // 高度或Z坐标

	// 外观描述
	Appearance string `gorm:"type:text" json:"appearance"` // 外观描述
	IconURL    string `gorm:"size:500" json:"icon_url"`    // 地点图标

	// 备注
	Notes string `gorm:"type:text" json:"notes"`
}

func init() {
	Models = append(Models, &Location{})
}
