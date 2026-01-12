package model

import "gorm.io/datatypes"

// StoryEvent 剧情节点
// 是小说中的关键节点，通常由一个或多个角色参与，描述了角色之间的互动和情节的发展
// 同时作为CharacterState和RelationshipState的时间锚点，用于记录状态变化
type StoryEvent struct {
	BaseModel
	ProjectID      uint64 `gorm:"not null;index" json:"project_id,string"`
	Seq            int    `gorm:"not null;index" json:"seq"`               // 顺序号
	Title          string `gorm:"size:200" json:"title"`                   // 标题
	Summary        string `gorm:"type:text" json:"summary"`                // 场景摘要
	PovCharacterID uint64 `gorm:"not null" json:"pov_character_id,string"` // 视角角色ID
	Goal           string `gorm:"type:text" json:"goal"`                   // POV角色在此场景的目标
	Conflict       string `gorm:"type:text" json:"conflict"`               // 角色冲突/阻碍
	Outcome        string `gorm:"type:text" json:"outcome"`                // 场景结局

	IsUsed bool `gorm:"not null;index" json:"is_used"` // 是否使用过(写进章节)

	// 扩展字段（可选）
	ChapterID               *uint64        `gorm:"index" json:"chapter_id,string"`             // 所属章节ID（可选）
	SceneType               string         `gorm:"size:50" json:"scene_type"`                  // 场景类型（室内/室外/虚拟/梦境等）
	Location                string         `gorm:"size:255" json:"location"`                   // 地理位置
	TimeSetting             string         `gorm:"size:100" json:"time_setting"`               // 时间设定
	EmotionalTone           string         `gorm:"size:100" json:"emotional_tone"`             // 情绪基调（愉快/紧张/浪漫/残酷等）
	ParticipatingCharacters datatypes.JSON `gorm:"type:jsonb" json:"participating_characters"` // 参与角色ID列表
}

func init() {
	Models = append(Models, &StoryEvent{})
}
