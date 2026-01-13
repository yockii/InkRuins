package model

// ItemState 道具状态快照（EAV模式）
// 设计理念：记录道具在特定剧情节点的属性值
//
// 使用示例：
// - ItemState{ItemID: 宝剑, StoryEventID: 第3章, TraitID: 攻击力, TraitValue: 100}
// - ItemState{ItemID: 宝剑, StoryEventID: 第5章, TraitID: 攻击力, TraitValue: 150}
// 可以看到宝剑在第3章到第5章之间获得了强化，攻击力从100提升到150
type ItemState struct {
	BaseModel
	ProjectID    uint64  `gorm:"not null;index" json:"project_id,string"`
	ItemID       uint64  `gorm:"not null;index" json:"item_id,string"`        // 道具ID
	StoryEventID uint64  `gorm:"not null;index" json:"story_event_id,string"` // 剧情节点ID
	TraitID      uint64  `gorm:"not null;index" json:"trait_id,string"`       // 属性ID
	TraitValue   float64 `gorm:"not null;default:0" json:"trait_value"`       // 属性值
	Note         string  `gorm:"size:500" json:"note"`                        // 备注，记录变化原因
}

func init() {
	Models = append(Models, &ItemState{})
}
