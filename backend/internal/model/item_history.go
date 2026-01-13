package model

// ItemHistory 道具流转历史
// 设计理念：记录道具的创建、获得、失去、转移等历史事件
type ItemHistory struct {
	BaseModel
	ProjectID    uint64 `gorm:"not null;index" json:"project_id,string"`
	ItemID       uint64 `gorm:"not null;index" json:"item_id,string"`        // 道具ID
	StoryEventID uint64 `gorm:"not null;index" json:"story_event_id,string"` // 剧情节点ID（时间锚点）

	// 动作类型
	Action string `gorm:"not null;index" json:"action"`
	// created(创建)/obtained(获得)/lost(失去)/transferred(转移)/destroyed(销毁)/upgraded(升级)/enchanted(附魔)

	// 所有者变更
	FromOwnerType string  `gorm:"size:50" json:"from_owner_type"` // 原所有者类型
	FromOwnerID   *uint64 `json:"from_owner_id,string"`           // 原所有者ID
	ToOwnerType   string  `gorm:"size:50" json:"to_owner_type"`   // 新所有者类型
	ToOwnerID     *uint64 `json:"to_owner_id,string"`             // 新所有者ID

	// 数量（针对可堆叠的道具）
	Quantity int `gorm:"default:1" json:"quantity"`

	// 备注
	Note string `gorm:"type:text" json:"note"` // 事件描述
}

func init() {
	Models = append(Models, &ItemHistory{})
}
