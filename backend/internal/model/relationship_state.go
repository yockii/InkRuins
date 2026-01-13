package model

// RelationshipState 关系状态快照
// 主要对关系的各个维度值进行记录，支持完整的历史追溯和趋势分析
type RelationshipState struct {
	BaseModel
	ProjectID      uint64   `gorm:"not null;index" json:"project_id,string"`
	RelationshipID uint64   `gorm:"not null;index" json:"relationship_id,string"` // 关联的CharacterRelationship
	StoryEventID   uint64   `gorm:"not null;index" json:"story_event_id,string"`  // 关联的剧情节点ID
	TraitID        uint64   `gorm:"not null;index" json:"trait_id,string"`        // 关系维度ID（指向RelationshipTrait）
	TraitValue     *float64 `gorm:"not null;default:50" json:"trait_value"`       // 维度值（如亲密度、仇恨度等）
	Note           string   `gorm:"size:500" json:"note"`                         // 备注，记录变化原因
}

func init() {
	Models = append(Models, &RelationshipState{})
}
