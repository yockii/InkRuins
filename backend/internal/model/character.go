package model

// Character 角色表
type Character struct {
	BaseDeleteModel

	ProjectID uint64 `json:"project_id,string" gorm:"not null;index"`
	Name      string `json:"name" gorm:"not null;size:255"`
	RoleType  string `json:"role_type" gorm:"index"` // 角色类型：主角/配角/反派
	Avatar    string `gorm:"size:500" json:"avatar"` // 可选头像URL

	// ⭐ 数值化属性 (0-100)
	PopularityScore int `gorm:"index;default:0" json:"popularity_score"` // 人气值
	ImportanceScore int `gorm:"index;default:0" json:"importance_score"` // 重要度
	ActivityScore   int `gorm:"index;default:0" json:"activity_score"`   // 活跃度

	// 雪花写作法字段（第三阶段：人物设定）
	Personality string `gorm:"type:text" json:"personality"`  // 性格特点
	Motivation  string `gorm:"type:text" json:"motivation"`   // 动机
	Goal        string `gorm:"type:text" json:"goal"`         // 目标
	Conflict    string `gorm:"type:text" json:"conflict"`     // 冲突
	Epiphany    string `gorm:"type:text" json:"epiphany"`     // 顿悟
	DetailedBio string `gorm:"type:text" json:"detailed_bio"` // 人物宝典（背景、性格、外貌等详细描述）
	Notes       string `gorm:"type:text" json:"notes"`        // 备注

}

func init() {
	Models = append(Models, &Character{})
}
