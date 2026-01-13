package model

// LocationState 地点状态快照（EAV模式）
// 设计理念：记录地点在特定剧情节点的属性值
//
// 使用示例：
// - LocationState{LocationID: 京城, StoryEventID: 第5章, TraitID: 人口, TraitValue: 100000}
// - LocationState{LocationID: 京城, StoryEventID: 第10章, TraitID: 人口, TraitValue: 80000}
// 可以看到京城在第5章到第10章之间发生了战争，人口从10万减少到8万
type LocationState struct {
	BaseModel
	ProjectID    uint64  `gorm:"not null;index" json:"project_id,string"`
	LocationID   uint64  `gorm:"not null;index" json:"location_id,string"`    // 地点ID
	StoryEventID uint64  `gorm:"not null;index" json:"story_event_id,string"` // 剧情节点ID
	TraitID      uint64  `gorm:"not null;index" json:"trait_id,string"`       // 属性ID
	TraitValue   float64 `gorm:"not null;default:0" json:"trait_value"`       // 属性值
	Note         string  `gorm:"size:500" json:"note"`                        // 备注，记录变化原因
}

func init() {
	Models = append(Models, &LocationState{})
}
