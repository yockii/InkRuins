package service

import "github.com/yockii/inkruins/internal/model"

type CharacterStateService interface {
	CreateCharacterState(characterState *model.CharacterState) error
	GetCharacterStateByID(id uint64) (*model.CharacterState, error)
	GetCharacterStateList(condition *model.CharacterState, page, pageSize int) ([]*model.CharacterState, int64, error)
	UpdateCharacterState(characterState *model.CharacterState) error
	DeleteCharacterState(id uint64) error
}
