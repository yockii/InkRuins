package model

// CharacterRelationship 角色关系定义
// 说明：此表仅定义"A与B有关系"，具体的关系类型和维度值在RelationshipState中记录
type CharacterRelationship struct {
	BaseModel
	ProjectID      uint64 `gorm:"not null;index" json:"project_id,string"`
	CharacterID    uint64 `gorm:"not null;index" json:"character_id,string"`     // 角色A
	RelationCharID uint64 `gorm:"not null;index" json:"relation_char_id,string"` // 角色B

	// 可选的关系描述（用于记录关系的整体说明）
	Description string `gorm:"type:text" json:"description"`
}

func init() {
	Models = append(Models, &CharacterRelationship{})
}
