package impl

import (
	"errors"

	"github.com/yockii/inkruins/internal/dao"
	"github.com/yockii/inkruins/internal/database"
	"github.com/yockii/inkruins/internal/model"
)

type CharacterStateServiceImpl struct{}

func NewCharacterStateService() *CharacterStateServiceImpl {
	return &CharacterStateServiceImpl{}
}

func (s *CharacterStateServiceImpl) CreateCharacterState(characterState *model.CharacterState) error {
	if characterState == nil {
		return errors.New("characterState is nil")
	}
	if characterState.ProjectID == 0 {
		return errors.New("project_id is required")
	}
	if characterState.CharacterID == 0 {
		return errors.New("character_id is required")
	}
	if characterState.TraitID == 0 {
		return errors.New("trait_id is required")
	}
	if characterState.TraitValue == nil {
		return errors.New("trait_value is required")
	}
	return database.DB.Create(characterState).Error
}

func (s *CharacterStateServiceImpl) GetCharacterStateByID(id uint64) (*model.CharacterState, error) {
	if id == 0 {
		return nil, errors.New("id is required")
	}
	var characterState model.CharacterState
	if err := database.DB.Where(dao.BaseModel.ID.Eq(id)).First(&characterState).Error; err != nil {
		return nil, err
	}
	return &characterState, nil
}

func (s *CharacterStateServiceImpl) GetCharacterStateList(condition *model.CharacterState, page, pageSize int) ([]*model.CharacterState, int64, error) {
	if condition == nil {
		condition = &model.CharacterState{}
	}
	if page <= 0 {
		page = 1
	}
	if pageSize <= 0 {
		pageSize = 10
	}
	tx := database.DB.Model(&model.CharacterState{})
	if condition.ProjectID > 0 {
		tx = tx.Where(dao.CharacterState.ProjectID.Eq(condition.ProjectID))
	}
	if condition.CharacterID > 0 {
		tx = tx.Where(dao.CharacterState.CharacterID.Eq(condition.CharacterID))
	}
	if condition.TraitID > 0 {
		tx = tx.Where(dao.CharacterState.TraitID.Eq(condition.TraitID))
	}
	var total int64
	if err := tx.Count(&total).Error; err != nil {
		return nil, 0, err
	}
	if total == 0 {
		return nil, 0, nil
	}
	var characterStates []*model.CharacterState
	if err := tx.Order(dao.BaseModel.ID.Desc()).Offset((page - 1) * pageSize).Limit(pageSize).Find(&characterStates).Error; err != nil {
		return nil, 0, err
	}
	return characterStates, total, nil
}
