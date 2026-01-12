package service

import "github.com/yockii/yunjin/internal/model"

type CharacterService interface {
	CreateCharacter(character *model.Character) error
	GetCharacterByID(id uint64) (*model.Character, error)
	GetCharacterList(projectID uint64, page, pageSize int) ([]*model.Character, int64, error)
	UpdateCharacter(character *model.Character) error
	DeleteCharacter(id uint64) error
}
