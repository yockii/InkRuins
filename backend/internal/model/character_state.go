package model

// CharacterState 角色状态快照
// 记录角色在特定剧情节点（StoryEvent）结束后的属性值，支持完整的历史追溯
type CharacterState struct {
	BaseModel
	ProjectID    uint64  `gorm:"not null;index" json:"project_id,string"`
	CharacterID  uint64  `gorm:"not null;index" json:"character_id,string"`
	StoryEventID uint64  `gorm:"not null;index" json:"story_event_id,string"` // 关联的剧情节点ID
	TraitID      uint64  `gorm:"not null;index" json:"trait_id,string"`       // 属性ID（指向CharacterTrait）
	TraitValue   float64 `gorm:"not null;default:50" json:"trait_value"`      // 属性值
	Note         string  `gorm:"size:500" json:"note"`                        // 备注，记录变化原因
}

func init() {
	Models = append(Models, &CharacterState{})
}
