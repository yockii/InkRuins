package impl

import (
	"errors"

	"github.com/yockii/inkruins/internal/dao"
	"github.com/yockii/inkruins/internal/database"
	"github.com/yockii/inkruins/internal/model"
)

type CharacterTraitServiceImpl struct{}

func NewCharacterTraitService() *CharacterTraitServiceImpl {
	return &CharacterTraitServiceImpl{}
}

func (s *CharacterTraitServiceImpl) CreateCharacterTrait(characterTrait *model.CharacterTrait) error {
	if characterTrait == nil {
		return errors.New("characterTrait is nil")
	}
	if characterTrait.ProjectID == 0 {
		return errors.New("project_id is required")
	}
	if characterTrait.Name == "" {
		return errors.New("name is required")
	}

	if *characterTrait.MinValue > *characterTrait.MaxValue {
		return errors.New("min_value must be less than max_value")
	}
	if *characterTrait.DefaultValue < *characterTrait.MinValue || *characterTrait.DefaultValue > *characterTrait.MaxValue {
		return errors.New("default_value must be between min_value and max_value")
	}

	return database.DB.Create(characterTrait).Error
}

func (s *CharacterTraitServiceImpl) GetCharacterTraitByID(id uint64) (*model.CharacterTrait, error) {
	if id == 0 {
		return nil, errors.New("id is required")
	}
	var characterTrait model.CharacterTrait
	if err := database.DB.Where(dao.BaseModel.ID.Eq(id)).First(&characterTrait).Error; err != nil {
		return nil, err
	}
	return &characterTrait, nil
}

func (s *CharacterTraitServiceImpl) GetCharacterTraitList(condition *model.CharacterTrait, page, pageSize int) ([]*model.CharacterTrait, int64, error) {
	if condition == nil {
		condition = &model.CharacterTrait{}
	}
	if page <= 0 {
		page = 1
	}
	if pageSize <= 0 {
		pageSize = 10
	}
	tx := database.DB.Model(&model.CharacterTrait{})
	if condition.ProjectID > 0 {
		tx.Where(dao.CharacterTrait.ProjectID.Eq(condition.ProjectID))
	}
	if condition.Name != "" {
		tx.Where(dao.CharacterTrait.Name.Like("%" + condition.Name + "%"))
	}
	var total int64
	if err := tx.Count(&total).Error; err != nil {
		return nil, 0, err
	}
	if total == 0 {
		return nil, 0, nil
	}

	var characterTraits []*model.CharacterTrait
	if err := tx.Order(dao.BaseModel.ID.Desc()).
		Offset((page - 1) * pageSize).
		Limit(pageSize).
		Find(&characterTraits).Error; err != nil {
		return nil, 0, err
	}
	return characterTraits, total, nil
}

func (s *CharacterTraitServiceImpl) UpdateCharacterTrait(characterTrait *model.CharacterTrait) error {
	if characterTrait == nil {
		return errors.New("characterTrait is nil")
	}
	if characterTrait.ID == 0 {
		return errors.New("id is required")
	}
	oldTrait, err := s.GetCharacterTraitByID(characterTrait.ID)
	if err != nil {
		return err
	}
	minValue := oldTrait.MinValue
	if characterTrait.MinValue != nil {
		minValue = characterTrait.MinValue
	}
	maxValue := oldTrait.MaxValue
	if characterTrait.MaxValue != nil {
		maxValue = characterTrait.MaxValue
	}
	defaultValue := oldTrait.DefaultValue
	if characterTrait.DefaultValue != nil {
		defaultValue = characterTrait.DefaultValue
	}

	if *minValue > *maxValue {
		return errors.New("min_value must be less than max_value")
	}
	if *defaultValue < *minValue || *defaultValue > *maxValue {
		return errors.New("default_value must be between min_value and max_value")
	}
	return database.DB.Save(characterTrait).Error
}

func (s *CharacterTraitServiceImpl) DeleteCharacterTrait(id uint64) error {
	if id == 0 {
		return errors.New("id is required")
	}
	if err := database.DB.Delete(&model.CharacterTrait{}, id).Error; err != nil {
		return err
	}
	return nil
}
