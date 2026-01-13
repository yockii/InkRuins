package service

import "github.com/yockii/inkruins/internal/model"

type CharacterTraitService interface {
	CreateCharacterTrait(characterTrait *model.CharacterTrait) error
	GetCharacterTraitByID(id uint64) (*model.CharacterTrait, error)
	GetCharacterTraitList(condition *model.CharacterTrait, page, pageSize int) ([]*model.CharacterTrait, int64, error)
	UpdateCharacterTrait(characterTrait *model.CharacterTrait) error
	DeleteCharacterTrait(id uint64) error
}
