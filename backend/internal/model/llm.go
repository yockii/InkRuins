package model

type Llm struct {
	BaseModel
	ProviderID uint64 `gorm:"not null;index" json:"provider_id,string"`
	ModelName  string `gorm:"not null" json:"model_name"`
	MaxContext *int   `gorm:"not null" json:"max_context"`
	MaxOutput  *int   `gorm:"not null" json:"max_output"`
	// 向量模型有向量维度
	MinEmbeddingDim *int `gorm:"not null" json:"min_embedding_dim"` // 最小向量维度
	MaxEmbeddingDim *int `gorm:"not null" json:"max_embedding_dim"` // 最大向量维度
	// 视觉、音频等模型
	SupportMultimodal int `gorm:"not null" json:"support_multimodal"` // 多模态支持能力 0-不支持 1-图像 2-音频 4-视频
	// 输出内容类型 0-文本 1-图片 2-音频 3-视频
	OutputType int `gorm:"not null;default:0" json:"output_type"` // 输出内容类型 0-文本 1-图片 2-音频 3-视频
}

func init() {
	Models = append(Models, &Llm{})
}
