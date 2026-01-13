package model

// Item 道具物品表
// 设计理念：管理小说中的各种道具物品
//
// 道具类型示例：
// - 武器装备（剑、刀、枪、盔甲等）
// - 消耗品（丹药、食物、魔法卷轴等）
// - 宝物神器（魔法物品、神器、圣物等）
// - 剧情道具（信物、钥匙、地图等）
// - 资源材料（矿石、草药、金币等）
type Item struct {
	BaseDeleteModel
	ProjectID uint64 `gorm:"not null;index" json:"project_id,string"`

	// 基本信息
	Name        string `gorm:"size:200;not null" json:"name"` // 道具名称
	Type        string `gorm:"size:50;index" json:"type"`     // 道具类型
	SubType     string `gorm:"size:50" json:"sub_type"`       // 子类型
	Rarity      int    `gorm:"default:0" json:"rarity"`       // 稀有度：数值越高越稀有
	Description string `gorm:"type:text" json:"description"`  // 道具描述

	// 当前所有者
	OwnerType string  `gorm:"index;size:50" json:"owner_type"` // "character" 或 "organization" 或 "location" 或 空字符串
	OwnerID   *uint64 `gorm:"index" json:"owner_id,string"`    // 所有者ID

	// 外观
	Appearance string `gorm:"type:text" json:"appearance"` // 外观描述
	IconURL    string `gorm:"size:500" json:"icon_url"`    // 图标URL

	// 备注
	Notes string `gorm:"type:text" json:"notes"`
}

func init() {
	Models = append(Models, &Item{})
}
